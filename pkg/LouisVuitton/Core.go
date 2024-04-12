package louisvuitton

import (
	"errors"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// Ядро парсинга
type Core struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// СОздать ядро парсинга с вводными данными
func NewCore(ClientID, ClientSecret string) *Core {
	return &Core{ClientID: ClientID, ClientSecret: ClientSecret}
}

var error_authentication_denied error = errors.New("Authentication denied.") // Нет ClientID
var error_invalid_client error = errors.New("Invalid Client")                // Неверный ClientID

// Обновить данные в ядре парсинга
func (cr *Core) UpdateCore() *Core {
	var ClientID string
	var ClientSecret string

	// update
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36"
	c.SetRequestTimeout(time.Minute)
	c.OnHTML("body>script:first-of-type", func(e *colly.HTMLElement) {

		data, _ := e.DOM.Html() // сохраняем текст элемента

		script := string(data)

		IndexStart := strings.Index(script, `&#34;Часы&#34;,&#34;`) + 24
		IndexEnd := strings.Index(script[IndexStart:], `&#34;,&#34;`) + IndexStart
		ClientID = script[IndexStart:IndexEnd]

		IndexStart = IndexEnd + 11
		IndexEnd = strings.Index(script[IndexStart:], `&#34;,&#34;`) + IndexStart
		ClientSecret = script[IndexStart:IndexEnd]

		// fmt.Println(ClientID)
		// fmt.Println(ClientSecret)

	})
	c.Visit("https://ru.louisvuitton.com/rus-ru/homepage")

	// save
	cr.ClientID = ClientID
	cr.ClientSecret = ClientSecret

	return cr
}
