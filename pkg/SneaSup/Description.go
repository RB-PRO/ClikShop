package sneaksup

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// Сделать запрос на описание товара
func Description(url string) (Description string, ErrorResp error) {
	// Create a collector
	c := colly.NewCollector()

	// Set HTML callback
	c.OnHTML(`div[class$=product-detail-tabs]>div[id=menu4]>p`, func(e *colly.HTMLElement) {
		Description += e.DOM.Text() + "\n"
	})

	c.OnHTML(`div[class$=product-detail-tabs]>div[id=menu4]>div[class=mt-2]`, func(e *colly.HTMLElement) {
		txt := e.DOM.Text()
		txt = strings.TrimSpace(txt)
		txt = strings.ReplaceAll(txt, "\t", " ")
		txt = strings.ReplaceAll(txt, "  ", " ")
		txt = strings.ReplaceAll(txt, "\n\n", "\n")
		txt = strings.TrimSpace(txt)

		Description += txt
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		ErrorResp = fmt.Errorf("Description Error: %v", err)
	})

	// Start scraping
	c.Visit(url)

	return Description, ErrorResp
}
