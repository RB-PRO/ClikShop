package bitrixupdate

import (
	"fmt"
	"strconv"
	"strings"

	massimodutti "github.com/RB-PRO/SanctionedClothing/pkg/MassimoDutti"
	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

func Start() {
	bx := NewBitrixUser()

	// Загружаем цены
	Coasts, ErrCoasts := bx.Coasts()
	if ErrCoasts != nil {
		panic(ErrCoasts)
	}
	fmt.Println(Coasts)

	// Получаем списки товаров
	ProductsID, ErrProducts := bx.Products()
	if ErrProducts != nil {
		panic(ErrProducts)
	}
	fmt.Println("В Bitrix всего", len(ProductsID), "товаров.")

	// Цикл по всем товарам
	for _, ProductID := range ProductsID {

		// Обновляем данные по товару
		ErrUpdateProduct := bx.UpdateProduct(ProductID)
		if ErrUpdateProduct != nil {
			fmt.Println(ErrUpdateProduct)
		}

		break
	}
}

// Обновить цены и наличие по ОДНОМУ товару
func (bx *BitrixUser) UpdateProduct(ProductID string) error {

	// Получить подробнее о товаре
	ProductsDetail, ErrProduct := bx.Product([]string{ProductID})
	if ErrProduct != nil {
		return fmt.Errorf("bitrix: UpdateProduct: %w", ErrProduct)
	}
	if len(ProductsDetail.Products) == 0 {
		return fmt.Errorf("bitrix: UpdateProduct: %w", ErrProduct)
	}
	fmt.Printf("%+v\n\n", ProductsDetail)

	Link := ProductsDetail.Products[0].Link
	switch {
	case strings.Contains(Link, "massimodutti"):
		// Получение ID товара в системе massimodutti. Оно же Toucher
		Link = strings.ReplaceAll(Link, "https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/", "")
		Link = strings.ReplaceAll(Link, "/detail?languageId=-1&appId=1", "")
		ID, ErrAtoi := strconv.Atoi(Link)
		if ErrAtoi != nil {
			return fmt.Errorf("bitrix: UpdateProduct: massimodutti: Atoi: %w", ErrAtoi)
		}

		// Делаем запрос на получение данных
		touch, ErrToucher := massimodutti.Toucher(ID)
		if ErrToucher != nil {
			return fmt.Errorf("bitrix: UpdateProduct: massimodutti: Toucher: %w", ErrToucher)
		}
		var Product bases.Product2
		Product = massimodutti.Touch2Product2(Product, touch)

		// Алгоритм обхода по результатам bx.Product в соответствии с massimodutti.Toucher
		// с целью созданию нового запросника для обновления данных в bitrix. Сложность o(n*n) - ужасная
		variationReq := make([]Variation_Request, 0)

		for _, BXproduct := range ProductsDetail.Products[0].Colors {
			for _, ProdItem := range Product.Item {
				fmt.Println("BXproduct.ColorEng, ProdItem.ColorEng", BXproduct.ColorEng, ProdItem.ColorEng)
				if strings.Contains(BXproduct.ColorEng, ProdItem.ColorEng) {
					for _, ProdColor := range ProdItem.Size {
						fmt.Println("BXproduct.Size, ProdColor.Val", BXproduct.Size, ProdColor.Val)
						if strings.Contains(BXproduct.Size, ProdColor.Val) {
							variationReq = append(variationReq, Variation_Request{
								ID: BXproduct.ID,
								Price: ProdItem.Price*bx.MapCoast[Product.Manufacturer].Walrus +
									float64(bx.MapCoast[Product.Manufacturer].Delivery),
								Availability: ProdColor.IsExit,
							})
						}
					}
				}
			}
		}
		fmt.Printf("%+v\n", variationReq)

		// VariationResp, ErrVariation := bx.Variation(variationReq)
		// if ErrVariation != nil {
		// 	t.Error(ErrVariation)
		// }
	default:
		return fmt.Errorf("bitrix: UpdateProduct: Не знаю, какую логику применить к '%s'", ProductsDetail.Products[0].Link)
	}

	return nil
}
