package zaratr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/playwright-community/playwright-go"
)

// Некоторые товары сменили свою ссылку. Эта функция позволяет обойти это ограничение
// и в случае необходимости повторить запрос до конечной ручки.
//
// Пример:
//
// https://www.zara.com/tr/en/short-trench-coat-p07999818.html?ajax=true
func LoadFantomTouch(id string) (Product bases.Product2, Err error) {
	// Specify the URL you want to access
	url := "https://www.zara.com/tr/en/running-trainers-with-chunky-soles-p12341120.html?ajax=true"
	// url := "https://www.zara.com/tr/en/running-trainers-with-chunky-soles-p12341120.html?ajax=true"
	url = fmt.Sprintf("https://www.zara.com/tr/en/%s.html?ajax=true", id)
	// fmt.Println(url)
	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: time.Second * 10}

	// Create an HTTP request
	req, ErrNewRequest := http.NewRequest("GET", url, nil)
	if ErrNewRequest != nil {
		return bases.Product2{}, fmt.Errorf("http.NewRequest: Error creating request: %v", ErrNewRequest)
	}

	// Set headers to mimic a real browser request
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	// Set any additional headers if needed

	// Set cookies if needed
	// req.Header.Set("Cookie", "key1=value1; key2=value2")

	// Make the HTTP request
	resp, ErrDo := client.Do(req)
	if ErrDo != nil {
		return bases.Product2{}, fmt.Errorf("client.Do: %v", ErrDo)
	}
	defer resp.Body.Close()

	// Read the response body
	body, ErrReadAll := ioutil.ReadAll(resp.Body)
	if ErrReadAll != nil {
		return bases.Product2{}, fmt.Errorf("ioutil.ReadAll: Error reading response body: %v", ErrReadAll)
	}

	// Print the response body
	// fmt.Println(string(body))

	var responseData map[string]interface{}

	// Unmarshal JSON
	ErrUnmarshal := json.Unmarshal(body, &responseData)
	if ErrUnmarshal != nil {
		return bases.Product2{}, fmt.Errorf("son.Unmarshal: Error decoding JSON: %v", ErrUnmarshal)
	}

	// Get the "location" data
	location, ok := responseData["location"].(string)
	if !ok {
		// fmt.Println("error: 'location' field not found or not a string")
		location = url
		// return bases.Product2{}, fmt.Errorf("error: 'location' field not found or not a string")
	}

	// Print the location data
	// fmt.Println("Location:", location)
	bodyBytes, ErrDetals := detals(location)
	if ErrDetals != nil {
		return bases.Product2{}, fmt.Errorf("detals: %v", ErrDetals)
	}

	var tou Touch
	ErrUnmarshalTouch := json.Unmarshal(bodyBytes, &tou)
	if ErrUnmarshalTouch != nil {
		return bases.Product2{}, fmt.Errorf("json.Unmarshal: %v", ErrUnmarshalTouch)
	}

	return Touch2Product2(tou), nil
}

func detals(url string) ([]byte, error) {

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: time.Second * 10}

	// Create an HTTP request
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return nil, fmt.Errorf("client.Do: Error creating request: %v", ErrNewRequest)
	}

	// Set headers to mimic a real browser request
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	// Make the HTTP request
	resp, ErrDo := client.Do(req)
	if ErrDo != nil {
		return nil, fmt.Errorf("client.Do: Error making request: %v", ErrDo)
	}
	defer resp.Body.Close()

	// Read the response body
	body, ErrReadAll := ioutil.ReadAll(resp.Body)
	if ErrReadAll != nil {
		return nil, fmt.Errorf(" ioutil.ReadAll: Error reading response body: %v", ErrReadAll)
	}

	// fmt.Println(string(body))
	return body, nil
}

////////////////////////
///////////////////////////

// Спасить со страницы фронта
func parseFront(id string) (Product bases.Product2, Err error) {
	// c := colly.NewCollector() // Create a collector
	// c.UserAgent = ("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.837 YaBrowser/23.9.4.837 Yowser/2.5 Safari/537.36")

	// // Set HTML callback
	// // Won't be called if error occurs
	// c.OnHTML("*", func(e *colly.HTMLElement) {
	// 	// fmt.Println(e)
	// })

	// // Set error handler
	// c.OnError(func(r *colly.Response, err error) {
	// 	Err = fmt.Errorf("request URL: %v failed with response: %v nError: %v",
	// 		r.Request.URL, r, err)
	// })

	// url := fmt.Sprintf("%s/tr/en/%s.html", URL, id)
	// Err = c.Visit(url)

	// Launching the driver internally
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not launch playwright: %v", err)
	}
	// Start the Chromium browser
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		log.Fatalf("could not launch Chromium: %v", err)
	}
	// Creates internally a context and a new page
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	// Visit the website and wait for a network idle for at least 500ms
	url := fmt.Sprintf("https://www.zara.com/tr/en/%s.html?ajax=true", id)

	fmt.Println(url)
	url = "https://www.zara.com/tr/en/running-trainers-with-chunky-soles-p12341120.html?ajax=true"
	url = URL + "/tr/en/running-trainers-with-chunky-soles-p12341120.html?ajax=true&bm-verify=AAQAAAAI_____5lwzlu22VtQaoh-9PiUIhyOgcPo7q7uLb9PJnFXoQHFkGw6azaqj3INCtml-ncQeHVuacOgEzL_lQD66tShwJ1bFmVhHlhO_yF5UW6CJQs8SNLQnr3veBGv5uja4HD9xTnSjnCrZWKCqV9kxqY2m9knCEVeFs99iGTMb85zlsY-mQUnnbqOurvcWrU6EYJ1cAc2CbRAp5ZwpMusWOxQgzo8GAGI5n91e-bLbtsaeMWKg0fYV8q9v7NGYW_hg8el_0Pp0XFonCncRls84_pdM3sE_hjPYptE3bTGDnJHGPc3XM6NSMk7O7jaWeMUPPo"

	fmt.Println(url)
	if _, err = page.Goto(url); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	time.Sleep(time.Second)
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

	return Product, Err
}
