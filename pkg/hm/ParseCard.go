package hm

import (
	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/gocolly/colly"
)

func Product(url string) (Prod bases.Product2, ErrParse error) {
	c := colly.NewCollector()

	// Цвета
	c.OnHTML("li[class=list-item]>a", func(e *colly.HTMLElement) {
		Color, ColorIsExit := e.DOM.Attr("title")
		Link, LinkIsExit := e.DOM.Attr("href")
		if ColorIsExit && LinkIsExit {
			Prod.Item = append(Prod.Item, bases.ColorItem{
				Link:     Link,
				ColorEng: Color,
			})
		}
	})

	c.Visit(url)

	return Prod, nil
}
