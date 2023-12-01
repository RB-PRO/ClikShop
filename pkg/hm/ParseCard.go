package hm

import (
	"fmt"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/playwright-community/playwright-go"
)

type ParsingCard struct {
	pw      *playwright.Playwright
	browser playwright.Browser
	page    playwright.Page
}

func NewParsingCard() (*ParsingCard, error) {

	pw, err := playwright.Run(&playwright.RunOptions{})
	if err != nil {
		return nil, err
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)})
	if err != nil {
		return nil, err
	}

	page, err := browser.NewPage(playwright.BrowserNewContextOptions{})
	if err != nil {
		return nil, err
	}
	// page.SetDefaultTimeout(100000.0)
	// page.SetDefaultNavigationTimeout(60)

	return &ParsingCard{
		pw:      pw,
		browser: browser,
		page:    page,
	}, nil
}

// Остановить ядро парсинга
func (core *ParsingCard) Stop() error {

	if err := core.browser.Close(); err != nil {
		return err // could not close browser
	}

	if err := core.pw.Stop(); err != nil {
		return err // could not stop Playwright
	}

	return nil
}

// Пропарсить карточку товара со всеми цветами
func (core *ParsingCard) VariableProduct3(Prod *bases.Product2, IndexItem int) (ErrParseProduct error) {
	Prod.Specifications = make(map[string]string)

	// page, err := core.browser.NewPage()
	// if err != nil {
	// 	return err
	// }

	// Переходим по ссылке с запроса
	core.page.Goto(Prod.Item[IndexItem].Link)

	// Картинки
	core.page.Locator("div[class=sticky-candidate] img", playwright.PageLocatorOptions{})
	//core.page.WaitForSelector("div[class=sticky-candidate] img", playwright.PageWaitForSelectorOptions{Timeout: playwright.Float(30)}) // С этой строкой он ждёт загрузку всех картинок
	core.page.WaitForSelector("div[class=sticky-candidate] > figure > img", playwright.PageWaitForSelectorOptions{})

	images, ErrImages := core.page.QuerySelectorAll("div[class=sticky-candidate] img")
	if ErrImages != nil {
		fmt.Println(ErrImages)
	}
	for _, ImageSelector := range images {
		ImageLinkMain, _ := ImageSelector.GetAttribute("src")
		ImageLinkMain = strings.ReplaceAll(ImageLinkMain, "[file:/product/fullscreen]", "[file:/product/main]")
		ImageLinkMain = strings.ReplaceAll(ImageLinkMain, "[file:/product/fullscreen", "[file:/product/main]")
		ImageLinkMain = strings.ReplaceAll(ImageLinkMain, "%2Fmain%5D", "%2Fmain%5D")
		ImageLinkMain = strings.ReplaceAll(ImageLinkMain, "call=url[file:/product/main]", "call=url%5Bfile%3A%2Fproduct%2Fmain%5D")
		ImageLinkMain = strings.ReplaceAll(ImageLinkMain, "call=url[file:/product/main", "call=url%5Bfile%3A%2Fproduct%2Fmain%5D")
		ImageLinkMain = strings.ReplaceAll(ImageLinkMain, "&call=url[file:/product/main]", "&call=url%5Bfile%3A%2Fproduct%2Fmain%5D")
		ImageLinkMain = strings.ReplaceAll(ImageLinkMain, "&call=url[file:/product/main", "&call=url%5Bfile%3A%2Fproduct%2Fmain%5D")
		// ImageLinkMain = strings.ReplaceAll(ImageLinkMain, "set=format%5Bwebp%5D", "set=format%5Bjpeg%5D")

		ImageLinkMain = "https:" + ImageLinkMain
		Prod.Item[IndexItem].Image = append(Prod.Item[IndexItem].Image, ImageLinkMain)
	}

	// sel, _ := core.page.QuerySelector("div[class=column1]")
	// html, _ := sel.InnerHTML()
	// f, _ := os.Create("site3.html")
	// f.WriteString(html)
	// f.Close()

	// Размеры
	sizes, ErrSizes := core.page.QuerySelectorAll("ul[class^=ListGrid-module--listGrid] label")
	if ErrSizes != nil {
		return ErrSizes
	}
	for _, SizeSelector := range sizes {
		var BaseSize bases.Size
		BaseSize.Val, _ = SizeSelector.GetAttribute("for")
		if Hidden, _ := SizeSelector.GetAttribute("aria-disabled"); Hidden != "true" {
			BaseSize.IsExit = true
		}
		Prod.Item[IndexItem].Size = append(Prod.Item[IndexItem].Size, BaseSize)
	}

	// Описание
	Description, ErrDeskription := core.page.QuerySelector("div[id=section-descriptionAccordion] p")
	if ErrDeskription != nil {
		fmt.Println(ErrDeskription)
	}
	Prod.Description.Eng, _ = Description.InnerText()

	// Вторичное описание
	Description2, ErrDeskription2 := core.page.QuerySelectorAll("div[id=section-descriptionAccordion]>div>div>dl>div")
	if ErrDeskription2 != nil {
		fmt.Println(ErrDeskription2)
	}
	for _, div := range Description2 {
		dt, _ := div.QuerySelector("dt")
		dd, _ := div.QuerySelector("dd")
		dt_test, _ := dt.TextContent()
		dd_test, _ := dd.TextContent()
		dt_test = strings.ReplaceAll(dt_test, ":", "")
		dt_test = strings.TrimSpace(dt_test)
		Prod.Description.Eng += "\n" + dt_test + " - " + dd_test
		Prod.Specifications[dt_test] = dd_test
	}

	return nil
}

// сфорировать гендер специально для загрузки на сайт
func GenderForming(str string) string {
	switch str {
	case "kadin":
		return "woman"
	case "erkek":
		return "man"
	case "bebek":
		return "kind"
	default:
		return ""
	}
}
