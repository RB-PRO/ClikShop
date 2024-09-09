package hm

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"

	"ClikShop/common/bases"
	"github.com/gocolly/colly"
)

// Структура наличия товара на складе H&M
type AvailabilityStruct struct {
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
func (s *Service) Availability(link string) ([]string, error) {
	c, err := s.NewServiceCollector()
	if err != nil {
		return nil, errors.Wrap(err, "create service collector: ")
	}

	sku := strings.ReplaceAll(link, "https://www2.hm.com/tr_tr/productpage.", "")
	sku = strings.ReplaceAll(sku, ".html", "")
	url := fmt.Sprintf("https://www2.hm.com/hmwebservices/service/product/tr/availability/%s.json", sku[:7])

	// s.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36"
	// req.Header.Add("x-requested-with", "XMLHttpRequest")

	headers := http.Header{}
	headers.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	var response AvailabilityStruct
	c.OnResponse(func(r *colly.Response) {
		if err := json.Unmarshal(r.Body, &response); err != nil {
			log.Println("ERROR:500:", err)
			return
		}
	})
	return append(response.Availability, response.FewPieceLeft...), c.Request(http.MethodGet, url, nil, nil, headers)
}

// Получить мапу, в которой будет продемонтрировано, что для этого товара является актуальным в размерах
func (s *Service) AavailabilityMap(Link string) (map[string]string, error) {
	c, err := s.NewServiceCollector()
	if err != nil {
		return nil, errors.Wrap(err, "create service collector: ")
	}

	MapSize := make(map[string]string)
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
		html = html[a:b]
		html = strings.ReplaceAll(html, "&#34;", `"`)
		html = strings.ReplaceAll(html, "&#39;", `"`)
		type SizeStruct struct {
			SizeCode      string `json:"sizeCode"`
			Size          string `json:"size"`
			Name          string `json:"name"`
			SizeScaleCode string `json:"sizeScaleCode"`
		}
		var RespData []SizeStruct

		if err := json.Unmarshal([]byte(html), &RespData); err != nil {
			e.Response.StatusCode = 500
			return
		}
		// fmt.Println(RespData)
		for _, data := range RespData {
			MapSize[data.Size] = data.Name
		}
	}) //1183407001003
	return MapSize, c.Visit(Link)
}

// Получить данные размеров товаров
func (s *Service) AvailabilityProduct(Product bases.Product2) (bases.Product2, error) {

	sku := Product.Article
	if len(Product.Article) == 10 {
		sku = Product.Article[:7]
	}

	// Получить все артикулы присутствубщих товаров
	IsLiveSKUs, ErrAvailability := s.Availability(sku)
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
