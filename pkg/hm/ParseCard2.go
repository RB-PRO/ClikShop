package hm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/gocolly/colly"
)

// Пропарсить товар и получить возмодные размеры и картинки
//
// Example:
//
//	`https://www2.hm.com/tr_tr/productpage/_jcr_content/product.quickbuy.1157823001.html`
func VariableProduct2(Product bases.Product2) (bases.Product2, error) {
	var Err error
	var Index int
	var TecalSKU string // Текущий артикул для цвета
	// var Sizes []bases.Size
	for Index = range Product.Item {
		Product.Item[Index].Size = []bases.Size{}
	}

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36"

	// Картинки
	c.OnHTML(`div[class="product-detail-thumbnails"]>ul>li>img`, func(e *colly.HTMLElement) {
		ImageLink := e.Attr("src")

		ImageLink = strings.ReplaceAll(ImageLink, "&call=url[file:/product/quickthumb]", "")

		ImageLink = "https:" + ImageLink + "&call=url[file:/product/main]"
		ImageLink = strings.ReplaceAll(ImageLink, "\u0026", "&")
		ImageLink = strings.ReplaceAll(ImageLink, `\u0026`, "&")
		ImageLink = strings.ReplaceAll(ImageLink, "u0026", "&")

		Product.Item[Index].Image = append(Product.Item[Index].Image, ImageLink)
	})

	// Допустимые размеры
	c.OnHTML(`select[data-sizelist]>option[data-code]`, func(e *colly.HTMLElement) {
		if Val, IsExit := e.DOM.Parent().Attr("data-sizelist"); IsExit && TecalSKU == Val {
			Product.Item[Index].Size = append(Product.Item[Index].Size, bases.Size{Val: e.Attr("value"), DataCode: e.Attr("data-code")})
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		Err = fmt.Errorf("LineUrl2: Request URL: %v failed with response: Error: %v", r.Request.URL, err)
	})

	for Index = range Product.Item {
		TecalSKU = Product.Item[Index].Link
		TecalSKU = strings.ReplaceAll(TecalSKU, URL, "")
		TecalSKU = strings.ReplaceAll(TecalSKU, "/tr_tr/productpage.", "")
		TecalSKU = strings.ReplaceAll(TecalSKU, ".html", "")

		// fmt.Println(URL + "/tr_tr/productpage/_jcr_content/product.quickbuy." + TecalSKU + ".html")
		c.Visit(URL + "/tr_tr/productpage/_jcr_content/product.quickbuy." + TecalSKU + ".html")
		if Err != nil {
			return Product, Err
		}
	}

	return Product, nil
}

// получить цену по SKU в 7 символов
func VariablePrice2(sku string) (Price float64, Err error) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36"

	// Допустимые размеры
	c.OnHTML(`span[class="price-value"]`, func(e *colly.HTMLElement) {
		PriceStr := e.DOM.Text()
		PriceStr = strings.ReplaceAll(PriceStr, "TL", "")
		PriceStr = strings.ReplaceAll(PriceStr, ",", "")
		PriceStr = strings.ReplaceAll(PriceStr, ".", "")
		PriceStr = strings.TrimSpace(PriceStr)
		// fmt.Printf("PriceStr '%s'", PriceStr)
		Price, _ = strconv.ParseFloat(PriceStr, 64)
		Price /= 100.0
	})

	Err = c.Visit(URL + "/tr_tr/productpage/_jcr_content/product.quickbuy." + sku + ".html")
	return Price, Err
}

// Пропарсить товар по классической [ссылке] и получить его описание вместе с дополнительными полями
//
// [ссылке]: https://www2.hm.com/tr_tr/productpage.1205348002.html
func VariableDescription2(Product bases.Product2) (bases.Product2, error) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36"
	Product.Specifications = make(map[string]string)

	// Описание
	c.OnHTML(`meta[name=description]`, func(e *colly.HTMLElement) {
		Product.Description.Eng, _ = e.DOM.Attr("content")
	})

	// Вторичное описание
	// c.OnHTML("div[id^=section-descriptionAccordion]>dl>div", func(e *colly.HTMLElement) {
	// 	dt := e.DOM.Find("dt").Text()
	// 	dd := e.DOM.Find("dd").Text()

	// 	fmt.Println(dt, dd)

	// 	dt = strings.ReplaceAll(dt, ":", "")
	// 	dt = strings.TrimSpace(dt)
	// 	// Product.Description.Eng += "\n" + dt + " - " + dd
	// 	Product.Specifications[dt] = dd
	// })
	c.OnHTML(`div[class="content pdp-text pdp-content"]`, func(e *colly.HTMLElement) {
		// fmt.Println(e.DOM.Find("noscript").SetHtml(e.DOM.Find("noscript").Text()))
		html, _ := e.DOM.Html()
		d, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
		d.Find("noscript").SetHtml(d.Find("noscript").Text())
		// fmt.Println(d.Html())
		d.Find("noscript").Each(func(i int, s *goquery.Selection) {
			s.ReplaceWithHtml(s.Text())
		})
		d.Find("div[id=section-descriptionAccordion]>dl>div").Each(func(i int, s *goquery.Selection) {
			dt := s.Find("dt").Text()
			dd := s.Find("dd").Text()

			dt = strings.ReplaceAll(dt, ":", "")
			dt = strings.ReplaceAll(dt, "\n", " ")
			dt = strings.ReplaceAll(dt, "\t", " ")
			dt = strings.ReplaceAll(dt, "  ", " ")
			dt = strings.TrimSpace(dt)
			dd = strings.ReplaceAll(dd, "\n", " ")
			dd = strings.ReplaceAll(dd, "\t", " ")
			dd = strings.ReplaceAll(dd, "  ", " ")
			dd = strings.TrimSpace(dd)
			// Product.Description.Eng += "\n" + dt + " - " + dd
			Product.Specifications[dt] = dd
		})

		// fmt.Println(d.Find("div[id=section-descriptionAccordion]>dl>div").Html())
	})

	c.Visit(fmt.Sprintf("%s/tr_tr/productpage.%s.html", URL, Product.Article))
	return Product, nil
}
