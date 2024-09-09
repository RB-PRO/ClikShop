package updator

import (
	sneaksup "ClikShop/common/SneaSup"
	"ClikShop/common/apibitrix"
	"ClikShop/common/bases"
	"fmt"
)

func (s *Service) UpdateSS(ProductsDetail apibitrix.Product_Response, priceFunc func(brand string, price float64) float64) (variationReq []apibitrix.Variation_Request, Err error) {

	// Получить мапу ссылок
	colorsItem, err := sneaksup.Aavailability(ProductsDetail.Products[0].Link)
	if err != nil {
		return nil, fmt.Errorf("sneaksup.Aavailability: %w", err)
	}

	// Решение задачи сличения данных из битрикса и из донора

	// Мапа вариаций, котоыре лежат в битиксе, пара значений размер+цвет обозначают каждую вариацию
	// Правда вмето size по факту у меня 10 символов SKU с HM
	BxMap := make(map[key]apibitrix.Variation_Request)
	for _, Prod := range ProductsDetail.Products[0].Colors {
		keyColorSize := key{color: bases.Name2Slug(Prod.ColorEng), size: naakt(bases.Name2Slug(Prod.Size)).String()}
		BxMap[keyColorSize] =
			apibitrix.Variation_Request{
				ID:    Prod.ID,
				Price: Prod.Price,
			}
	}

	// Теперь донорская мапа с данными по товарами со специфичной структурой в качестве ключа
	DonMap := make(map[key]apibitrix.Variation_Request)
	for _, Item := range colorsItem {
		for _, Size := range Item.Size {
			keyColorSize := key{color: bases.Name2Slug(Item.ColorEng), size: naakt(bases.Name2Slug(Size.Val)).String()}
			DonMap[keyColorSize] = apibitrix.Variation_Request{
				Price:        priceFunc("ss", Item.Price),
				Availability: Size.IsExit,
			}
		}
	}

	// Теперь объединяется всё в единую мапу битрикса
	for BxKey, BxVal := range BxMap {
		BxVal.Availability = DonMap[BxKey].Availability
		BxVal.Price = DonMap[BxKey].Price
		BxMap[BxKey] = BxVal
	}

	// Алгоритм обхода по результатам bx.Product в соответствии с massimodutti.Toucher
	// с целью созданию нового запросника для обновления данных в bitrix. Сложность o(n*n) - ужасная
	variationReq = make([]apibitrix.Variation_Request, 0)
	// формирование слайза запроса на обновление данных со всеми входными характеристиками
	for _, BxVal := range BxMap {
		variationReq = append(variationReq, BxVal)
	}

	return variationReq, nil
}
