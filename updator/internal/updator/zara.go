package updator

import (
	"ClikShop/common/apibitrix"
	"ClikShop/common/bases"
	"fmt"
	"strings"
)

func (s *Service) UpdateZara(ProductsDetail apibitrix.Product_Response, priceFunc func(brand string, price float64) float64) ([]apibitrix.Variation_Request, error) {

	// Ссылки на все вариации в подтоваре
	link := ProductsDetail.Products[0].Link
	code := strings.ReplaceAll(link, "https://www.zara.com/tr/en/", "")
	code = strings.ReplaceAll(code, ".html?ajax=true", "")

	Prod2, ErrTouch := s.zaraService.LoadFantomTouch(code) // Выполняем запрос
	if ErrTouch != nil {
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

	// Теперь донорская мапа с данными по товарами со специфичной структурой в качестве ключа
	DonMap := make(map[key]apibitrix.Variation_Request)
	for _, Item := range Prod2.Item {
		for _, Size := range Item.Size {
			price := priceFunc("zara", Item.Price)
			DonMap[key{color: bases.Name2Slug(Item.ColorEng), size: bases.Name2Slug(Size.Val)}] = apibitrix.Variation_Request{
				Price:        price,
				Availability: Size.IsExit,
			}

			if Size.Val == "XXL" {
				DonMap[key{color: bases.Name2Slug(Item.ColorEng), size: bases.Name2Slug("xxxl")}] = apibitrix.Variation_Request{
					Price:        price,
					Availability: Size.IsExit,
				}
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
	variationReq := make([]apibitrix.Variation_Request, 0)
	// формирование слайза запроса на обновление данных со всеми входными характеристиками
	for _, BxVal := range BxMap {
		variationReq = append(variationReq, BxVal)
	}

	s.Gol.Info(fmt.Sprintf("Zara: В товаре %s(%s) на обвновление идут %d товара",
		ProductsDetail.Products[0].ID, link, len(variationReq)))
	return variationReq, nil
}
