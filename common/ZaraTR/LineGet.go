package zaratr

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

// Загрузить все товары категории
func (s *Service) LoadLine(id string) (Line, error) {
	c, err := s.NewServiceCollector()
	if err != nil {
		return Line{}, errors.Wrap(err, "create service collector: ")
	}

	url := fmt.Sprintf(LineURL, id)

	headers := http.Header{}
	headers.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	headers.Add("accept-encoding", "application/json; charset=utf-8")
	headers.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	headers.Add("if-none-match", "W/\"c2b35-nPn7rjU78OjmncJadKH5VqRw6b8\"")
	headers.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.895 Yowser/2.5 Safari/537.36")

	var response Line
	c.OnResponse(func(r *colly.Response) {
		if err := json.Unmarshal(r.Body, &response); err != nil {
			log.Println("ERROR:500:", err)
			return
		}
	})

	return response, c.Request(http.MethodGet, url, nil, nil, headers)
}
