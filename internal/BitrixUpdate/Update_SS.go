package bitrixupdate

import (
	"fmt"

	sneaksup "ClikShop/common/SneaSup"
	"ClikShop/common/apibitrix"
	"ClikShop/common/bases"
)

// const URL string = "https://www.sneaksup.com"

func (bx *BitrixUpdator) UpdateSS(ProductsDetail apibitrix.Product_Response) (variationReq []apibitrix.Variation_Request, Err error) {

	// Получить мапу ссылок
	ColorsItem, ErrAavailability := sneaksup.Aavailability(ProductsDetail.Products[0].Link)
	if ErrAavailability != nil {
		return nil, fmt.Errorf("sneaksup.Aavailability: %v", ErrAavailability)
	}
	// fmt.Println(ProductsDetail)
	// fmt.Println(ColorsItem)

	// Решение задачи сличения данных из битрикса и из донора

	// Мапа вариаций, котоыре лежат в битиксе, пара значений размер+цвет обозначают каждую вариацию
	// Правда вмето size по факту у меня 10 символов SKU с HM
	BxMap := make(map[key]apibitrix.Variation_Request)
	for _, Prod := range ProductsDetail.Products[0].Colors {
		BxMap[key{size: naaktstring(bases.Name2Slug(Prod.Size)), color: bases.Name2Slug(Prod.ColorEng)}] =
			apibitrix.Variation_Request{
				ID:    Prod.ID,
				Price: Prod.Price,
			}
	}
	// 	fmt.Println("BxMap", BxMap)

	// Теперь донорская мапа с данными по товарами со специфичной структурой в качестве ключа
	DonMap := make(map[key]apibitrix.Variation_Request)
	for _, Item := range ColorsItem {
		for _, Size := range Item.Size {
			Price := bases.EditDecadense((bx.BX.CB.Data.Valute.Try.Value/10)*Item.Price*bx.BX.MapCoast["ss"].Walrus +
				float64(bx.BX.MapCoast["ss"].Delivery))
			DonMap[key{color: (bases.Name2Slug(Item.ColorEng)), size: naaktstring(bases.Name2Slug(Size.Val))}] = apibitrix.Variation_Request{
				Price:        Price,
				Availability: Size.IsExit,
			}
		}
	}
	// fmt.Println("DonMap", DonMap)

	// Теперь объединяется всё в единую мапу битрикса
	for BxKey, BxVal := range BxMap {
		BxVal.Availability = DonMap[BxKey].Availability
		BxVal.Price = DonMap[BxKey].Price
		BxMap[BxKey] = BxVal
	}
	// fmt.Println("BxMap", BxMap)

	// Алгоритм обхода по результатам bx.Product в соответствии с massimodutti.Toucher
	// с целью созданию нового запросника для обновления данных в bitrix. Сложность o(n*n) - ужасная
	variationReq = make([]apibitrix.Variation_Request, 0)
	// формирование слайза запроса на обновление данных со всеми входными характеристиками
	for _, BxVal := range BxMap {
		variationReq = append(variationReq, BxVal)
		// fmt.Printf("%+v\n", BxVal)
	}

	return variationReq, nil
}
