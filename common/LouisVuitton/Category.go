package louisvuitton

import (
	"strings"

	"github.com/gocolly/colly"
)

type Categ struct {
	CategoryTag string // Ссылка на категорию
	Path        string // Путь категории
}

// Загрузить список категорий товаров
func Category() (Categs []Categ, Err error) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36"

	c.OnHTML(`ul[class="lv-megamenu-list__items lv-list -level1"] a[class^=lv-smart-link]`, func(e *colly.HTMLElement) { // class^=lv-smart-link

		// Получить ссылку на категорию
		link, _ := e.DOM.Attr("href")

		// При необходимости добавляем категорию
		categ := Link2Cat(link)
		if categ.CategoryTag != "" {
			Categs = append(Categs, categ)
		}
	})

	// c.OnHTML("div[class=lv-mega-menu__lock]", func(e *colly.HTMLElement) {
	// 	// fmt.Println(e.DOM.Html())
	// 	data, _ := e.DOM.Html() // сохраняем текст элемента и переходим на новую строку

	// 	file, err := os.Create("listing.html")
	// 	if err != nil {
	// 		fmt.Println("Ошибка при создании файла:", err)
	// 		return
	// 	}
	// 	defer file.Close()

	// 	_, err = file.WriteString(data)
	// 	if err != nil {
	// 		fmt.Println("Ошибка при записи в файл:", err)
	// 		return
	// 	}
	// })

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		Err = err
	})

	c.Visit("https://ru.louisvuitton.com/rus-ru/homepage")

	return Categs, nil
}

// Пересобрать путь до категории на сайте в путь до папки с фото на локальной машине
func Link2Cat(link string) (categ Categ) {
	ll := strings.Split(link, "/")

	if len(ll) == 7 { // /rus-ru/art-of-living/books-and-stationery/hard-cover-books/_/N-t1134j9w
		categ.CategoryTag = ll[6]
		categ.CategoryTag = categ.CategoryTag[2:]
		categ.Path = "LV/" + strings.Join(ll[2:5], "/") + "/"
	}
	if len(ll) == 6 { // /rus-ru/women/handbags/_/N-tfr7qdp
		categ.CategoryTag = ll[5]
		categ.CategoryTag = categ.CategoryTag[2:]
		categ.Path = "LV/" + strings.Join(ll[2:4], "/") + "/"
	}
	if len(ll) == 4 { // /rus-ru/stories/gifting
		categ.CategoryTag = ""
		categ.Path = "LV/" + strings.Join(ll[2:], "/") + "/"
	}
	if len(ll) == 3 { // /rus-ru/magazine
		categ.CategoryTag = ""
		categ.Path = "LV/" + ll[2] + "/"
	}
	return categ
}
