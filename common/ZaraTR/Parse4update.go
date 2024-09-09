package zaratr

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"

	"ClikShop/common/bases"
	"github.com/playwright-community/playwright-go"
)

// Некоторые товары сменили свою ссылку. Эта функция позволяет обойти это ограничение
// и в случае необходимости повторить запрос до конечной ручки.
//
// Пример:
//
// https://www.zara.com/tr/en/short-trench-coat-p07999818.html?ajax=true
func (s *Service) LoadFantomTouch(id string) (bases.Product2, error) {
	c, err := s.NewServiceCollector()
	if err != nil {
		return bases.Product2{}, errors.Wrap(err, "create service collector: ")
	}

	url := fmt.Sprintf("https://www.zara.com/tr/en/%s.html?ajax=true", id)

	// TODO: Timeout: time.Second * 10
	// s.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"

	headers := http.Header{}
	headers.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	var responseProduct bases.Product2
	c.OnResponse(func(r *colly.Response) {
		response := make(map[string]interface{})
		if err := json.Unmarshal(r.Body, &response); err != nil {
			log.Println("ERROR:500:", err)
			return
		}

		// Get the "location" data
		location, ok := response["location"].(string)
		if !ok {
			// fmt.Println("error: 'location' field not found or not a string")
			location = url
			// return bases.Product2{}, fmt.Errorf("error: 'location' field not found or not a string")
		}

		// Print the location data
		// fmt.Println("Location:", location)
		bodyBytes, err := s.detals(location)
		if err != nil {
			// return bases.Product2{}, fmt.Errorf("detals: %v", err)
			log.Println("ERROR:500:", err)
			return
		}

		var tou Touch
		if err := json.Unmarshal(bodyBytes, &tou); err != nil {
			// return bases.Product2{}, fmt.Errorf("json.Unmarshal: %v", ErrUnmarshalTouch)
			log.Println("ERROR:500:", err)
			return
		}
		responseProduct = Touch2Product2(tou)
	})

	return responseProduct, c.Request(http.MethodGet, url, nil, nil, headers)
}

func (s *Service) detals(url string) ([]byte, error) {
	c, err := s.NewServiceCollector()
	if err != nil {
		return nil, errors.Wrap(err, "create service collector: ")
	}

	// TODO: Timeout: time.Second * 10
	// s.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"

	headers := http.Header{}
	headers.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	var responseBytes []byte
	c.OnResponse(func(r *colly.Response) {
		responseBytes = r.Body
	})

	return responseBytes, c.Request(http.MethodGet, url, nil, nil, headers)
}

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
