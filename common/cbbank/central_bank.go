package cbbank

import (
	"encoding/json"
	"io"
	"net/http"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Lira() (float64, error) {

	resp, err := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
	if err != nil {
		return 0.0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}

	var bankResponse Bank
	if json.Unmarshal(body, &bankResponse) != nil { // Если ошибка распарсивания в структуру данных
		return 0.0, err
	}

	return bankResponse.Valute.Try.Value / 10, nil
}
