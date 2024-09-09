package massimodutti

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

// Получить список ID товаров по входной ID категории
func (s *Service) SKUs(id_category int) (ID, error) {
	c, err := s.NewServiceCollector()
	if err != nil {
		return ID{}, errors.Wrap(err, "create service collector: ")
	}

	url := fmt.Sprintf("https://www.massimodutti.com/itxrest/3/catalog/store/34009471/30359503/category/%v/product?languageId=-1&appId=1&showProducts=false", id_category)

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

	var response ID
	c.OnResponse(func(r *colly.Response) {
		if err := json.Unmarshal(r.Body, &response); err != nil {
			log.Println("ERROR:500:", err)
			return
		}
	})

	err = c.Request(http.MethodGet, url, nil, nil, headers)
	return response, err
}
