package hm

import (
	"fmt"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/gocolly/colly"
)

func Product(url string) (Prod bases.Product2, ErrParse error) {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	return Prod, nil
}
