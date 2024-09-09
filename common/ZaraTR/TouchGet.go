package zaratr

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

// Загрузить информацию по товару
func (s *Service) LoadTouch(id string) (tou Touch, ErrorLine error) {
	c, err := s.NewServiceCollector()
	if err != nil {
		return Touch{}, errors.Wrap(err, "create service collector: ")
	}

	url := fmt.Sprintf(TouchURL, id)

	// TODO: Timeout: time.Second * 10
	//s.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"

	headers := http.Header{}
	headers.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	var response Touch
	c.OnResponse(func(r *colly.Response) {
		if err := json.Unmarshal(r.Body, &response); err != nil {
			log.Println("ERROR:500:", err)
			return
		}
	})

	return response, c.Request(http.MethodGet, url, nil, nil, headers)
}
