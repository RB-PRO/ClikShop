package apibitrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Стуктура запроса метода getProduct
type product_Request struct {
	ProductsID []string `json:"products_id"`
}

// Стуктура ответа метода getProduct
type Product_Response struct {
	Products []struct {
		ID     string `json:"id"`
		Link   string `json:"link"`
		Colors []struct {
			ID       string  `json:"id"`
			ColorEng string  `json:"ColorEng"`
			Link     string  `json:"link"`
			Price    float64 `json:"price"`
			Size     string  `json:"size"`
		} `json:"colors"`
	} `json:"products"`
	Error []string `json:"error"`
}

// Получить информацию по определённым товарам
func (user *BitrixUser) Product(Values []string) (ProdResp Product_Response, Err error) {

	// Преобразуем структуру в массив байтов
	bodyBytesReq, errMarshal := json.Marshal(product_Request{ProductsID: Values})
	if errMarshal != nil {
		return Product_Response{}, fmt.Errorf("bitrix: Product: Не смог запаковать структуру в массив байтов: %w", errMarshal)
	}

	// Создаём запрос
	req, errNewRequest := http.NewRequest(http.MethodPost, fmt.Sprintf(bitrixURL, "getProduct"), bytes.NewReader(bodyBytesReq))
	if errNewRequest != nil {
		return Product_Response{}, fmt.Errorf("bitrix: Product: Не смог создать запрос: %w", errNewRequest)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "BITRIX_SM_GUEST_ID=5009; BITRIX_SM_LAST_VISIT=07.08.2023%2001%3A19%3A12; BITRIX_SM_SALE_UID=1524ac0f1701198a7380ac70768d3606; PHPSESSID=kbNuDok3oE8R6fJ7dExSGO8fbympRahj")

	// Выполняем запрос
	resp, errDo := http.DefaultClient.Do(req)
	if errDo != nil {
		return Product_Response{}, fmt.Errorf("bitrix: Product: Не смог выполнить запрос: %w", errDo)
	}
	defer resp.Body.Close()

	// Считываем ответ из потка
	bodyBytesRes, errReadAll := io.ReadAll(resp.Body)
	if errReadAll != nil {
		return Product_Response{}, fmt.Errorf("bitrix: Product: Не смог считать ответ из потока: %w", errReadAll)
	}

	// Распарсить полученные данные
	responseErrorUnmarshal := json.Unmarshal(bodyBytesRes, &ProdResp)
	if responseErrorUnmarshal != nil {
		return Product_Response{}, fmt.Errorf("bitrix: Product: Не смог распарсить данные из ответа: %w", responseErrorUnmarshal)
	}

	return ProdResp, nil
}
