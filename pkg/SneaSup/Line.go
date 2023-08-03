package sneaksup

import (
	"encoding/json"
	"net/http"
)

func LinePost(link, pagenumber string) (Line LineStruct, ErrLine error) {

	url := "https://www.sneaksup.com/kadin-ayakkabi-sneaker?paginationType=20&orderby=0&pagenumber=1"

	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return Line, ErrNewRequest
	}
	req.Header.Add("authority", "www.sneaksup.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("referer", "https://www.sneaksup.com/kadin-ayakkabi-sneaker")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36")
	req.Header.Add("Cookie", "inCommerce.customer.info=c1d75719-734a-45a9-9fbd-15298b12b72f; inveonSessionId=11dw5bf3nuzj1ht1q1ix4ojg")

	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return Line, ErrDo
	}
	defer res.Body.Close()

	ErrNewDecoder := json.NewDecoder(res.Body).Decode(&Line)
	if ErrNewDecoder != nil {
		return Line, ErrNewDecoder
	}

	return Line, ErrNewRequest
}
