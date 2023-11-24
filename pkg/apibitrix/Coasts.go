package apibitrix

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// "https://clikshop.ru/parser/Update.json"

type coast struct {
	Main []struct {
		Market   string  `json:"market"`
		Walrus   float64 `json:"walrus"`
		Delivery int     `json:"delivery"`
	} `json:"main"`
}
type CoastMap struct {
	Walrus   float64
	Delivery int
}

// Получить цены на все товары, которые только могут быть
func (user *BitrixUser) Coasts() (map[string]CoastMap, error) {

	// Создаём запрос
	req, errNewRequest := http.NewRequest(http.MethodPost, "https://clikshop.ru/parser/Update.json", nil)
	if errNewRequest != nil {
		return nil, fmt.Errorf("bitrix: Coasts: Не смог создать запрос: %w", errNewRequest)
	}

	// Выполняем запрос
	resp, errDo := http.DefaultClient.Do(req)
	if errDo != nil {
		return nil, fmt.Errorf("bitrix: Coasts: Не смог выполнить запрос: %w", errDo)
	}
	defer resp.Body.Close()

	// Считываем ответ из потка
	bodyBytesRes, errReadAll := io.ReadAll(resp.Body)
	if errReadAll != nil {
		return nil, fmt.Errorf("bitrix: Coasts: Не смог считать ответ из потока: %w", errReadAll)
	}

	// Распарсить полученные данные
	var CoastsResp coast
	responseErrorUnmarshal := json.Unmarshal(bodyBytesRes, &CoastsResp)
	if responseErrorUnmarshal != nil {
		return nil, fmt.Errorf("bitrix: Coasts: Не смог распарсить данные из ответа: %w", responseErrorUnmarshal)
	}

	if len(CoastsResp.Main) == 0 {
		return nil, fmt.Errorf("bitrix: Coasts: Получено ноль различных цен для разных маркетов\n" +
			"Проверь https://clikshop.ru/parser/Update.json")
	}

	CM := make(map[string]CoastMap, len(CoastsResp.Main))
	for _, market := range CoastsResp.Main {
		CM[market.Market] = CoastMap{Walrus: market.Walrus, Delivery: market.Delivery}
	}

	// Сохранение результата
	user.MapCoast = CM

	return CM, nil
}
