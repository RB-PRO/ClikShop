package actualizer

import (
	"fmt"
	"time"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/trendyol"
	"github.com/cheggaaa/pb"
)

func (bx *bitrixActualizer) trendyol(folder string) error {

	MakeDir(folder)

	// SneakSupCategory := sneasup.Category()

	// for iCategory, Category := range SneakSupCategory {
	// 	Namecategory := strings.ReplaceAll(Category.Link, "https://public.trendyol.com/", "") // Название категории

	// 	// Загружаем line по категории
	// 	line, ErrLines := sneasup.Lines(Category.Link)
	// 	if ErrLines != nil {
	// 		fmt.Println(ErrLines)
	// 	}

	// 	// Переводим в свою структур товаров
	// 	Products := sneasup.Line2Product(line, Category)

	// 	// Редактирование товара, согласно требованиям
	// 	BarProducts := pb.StartNew(len(Products))
	// 	BarProducts.Prefix(fmt.Sprintf("[%d/%d]", iCategory+1, len(SneakSupCategory)))
	// 	for iProd := range Products {
	// 		// BarCategory.Prefix(fmt.Sprintf("%s [%d/%d]", Namecategory, iProd, len(Products)))
	// 		for iItem := range Products[iProd].Item {
	// 			for iSize := range Products[iProd].Item[iItem].Size {
	// 				Products[iProd].Item[iItem].Size[iSize].Val = strings.TrimSpace(Products[iProd].Item[iItem].Size[iSize].Val)
	// 				Products[iProd].Item[iItem].Size[iSize].Val = strings.ReplaceAll(Products[iProd].Item[iItem].Size[iSize].Val, "Yaş", "лет")
	// 				Products[iProd].Item[iItem].Size[iSize].Val = strings.ReplaceAll(Products[iProd].Item[iItem].Size[iSize].Val, "YAŞ", "лет")
	// 			}
	// 		}
	// 		Products[iProd].Size = bases.EditProdSize(Products[iProd])
	// 		Products[iProd].Description.Eng, _ = sneasup.Description(Products[iProd].Link)
	// 		Products[iProd].Img = bases.EditIMG(Products[iProd])
	// 		img := make([]string, 0)
	// 		for iItem := range Products[iProd].Item {
	// 			img = append(img, Products[iProd].Item[iItem].Image...)
	// 		}
	// 		Products[iProd].Img = img
	// 		BarProducts.Increment()
	// 	}
	// 	BarProducts.Finish()

	// 	// Сохранить в файл
	// 	bases.Variety2{Product: Products}.SaveJson(fmt.Sprintf("%s/ss_%d_%s",
	// 		folder, iCategory, Namecategory))

	// }

	ShopID := 106871 // SneakSup
	ProductGroupIDs, ErrGroup := trendyol.Pages(ShopID)
	if ErrGroup != nil {
		return fmt.Errorf("trendyol.Pages: %v", ErrGroup)
	}

	BarProducts := pb.StartNew(len(ProductGroupIDs))
	BarProducts.Prefix("trendyol")
	var Products bases.Variety2
	for _, ProductGroupID := range ProductGroupIDs {
		// fmt.Println(iProductGroupID, len(ProductGroupIDs))
		Product, ErrProduct := trendyol.Product(ProductGroupID, ShopID)
		if ErrProduct != nil {
			// panic(ErrProduct)
			fmt.Println(ErrProduct)
			continue
		}

		// Да-да, может быть такое, что вариаций у товара не будет.
		// Например это может возникнуть, когда продавец вариаций товаров не оригинальный
		if len(Product.Item) != 0 {
			// Product
			Product = bases.EditDoubleColors(Product)
			Product.Size = bases.EditProdSize(Product)
			Product.Img = bases.EditIMG(Product)

			Products.Product = append(Products.Product, Product)
		}
		BarProducts.Increment()
		time.Sleep(time.Millisecond * 200)
	}
	BarProducts.Finish()
	Products.SaveJson(fmt.Sprintf("%s/trendyol_%d", folder, ShopID))

	return nil
}
