package sneaksup

import (
	"fmt"
	"strconv"
	"strings"

	"ClikShop/common/bases"
	"github.com/gocolly/colly"
)

func Aavailability(Link string) (Colors []bases.ColorItem, Err error) {

	// Получить мапу ссылок
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36"
	var Price float64
	var Color string

	// Цена
	c.OnHTML(`ul[class="price-list list-inline mb-30 mobile-col-price"]>li:last-of-type`,
		func(e *colly.HTMLElement) {
			text := e.DOM.Text()
			text = strings.ReplaceAll(text, "TL", "")
			text = strings.ReplaceAll(text, ",", ".")
			text = strings.TrimSpace(text)
			Price, _ = strconv.ParseFloat(text, 64)
		})

	// Цвет
	c.OnHTML(`div[id=menu4]>p[class="m-0 mr-2"]`, func(e *colly.HTMLElement) {
		if e.DOM.Find("span:first-child").Text() == "Renk" {
			Color = e.DOM.Find("span:last-child").Text()
			Color = strings.TrimSpace(Color)

		}
	})

	// Наличие
	c.OnHTML(`div[class="raffle-modal-attributes-wrapper"]>div[class="size-select-container"]>ul`, func(e *colly.HTMLElement) {

		sizes := make([]bases.Size, 0)
		e.ForEach(`li[class="size-options-item position-relative in-stock-attribute-item"]>label>span:first-child`, func(i int, e *colly.HTMLElement) {
			size := e.DOM.Text()
			sizes = append(sizes, bases.Size{
				Val:    bases.Name2Slug(size),
				IsExit: true,
			})
		})

		Colors = append(Colors, bases.ColorItem{
			Price:    Price,
			ColorEng: bases.Name2Slug(Color),
			Size:     sizes,
		})
	})

	// Допустимые цвета
	c.OnHTML(`ul[class="detail-color-list list-inline d-md-none"]>li>a`, func(e *colly.HTMLElement) {
		if Link != URL+e.Attr("href") {
			e.Request.Visit(URL + e.Attr("href"))
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		Err = fmt.Errorf("LineUrl2: Request URL: %v failed with response: Error: %v", r.Request.URL, err)
	})

	Err = c.Visit(Link)

	// variationReq := make([]Variation_Request, 0)

	return Colors, nil
}
