package massimodutti

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"

	"ClikShop/common/bases"
)

// Получить полее подробный список по товарам, путём выполнения [запроса].
//
// # На вход нужно подать слайс с артикулами, информацию по которым и нужно будет получить в итоге
//
// [запроса]: https://www.massimodutti.com/itxrest/3/catalog/store/34009471/30359503/productsArray?languageId=-1&productIds=28576969%2C27186929%2C27186928%2C28697808%2C27491405%2C27491406%2C27491404%2C27181446%2C27286034%2C28793012%2C30276951
func (s *Service) Lines(ids []int) (line Line, respErr error) {
	c, err := s.NewServiceCollector()
	if err != nil {
		return Line{}, errors.Wrap(err, "create service collector: ")
	}

	url := "https://www.massimodutti.com/itxrest/3/catalog/store/34009471/30359503/productsArray?languageId=-1&productIds=" + JoinInt(ids, "%2C")

	headers := http.Header{}
	headers.Add("authority", "www.massimodutti.com")
	headers.Add("accept", "application/json, text/plain, */*")
	headers.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	headers.Add("referer", "https://www.massimodutti.com/")
	headers.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	headers.Add("sec-ch-ua-mobile", "?0")
	headers.Add("sec-ch-ua-platform", "\"Linux\"")
	headers.Add("sec-fetch-dest", "empty")
	headers.Add("sec-fetch-mode", "cors")
	headers.Add("sec-fetch-site", "same-origin")
	headers.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36")

	var response Line
	c.OnResponse(func(r *colly.Response) {
		if respErr = json.Unmarshal(r.Body, &response); err != nil {
			log.Println("ERROR:500:", respErr)
			return
		}
	})

	err = c.Request(http.MethodGet, url, nil, nil, headers)
	return response, err
}

// Объединить слайс int с разделителем sep
//
// Аналог функции strings.Join() только для слайса []int
func JoinInt(input []int, sep string) (result string) {
	if len(input) == 0 {
		return ""
	}
	for _, val := range input {
		result += sep + strconv.Itoa(val)
	}
	return result[len([]rune(sep)):]
}

// Преобразовать line в стандартную структуру данных
func Line2Product2(line Line, cats []bases.Cat) (prod []bases.Product2) {

	//
	for _, LineProd := range line.Products {
		var AddingProduct bases.Product2

		AddingProduct.Article = strconv.Itoa(LineProd.ID) // Артикулы
		AddingProduct.Name = LineProd.Name                // Название товара
		AddingProduct.Cat = cats

		for _, Itemproduct := range LineProd.BundleColors {
			AddingProduct.Item = append(AddingProduct.Item, bases.ColorItem{
				ColorEng:  Itemproduct.Name,
				ColorCode: bases.Name2Slug(Itemproduct.Name),
				Link:      LineProd.ProductURL,
			})
		}

		prod = append(prod, AddingProduct)
	}

	return prod
}
