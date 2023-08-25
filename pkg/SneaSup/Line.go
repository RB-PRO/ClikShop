package sneaksup

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

// Получить список товаров в результате выполнения ф-й Line
func Lines(link string) (Products []bases.Product2, Err error) {

	// Цикл по всем страницам
	for i := 1; ; i++ {

		// Загружаем данные
		line, ErrLinePost := LinePost(link, i)
		if ErrLinePost != nil {
			return Products, ErrLinePost
		}

		// Выход по причине того, что товаров больше нет
		if line.Pager.PageIndex+1 >= line.Pager.TotalPages {
			break
		}
	}

	return Products, Err
}

// Загрузить список товаров
func LinePost(link string, pagenumber int) (Line LineStruct, ErrLine error) {

	url, errMakeLink := linkTranstore(link, pagenumber)
	if errMakeLink != nil {
		return Line, errMakeLink
	}

	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return Line, ErrNewRequest
	}
	req.Header.Add("authority", "www.sneaksup.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("referer", "https://www.sneaksup.com/kadin-ayakkabi-sneaker")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36")

	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return Line, ErrDo
	}
	defer res.Body.Close()

	ErrNewDecoder := json.NewDecoder(res.Body).Decode(&Line)
	if ErrNewDecoder != nil {
		return Line, ErrNewDecoder
	}

	return Line, ErrNewRequest
}

// Преобразовать ссылку на парсинга line
func linkTranstore(link string, pagenumber int) (string, error) {

	// Парсим ссылку в формат отдачи json
	u, ErrParse := url.Parse(link)
	if ErrParse != nil {
		return "", ErrParse
	}

	// Добавляем аттрибусы, которые соответствуют запросу, который отдаёт json
	q := u.Query()
	q.Set("paginationType", "20")
	q.Set("orderby", "0")
	q.Set("pagenumber", strconv.Itoa(pagenumber))
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// Перевести из line в структуру product2
func Line2Product(line LineStruct) (prods []bases.Product2) {
	for _, LineProd := range line.Products {
		var prod bases.Product2

		prod.Name = LineProd.Name
		prod.Link = URL + LineProd.URL
		prod.Article = LineProd.Sku
		prod.FullName = LineProd.DefaultPictureModel.AlternateText
		prod.Manufacturer = LineProd.ManufacturerName

		// for _, Sibligs := range prod.Sibligs {

		// }

		prods = append(prods, prod)
	}
	return prods
}
