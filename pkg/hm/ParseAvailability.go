package hm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

// Структура наличия товара на складе H&M
type Availability struct {
	Availability []string `json:"availability"` // В наличии
	FewPieceLeft []string `json:"fewPieceLeft"` // Осталось немножка(Красная точка рядом с размером)
}

// Загрузить сведения по размерам, а точнее по их пристутстию
//
// Эти сведения необходимы для уточнения остатка по размером для товара.
//
// Если артикул: 1157823001002, то
//   - фактический артикул(7 цифр) - 1157823
//   - подъартикул(3 цифры) - 001
//   - подцвет(3 цифры) - 002
//
// Example:
//
//	`https://www2.hm.com/hmwebservices/service/product/tr/availability/1157823.json`
func Aavailability(SKU string) ([]string, error) {
	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www2.hm.com/hmwebservices/service/product/tr/availability/%s.json", SKU), nil)
	if ErrNewRequest != nil {
		return nil, ErrNewRequest
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	// Выполнить запорос
	Response, ErrDo := client.Do(req)
	if ErrDo != nil {
		return nil, ErrDo
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return nil, ErrorReadAll
	}

	// Распарсить полученный json в структуру
	var RespData Availability
	ErrorUnmarshal := json.Unmarshal(BodyPage, &RespData)
	if ErrorUnmarshal != nil {
		return nil, ErrorUnmarshal
	}

	return append(RespData.Availability, RespData.FewPieceLeft...), nil
}

// Получить данные размеров товаров
func AvailabilityProduct(Product bases.Product2) (bases.Product2, error) {

	sku := Product.Article
	if len(Product.Article) == 10 {
		sku = Product.Article[:7]
	}

	// Получить все артикулы присутствубщих товаров
	IsLiveSKUs, ErrAvailability := Aavailability(sku)
	if ErrAvailability != nil {
		return Product, ErrAvailability
	}

	// Теперь надо перебрать все возможные размеры и есть размер есть
	// в массиве артикулов имеющихся в наличии то выставляем true
	for i := range Product.Item { // Цикл по всем цветам
		for j := range Product.Item[i].Size { //  цикл по всем размерам цвета
			for _, ValExitSKU := range IsLiveSKUs { // Цикл по всем размерам в наличии
				if Product.Item[i].Size[j].DataCode == ValExitSKU {
					Product.Item[i].Size[j].IsExit = true
				}
			}
		}
	}
	return Product, nil
}
