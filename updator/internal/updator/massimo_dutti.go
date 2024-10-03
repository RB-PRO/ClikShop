package updator

import (
	massimodutti "ClikShop/common/MassimoDutti"
	"ClikShop/common/apibitrix"
	"ClikShop/common/bases"
	"fmt"
	"strconv"
	"strings"
)

// Обновить цены и наличие по ОДНОМУ товару
func (s *Service) updateMassimoDutti(ProductsDetail apibitrix.Product_Response, priceFunc func(brand string, price float64) float64) ([]apibitrix.Variation_Request, error) {

	Link := ProductsDetail.Products[0].Link // Основная ссылка на товар

	// Получение ID товара в системе massimodutti. Оно же Toucher
	Link = strings.ReplaceAll(Link, "https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/", "")
	Link = strings.ReplaceAll(Link, "/detail?languageId=-1&appId=1", "")
	ID, ErrAtoi := strconv.Atoi(Link)
	if ErrAtoi != nil {
		return nil, fmt.Errorf("atoi: %w", ErrAtoi)
	}

	// Делаем запрос на получение данных
	touch, err := s.mdService.Toucher(ID)
	if err != nil {
		return nil, fmt.Errorf("toucher: %w", err)
	}
	var Product bases.Product2
	Product = massimodutti.Touch2Product2(Product, touch)
	variationReq := make([]apibitrix.Variation_Request, 0)

	// Решение задачи сличения данных из битрикса и из донора

	// Мапа вариаций, котоыре лежат в битиксе, пара значений размер+цвет обозначают каждую вариацию
	// Правда вмето size по факту у меня 10 символов SKU с HM
	BxMap := make(map[key]apibitrix.Variation_Request)
	for _, Prod := range ProductsDetail.Products[0].Colors {
		if _, ok := BxMap[key{size: bases.Name2Slug(Prod.Size), color: bases.Name2Slug(Prod.ColorEng)}]; ok {
			variationReq = append(variationReq, apibitrix.Variation_Request{
				Availability: false,
				ID:           Prod.ID,
			})
			continue
		}
		BxMap[key{size: bases.Name2Slug(Prod.Size), color: bases.Name2Slug(Prod.ColorEng)}] =
			apibitrix.Variation_Request{
				ID:    Prod.ID,
				Price: Prod.Price,
			}
	}

	// Теперь донорская мапа с данными по товарам со специфичной структурой в качестве ключа
	DonMap := make(map[key]apibitrix.Variation_Request)
	for _, Item := range Product.Item {
		for _, Size := range Item.Size {
			Price := priceFunc("Massimo Dutti", Item.Price)
			DonMap[key{color: bases.Name2Slug(Item.ColorEng), size: bases.Name2Slug(Size.Val)}] = apibitrix.Variation_Request{
				Price:        Price,
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

	// формирование слайза запроса на обновление данных со всеми входными характеристиками
	for _, BxVal := range BxMap {
		variationReq = append(variationReq, BxVal)
	}

	s.Gol.Info(fmt.Sprintf("В товаре %s  на обвновление идут %d товара",
		ProductsDetail.Products[0].ID, len(variationReq)))
	return variationReq, nil
}
