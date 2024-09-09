package apibitrix

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const costUrl = "https://clikshop.ru/parser/Update.json"

type CoastMap struct {
	Walrus   float64
	Delivery int
}

func (s *Service) Coasts() (map[string]CoastMap, error) {

	// Создаём запрос
	req, err := http.NewRequest(http.MethodPost, costUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("bitrix: Coasts: Не смог создать запрос: %w", err)
	}

	// Выполняем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("bitrix: Coasts: Не смог выполнить запрос: %w", err)
	}
	defer resp.Body.Close()

	bodyBytesRes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("bitrix: Coasts: Не смог считать ответ из потока: %w", err)
	}

	CoastsResp := struct {
		Main []struct {
			Market   string  `json:"market"`
			Walrus   float64 `json:"walrus"`
			Delivery int     `json:"delivery"`
		} `json:"main"`
	}{}
	if json.Unmarshal(bodyBytesRes, &CoastsResp) != nil {
		return nil, fmt.Errorf("bitrix: Coasts: Не смог распарсить данные из ответа: %w", err)
	}

	if len(CoastsResp.Main) == 0 {
		return nil, fmt.Errorf("bitrix: Coasts: Получено ноль различных цен для разных маркетов\n" +
			"Проверь https://clikshop.ru/parser/Update.json")
	}

	coastMap := make(map[string]CoastMap, len(CoastsResp.Main))
	for _, market := range CoastsResp.Main {
		coastMap[market.Market] = CoastMap{Walrus: market.Walrus, Delivery: market.Delivery}
	}

	return coastMap, nil
}
