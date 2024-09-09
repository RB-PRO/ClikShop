package updator

import (
	"ClikShop/common/apibitrix"
	"fmt"
	"strings"
	"time"
)

func (s *Service) UpdateHandM(productsDetail apibitrix.Product_Response, priceFunc func(brand string, price float64) float64) ([]apibitrix.Variation_Request, error) {

	// Запрос данных по наличию товаров на HM
	SKUhm := productsDetail.Products[0].Link
	SKUhm = strings.ReplaceAll(SKUhm, "https://www2.hm.com/tr_tr/productpage.", "")
	SKUhm = strings.ReplaceAll(SKUhm, ".html", "")
	availability, err := s.hmService.Availability(productsDetail.Products[0].Link)

	// Тут надо хендлить ошибки, чтобы отделить ошибки сети от товаров которых просто нет
	if err != nil {
		return nil, fmt.Errorf("hm.Availability: Не получилось получить данные по артикулу %s из ссылки %s: %w", SKUhm[:7], productsDetail.Products[0].Link, err)
	}

	time.Sleep(100 * time.Millisecond)

	availableMap, err := s.hmService.AavailabilityMap(productsDetail.Products[0].Link)
	// и тут бы тоже хенлдить ошибки нормально
	if err != nil {
		return nil, fmt.Errorf("hm.Aavailability2: %v: Не получилось получить данные по артикулу %s из ссылки %s", err, SKUhm[:7], productsDetail.Products[0].Link)
	}
	time.Sleep(100 * time.Millisecond)

	// Формирование мапы наличия для каждой вариации
	variableAvailableMap := make(map[key]bool)
	for _, available := range availability {
		lenStr := len([]byte(available))
		if lenStr == 13 {
			variableAvailableMap[key{size: availableMap[available[lenStr-3:]], color: available[:10]}] = true
		}
	}

	// Мапа вариаций, котоыре лежат в битиксе, пара значений размер+цвет обозначают каждую вариацию
	// Правда вмето size по факту у меня 10 символов SKU с HM
	bxMap := make(map[key]apibitrix.Variation_Request)
	for _, Prod := range productsDetail.Products[0].Colors {
		sku := Prod.Link
		sku = strings.ReplaceAll(sku, "https://www2.hm.com/tr_tr/productpage.", "")
		sku = strings.ReplaceAll(sku, ".html", "")
		bxMap[key{size: Prod.Size, color: sku[:10]}] =
			apibitrix.Variation_Request{
				ID: Prod.ID,
			}
	}

	// Делаем мапу цен, где в качестве ключа используется артикул товара(7 символов)
	priceMap := make(map[string]float64)
	for _, prod := range productsDetail.Products[0].Colors {
		sku := strings.ReplaceAll(prod.Link, "https://www2.hm.com/tr_tr/productpage.", "")
		sku = strings.ReplaceAll(sku, ".html", "")
		if _, ok := priceMap[sku]; !ok {
			Price, ErrVariablePrice2 := s.hmService.VariablePrice2(sku)
			// handle error pls me
			if ErrVariablePrice2 != nil {
				//return nil, fmt.Errorf("hm.VariablePrice2: Не получилось получить данные цене по артикулу %s из ссылки https://www2.hm.com/tr_tr/productpage/_jcr_content/product.quickbuy.%s.html", SKUhm, SKUhm)
			}
			time.Sleep(1500 * time.Millisecond)
			priceMap[sku] = Price
		}
	}

	// Алгоритм обхода по результатам bx.Product в соответствии с massimodutti.Toucher
	// с целью созданию нового запросника для обновления данных в bitrix. Сложность o(n*n) - ужасная
	variationReq := make([]apibitrix.Variation_Request, 0)
	for bxKey, bxVal := range bxMap {
		if hmVal, ok := variableAvailableMap[bxKey]; ok {
			variationReq = append(variationReq, apibitrix.Variation_Request{
				ID:           bxVal.ID,
				Availability: hmVal,
				Price:        priceFunc("H&M", priceMap[bxKey.color]),
			})
			delete(bxMap, bxKey)
		}
	}

	// Теперь готовим обновление по товарам, котоыре недоступны
	for bxKey, bxVal := range bxMap {
		variationReq = append(variationReq, apibitrix.Variation_Request{
			ID:           bxVal.ID,
			Availability: false,
			Price:        priceFunc("H&M", priceMap[bxKey.color]),
		})
	}

	s.Gol.Info(fmt.Sprintf("HM: В товаре %s  на обвновление идут %d товара",
		productsDetail.Products[0].ID, len(variationReq)))
	return variationReq, nil
}
