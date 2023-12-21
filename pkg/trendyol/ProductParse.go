package trendyol

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const product_URL string = "https://public.trendyol.com/discovery-web-productgw-service/api/productDetail/%d?storefrontId=1"

func ParseProduct(ProductID int) (pg ProductStruct, Err error) {
	url := fmt.Sprintf(product_URL, ProductID) // Рабочая ссылка для парсинга
	// fmt.Println("Lines:", url)
	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return ProductStruct{}, ErrNewRequest
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36")

	// Выполнить запорос
	Response, ErrDo := client.Do(req)
	if ErrDo != nil {
		return ProductStruct{}, ErrDo
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return ProductStruct{}, ErrorReadAll
	}

	// Распарсить полученный json в структуру
	ErrorUnmarshal := json.Unmarshal(BodyPage, &pg)
	if ErrorUnmarshal != nil {
		return ProductStruct{}, ErrorUnmarshal
	}

	return pg, nil
}
