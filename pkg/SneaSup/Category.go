package sneaksup

import (
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/gocolly/colly"
)

const URL string = "https://www.sneaksup.com"

// Внутренняя категория разделов для подкатегорий
type SScat struct {
	// Name string // название категории, он же пол
	// Slug string // Обозначение категории транслитом
	Link string // Ссылка на категорию
	Cat  []bases.Cat
}

func Category() (cc []SScat) {
	c := colly.NewCollector()

	var RealChildParentName string
	// Find and visit all links
	c.OnHTML(`div[class="MenuLink spec-list"]>a`, func(e *colly.HTMLElement) {

		// Название самой категории
		Name := e.DOM.Text()
		Name = strings.TrimSpace(Name)
		Link, _ := e.DOM.Attr("href")

		// Главная категория - Мужчины, женщины и тд
		ManParentName := e.DOM.Parent().Parent().Parent().Parent().Parent().Parent().Parent().Parent().Parent().Parent().Parent().Find("a>h2").Text()
		ManParentName = strings.TrimSpace(ManParentName)
		ManParentLink, _ := e.DOM.Parent().Parent().Parent().Parent().Parent().Parent().Parent().Parent().Parent().Parent().Parent().Find("a").Attr("href")

		// Подопечная подкатегория категория - кросовки и тд
		ChildParentName := e.DOM.Parent().Parent().Parent().Find("a:first-of-type>span").Text()
		ChildParentName = strings.TrimSpace(ChildParentName)
		ChildParentLink, _ := e.DOM.Parent().Parent().Parent().Find("a:first-of-type").Attr("href")

		if ChildParentName != "" {
			RealChildParentName = ChildParentName
		}

		cc = append(cc, SScat{
			Cat: []bases.Cat{{Name: "sneaksup", Slug: "sneaksup"},
				{Name: ManParentName, Slug: Name2Slug(ManParentLink)},
				{Name: RealChildParentName, Slug: Name2Slug(ChildParentLink)},
				{Name: Name, Slug: Name2Slug(Link)}},
			Link: URL + Link,
		})
	})

	c.Visit(URL)

	return cc
}

// Преобразовать название в путь для ссылки, он же ярлык
func Name2Slug(str string) string {
	str = strings.ReplaceAll(str, "/pages/", "")
	str = strings.ReplaceAll(str, "/", "")
	return str
}
