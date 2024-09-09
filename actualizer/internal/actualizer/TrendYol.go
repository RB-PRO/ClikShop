package actualizer

import (
	"fmt"
	"strconv"
	"time"

	"ClikShop/common/bases"
	"ClikShop/common/trendyol"
	"github.com/cheggaaa/pb"
)

// Структура HM для парсинга
type TY struct {
}

func NewTY() *TY {
	return &TY{}
}

// Парсинг данных и сохранение их в файлы
//
//	Заменить во всех файлах нужно символы '\u0026' на '&'
func (s *TY) Scraper() (string, error) {
	folder := "ty"
	_ = ReMakeDir(folder)

	ShopIDs := []int{
		332585, // Levi's
		107483, // Aktaş Sport AS
		815951, // HUGO
		804476, // BOSS
		742918, // Victoria's secret
		106871, // SneakSup

		110890, // New Balance
		230712, // Tommy Hilfiger
		194194, // Guess
		1920,   // Lacoste
		104961, // BERSHKA
		112044, // Pull&Bear
		104723, // MANGO
		155194, // ExxeSelection

	}

	for _, ShopID := range ShopIDs {
		_ = s.trendyolOne(folder, ShopID)
	}

	return folder, nil
}

func (s *TY) trendyolOne(folder string, ShopID int) error {

	ProductGroupIDs, ErrGroup := trendyol.Pages(ShopID)
	if ErrGroup != nil {
		return fmt.Errorf("trendyol.Pages: %v", ErrGroup)
	}

	BarProducts := pb.StartNew(len(ProductGroupIDs))
	defer BarProducts.Finish()
	BarProducts.Prefix(strconv.Itoa(ShopID))
	var Products bases.Variety2
	for _, ProductGroupID := range ProductGroupIDs {

		// Спарсить информацию по товару
		Product, ErrProduct := trendyol.Product(ProductGroupID, ShopID)
		if ErrProduct != nil {
			// fmt.Println(ErrProduct)
			// s.Gol.Warn(fmt.Sprintf("trendyol Product: %v", ErrProduct))
			continue
		}

		// Да-да, может быть такое, что вариаций у товара не будет.
		// Например это может возникнуть, когда продавец вариаций товаров не оригинальный
		if len(Product.Item) != 0 {
			Product.Cat = append([]bases.Cat{
				{
					Name: "trendyol",
					Slug: bases.Name2Slug("trendyol"),
				},
				{
					Name: strconv.Itoa(ShopID),
					Slug: bases.Name2Slug(strconv.Itoa(ShopID)),
				},
			}, Product.Cat...)
			Product = bases.EditOneSize(Product)
			Product = bases.EditDoubleColors(Product)
			Product.Size = bases.EditProdSize(Product)
			Product.Img = bases.EditIMG(Product)
			Products.Product = append(Products.Product, Product)
		}
		BarProducts.Increment()
		time.Sleep(time.Millisecond * 50)
	}

	Products.SaveJson(fmt.Sprintf("%s/trendyol_%d", folder, ShopID))

	return nil
}
