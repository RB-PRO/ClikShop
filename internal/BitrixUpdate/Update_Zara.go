package bitrixupdate

import (
	"fmt"
	"strings"

	zaratr "github.com/RB-PRO/ClikShop/pkg/ZaraTR"
	"github.com/RB-PRO/ClikShop/pkg/apibitrix"
	"github.com/RB-PRO/ClikShop/pkg/bases"
)

// Обновить цены и наличие по ОДНОМУ товару
func (bx *BitrixUpdator) UpdateZara(ProductsDetail apibitrix.Product_Response) ([]apibitrix.Variation_Request, error) {

	// Ссылки на все вариации в подтоваре
	Link := ProductsDetail.Products[0].Link
	Code := strings.ReplaceAll(Link, "https://www.zara.com/tr/en/", "")
	Code = strings.ReplaceAll(Code, ".html?ajax=true", "")

	Prod2, ErrTouch := zaratr.LoadFantomTouch(Code) // Выполняем запрос
	if ErrTouch != nil {
		// fmt.Println(fmt.Errorf("touch: %s", ErrTouch))
		return nil, fmt.Errorf("touch: %s", ErrTouch)
	}

	// Решение задачи сличения данных из битрикса и из донора

	// Мапа вариаций, котоыре лежат в битиксе, пара значений размер+цвет обозначают каждую вариацию
	// Правда вмето size по факту у меня 10 символов SKU с HM
	BxMap := make(map[key]apibitrix.Variation_Request)
	for _, Prod := range ProductsDetail.Products[0].Colors {
		BxMap[key{size: bases.Name2Slug(Prod.Size), color: bases.Name2Slug(Prod.ColorEng)}] =
			apibitrix.Variation_Request{
				ID:    Prod.ID,
				Price: Prod.Price,
			}
	}
	// fmt.Println("BxMap", BxMap)

	// Теперь донорская мапа с данными по товарами со специфичной структурой в качестве ключа
	DonMap := make(map[key]apibitrix.Variation_Request)
	for _, Item := range Prod2.Item {
		for _, Size := range Item.Size {
			Price := bases.EditDecadense((bx.BX.CB.Data.Valute.Try.Value/10)*Item.Price*bx.BX.MapCoast["zara"].Walrus +
				float64(bx.BX.MapCoast["zara"].Delivery))
			DonMap[key{color: bases.Name2Slug(Item.ColorEng), size: bases.Name2Slug(Size.Val)}] = apibitrix.Variation_Request{
				Price:        Price,
				Availability: Size.IsExit,
			}

			if Size.Val == "XXL" {
				DonMap[key{color: bases.Name2Slug(Item.ColorEng), size: bases.Name2Slug("xxxl")}] = apibitrix.Variation_Request{
					Price:        Price,
					Availability: Size.IsExit,
				}
			}
		}
	}
	// fmt.Println("BxMap", DonMap)

	// Теперь объединяется всё в единую мапу битрикса
	for BxKey, BxVal := range BxMap {
		BxVal.Availability = DonMap[BxKey].Availability
		BxVal.Price = DonMap[BxKey].Price
		BxMap[BxKey] = BxVal
	}
	// fmt.Println("BxMap", BxMap)

	// Алгоритм обхода по результатам bx.Product в соответствии с massimodutti.Toucher
	// с целью созданию нового запросника для обновления данных в bitrix. Сложность o(n*n) - ужасная
	variationReq := make([]apibitrix.Variation_Request, 0)
	// формирование слайза запроса на обновление данных со всеми входными характеристиками
	for _, BxVal := range BxMap {
		variationReq = append(variationReq, BxVal)
		// fmt.Printf("%+v\n", BxVal)
	}

	bx.BX.Log.Info(fmt.Sprintf("Zara: В товаре %s(%s) на обвновление идут %d товара",
		ProductsDetail.Products[0].ID, Link, len(variationReq)))
	return variationReq, nil
}
