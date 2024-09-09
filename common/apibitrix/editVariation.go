package apibitrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Стуктура запроса метода variation
type VariationSize_Request struct {
	ID   string `json:"id"`
	Size string `json:"size"`
}

type variationProdSize_Request struct {
	Products []VariationSize_Request `json:"products"`
}

// Обновить информацию по определённой категории
func (s *Service) UpdateSizeVariation(VrtReq []VariationSize_Request) (Err error) {

	// Преобразуем структуру в массив байтов
	bodyBytesReq, errMarshal := json.Marshal(variationProdSize_Request{VrtReq})
	if errMarshal != nil {
		return fmt.Errorf("bitrix: Product: Не смог запаковать структуру в массив байтов: %w", errMarshal)
	}

	// Создаём запрос
	req, errNewRequest := http.NewRequest(http.MethodPost, fmt.Sprintf(bitrixURL, "editVariation"), bytes.NewReader(bodyBytesReq))
	if errNewRequest != nil {
		return fmt.Errorf("bitrix: Product: Не смог создать запрос: %w", errNewRequest)
	}
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Cookie", "BITRIX_SM_GUEST_ID=5009; BITRIX_SM_LAST_VISIT=07.08.2023%2001%3A19%3A12; BITRIX_SM_SALE_UID=1524ac0f1701198a7380ac70768d3606; PHPSESSID=kbNuDok3oE8R6fJ7dExSGO8fbympRahj")

	// Выполняем запрос
	resp, errDo := http.DefaultClient.Do(req)
	if errDo != nil {
		return fmt.Errorf("bitrix: Product: Не смог выполнить запрос: %w", errDo)
	}
	defer resp.Body.Close()

	return nil
}
