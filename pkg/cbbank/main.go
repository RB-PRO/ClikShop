package cbbank

import (
	"encoding/json"
	"io"
	"net/http"
)

// Главная структура центрального банка
type CentrakBank struct {
	Data Bank
}

// Создать экземпляр запроса на получение курса валют
func New() (*CentrakBank, error) {

	// Получить ответ от сервера
	resp, err := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
	if err != nil {
		return &CentrakBank{}, err
	}

	// Читаем ответ
	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		return &CentrakBank{}, errBody
	}

	// Распарсить ответ
	var BankResponse CentrakBank
	errUnmarshal := json.Unmarshal(body, &BankResponse.Data)
	if errUnmarshal != nil { // Если ошибка распарсивания в структуру данных
		return &CentrakBank{}, errUnmarshal
	}

	return &BankResponse, nil
}
