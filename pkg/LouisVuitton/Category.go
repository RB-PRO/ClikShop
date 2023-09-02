package louisvuitton

import (
	"fmt"
	"os"
	"time"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/gocolly/colly"
)

type Categ struct {
	Link string      // Ссылка на категорию
	Cat  []bases.Cat // Путь категории
}

func Category() (Categs []Categ, Err error) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36"
	c.SetRequestTimeout(time.Minute)

	c.OnHTML("div[class=lv-mega-menu__lock]", func(e *colly.HTMLElement) {
		// fmt.Println(e.DOM.Html())
		data, _ := e.DOM.Html() // сохраняем текст элемента и переходим на новую строку

		file, err := os.Create("listing.html")
		if err != nil {
			fmt.Println("Ошибка при создании файла:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(data)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	})

	c.OnHTML(`ul[class="lv-megamenu-list__items lv-list -level1"] a[class^=lv-smart-link]`, func(e *colly.HTMLElement) { // class^=lv-smart-link

		var cat []bases.Cat
		cat = append(cat, bases.Cat{Name: e.Text}) // Сама категория. Самая нижняя
		link, _ := e.DOM.Attr("href")

		Categs = append(Categs, Categ{Link: link, Cat: cat})
	})

	// // Set error handler
	// c.OnError(func(r *colly.Response, err error) {
	// 	Err = err
	// })

	c.Visit("https://ru.louisvuitton.com/rus-ru/homepage")

	fmt.Println("Листинг сайта сохранен в файл listing.txt")

	return Categs, nil
}
