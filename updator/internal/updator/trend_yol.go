package updator

import (
	"ClikShop/common/apibitrix"
	"ClikShop/common/bases"
	"ClikShop/common/trendyol"
	"fmt"
)

func (s *Service) UpdateTrendYol(ProductsDetail apibitrix.Product_Response, priceFunc func(brand string, price float64) float64) ([]apibitrix.Variation_Request, error) {

	// Решение задачи сличения данных из битрикса и из донора

	// Все уникальнейшие ссылки на вариации товаров
	links := make(map[string]bool)
	for _, Prod := range ProductsDetail.Products[0].Colors {
		links[Prod.Link] = true
	}

	// Создаём
	ColorPriceExit := make([]trendyol.ColorPriceExit, 0, len(ProductsDetail.Products))
	for link := range links {
		var productID int
		_, err := fmt.Sscanf(link, trendyol.Product_URL, &productID)
		if err != nil {
			return nil, fmt.Errorf("parse link '%s': %w", link, err)
		}

		pg, err := trendyol.ParseProduct(productID)
		if err != nil {
			return nil, fmt.Errorf("trendyol.ParseProduct: %w", err)
		}
		ColorPriceExit = append(ColorPriceExit, trendyol.Touch2ColorPriceExit(pg)...)
	}

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
	for _, Item := range ColorPriceExit {
		keyColorSize := key{color: bases.Name2Slug(Item.Color), size: naakt(bases.Name2Slug(Item.Size)).String()}
		DonMap[keyColorSize] = apibitrix.Variation_Request{
			Price:        priceFunc("trendyol", Item.Price),
			Availability: Item.IsExit,
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
	variationReq := make([]apibitrix.Variation_Request, 0, len(BxMap))
	// формирование слайза запроса на обновление данных со всеми входными характеристиками
	for _, BxVal := range BxMap {
		variationReq = append(variationReq, BxVal)
	}

	s.Gol.Info(fmt.Sprintf("TrandYol: В товаре %s  на обвновление идут %d товара",
		ProductsDetail.Products[0].ID, len(variationReq)))
	return variationReq, nil
}
