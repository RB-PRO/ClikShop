package hm

import (
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"strings"

	"ClikShop/common/bases"
	"github.com/gocolly/colly"
)

const URL string = "https://www2.hm.com"

type CategorysCat struct {
	Link      string
	Cat       []bases.Cat
	GendetTag string
}

// Получить список всех категорий и ссылки на эти категории
func (s *Service) Categorys() (Category []CategorysCat, ErrParse error) {
	const productCategory string = "Ürüne göre satın al" // Константа, которая содержит название на Турецком языке, которое означает все продукты

	c, err := s.NewServiceCollector()
	if err != nil {
		return nil, errors.Wrap(err, "create service collector: ")
	}

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 YaBrowser/23.5.2.625 Yowser/2.5 Safari/537.36"

	// Проверка на ошибку
	c.OnError(func(r *colly.Response, e error) {
		ErrParse = fmt.Errorf("categorys: Error http reequest %d", r.StatusCode)
	})

	// Категории товаров
	c.OnHTML("ul[class] li div ul li ul li a", func(e *colly.HTMLElement) {
		Link, LinkIsExit := e.DOM.Attr("href") // Ссылка на категорию
		if LinkIsExit {                        // Если существует некоторый тег с ссылкой на подкатегорию
			if SelectorSpan := e.DOM.Parent().Parent().Parent().Find("span"); SelectorSpan.Text() == productCategory {
				if cat, ErrorParseCat := PullOutCat(URL + Link); ErrorParseCat == nil {
					if Filter(cat[0].Slug) {
						NewCategorys := append([]bases.Cat{{Name: "HM", Slug: "hm", ID: 0}}, cat...)
						Category = append(Category, CategorysCat{Link: Link, Cat: NewCategorys, GendetTag: TransCategoryTR(NewCategorys[1].Slug)})
					}
				}
			}
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		ErrParse = fmt.Errorf("categorys: Request URL: %v failed with response: %v Error: %v", r.Request.URL, r, err)
	})

	return Category, c.Visit("https://www2.hm.com/tr_tr/index.html")
}

// Фильтр того, что будем учитывать, а что нет.
func Filter(str string) bool {
	switch str {
	case "kadin":
		return true
	case "erkek":
		return true
	case "bebek":
		return true
	case "cocuk":
		return true
	default:
		return false
	}
}

// Перевести с Турецкого на Английский по категориям
func TransCategoryTR(str string) string {
	switch str {
	case "kadin":
		return "woman"
	case "erkek":
		return "man"
	case "bebek":
		return "boy"
	case "cocuk":
		return "girl"
	default:
		return "unisex"
	}
}

// Перевести ссылку в массив категорий
//
//	Пример:
//
//	`https://www2.hm.com/tr_tr/home/urune-gore-satin-al/dekorasyon.html`
//
// в
//
//	`[]bases.Cat{{Name: "Home", Slug: "home"}, {Name: "Urune Gore Satin Al", Slug: "urune-gore-satin-al"}, {Name: "Dekorasyon", Slug: "dekorasyon"}}`
func PullOutCat(link string) ([]bases.Cat, error) {

	// Парсим ссылку
	u, ErrParse := url.Parse(link)
	if ErrParse != nil {
		return nil, ErrParse
	}

	// Берём "ручку"
	Path := u.Path
	paths := strings.Split(Path, "/")
	if len(paths) != 5 {
		return nil, fmt.Errorf("PullOutCat: len of '%s' is %d. Correct - 5", paths, len(paths))
	}

	// Делаем слайс категорий
	// cat := make([]bases.Cat, 3)
	// for i := 0; i < len(cat); i++ {
	// 	CategoryName := strings.ReplaceAll(paths[i+2], ".html", "")
	// 	// if CategoryName == "urune-gore-satin-al" {
	// 	// 	continue
	// 	// }
	// 	cat[i].Slug = CategoryName
	// 	cat[i].Name = bases.Slug2Name(cat[i].Slug)
	// }
	cat := make([]bases.Cat, 2)
	cat[0].Slug = strings.ReplaceAll(paths[2], ".html", "")
	cat[0].Name = bases.Slug2Name(cat[0].Slug)
	cat[1].Slug = strings.ReplaceAll(paths[4], ".html", "")
	cat[1].Name = bases.Slug2Name(cat[1].Slug)

	return cat, nil
}
