package hm

import (
	"fmt"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/gocolly/colly"
)

const URL string = "https://www2.hm.com"

type CategorysCat struct {
	Link string
	bases.Cat
}

// Получить список всех категорий и ссылки на эти категории
func Categorys() (Category []CategorysCat, ErrParse error) {
	const ProductCategory string = "Ürüne göre satın al" // Константа, которая содержит название на Турецком языке, которое означает все продукты

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36"

	// Проверка на ошибку
	c.OnError(func(r *colly.Response, e error) {
		ErrParse = fmt.Errorf("categorys: Error http reequest %d", r.StatusCode)
	})

	// Категории товаров
	c.OnHTML("ul[class=MLEL] li div ul li ul li a", func(e *colly.HTMLElement) {
		Link, LinkIsExit := e.DOM.Attr("href") // Ссылка на категорию
		if LinkIsExit {
			HeaderLine := e.DOM.Parent().Parent().Parent().Parent().Parent().Parent().Find("a").Text()
			HeaderLine = e.DOM.Parent().Parent().Parent().Parent().Parent().Parent().Empty().Children().Filter("a").Text()
			HeaderLine, _ = e.DOM.Parent().Parent().Parent().Parent().Parent().Parent().First().Html()

			if e.DOM.Parent().Parent().Parent().Find("span").Text() == ProductCategory {
				fmt.Println(e.DOM.Text(), ">"+HeaderLine+"<", Filter(HeaderLine), URL+Link)
				fmt.Println(">" + e.DOM.Parent().Parent().Parent().Parent().Parent().Children().Text())
				fmt.Println(">" + e.DOM.Parent().Parent().Parent().Parent().Parent().Find("a:nth-child(1)").Text())
				fmt.Println(">" + e.DOM.Parent().Parent().Parent().Parent().Parent().Find("a:nth-of-type(1)").Text())
				fmt.Println(">" + e.DOM.Parent().Parent().Parent().Parent().Parent().Find("a:only-of-type").Text())
				fmt.Println(">" + e.DOM.Parent().Parent().Parent().Parent().Parent().Find("a:first-child").Text())
				fmt.Println(">" + e.DOM.Parent().Parent().Parent().Parent().Parent().Find("a:empty").Text())
				fmt.Println(">" + e.DOM.Parent().Parent().Parent().Parent().Parent().Contents().Find("a:first-child").Text())
				fmt.Println()
			}
		}
	})

	c.Visit("https://www2.hm.com/tr_tr/index.html")

	return Category, nil
}

func Filter(str string) bool {
	switch str {
	case "H&M HOME":
		return false
	default:
		return true
	}
}
