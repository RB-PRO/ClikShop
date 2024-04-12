package hm

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/gocolly/colly"
)

// Получить актуальные данные по товарам.
// По [ссылке] получить данные, как по размерам, так и по ценам
//
// Выведено из использования, в виду того, что загружается только часть тега,
// остальная часть содержимого загружается с помощью js
//
// [ссылке]: https://www2.hm.com/tr_tr/productpage.1163274001.html
func VariableActual2(url string) (Item bases.ColorItem, Err error) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 " +
		"(KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36"

	type ShemasProduct struct {
		Context       string `json:"@context"`
		Type          string `json:"@type"`
		ItemCondition string `json:"itemCondition"`
		Image         string `json:"image"`
		Name          string `json:"name"`
		Color         string `json:"color"`
		Description   string `json:"description"`
		Sku           string `json:"sku"`
		Brand         struct {
			Type string `json:"@type"`
			Name string `json:"name"`
		} `json:"brand"`
		Category struct {
			Type string `json:"@type"`
			Name string `json:"name"`
		} `json:"category"`
		Offers []struct {
			Sku           string `json:"SKU"`
			Type          string `json:"@type"`
			URL           string `json:"url"`
			PriceCurrency string `json:"priceCurrency"`
			Price         string `json:"price"`
			Availability  string `json:"availability"`
			Seller        struct {
				Type string `json:"@type"`
				Name string `json:"name"`
			} `json:"seller"`
		} `json:"offers"`
		AggregateRating struct {
			Type        string  `json:"@type"`
			RatingValue float64 `json:"ratingValue"`
			ReviewCount int     `json:"reviewCount"`
		} `json:"aggregateRating"`
	}

	c.OnHTML(`script[id=product-schema]`, func(e *colly.HTMLElement) {
		jsonTxt := e.DOM.Text()

		// Парсарсить структуру продукта
		var Product ShemasProduct
		ErrUnmarshal := json.Unmarshal([]byte(jsonTxt), &Product)
		if ErrUnmarshal != nil {
			Err = fmt.Errorf("VariableActual2: Unmarshal: %s", ErrUnmarshal)
		}
		// Items = make([]bases.ColorItem, len(Product.Offers))
		// fmt.Printf("%+v\n", Product)
		Price := 0.0

		fmt.Println("len(Product.Offers)", len(Product.Offers))
		fmt.Println(jsonTxt)

		Size := make([]bases.Size, 0)
		for _, offer := range Product.Offers {
			SizeVal := StrFromSKU(offer.Sku)
			fmt.Println("offer.Sku", offer.Sku, "--------", SizeVal)
			IsExit := strings.Contains(offer.Availability, "InStock")
			Price, _ = strconv.ParseFloat(offer.Price, 64)
			if SizeVal != "" {
				Size = append(Size, bases.Size{
					Val:    SizeVal,
					IsExit: IsExit,
				})
			}
		}
		Item.Price = Price
		Item.Size = Size
		Item.ColorEng = bases.Name2Slug(Product.Color)
		Item.ColorCode = bases.KeepLettersAndSpaces(bases.Translit(Product.Color))
	})

	c.Visit(url)
	if Err != nil {
		return bases.ColorItem{}, fmt.Errorf("VariableActual2: %s", Err)
	}
	fmt.Printf("%+v\n", Item)
	return Item, nil
}

func StrFromSKU(str string) string {
	lenstr := len([]byte(str))
	if lenstr == 13 {
		switch str[lenstr-3:] {
		case "001":
			return "XXS"
		case "002":
			return "XS"
		case "003":
			return "S"
		case "004":
			return "M"
		case "005":
			return "L"
		case "006":
			return "XL"
		case "007":
			return "XXL"
		case "016":
			return "3XL"
		case "015":
			return "4XL"
		default:
			return "NOSIZE"
		}
	}
	return ""
}

// Получить строку по размеру
func StrFromSKU2(str string) string {
	lenstr := len([]byte(str))
	if lenstr == 13 {
		switch str[lenstr-3:] {
		case "001":
			return "32"
		case "002":
			return "34"
		case "003":
			return "36"
		case "004":
			return "38"
		case "005":
			return "40"
		case "006":
			return "42"
		case "007":
			return "44"
		case "008":
			return "46"
		case "009":
			return "48"
		case "010":
			return "50"
		default:
			return "NOSIZE"
		}
	}
	return ""
}
func StrFromSKU3(str string) string {
	lenstr := len([]byte(str))
	if lenstr == 13 {
		switch str[lenstr-3:] {
		case "001":
			return "22"
		case "002":
			return "23"
		case "003":
			return "24"
		case "004":
			return "25"
		case "005":
			return "26"
		case "006":
			return "27"
		case "007":
			return "28"
		case "008":
			return "29"
		case "009":
			return "30"
		case "010":
			return "31"
		case "011":
			return "32"
		case "012":
			return "33"
		case "013":
			return "34"
		case "014":
			return "35"
		default:
			return "NOSIZE"
		}
	}
	return ""
}
