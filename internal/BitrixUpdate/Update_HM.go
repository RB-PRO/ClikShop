package bitrixupdate

import (
	"fmt"
	"strings"
	"time"

	"ClikShop/common/apibitrix"
	"ClikShop/common/bases"
	hm "ClikShop/common/hm"
)

// Ключ для хэш-мапы для определения каждой вариации
type key struct {
	size  string
	color string
}

// Обновить цены и наличие по ОДНОМУ товару
func (bx *BitrixUpdator) UpdateHandM(ProductsDetail apibitrix.Product_Response) ([]apibitrix.Variation_Request, error) {

	// Запрос данных по наличию товаров на HM
	SKUhm := ProductsDetail.Products[0].Link
	SKUhm = strings.ReplaceAll(SKUhm, "https://www2.hm.com/tr_tr/productpage.", "")
	SKUhm = strings.ReplaceAll(SKUhm, ".html", "")
	Avalibs, Aavailability := hm.Availability(ProductsDetail.Products[0].Link)
	// Тут надо хендлить ошибки, чтобы отделить ошибки сети от товаров которых просто нет
	if Aavailability != nil {
		return nil, fmt.Errorf("hm.Availability: Не получилось получить данные по артикулу %s из ссылки %s", SKUhm[:7], ProductsDetail.Products[0].Link)
	}
	// fmt.Println(Avalibs)
	time.Sleep(1300 * time.Millisecond)

	AvalibMap, ErrAavailability2 := hm.AavailabilityMap(ProductsDetail.Products[0].Link)
	// и тут бы тоже хенлдить ошибки нормально
	if ErrAavailability2 != nil {
		return nil, fmt.Errorf("hm.Aavailability2: %v: Не получилось получить данные по артикулу %s из ссылки %s", ErrAavailability2, SKUhm[:7], ProductsDetail.Products[0].Link)
	}
	time.Sleep(1300 * time.Millisecond)

	// Формирование мапы наличия для каждой вариации
	MapAvalibs := make(map[key]bool)
	for _, Avalib := range Avalibs {
		lenstr := len([]byte(Avalib))
		if lenstr == 13 {
			MapAvalibs[key{size: AvalibMap[Avalib[lenstr-3:]], color: Avalib[:10]}] = true
		}
		// MapAvalibs[key{size: hm.StrFromSKU(Avalib), color: Avalib[:10]}] = true
		// MapAvalibs[key{size: hm.StrFromSKU2(Avalib), color: Avalib[:10]}] = true
		// MapAvalibs[key{size: hm.StrFromSKU3(Avalib), color: Avalib[:10]}] = true
	}

	// Мапа вариаций, котоыре лежат в битиксе, пара значений размер+цвет обозначают каждую вариацию
	// Правда вмето size по факту у меня 10 символов SKU с HM
	BxMap := make(map[key]apibitrix.Variation_Request)
	for _, Prod := range ProductsDetail.Products[0].Colors {
		SKU := Prod.Link
		SKU = strings.ReplaceAll(SKU, "https://www2.hm.com/tr_tr/productpage.", "")
		SKU = strings.ReplaceAll(SKU, ".html", "")
		BxMap[key{size: Prod.Size, color: SKU[:10]}] =
			apibitrix.Variation_Request{
				ID: Prod.ID,
			}
	}
	// fmt.Println(BxMap)

	// Делаем мапу цен, где в качестве ключа используется артикул товара(7 символов)
	PriceMap := make(map[string]float64)
	for _, Prod := range ProductsDetail.Products[0].Colors {
		SKU := Prod.Link
		SKU = strings.ReplaceAll(SKU, "https://www2.hm.com/tr_tr/productpage.", "")
		SKU = strings.ReplaceAll(SKU, ".html", "")
		if _, ok := PriceMap[SKU]; !ok {
			Price, ErrVariablePrice2 := hm.VariablePrice2(SKU)
			// handle error pls me
			if ErrVariablePrice2 != nil {
				//return nil, fmt.Errorf("hm.VariablePrice2: Не получилось получить данные цене по артикулу %s из ссылки https://www2.hm.com/tr_tr/productpage/_jcr_content/product.quickbuy.%s.html", SKUhm, SKUhm)
			}
			time.Sleep(1500 * time.Millisecond)
			PriceMap[SKU] = Price
		}
	}
	// fmt.Println(PriceMap)

	// Алгоритм обхода по результатам bx.Product в соответствии с massimodutti.Toucher
	// с целью созданию нового запросника для обновления данных в bitrix. Сложность o(n*n) - ужасная
	variationReq := make([]apibitrix.Variation_Request, 0)
	for BxKey, BxVal := range BxMap {
		if hmVal, ok := MapAvalibs[BxKey]; ok {
			// fmt.Println(BxKey.color, "PRICE", (bx.cb.Data.Valute.Try.Value / 10), PriceMap[BxKey.color], bx.MapCoast["H&M"].Walrus)
			Price := bases.EditDecadense((bx.BX.CB.Data.Valute.Try.Value/10)*PriceMap[BxKey.color]*bx.BX.MapCoast["H&M"].Walrus +
				float64(bx.BX.MapCoast["H&M"].Delivery))
			variationReq = append(variationReq, apibitrix.Variation_Request{
				ID:           BxVal.ID,
				Availability: hmVal,
				Price:        Price,
			})
			delete(BxMap, BxKey)
		}
	}

	// Теперь готовим обновление по товарам, котоыре недоступны
	for BxKey, BxVal := range BxMap {
		Price := bases.EditDecadense((bx.BX.CB.Data.Valute.Try.Value/10)*PriceMap[BxKey.color]*bx.BX.MapCoast["H&M"].Walrus +
			float64(bx.BX.MapCoast["H&M"].Delivery))
		variationReq = append(variationReq, apibitrix.Variation_Request{
			ID:           BxVal.ID,
			Availability: false,
			Price:        Price,
		})
	}

	bx.BX.Log.Info(fmt.Sprintf("HM: В товаре %s  на обвновление идут %d товара",
		ProductsDetail.Products[0].ID, len(variationReq)))
	return variationReq, nil
}
