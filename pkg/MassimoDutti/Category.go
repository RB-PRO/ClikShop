package massimodutti

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Получить список категорий массимо дутти, путём выполнения [запроса]
//
// [запроса]: https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category?languageId=-1&typeCatalog=1&appId=1
func Category() (categs Categories, ErrCategory error) {
	url := "https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category?languageId=-1&typeCatalog=1&appId=1"

	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return categs, ErrNewRequest
	}
	req.Header.Add("authority", "www.massimodutti.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("referer", "https://www.massimodutti.com/")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36")

	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return categs, ErrDo
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		ErrNewDecoder := json.NewDecoder(res.Body).Decode(&categs)
		if ErrNewDecoder != nil {
			return categs, ErrNewDecoder
		}
	} else {
		return categs, errors.New("Category: http.Status is not ok")
	}

	return categs, nil
}
