package hm

import (
	"fmt"
	"strings"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
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
	Product.Specifications = make(map[string]string)

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36"

	// Цвета
	c.OnHTML(`div[class="product-detail-thumbnails"]>ul>li>img`, func(e *colly.HTMLElement) {
		ImageLink := e.Attr("src")
		ImageLink = strings.ReplaceAll(ImageLink, "[file:/product/quickthumb]", "[file:/product/main]")
		Product.Item[Index].Image = append(Product.Item[Index].Image, "	https:"+ImageLink)
	})

	// Допустимые размеры
	c.OnHTML(`select[data-sizelist]>option[data-code]`, func(e *colly.HTMLElement) {
		if Val, IsExit := e.DOM.Parent().Attr("data-sizelist"); IsExit && TecalSKU == Val {
			Product.Item[Index].Size = append(Product.Item[Index].Size, bases.Size{Val: e.Attr("value")})
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

		c.Visit(URL + "/tr_tr/productpage/_jcr_content/product.quickbuy." + TecalSKU + ".html")
		if Err != nil {
			return Product, Err
		}
	}

	return Product, nil
}
