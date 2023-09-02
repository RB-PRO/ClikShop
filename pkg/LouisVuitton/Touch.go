package louisvuitton

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Запросить данные по товару, например: nvprod3900007v
//
// Работает с [ссылками]
//
//	[ссылками]: https://api.louisvuitton.com/eco-eu/catalog-lvcom/v1/rus-ru/product/nvprod3900007v
func (cr *Core) Toucher(ProductTag string) (Touch TouchResponse, ErrTouch error) {

	url := fmt.Sprintf("https://api.louisvuitton.com/eco-eu/catalog-lvcom/v1/rus-ru/product/%s", ProductTag)

	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return TouchResponse{}, ErrNewRequest
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
		return TouchResponse{}, err
	}
	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return TouchResponse{}, fmt.Errorf("Toucher: wrong status code: %d", res.StatusCode)
	}

	// Читаем ответ в массив байтов
	bodyBytes, ErrReadAll := io.ReadAll(res.Body)
	if ErrReadAll != nil {
		return TouchResponse{}, ErrReadAll
	}

	// fmt.Println(string(bodyBytes))

	// Декодируем полученный json и получаем данные
	ErrUnmarshal := json.Unmarshal(bodyBytes, &Touch)
	if ErrUnmarshal != nil {
		return TouchResponse{}, ErrUnmarshal
	}

	return Touch, nil
}
