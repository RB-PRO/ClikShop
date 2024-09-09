package bitrixupdate

import (
	"fmt"

	"ClikShop/common/apibitrix"
	"ClikShop/common/bases"
	"ClikShop/common/trendyol"
)

// Обновить цены и наличие по ОДНОМУ товару
func (bx *BitrixUpdator) UpdateTrandYol(ProductsDetail apibitrix.Product_Response) ([]apibitrix.Variation_Request, error) {

	// Решение задачи сличения данных из битрикса и из донора

	// Все уникальнейшие ссылки на вариации товаров
	links := make(map[string]bool)
	for _, Prod := range ProductsDetail.Products[0].Colors {
		links[Prod.Link] = true
	}

	// Создаём
	ColorPriceExit := make([]trendyol.ColorPriceExit, 0, len(ProductsDetail.Products))
	for link := range links {
		var ProductID int
		fmt.Sscanf(link, trendyol.Product_URL, &ProductID)
		pg, ErrPP := trendyol.ParseProduct(ProductID)
		if ErrPP != nil {
			return nil, fmt.Errorf("trendyol.ParseProduct: %v", ErrPP)
		}
		ColorPriceExit = append(ColorPriceExit, trendyol.Touch2ColorPriceExit(pg)...)
	}
	// fmt.Println("ColorPriceExit", ColorPriceExit, "\n ")

	// Мапа вариаций, котоыре лежат в битиксе, пара значений размер+цвет обозначают каждую вариацию
	// Правда вмето size по факту у меня 10 символов SKU с HM
	BxMap := make(map[key]apibitrix.Variation_Request)
	for _, Prod := range ProductsDetail.Products[0].Colors {

		//fmt.Printf("$+$ '%s'. - '%s' '%s'. '%s'\n", bases.Name2Slug(Prod.ColorEng), Prod.Size, bases.Name2Slug(Prod.Size), naaktstring(bases.Name2Slug(Prod.Size)))

		BxMap[key{size: naaktstring(bases.Name2Slug(Prod.Size)), color: bases.Name2Slug(Prod.ColorEng)}] =
			apibitrix.Variation_Request{
				ID:    Prod.ID,
				Price: Prod.Price,
			}
	}
	// fmt.Println("BxMap 1", BxMap, "\n ")

	// Теперь донорская мапа с данными по товарами со специфичной структурой в качестве ключа
	DonMap := make(map[key]apibitrix.Variation_Request)
	for _, Item := range ColorPriceExit {
		Price := bases.EditDecadense((bx.BX.CB.Data.Valute.Try.Value/10)*Item.Price*bx.BX.MapCoast["trendyol"].Walrus +
			float64(bx.BX.MapCoast["trendyol"].Delivery))
		DonMap[key{color: (bases.Name2Slug(Item.Color)), size: naaktstring(bases.Name2Slug(Item.Size))}] = apibitrix.Variation_Request{
			Price:        Price,
			Availability: Item.IsExit,
		}
	}
	// fmt.Println("DonMap	", DonMap, "\n ")

	// Теперь объединяется всё в единую мапу битрикса
	for BxKey, BxVal := range BxMap {
		BxVal.Availability = DonMap[BxKey].Availability
		BxVal.Price = DonMap[BxKey].Price
		BxMap[BxKey] = BxVal
	}
	// fmt.Println("BxMap 2", BxMap, "\n ")

	// Алгоритм обхода по результатам bx.Product в соответствии с massimodutti.Toucher
	// с целью созданию нового запросника для обновления данных в bitrix. Сложность o(n*n) - ужасная
	variationReq := make([]apibitrix.Variation_Request, 0, len(BxMap))
	// формирование слайза запроса на обновление данных со всеми входными характеристиками
	for _, BxVal := range BxMap {
		variationReq = append(variationReq, BxVal)
		// fmt.Printf("variationReq: %+v\n", BxVal)
	}

	// bx.BX.Log.Info(fmt.Sprintf("TrandYol: В товаре %s  на обвновление идут %d товара",
	// 	ProductsDetail.Products[0].ID, len(variationReq)))
	return variationReq, nil
}
