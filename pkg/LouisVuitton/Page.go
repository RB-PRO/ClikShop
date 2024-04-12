package louisvuitton

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Получить список всех товаров без подробностей
//
// Передай ей примерный тег категории, например: "t1z0ff7q", и она тебе выдаст список товаров с внутренними артикулами.
//
// Работает с [ручками]. Пример на [PostMan].
//
// Замкнутая функция. Должна сама решать свои проблемы в ином случае паникует
//
// [ручками]: https://api.louisvuitton.com/eco-eu/search-merch-eapi/v1/rus-ru/plp/products/t1z0ff7q?page=0
// [PostMan]:
func (cr *Core) Pages(CategoryTag string) (Products []HitsPage) {

	var CoutPages int = 1
	for iPage := 0; iPage < CoutPages; iPage++ {

		// Делаем запрос
		PageRes, ErrPages := cr.PageSingle(CategoryTag, iPage)

		// Обработка ошибок, связанной с неправильными данными ClientID ClientSecret
		if errors.Is(ErrPages, error_authentication_denied) || errors.Is(ErrPages, error_invalid_client) {
			cr.UpdateCore()
			iPage--
			continue
		}

		// Если иная ошибка, то паникуем
		if ErrPages != nil {
			panic(errors.New(fmt.Sprintf("[%d/%d]: %v", iPage, CoutPages, ErrPages)))
		}

		// Актуализируем данные по к-ву страниц с товарами
		CoutPages = PageRes.NbPages

		// Если всё хорошо, то добавляем товары в слайс
		Products = append(Products, PageRes.Hits...)
	}

	return Products
}

// Запросить одну страницу с номером page из категории CategoryTag
func (cr *Core) PageSingle(CategoryTag string, page int) (PageRes PageResponse, ErrPages error) {

	url := fmt.Sprintf("https://api.louisvuitton.com/eco-eu/search-merch-eapi/v1/rus-ru/plp/products/%s?page=%d", CategoryTag, page)

	// fmt.Println(url)

	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return PageResponse{}, ErrNewRequest
	}

	req.Header.Add("authority", "api.louisvuitton.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("client_id", cr.ClientID)
	req.Header.Add("client_secret", cr.ClientSecret)
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36")

	// Выполнить запрос
	res, err := client.Do(req)
	if err != nil {
		return PageResponse{}, err
	}
	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return PageResponse{}, fmt.Errorf("PageSingle: wrong status code: %d", res.StatusCode)
	}

	// Читаем ответ в массив байтов
	bodyBytes, ErrReadAll := io.ReadAll(res.Body)
	if ErrReadAll != nil {
		return PageResponse{}, ErrReadAll
	}

	// fmt.Println(string(bodyBytes))

	// Декодируем полученный json и получаем данные
	ErrUnmarshal := json.Unmarshal(bodyBytes, &PageRes)
	if ErrUnmarshal != nil {
		return PageResponse{}, ErrUnmarshal
	}

	return PageRes, nil
}
