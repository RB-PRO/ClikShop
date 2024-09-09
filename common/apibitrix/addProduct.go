package apibitrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"ClikShop/common/bases"
)

// Получить ID всех товаров
func (s *Service) AddProduct(prod bases.Product2) (response int, Err error) {

	// Преобразуем структуру в массив байтов
	bodyBytesReq, errMarshal := json.Marshal(prod)
	if errMarshal != nil {
		return 0, fmt.Errorf("bitrix: addProduct: Не смог запаковать структуру в массив байтов: %w", errMarshal)
	}

	// Создаём запрос
	req, errNewRequest := http.NewRequest(http.MethodPost,
		fmt.Sprintf(bitrixURL, "addProduct"), bytes.NewBuffer(bodyBytesReq))
	if errNewRequest != nil {
		return 0, fmt.Errorf("bitrix: AddProducts: Не смог создать запрос: %w", errNewRequest)
	}
	req.Header.Add("Cookie", "BITRIX_SM_GUEST_ID=5009; BITRIX_SM_LAST_VISIT=07.08.2023%2001%3A19%3A12; BITRIX_SM_SALE_UID=1524ac0f1701198a7380ac70768d3606; PHPSESSID=kbNuDok3oE8R6fJ7dExSGO8fbympRahj")

	// Выполняем запрос
	client := http.Client{Timeout: 60 * time.Second}
	resp, errDo := client.Do(req)
	if errDo != nil {
		return 0, fmt.Errorf("bitrix: AddProducts: Не смог выполнить запрос: %w", errDo)
	}
	defer resp.Body.Close()

	bodyBytesRes, errReadAll := io.ReadAll(resp.Body)
	if errReadAll != nil {
		return 0, fmt.Errorf("bitrix: AddProducts: Не смог считать ответ из потока: %w", errReadAll)
	}

	responseErrorUnmarshal := json.Unmarshal(bodyBytesRes, &response)
	if responseErrorUnmarshal != nil {
		return 0, fmt.Errorf("bitrix: AddProducts: Не смог распарсить данные из ответа: %w", responseErrorUnmarshal)
	}

	return response, nil
}
