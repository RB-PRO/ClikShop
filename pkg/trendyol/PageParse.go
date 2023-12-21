package trendyol

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Спарсить страницы товаров
func Pages(ShopID string, page int) (ProductGroupIDs []int, Err error) {

	// Количество страниц в товаре
	var CoutPages int = 1

	// Цикл по всем страницам
	for iPage := 1; iPage <= CoutPages; iPage++ {

		// Сбор информации о товарах на странице iPage
		pg, ErrPage := ParsePage(ShopID, iPage)
		if Err != nil {
			return nil, fmt.Errorf("parsePage: %v", ErrPage)
		}
		CoutPages = RoundUp(pg.Result.TotalCount, 24)

		// Сохраняем ссылки на группы товаров
		for _, Product := range pg.Result.Products {
			ProductGroupIDs = append(ProductGroupIDs, Product.ProductGroupID)
		}

	}

	return ProductGroupIDs, Err
}

const page_URL string = "https://public.trendyol.com/discovery-web-searchgw-service/v2/api/infinite-scroll/sr?mid=%s&os=1&pi=%d"

// https://public.trendyol.com/discovery-web-searchgw-service/v2/api/infinite-scroll/sr?mid=106871&os=1&pi=2
func ParsePage(ShopID string, page int) (pg PageStruct, Err error) {
	url := fmt.Sprintf(page_URL, ShopID, page) // Рабочая ссылка для парсинга
	// fmt.Println("Lines:", url)
	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return PageStruct{}, ErrNewRequest
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36")

	// Выполнить запорос
	Response, ErrDo := client.Do(req)
	if ErrDo != nil {
		return PageStruct{}, ErrDo
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return PageStruct{}, ErrorReadAll
	}

	// Распарсить полученный json в структуру
	ErrorUnmarshal := json.Unmarshal(BodyPage, &pg)
	if ErrorUnmarshal != nil {
		return PageStruct{}, ErrorUnmarshal
	}

	return pg, nil
}

// Округление в большую сторону с переменным делителем
func RoundUp(value int, slice int) int {
	b := value / slice
	if value-b*slice != 0 {
		return b + 1
	}
	return b
}
