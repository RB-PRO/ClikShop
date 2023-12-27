package trendyol

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Структура которая содержит как группу товара, так и его ID
type Groupeng struct {
	ID             int
	ProductGroupID int
}

// Спарсить страницы товаров
func Pages(ShopID int) (ProductGroupIDs []Groupeng, Err error) {

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
		// fmt.Printf("iPage: %d, pg.Result.TotalCount %d, CoutPages %d\n", iPage, pg.Result.TotalCount, CoutPages)

		// Сохраняем ссылки на группы товаров
		for _, Product := range pg.Result.Products {
			ProductGroupIDs = append(ProductGroupIDs, Groupeng{
				ID:             Product.ID,
				ProductGroupID: Product.ProductGroupID,
			})
		}

		time.Sleep(time.Millisecond * 200)

		// break
	}

	return ProductGroupIDs, Err
}

const page_URL string = "https://public.trendyol.com/discovery-web-searchgw-service/v2/api/infinite-scroll/sr?mid=%d&os=1&pi=%d"

// Пропарсить страницу товаров для получения списка ID групп
func ParsePage(ShopID int, page int) (pg PageStruct, Err error) {
	url := fmt.Sprintf(page_URL, ShopID, page) // Рабочая ссылка для парсинга
	// fmt.Println("Lines:", url)
	client := &http.Client{Timeout: time.Second * 60}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return PageStruct{}, fmt.Errorf("http.NewRequest: %v", ErrNewRequest)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 YaBrowser/23.11.0.0 Safari/537.36")

	// Выполнить запорос
	Response, ErrDo := client.Do(req)
	if ErrDo != nil {
		return PageStruct{}, fmt.Errorf("client.Do: %v", ErrDo)
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return PageStruct{}, fmt.Errorf("io.ReadAll: %v", ErrorReadAll)
	}

	// Распарсить полученный json в структуру
	if ErrorUnmarshal := json.Unmarshal(BodyPage, &pg); ErrorUnmarshal != nil {
		return PageStruct{}, fmt.Errorf("json.Unmarshal: %v", ErrorUnmarshal)
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
