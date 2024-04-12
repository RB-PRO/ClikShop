package louisvuitton

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Структура ответа на запрос цен
type PriceResponse struct {
	SkuList []struct {
		Offers struct {
			PriceSpecification struct {
				PriceCurrency string `json:"priceCurrency,omitempty"`
				Price         int    `json:"price,omitempty"`
			} `json:"priceSpecification,omitempty"`
			Type  string `json:"@type,omitempty"`
			Price string `json:"price,omitempty"`
		} `json:"offers,omitempty"`
		ProductID  string `json:"productId,omitempty"`
		Identifier string `json:"identifier,omitempty"`
	} `json:"skuList,omitempty"`
	SkuListSize int `json:"skuListSize,omitempty"`
}

func Price(SKUs []string, countrylanguage, region string) (map[string]float64, error) {
	price := make(map[string]float64)

	url := fmt.Sprintf("https://api.louisvuitton.com/api/%s/catalog/skus/%s/price?dispatchCountry=%s", countrylanguage, strings.Join(SKUs, ","), region)

	// fmt.Println(url)

	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return nil, ErrNewRequest
	}

	req.Header.Add("authority", "api.louisvuitton.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("authorization", "Basic VnVpdHRvbjpSdjY1bEQzUw==")
	req.Header.Add("content-type", "text/plain")
	req.Header.Add("origin", "https://me.louisvuitton.com")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36")

	// Выполнить запрос
	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return nil, ErrDo
	}
	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Price: wrong status code: %d", res.StatusCode)
	}

	// Читаем ответ в массив байтов
	bodyBytes, ErrReadAll := io.ReadAll(res.Body)
	if ErrReadAll != nil {
		return nil, ErrReadAll
	}

	// fmt.Println(string(bodyBytes))

	// Декодируем полученный json и получаем данные
	var PriceResp PriceResponse
	ErrUnmarshal := json.Unmarshal(bodyBytes, &PriceResp)
	if ErrUnmarshal != nil {
		return nil, ErrUnmarshal
	}

	// Происходит полный распарсинг значений в мою структур данных
	for _, val := range PriceResp.SkuList {
		price[val.Identifier] = float64(val.Offers.PriceSpecification.Price)
	}

	return price, nil
}
