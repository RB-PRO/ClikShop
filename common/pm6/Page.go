package pm6

import (
	"fmt"
	"strconv"
	"time"

	"ClikShop/common/bases"
	"github.com/gocolly/colly"
)

const URL = "https://www.6pm.com"

// Пропасить стрицы товара и записать данные в канал
func ParsePages(ProductChan chan<- bases.Product2) {
	ProductChan <- bases.Product2{Name: "123"}
	time.Sleep(time.Second) // Задержка
}

// Пропасить стрицу с товарами
func ParsePage(link string, page int) (prod []bases.Product2) {
	c := colly.NewCollector()
	c.OnHTML("article[class='qda-z']", func(e *colly.HTMLElement) {
		link, ErrorAttrLink := e.DOM.Find("a[class='OP-z']").Attr("href") // Получить ссылку на товар
		if ErrorAttrLink {
			color := e.DOM.Find("dd[class='UP-z']").Text() // Получить цвет товара
			color = bases.FormingColorEng(color)           // Преобразовать в ссылку
			prod = append(prod, bases.Product2{
				Link: link,
				Name: e.DOM.Find("dd[class='SP-z']").Text(),
				// Item: make(map[string]bases.ProdParam),
				Item: []bases.ColorItem{},
			})
			// prod[len(prod)-1].Item[color] = bases.ProdParam{}
			prod[len(prod)-1].Item = append(prod[len(prod)-1].Item, bases.ColorItem{ColorCode: color})
		}

	})
	c.Visit(URL + link + "&p=" + strconv.Itoa(page))
	return prod
}

// Пропасить стрицу с товарами
func (pmm pm) ParsePageWithVarienty(varient bases.Variety2, link string, page int) bases.Variety2 {
	var TecalName string             // Текущее имя, используемое для осозначения цветов
	LenProds := len(varient.Product) // Получаем к-во товаров
	if LenProds != 0 {
		TecalName = varient.Product[LenProds-1].Name
	}

	var c *colly.Collector
	if pmm.proxy == "" { // Если без прокси
		c = colly.NewCollector() // Instantiate default collector
	} else {
		c = colly.NewCollector(colly.AllowURLRevisit()) // Instantiate default collector
		c.SetProxy(pmm.proxy)                           // Set Proxy
	}
	c.UserAgent = "Golang"
	c.SetRequestTimeout(30 * time.Second) // Установить таймауты 30 секунд

	// Поиск и добавление самой ссылки на товар
	c.OnHTML("div[id=products] article", func(e *colly.HTMLElement) {
		LenProds := len(varient.Product)                                   // Получаем к-во товаров
		link, ErrorAttrLink := e.DOM.Find("a[tabindex='-1']").Attr("href") // Получить ссылку на товар
		if ErrorAttrLink {
			name := e.DOM.Find("dd[itemprop='name']").Text() // Название товара

			color := e.DOM.Find("dd[itemprop='color']").Text() // Получить цвет товара
			color = bases.FormingColorEng(color)               // Преобразовать в ссылку

			//fmt.Println(">"+TecalName+"<", ">"+name+"<")
			// Если нужно дозаписать подтовар
			if TecalName == name {
				_, FindIdError := FindFirstNameProducts(varient.Product, name)
				if FindIdError != nil { // Если не найден такой товар по имени
					// varient.Product[LenProds-1].Item[color] = bases.ProdParam{Link: link, ColorEng: color}
					for index := range varient.Product[LenProds-1].Item {
						if color == varient.Product[LenProds-1].Item[index].ColorCode {
							varient.Product[LenProds-1].Item[index].Link = link
							varient.Product[LenProds-1].Item[index].ColorEng = color
						}
					}
				} else { // Если найден такой ID товара
					// varient.Product[FindIdInt].Item[color] = bases.ProdParam{Link: link, ColorEng: color}
					for index := range varient.Product[LenProds-1].Item {
						if color == varient.Product[LenProds-1].Item[index].ColorCode {
							varient.Product[LenProds-1].Item[index].Link = link
							varient.Product[LenProds-1].Item[index].ColorEng = color
						}
					}

				}
			} else { // Если нужно создать новый товар
				// Добавляем такой товар
				varient.Product = append(varient.Product, bases.Product2{
					Name: name,
					Link: link,
					// Item: make(map[string]bases.ProdParam),
					Item: make([]bases.ColorItem, 0),
				})
				// Добавляем подтовар
				// varient.Product[LenProds].Item[color] = bases.ProdParam{Link: link, ColorEng: color}
				for index := range varient.Product[LenProds-1].Item {
					if color == varient.Product[LenProds-1].Item[index].ColorCode {
						varient.Product[LenProds-1].Item[index].Link = link
						varient.Product[LenProds-1].Item[index].ColorEng = color
					}
				}
			}
			TecalName = name
		}
	})
	c.Visit(URL + link + "&p=" + strconv.Itoa(page))
	fmt.Println(URL + link + "&p=" + strconv.Itoa(page))
	return varient
}

// Получить все страницы с товарами на сайте
func (pmm pm) AllPages(link string) (pages int) {
	fmt.Println(pmm)

	var c *colly.Collector
	if pmm.proxy == "" { // Если без прокси
		c = colly.NewCollector() // Instantiate default collector
	} else {
		c = colly.NewCollector(colly.AllowURLRevisit()) // Instantiate default collector
		c.SetProxy(pmm.proxy)                           // Set Proxy
	}
	c.UserAgent = "Golang"
	c.SetRequestTimeout(30 * time.Second) // Установить таймауты 30 секунд

	//>div>span>a:last-of-type
	c.OnHTML("div[id='searchPagination']>div>span>a:last-of-type", func(e *colly.HTMLElement) {
		pagesStr := e.DOM.Text()
		fmt.Println(pagesStr)
		var ErrorPages error
		pages, ErrorPages = strconv.Atoi(pagesStr)
		if ErrorPages != nil {
			fmt.Println(ErrorPages)
		}
	})
	c.Visit(URL + link)
	return pages
}
