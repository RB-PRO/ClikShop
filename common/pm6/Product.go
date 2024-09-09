package pm6

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"ClikShop/common/bases"

	"github.com/gocolly/colly"
)

type pm struct {
	proxy string
}

// Создаём экземпляр структуры парсинга pm6 со встроенным proxy
func NewPM() (*pm, error) {
	proxy, ErrorFile := bases.DataFile("proxy")
	if ErrorFile != nil {
		return nil, errors.New("не удалось открыть файл \"proxy\"")
	}
	return &pm{proxy}, nil
}

// Парсинг страницы товара
// Парсинг будет выглядеть в виде редактирования структуры bases.Product2 со своевременным добавлением цвета
func (pmm pm) ParseProduct(prod *bases.Product2, ProductColorLink string) {
	var tecalColor string // Цвет текущей страницы
	var c *colly.Collector
	if pmm.proxy == "" { // Если без прокси
		c = colly.NewCollector() // Instantiate default collector
	} else {
		c = colly.NewCollector(colly.AllowURLRevisit()) // Instantiate default collector
		c.SetProxy(pmm.proxy)                           // Set Proxy
	}
	c.UserAgent = "Golang"
	c.SetRequestTimeout(30 * time.Second) // Установить таймауты 30 секунд

	// Создаём структуру цвета
	c.OnHTML("form[method='POST']>div[class]:first-of-type>div[class]>span:last-of-type", func(e *colly.HTMLElement) {
		ColorFull := e.DOM.Text()
		ColorFull = strings.TrimSpace(ColorFull)
		if ColorFull != "" {
			tecalColor = bases.FormingColorEng(ColorFull)
			// prod.Item[tecalColor] = bases.ProdParam{ColorEng: ColorFull}
			for index := range prod.Item {
				if prod.Item[index].ColorCode == tecalColor {
					prod.Item[index].ColorEng = ColorFull
				}
			}
			prod.Specifications = make(map[string]string)
		}
	})

	// Артикул
	c.OnHTML("div[role='region'] span[itemprop='sku']", func(e *colly.HTMLElement) {
		prod.Article = e.DOM.Text()
	})

	// описание товара
	c.OnHTML("div[role='region'] ul li", func(e *colly.HTMLElement) {

		if strings.Contains(e.DOM.Text(), "Measurements:") {
			// Обработка дополнительных полей
			// Вынес в отдельный обработчик.
			return
		}

		// Обработка обычного описания товара
		if prod.Description.Eng == "" {
			prod.Description.Eng = e.Text
		} else {
			prod.Description.Eng += "\n" + e.Text
		}
	})

	// Описание товара по ключам
	c.OnHTML("div[role='region'] ul li ul li", func(e *colly.HTMLElement) {
		KeyValStr := strings.Split(e.DOM.Text(), ":")
		if len(KeyValStr) == 2 {
			KeyValStr[0] = strings.TrimSpace(KeyValStr[0])
			KeyValStr[1] = strings.TrimSpace(KeyValStr[1])
			prod.Specifications[KeyValStr[0]] = KeyValStr[1]

		}
	})

	// Размеры товара
	c.OnHTML("form[id=buyBoxForm]>div>fieldset>div[class]>div[class]>input", func(e *colly.HTMLElement) {
		if attr, ok := e.DOM.Attr("data-label"); ok { // Если такой атрибут существует
			for index := range prod.Item {
				if prod.Item[index].ColorCode == tecalColor {
					prod.Size = append(prod.Size, attr)
					prod.Item[index].Size = append(prod.Item[index].Size, bases.Size{Val: attr, IsExit: true})
				}
			}

			// if entry, ok := prod.Item[tecalColor]; ok {
			// 	entry.Size = append(entry.Size, attr)
			// 	prod.Size = append(prod.Size, attr)
			// 	prod.Item[tecalColor] = entry
			// }
		}
	})

	// Картинки - Не работает.
	c.OnHTML("div[id=productThumbnails] div ul li button picture img", func(e *colly.HTMLElement) {
		if sourseValue, isFind := e.DOM.Attr("src"); isFind { // Если есть аттрибут src
			// if entry, oks := prod.Item[tecalColor]; oks { // То добавляем его
			// 	// Берём из общей ссылки на маленькую картинку, базовую ссылку на основную картинку
			// 	if sourseValue, exitImgCodeerror := PictureCode(sourseValue); exitImgCodeerror == nil {
			// 		// Добавляем картинку в массив
			// 		entry.Image = append(entry.Image, "https://m.media-amazon.com/images/I/"+sourseValue+".jpg")
			// 		prod.Item[tecalColor] = entry
			// 	}
			// }

			for index := range prod.Item {
				if prod.Item[index].ColorCode == tecalColor {
					// Берём из общей ссылки на маленькую картинку, базовую ссылку на основную картинку
					if sourseValue, exitImgCodeerror := PictureCode(sourseValue); exitImgCodeerror == nil {
						// Добавляем картинку в массив
						prod.Item[index].Image = append(prod.Item[index].Image, "https://m.media-amazon.com/images/I/"+sourseValue+".jpg")
					}
				}
			}
		}
	})

	// Категории + производитель
	c.OnHTML("div[id=breadcrumbs] div", func(e *colly.HTMLElement) {
		prod.Cat[1].Name = e.DOM.Find("a:nth-of-type(2)").Text()
		prod.Cat[1].Slug, _ = e.DOM.Find("a:nth-of-type(2)").Attr("href")
		prod.Cat[1].Slug = formSlump(prod.Cat[1].Slug, 1) // Редактирование ссылки

		prod.Cat[2].Name = e.DOM.Find("a:nth-of-type(3)").Text()
		prod.Cat[2].Slug, _ = e.DOM.Find("a:nth-of-type(3)").Attr("href")
		prod.Cat[2].Slug = formSlump(prod.Cat[2].Slug, 1) // Редактирование ссылки

		prod.Cat[3].Name = e.DOM.Find("a:nth-of-type(4)").Text()
		prod.Cat[3].Slug, _ = e.DOM.Find("a:nth-of-type(4)").Attr("href")
		prod.Cat[3].Slug = formSlump(prod.Cat[3].Slug, 2) // Редактирование ссылки
		prod.Manufacturer = prod.Cat[3].Name              // Производитель
	})

	// Название Товара
	c.OnHTML("div[id='productRecap'] div div div div h1 div span:last-of-type", func(e *colly.HTMLElement) {
		//c.OnHTML("meta[itemprop=name]", func(e *colly.HTMLElement) {
		prod.Name = e.DOM.Text()
	})

	// Полное название товара, оно же краткое описание товара
	c.OnHTML("div[role='region'] ul:first-of-type li[class]:first-of-type", func(e *colly.HTMLElement) {
		prod.FullName = e.DOM.Text()
	})

	// Ссылка на товар
	c.OnHTML("div[itemprop=offers] meta[itemprop=url]", func(e *colly.HTMLElement) {
		if link, linkFind := e.DOM.Attr("content"); linkFind {
			prod.Link = link // Записать ссылку в продукт
			// Если есть такой
			// if entry, oks := prod.Item[tecalColor]; oks { // То добавляем его
			// 	entry.Link = link
			// 	prod.Item[tecalColor] = entry
			// }

			for index := range prod.Item {
				if prod.Item[index].ColorCode == tecalColor {
					prod.Item[index].Link = link
				}
			}
		}
	})

	// Гендер товара
	c.OnHTML("form[id=buyBoxForm] div fieldset", func(e *colly.HTMLElement) {
		// В блоке размеров есть тег legend с аттрибутом id="sizingChooser"
		if _, isFind := e.DOM.Find("legend[id='sizingChooser']").Attr("id"); isFind {
			textSize := e.DOM.Find("legend span").Text()
			textSize = strings.ReplaceAll(textSize, "'s Sizes:", "") // Удалить лишнее из гендера
			textSize = strings.ReplaceAll(textSize, "s Sizes:", "")  // Удалить лишнее из гендера
			textSize = strings.ReplaceAll(textSize, "s Size:", "")   // Удалить лишнее из гендера
			textSize = strings.ReplaceAll(textSize, "Size:", "")     // Удалить лишнее из гендера
			textSize = strings.ReplaceAll(textSize, "'s sizes:", "") // Удалить лишнее из гендера
			textSize = strings.ReplaceAll(textSize, "s sizes:", "")  // Удалить лишнее из гендера
			textSize = strings.ReplaceAll(textSize, "s size:", "")   // Удалить лишнее из гендера
			textSize = strings.ReplaceAll(textSize, "size:", "")     // Удалить лишнее из гендера
			textSize = strings.ReplaceAll(textSize, "Little ", "")   // Удалить лишнее из гендера

			textSize = strings.TrimSpace(textSize)                               // Удалить лишнее из гендера
			prod.Cat[0].Name, textSize, _ = bases.GenderBook(textSize, textSize) // Название главной категории товара
			textSize = strings.ToLower(textSize)                                 // Понизить регистр
			prod.Cat[0].Slug = textSize                                          // Название главной ссылки категории товара
			prod.GenderLabel = textSize                                          // Заполнить гендер
		}

		// Для товара с одним цветом, по типу https://www.6pm.com/p/2xu-non-stirrup-calf-guard-white-white/product/7892154/color/1001
		//if _, isFind := e.DOM.Find("legend[id='screenReadersOnly']"); isFind {
		//}
	})

	// Цена
	c.OnHTML("span[itemprop=price]", func(e *colly.HTMLElement) {
		coast, findCoast := e.DOM.Attr("aria-label")
		if findCoast {
			coast = strings.ReplaceAll(coast, "$", "")
			floaCoast, errCoast := strconv.ParseFloat(coast, 64) // Преобразование типов
			if errCoast == nil {
				// if entry, oks := prod.Item[tecalColor]; oks { // То добавляем его
				// 	entry.Price = floaCoast
				// 	prod.Item[tecalColor] = entry
				// }

				for index := range prod.Item {
					if prod.Item[index].ColorCode == tecalColor {
						prod.Item[index].Price = floaCoast
					}
				}
			}
		}
	})

	// Размеры для товаров по типу https://www.6pm.com/p/2xu-non-stirrup-calf-guard-white-white/product/7892154/color/1001
	c.OnHTML("select[id='pdp-size-select']>option[value]", func(e *colly.HTMLElement) {
		// if entry, ok := prod.Item[tecalColor]; ok {
		// 	if e.DOM.Text() != "Select a Size" {
		// 		entry.Size = append(entry.Size, e.DOM.Text())
		// 		prod.Size = append(prod.Size, e.DOM.Text())
		// 		prod.Item[tecalColor] = entry
		// 	}
		// }

		for index := range prod.Item {
			if prod.Item[index].ColorCode == tecalColor {
				prod.Item[index].Size = append(prod.Item[index].Size, bases.Size{Val: e.DOM.Text(), IsExit: true})
				prod.Size = append(prod.Size, e.DOM.Text())
			}
		}
	})

	// Обработка ошибки
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Ошибка. Поэтому немного ждём. Ошибка", e)
		time.Sleep(5 * time.Second)
	})

	// Обработка ошибки после ответа сервера
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != http.StatusOK { // Если нет ответа
			fmt.Println("Ошибка. Поэтому немного ждём. Статус", r.StatusCode)
			time.Sleep(5 * time.Second)
		}
	})

	fmt.Println("Захожу:", URL+ProductColorLink)
	c.Visit(URL + ProductColorLink)

	prod.Link = ProductColorLink
	prod.Size = bases.RemoveDuplicateStr(prod.Size)
}

// Перевести /sweaters/CKvXARDQ1wHiAgIBAg.zso в sweaters
func formSlump(input string, selection int) (output string) {
	input = strings.ReplaceAll(input, ":", "")
	strs := strings.Split(input, "/")
	if len(strs) >= selection+1 {
		return strs[selection]
	}
	return ""
}

// Распечатать продукт
func PrintProduct2(prod bases.Product2) (output string) {
	output = "Название: " + prod.Name + ". Артикул: " + prod.Article + "\n" +
		"Производитель: " + prod.Manufacturer + ". Гендер: " + prod.GenderLabel + "\n" + " Название(Полн): " + prod.FullName + "\n" +
		"Ссылка: " + prod.Link + "\n" +
		"Размеры: " + strings.Join(prod.Size, ",") + "\n" +
		"Подкатегория: " + fmt.Sprintf("%+v", prod.Cat) + "\n" +
		"Описание(Рус): " + prod.Description.Rus + "\n" +
		"Описание(Eng): " + prod.Description.Eng + "\n" +
		"Дополнительные поля: " + fmt.Sprintf("%+v", prod.Specifications) + "\n" +
		"Подробнее по каждому цвету:\n" + PrintItems(prod.Item)

	return output
}
func PrintItems(items []bases.ColorItem) (output string) {
	for key, val := range items {
		var StrsSizes string
		for ind, valSize := range val.Size {
			StrsSizes += strconv.Itoa(ind) + ": " + valSize.Val + "," + strconv.FormatBool(valSize.IsExit) + "; "
		}
		output += strconv.Itoa(key) + " - " + val.ColorEng + "\n" +
			"\tЦена: " + fmt.Sprintf("%.2f", val.Price) + ". Ссылка: " + val.Link + "\n" +
			"\tРазмеры(" + strconv.Itoa(len(val.Size)) + "): " + StrsSizes + "\n" +
			"\tКартинка: " + strings.Join(val.Image, ",") + "\n"
	}
	// for key, val := range items {
	// 	output += key + " - " + val.ColorEng + "\n" +
	// 		"\tЦена: " + fmt.Sprintf("%.2f", val.Price) + ". Ссылка: " + val.Link + "\n" +
	// 		"\tРазмеры(" + strconv.Itoa(len(val.Size)) + "): " + strings.Join(val.Size, ",") + "\n" +
	// 		"\tКартинка: " + strings.Join(val.Image, ",") + "\n"
	// }

	return output
}

// Достать код картинки из ссылки
//
//	https://m.media-amazon.com/images/I/91GJ2hRcTeL._AC_SR58.88,73.60000000000001_.jpg > 91GJ2hRcTeL
func PictureCode(imgStr string) (code string, parseError error) {
	u, err := url.Parse(imgStr)
	if err != nil {
		return "", parseError
	}

	if !strings.Contains(u.Path, "/images/I/") {
		return "", errors.New("nothing \"/images/I/\" in url: " + imgStr)
	}
	code = strings.ReplaceAll(u.Path, "/images/I/", "") // /images/I/91GJ2hRcTeL._AC_SR58.88,73.60000000000001_.jpg > 91GJ2hRcTeL._AC_SR58.88,73.60000000000001_.jpg

	strs := strings.Split(code, ".")
	if len(strs) == 0 {
		return "", errors.New("null imgStr.split for url: " + imgStr)
	}

	return strs[0], nil
}
