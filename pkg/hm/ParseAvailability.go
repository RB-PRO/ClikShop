package hm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/gocolly/colly"
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
func Aavailability(Link string) ([]string, error) {

	SKUhm := strings.ReplaceAll(Link, "https://www2.hm.com/tr_tr/productpage.", "")
	SKUhm = strings.ReplaceAll(SKUhm, ".html", "")

	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www2.hm.com/hmwebservices/service/product/tr/availability/%s.json", SKUhm[:7]), nil)
	if ErrNewRequest != nil {
		return nil, fmt.Errorf("http.NewRequest: %v", ErrNewRequest)
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	// Выполнить запорос
	Response, ErrDo := client.Do(req)
	if ErrDo != nil {
		return nil, fmt.Errorf("client.Do: %v", ErrDo)
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return nil, fmt.Errorf("io.ReadAll: %v", ErrorReadAll)
	}

	// Распарсить полученный json в структуру
	var RespData Availability
	ErrorUnmarshal := json.Unmarshal(BodyPage, &RespData)
	if ErrorUnmarshal != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", ErrorUnmarshal)
	}

	return append(RespData.Availability, RespData.FewPieceLeft...), nil
}

// Получить мапу, в которой будет продемонтрировано, что для этого товара является актуальным в размерах
func AavailabilityMap(Link string) (MapSize map[string]string, Err error) {
	MapSize = make(map[string]string)
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.837 YaBrowser/23.9.4.837 Yowser/2.5 Safari/537.36"

	// // Find and visit all links
	// c.OnHTML(`ul[class^=picker-list]>li`, func(e *colly.HTMLElement) {
	// 	datacode := e.Attr("data-code")
	// 	fmt.Println("datacode", datacode)
	// 	if datacode != "" {
	// 		MapSize[datacode[10:]] = e.DOM.Find("div>button>span:first-of-type").Text()
	// 	}
	// }) //1183407001003

	c.OnHTML(`div[class="catalogwarning parbase"]`, func(e *colly.HTMLElement) {
		html, _ := e.DOM.Next().Html()
		var a, b int
		a = strings.Index(html, "&#39;sizes&#39;") + 16
		b = strings.Index(html[a:], "]") + 1
		b += a
		// fmt.Println(html[a:b])
		html = html[a:b]
		html = strings.ReplaceAll(html, "&#34;", `"`)
		html = strings.ReplaceAll(html, "&#39;", `"`)
		// html = `{"sizes":` + html + `}`
		// fmt.Println(html)
		type SizeStruct struct {
			SizeCode      string `json:"sizeCode"`
			Size          string `json:"size"`
			Name          string `json:"name"`
			SizeScaleCode string `json:"sizeScaleCode"`
		}
		var RespData []SizeStruct
		ErrorUnmarshal := json.Unmarshal([]byte(html), &RespData)
		if ErrorUnmarshal != nil {
			Err = fmt.Errorf("json.Unmarshal: %v", ErrorUnmarshal)
			return
		}
		// fmt.Println(RespData)
		for _, data := range RespData {
			MapSize[data.Size] = data.Name
		}
	}) //1183407001003
	Err = c.Visit(Link)
	return MapSize, Err
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
