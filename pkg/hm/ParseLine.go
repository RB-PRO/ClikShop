package hm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/gocolly/colly"
	"github.com/playwright-community/playwright-go"
)

type Line struct {
	Total      int `json:"total"`
	ItemsShown int `json:"itemsShown"`
	Filters    []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Filtervalues []struct {
			Name        string `json:"name"`
			ID          string `json:"id"`
			Color       string `json:"color"`
			Filtercount int    `json:"filtercount"`
			Code        string `json:"code"`
			Selected    bool   `json:"selected"`
			Disabled    bool   `json:"disabled"`
		} `json:"filtervalues"`
	} `json:"filters"`
	Sortby struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Filtervalues []struct {
			Name     string `json:"name"`
			ID       string `json:"id"`
			Code     string `json:"code"`
			Selected bool   `json:"selected"`
		} `json:"filtervalues"`
	} `json:"sortby"`
	Products []struct {
		ArticleCode string `json:"articleCode"`
		OnClick     string `json:"onClick"`
		Link        string `json:"link"`
		Title       string `json:"title"`
		Category    string `json:"category"`
		Image       []struct {
			Src          string `json:"src"`
			DataAltImage string `json:"dataAltImage"`
			Alt          string `json:"alt"`
			DataAltText  string `json:"dataAltText"`
		} `json:"image"`
		LegalText                 string `json:"legalText"`
		PromotionalMarkerText     string `json:"promotionalMarkerText"`
		ShowPromotionalClubMarker bool   `json:"showPromotionalClubMarker"`
		ShowPriceMarker           bool   `json:"showPriceMarker"`
		FavouritesTracking        string `json:"favouritesTracking"`
		FavouritesSavedText       string `json:"favouritesSavedText"`
		FavouritesNotSavedText    string `json:"favouritesNotSavedText"`
		MarketingMarkerText       string `json:"marketingMarkerText"`
		MarketingMarkerType       string `json:"marketingMarkerType"`
		MarketingMarkerCSS        string `json:"marketingMarkerCss"`
		Price                     string `json:"price"`
		RedPrice                  string `json:"redPrice"`
		YellowPrice               string `json:"yellowPrice"`
		BluePrice                 string `json:"bluePrice"`
		ClubPriceText             string `json:"clubPriceText"`
		SellingAttribute          string `json:"sellingAttribute"`
		SwatchesTotal             string `json:"swatchesTotal"`
		Swatches                  []struct {
			ColorCode   string `json:"colorCode"`
			ArticleLink string `json:"articleLink"`
			ColorName   string `json:"colorName"`
		} `json:"swatches"`
		PreAccessStartDate string `json:"preAccessStartDate"`
		PreAccessEndDate   string `json:"preAccessEndDate"`
		PreAccessGroups    []any  `json:"preAccessGroups"`
		OutOfStockText     string `json:"outOfStockText"`
		ComingSoon         string `json:"comingSoon"`
		BrandName          string `json:"brandName"`
		DamStyleWith       string `json:"damStyleWith"`
	} `json:"products"`
	Labels struct {
		FilterBy      string `json:"filterBy"`
		TotalCount    string `json:"totalCount"`
		ShowItemsText string `json:"showItemsText"`
		LoadMoreText  string `json:"loadMoreText"`
	} `json:"labels"`
	Datatracking struct {
		FilterUsed       string `json:"filterUsed"`
		FilterChanged    string `json:"filterChanged"`
		FilterRemoved    string `json:"filterRemoved"`
		LoadMoreProducts string `json:"loadMoreProducts"`
	} `json:"datatracking"`
}

// Вернуть все товары данной категории
func LinesAll(CategoryLink string) (Line, error) {
	// Получить к-во товаров
	Count, ErrLinesCount := LinesCount(CategoryLink)
	if ErrLinesCount != nil {
		return Line{}, ErrLinesCount
	}
	// Получить сами товары
	lines, ErrLines := Lines(CategoryLink, Count)
	if ErrLines != nil {
		return Line{}, ErrLines
	}
	return lines, nil
}

// Вернуть к-во товаров в этой категории
func LinesCount(CategoryLink string) (int, error) {
	lines, ErrLines := Lines(CategoryLink, 0)
	if ErrLines != nil {
		return 0, ErrLines
	}
	return lines.Total, nil
}

// Сделать запрос на загрузку списка товаров
func Lines(CategoryLink string, PageSize int) (Line, error) {
	url := fmt.Sprintf("%s%s?page-size=%d", URL, CategoryLink, PageSize) // Рабочая ссылка для парсинга
	// fmt.Println("Lines:", url)
	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return Line{}, ErrNewRequest
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	// Выполнить запорос
	Response, ErrDo := client.Do(req)
	if ErrDo != nil {
		return Line{}, ErrDo
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return Line{}, ErrorReadAll
	}

	// Распарсить полученный json в структуру
	var DataLine Line
	ErrorUnmarshal := json.Unmarshal(BodyPage, &DataLine)
	if ErrorUnmarshal != nil {
		return Line{}, ErrorUnmarshal
	}

	return DataLine, nil
}

// Конвертировать структуру Line в структу продуктов
func Line2Product2(line Line, cat []bases.Cat, GenderCatrogy string) []bases.Product2 {
	products := make([]bases.Product2, len(line.Products))
	for i := 0; i < len(line.Products); i++ {
		products[i].Name = line.Products[i].Title
		products[i].Link = URL + line.Products[i].Link
		products[i].Manufacturer = line.Products[i].BrandName
		products[i].Article = line.Products[i].ArticleCode
		products[i].GenderLabel = GenderCatrogy
		products[i].Cat = cat

		// Цена
		var PriceStr string
		if line.Products[i].RedPrice == "" {
			PriceStr = line.Products[i].Price
		} else {
			PriceStr = line.Products[i].RedPrice
		}
		PriceStr = strings.ReplaceAll(PriceStr, "TL", "")
		PriceStr = strings.ReplaceAll(PriceStr, ",", ".")
		PriceStr = strings.ReplaceAll(PriceStr, " ", "")
		// В случае нахождения определённого пробела в виде ASCLL символов 194 и 160, домается переврд во float
		// Например 3 299.00 не отработает.
		// Поэтому идём по всему слайсу байтов и отфильтровывем цифры и точки
		bytes := []byte(PriceStr)
		FilterBytes := make([]byte, 0)
		for _, v := range bytes {
			if v >= 46 && v <= 59 {
				FilterBytes = append(FilterBytes, v)
			}
		}
		price, _ := strconv.ParseFloat(string(FilterBytes), 64)

		// Цикл по всем цветам
		products[i].Item = make([]bases.ColorItem, len(line.Products[i].Swatches))
		for j := 0; j < len(line.Products[i].Swatches); j++ {
			// RGB - line.Products[i].Swatches[j].ColorCode
			products[i].Item[j].Link = URL + line.Products[i].Swatches[j].ArticleLink
			products[i].Item[j].ColorEng = line.Products[i].Swatches[j].ColorName
			products[i].Item[j].ColorCode = bases.KeepLettersAndSpaces(bases.Translit(line.Products[i].Swatches[j].ColorName))
			products[i].Item[j].Price = price
		}
	}
	return products
}

func (core *ParsingCard) LineUrl(link string) (string, error) {

	core.page.Goto(URL + link) // Переходим по ссылке с запроса
	// time.Sleep(5 * time.Second)

	// screenfile := strings.ReplaceAll(link, "/", "")
	// screenfile = strings.ReplaceAll(screenfile, ".html", ".png")
	// core.Screen(screenfile)
	Handle, _ := core.page.WaitForSelector("form[class=js-product-filter-form]",
		playwright.PageWaitForSelectorOptions{
			State:   playwright.WaitForSelectorStateAttached,
			Timeout: playwright.Float(30000),
		}) // WaitForSelector

	if Handle != nil {
		AttrLink, ErrAttr := Handle.GetAttribute("data-filtered-products-url")
		if ErrAttr != nil {
			return "", ErrAttr
		}
		return AttrLink, nil
	}

	return "", errors.New("LineUrl: Не дождался появления тега со ссылкой")
}

// Сделать запрос на загрузку списка товаров
func LineUrl2(link string) (FormLink string, Err error) {

	// var err error
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36"

	// Find and visit all links
	c.OnHTML("form[class=js-product-filter-form]", func(e *colly.HTMLElement) {
		FormLink = e.Attr("data-filtered-products-url")
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		Err = fmt.Errorf("LineUrl2: Request URL: %v failed with response: %v Error: %v", r.Request.URL, r, err)
	})

	c.Visit(URL + link)
	return FormLink, Err
}

// Сделать скриншот браузера
func (core *ParsingCard) Screen(FileName string) (ErrorScreen error) {
	_, ErrorScreen = core.page.Screenshot(playwright.PageScreenshotOptions{Path: playwright.String("tmp/" + FileName)})
	if ErrorScreen != nil {
		return ErrorScreen
	}
	return nil
}
