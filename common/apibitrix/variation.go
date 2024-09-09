package apibitrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Стуктура запроса метода variation
type Variation_Request struct {
	ID           string  `json:"id"`
	Price        float64 `json:"price"`
	Availability bool    `json:"availability"`
}

type variationProd_Request struct {
	Products []Variation_Request `json:"products"`
}

// Стуктура ответа метода variation
type Variation_Response struct {
	Error []string `json:"error"`
}

// Обновить информацию по определённой категории
func (s *Service) Variation(VrtReq []Variation_Request) (VrtResp Variation_Response, Err error) {

	// Преобразуем структуру в массив байтов
	bodyBytesReq, errMarshal := json.Marshal(variationProd_Request{VrtReq})
	if errMarshal != nil {
		return Variation_Response{}, fmt.Errorf("bitrix: Product: Не смог запаковать структуру в массив байтов: %w", errMarshal)
	}

	// Создаём запрос
	req, errNewRequest := http.NewRequest(http.MethodPost, fmt.Sprintf(bitrixURL, "variation"), bytes.NewReader(bodyBytesReq))
	if errNewRequest != nil {
		return Variation_Response{}, fmt.Errorf("bitrix: Product: Не смог создать запрос: %w", errNewRequest)
	}
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Cookie", "BITRIX_SM_GUEST_ID=5009; BITRIX_SM_LAST_VISIT=07.08.2023%2001%3A19%3A12; BITRIX_SM_SALE_UID=1524ac0f1701198a7380ac70768d3606; PHPSESSID=kbNuDok3oE8R6fJ7dExSGO8fbympRahj")

	// Выполняем запрос
	resp, errDo := http.DefaultClient.Do(req)
	if errDo != nil {
		return Variation_Response{}, fmt.Errorf("bitrix: Product: Не смог выполнить запрос: %w", errDo)
	}
	defer resp.Body.Close()

	// Считываем ответ из потка
	bodyBytesRes, errReadAll := io.ReadAll(resp.Body)
	if errReadAll != nil {
		return Variation_Response{}, fmt.Errorf("bitrix: Product: Не смог считать ответ из потока: %w", errReadAll)
	}

	// Распарсить полученные данные
	responseErrorUnmarshal := json.Unmarshal(bodyBytesRes, &VrtResp)
	if responseErrorUnmarshal != nil {
		return Variation_Response{}, fmt.Errorf("bitrix: Product: Не смог распарсить данные из ответа: %w", responseErrorUnmarshal)
	}

	return VrtResp, nil
}
