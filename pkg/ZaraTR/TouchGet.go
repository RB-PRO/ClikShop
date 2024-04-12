package zaratr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Загрузить информацию по товару
func LoadTouch(id string) (tou Touch, ErrorLine error) {

	// fmt.Println("URLS", fmt.Sprintf(TouchURL, id))

	// Делаем запрос на получение категорий
	url := fmt.Sprintf(TouchURL, id)
	// url = URL + "/tr/en/running-trainers-with-chunky-soles-p12341120.html?ajax=true&bm-verify=AAQAAAAI_____5lwzlu22VtQaoh-9PiUIhyOgcPo7q7uLb9PJnFXoQHFkGw6azaqj3INCtml-ncQeHVuacOgEzL_lQD66tShwJ1bFmVhHlhO_yF5UW6CJQs8SNLQnr3veBGv5uja4HD9xTnSjnCrZWKCqV9kxqY2m9knCEVeFs99iGTMb85zlsY-mQUnnbqOurvcWrU6EYJ1cAc2CbRAp5ZwpMusWOxQgzo8GAGI5n91e-bLbtsaeMWKg0fYV8q9v7NGYW_hg8el_0Pp0XFonCncRls84_pdM3sE_hjPYptE3bTGDnJHGPc3XM6NSMk7O7jaWeMUPPo"
	// fmt.Println("->>>>>>>>>>", url)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Touch{}, err
	}

	// Добавляем необходимые атрибуты
	// // req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	// req.Header.Add("accept-encoding", "application/json; charset=utf-8")
	// // req.Header.Add("Content-Type", "application/json")
	// // // req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.837 YaBrowser/23.9.4.837 Yowser/2.5 Safari/537.36")

	// // req.Header.Add("Cookie", "IDROSTA=64dc1351db62:268d3b4f6493ea3585dfd2b06; ITXDEVICEID=00b31a8146bfb20fa2acfe985759f4db; ITXSESSIONID=db01a42a6d6835973876737758b927af; _abck=5E8D2C29365BBC0C2A37267FDB23A6BD~-1~YAAQBc8ti9QFf7mLAQAAGqUwugpPQUyzsSXz2NdRPBW4xlzWPxQ1J3rBpqMSCq9xn1HnkkIai8lFhfZS2dBTyWAdmSbp9JTLczUicl+L4PfjRiuEz4DX3UzpgwAZPhukQ2gVXMo3WlQsm7+7jyh9aE+E58XGtjYWLNJygglFhVicXq0CYDAMyD1VEfnT4GInTEUDWqg1lzMB45gBnhmWVYi5Al42YpA8Fbm26oUB/VtSGnl+JjdIAsVKQzunvlZxgDcslyjjCQJ3HcugCLoHvW6uJsxHYdLaZXtAusOH5rWqn+AXLAIKvmULMNscs2GWu6euiJW/bs1p2bo2gt0Ior3VEhjFWx73g+V8IolLM1g9JsuTlDftDfhiz/eaTWqpcVKGR3ZKxSucgVC7IWT/yec6DxIkOg==~0~-1~-1; _ga=GA1.1.735583386.1699385774; ak_bmsc=E1B773D786B9C4BD192291F56AC418A9~000000000000000000000000000000~YAAQLs8tiysc4amLAQAAQTVquhUcRZ7mtGITtcvszMx2UwhuqZrd2U83sZ8TBWpN7HjtBOpuhW2qUGBCOh07mqsiTzUf83DCFCGyTSINtq/X67RnZuOg0PfExvGzLMLM9CZKRRTLZa8x/uHV0ECHfw+oY9NQUXwmAd5+0zlThP6V9a0zeQT1nr3qm9Sq+mvsl5GpjvnHYwRKazawA0VvWYPNnohkD0+RahBQOuSq43kv/w/sxg5EOIS8XuERtie7Agw0VLhAFPKgYKROPDOb8xMoGg4Xr8AwZH2qIhtFlypBDMCvzi3ZrK3emnisGZFkSUpps3lrH0I3KR3JsMQQ0Jjto0/A8WjOegOkPwfQZHHle10Pv1nwc/dX6aaFq311FdN4CBzI2XRItPQgai7OIuZkh19MoiTkscPihkIgDv4d04rW; bm_sz=0E3843A340FCBA196E91DB742EB29540~YAAQLM8tiz0FFKmLAQAA5r/buRXwVXXSFg5kmFnhaZd146fMLzpUsm+dmAjI1bXpo7D3aiu6KQAtsjQNNedfoZ+K3f/KnfshLUNTBoCOvI93Hg0x63AO9LrSqwouxpRLL0WDkqV+Tw+hwAjXWijZchc/ZatrcT0Enh7SWpItmxgt8uEOXiPKckID/gkwKRN1cXZDPepyZW6No5dZ86NDHDtm3vOtosKMfqNqZPH/WvECX94a9ph6WPPXIZNsh0RKrenOoZ55xraZfeF7K5jzaq/kfuurcse7qbdWn8n5KHez~4342838~4469049; TS0122c9b6=011f37387cd010dc88dc499de0f3ba8d88410afa3cc793d4bef1f21608946a3705bf0c0013bc4c5b4448cef9254ad00e4317a372b5")

	// // Добавляем необходимые атрибуты
	// req.Header.Add("Content-Type", "application/json; charset=utf-8")
	// // req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.895 Yowser/2.5 Safari/537.36")

	// req.Header.Add("authority", "www.zara.com")
	// req.Header.Add("accept", "application/json,text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	// req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	// req.Header.Add("cache-control", "max-age=0")
	// req.Header.Add("cookie", "ITXDEVICEID=521da461e07ebe692addd4a984daffb4; ITXSESSIONID=814f46867501b4a96e9dd26f9190e054; _abck=3178DE4EB0FC6D4D71BA33BDF2838B74~-1~YAAQFs8tizPDMaqLAQAAQxi2xwocSNabi3jk05LuGZXy0zdnXFPRD02oN8D18MG+suoeX7cNVlsgcyiFw0oVsoDJg9d1xbzgi1rejOFc0ttLcolMihtqRFxhpv0T5GvJC074bQZ1pj72Y4No9iwiAkFOYJjOkJSzubfO8may7WK4TG2fpsSPADqyBVP6VsqHbEyFbqBPEYb/kg9/9GeUWudVfc0qKxerUjG1A5Rk1KfN9q+Y7ueEP3DjtsMT/4aTo0M01kVD5rPeLy5XswF0jsCq8so+29IjrXoRtYCW92/YMlJMhCfKuAwNY58UVmf34OQQpp3noOgmc8fh0U2z8c6kFYu8mZzMRROKnqe2PctFYohtapBr9rEFMrN/BAnWrUbuXoutAe2R~-1~-1~-1; bm_sz=D4E3113CAC4C27E25ED82C2C4D7FEA67~YAAQFs8tizXDMaqLAQAAQxi2xxXerrp899029mxbJ0UbYKAPscPJ7MdFC8pyDWs3WtbAs/GAadgsWC3/o627XnxlxJU6XDBA79WbT3NQouW7FglFJb5TVQAES4ZX0Z7moTvh2EIJx+qhnCZ5yNsEMo6L8EZRJJYY89uI0HcJnGcZqt2NE7xcKYb9l5+kIVWwjPZJ2DwNgoKjeoMgasT+dsMa394ZWinXSTEGG2lv4JnhUvj6rr9wGN09swKfjoQswiQxWxW4jswLKSxEggKt22kdJIfBm1ojwQGsVBHPIC/4~3422277~3683639; ak_bmsc=0A6BF1EC015061BC427C318FD65E367F~000000000000000000000000000000~YAAQFs8ti0DDMaqLAQAARBm2xxX6ZT9n/+vJrZYqWgSW/7MVdKPQ/4brG5RunYitaVuVj9jb1sfci0w2CmW57JRIftT7XEG+vq951JAG30wpIkppCJgcu8pNjgGVs3wDNHKg26BywTa8QxY4S26x+1wYX77whZ8CocHoWXmF80olnRT/IivL+DHRkdJ9vEMJ6gJaGn6FMq1vzcSO0SCLoCE+mttlKxg7SrTnS2W09uskiEakbfnoaYmmldFqKIt0HAv4xxUlyz3EZQXhvlMjoeLMjLOTOz/ZFzILNDuBZFFToP9/8UKffJYM9mzWhjctWiWrDd5MZ075PZ2uGYv3KLVgWzz0oZqvlpgc6N+R5i4Hw6OlOrFvECgxHKxCZf5dXaXNvNlaMjyYHNTjJ13SnEDSXvRhnm+wDe0//vPIIcvEfXWG2Yb1VqcuipZn7py0Eh4pt2pOc92Y; TS0122c9b6=01113394fd40ceeff94985f8c16cc5b4370728355eb324f783b85551763dd6c53d14580fd25d4e5999c72f5e759afbd14ed8a92407; IDROSTA=1b976b492a26:22adffaf8d69ea60c453e4781; bm_sv=BA8A1AA8044F4F85CFE99FF632F120F3~YAAQNM8ti0MKfbCLAQAAEaa5xxUY0CndHDvxl9TjtMdXxjQhy6JcjMg7+gwRZN9Rmk0q4q7uP9P0aILH8N7I1TmsbXLt9NUIc26VkmOXKAT8ip1ZJjC+Z0XSG1pU20Lg/LDKPxVNGJmhTI0qBArMMube0KPgELZZOkbA9kV+F7ihKbJSCq0gUjUEc3YLE8zpNS8CaqdErtjP2AqZshJ4KmtaP1d99vkHiLeq0AsdikpsgQdNzf4KglcjBFlfUQ==~1")
	// req.Header.Add("sec-ch-ua", `"Chromium";v="116", "Not)A;Brand";v="24", "YaBrowser";v="23"`)
	// req.Header.Add("sec-ch-ua-mobile", "?0")
	// req.Header.Add("sec-ch-ua-platform", `"Windows"`)
	// req.Header.Add("sec-fetch-dest", "document")
	// req.Header.Add("sec-fetch-mode", "navigate")
	// req.Header.Add("sec-fetch-site", "none")
	// req.Header.Add("sec-fetch-user", "?1")
	// req.Header.Add("upgrade-insecure-requests", "1")
	// req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.837 YaBrowser/23.9.4.837 Yowser/2.5 Safari/537.36")

	// req.Header.Add("Cookie", "IDROSTA=c26e0553b1ee:2d4861fb6604c4e7594a43fa8; ITXDEVICEID=5c8042fcbdbca1a24dbf73189f1ffff5; ITXSESSIONID=fb84e41263cecf6bf04f73c1926b897c; _abck=5E8D2C29365BBC0C2A37267FDB23A6BD~-1~YAAQLs8ti/9QDaqLAQAAKQjZxwoy1VsABJdNhuWmnUuS/ADYq/whylyAm9s6OM6dE9eP5tV+f9n+1lAc8OMf8vGtJNL91Ql0JqzLZ0g79OL368P6k02DcU3jOYDpGe6NxQm49jJ3essVSHUk1WxXoc4mmchJiI4gqRfvxNa7zeO0WU7W/lra9B2YTkA6OAGt7MCFs5sOPn1Mj0mVlVxpFFQuLKd+s49IKcccmtELELlaRa1r7WAaHa06bFZBMahuTZSa6T779rBJuY5Duhwzb5h1QaRliweirh6IdCew9o+iZjaRWn7IVOBq+EttnuFqhherPOAQ1jEC9XZyETdJ5mjPJal+KGIb/6581e2/rHZjrsiqaVlTlqFcaTBA6TLPSGsrHVADOOxaH/XtEIL3sP30pT40M0g=~-1~-1~-1; _ga=GA1.1.735583386.1699385774; ak_bmsc=3C4501BD7616BEA5802143086D5DAD2F~000000000000000000000000000000~YAAQLs8tiwBRDaqLAQAAKQjZxxV/GY4/nIE8YiOzYL/0rkD35YKU9fvFcopRK0UKRO4Af75+1XlDqnmmTT6GHLNHZOw+d+VjpmS+035Wgui/J/UtQ2pPl242Jyzz7NkSBmnV6IFbkWQ97RUSvnTkVcAmKNXCl62tRCRzxH/RFAjBvyHORxwz0+DWPONYNJPzrvMrI7KrlSbncvAl1upF3JI5juigM7B6R8Sj6IQHjHqE6WIevlNYK+y/heG+opNaVakj166NKfDSEZLfm8NY3iPqrkOgZla4aR5DLHEEU/LTJZpLtf9wrLCki7xILnTo+1UhwVcB7ArYWk95c13aSC7j1Jhh02j1vWM3V7glUaQMYXw0URbQSxVt/6g7MnoqvhxPX3JrtGwFXqqOrQU=; bm_sz=6C71D4C36B5F639985E401EB056CD093~YAAQLs8tiwFRDaqLAQAAKQjZxxX2abdOdmawGIF680MjBeCheqJi5R72H8BKAIWwlizOB1ccepPSo22+8OkhgeJVlYWVaSYTHbUBTYYIWyaH0KvIqQfs9CZ9/a88MVETxyDqTHb9p/nQsJsFYE5QgLs83Y41k8lBwT0u4DgW4w1HgY+iXm0J+Sswu5lTUpqxQC6eVyoeA/asQ0NLQOADjyvuW3gHNRqV6324RErOLOvK286ic/RV6hhdS6yWerYR2z1JQUvD1bc8kOle84Xh4HtvgBvUkBU2ZeDufZ2vjNWw~4536646~3490118; TS0122c9b6=011f37387c09bf6bf5e6d2666c7855d7da4aabbcee8ff375251ced6a70187869b072814d13e105c5d65417a436beeb163b766ba2a0")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	// Выполнить запрос
	res, err := client.Do(req)
	if err != nil {
		return Touch{}, err
	}

	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return Touch{}, fmt.Errorf("LoadTouch: wrong status code: %d", res.StatusCode)
	}

	// Читаем ответ в массив байтов
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Touch{}, err
	}

	// fmt.Println(string(bodyBytes))

	// Декодируем полученный json и получаем данные
	err = json.Unmarshal(bodyBytes, &tou)
	if err != nil {
		return Touch{}, err
	}

	return tou, nil
}
