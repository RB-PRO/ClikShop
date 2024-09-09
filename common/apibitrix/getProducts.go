package apibitrix

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Стуктура ответа метода getProducts
type productsResponse struct {
	ProductsId []string `json:"products_id"`
}

// Получить ID всех товаров
func (s *Service) Products() ([]string, error) {

	// Создаём запрос
	req, errNewRequest := http.NewRequest(http.MethodPost, fmt.Sprintf(bitrixURL, "getProducts"), nil)
	if errNewRequest != nil {
		return nil, fmt.Errorf("bitrix: Products: Не смог создать запрос: %w", errNewRequest)
	}
	req.Header.Add("Cookie", "BITRIX_SM_GUEST_ID=5009; BITRIX_SM_LAST_VISIT=07.08.2023%2001%3A19%3A12; BITRIX_SM_SALE_UID=1524ac0f1701198a7380ac70768d3606; PHPSESSID=kbNuDok3oE8R6fJ7dExSGO8fbympRahj")

	// Выполняем запрос
	resp, errDo := http.DefaultClient.Do(req)
	if errDo != nil {
		return nil, fmt.Errorf("bitrix: Products: Не смог выполнить запрос: %w", errDo)
	}
	defer resp.Body.Close()

	// Считываем ответ из потка
	bodyBytesRes, errReadAll := io.ReadAll(resp.Body)
	if errReadAll != nil {
		return nil, fmt.Errorf("bitrix: Products: Не смог считать ответ из потока: %w", errReadAll)
	}

	// Распарсить полученные данные
	var ProdsResp productsResponse
	responseErrorUnmarshal := json.Unmarshal(bodyBytesRes, &ProdsResp)
	if responseErrorUnmarshal != nil {
		return nil, fmt.Errorf("bitrix: Products: Не смог распарсить данные из ответа: %w", responseErrorUnmarshal)
	}

	return ProdsResp.ProductsId, nil
}
