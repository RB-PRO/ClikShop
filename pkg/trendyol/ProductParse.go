package trendyol

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
)

const Product_URL string = "https://public.trendyol.com/discovery-web-productgw-service/api/productDetail/%d?storefrontId=1"

func ParseProduct(ProductID int) (pg ProductStruct, Err error) {
	url := fmt.Sprintf(Product_URL, ProductID) // Рабочая ссылка для парсинга
	// fmt.Println("Lines:", url)
	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return ProductStruct{}, fmt.Errorf("http.NewRequest: %v", ErrNewRequest)
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36")

	// Выполнить запорос
	Response, ErrDo := client.Do(req)
	if ErrDo != nil {
		return ProductStruct{}, fmt.Errorf("client.Do: %v", ErrDo)
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return ProductStruct{}, fmt.Errorf("io.ReadAll: %v", ErrorReadAll)
	}

	// Распарсить полученный json в структуру
	ErrorUnmarshal := json.Unmarshal(BodyPage, &pg)
	if ErrorUnmarshal != nil {
		return ProductStruct{}, fmt.Errorf("json.Unmarshal: %v", ErrorUnmarshal)
	}

	return pg, nil
}

// Перевести структуру товара сайта донора в слайс цветов
func Touch2ColorItem(pg ProductStruct) (color bases.ColorItem) {
	// colors := make([]bases.ColorItem, 0, len(pg.Result.AllVariants))

	// По цене немного не то. Тут она для каждого товара своя
	// сперва ищем максимальную цену
	var MaxPrice float64

	// Цикл по всем вариантам и формирование ценовой политике
	for _, variant := range pg.Result.AllVariants {

		// Поиск максимальной цены
		if variant.Price > MaxPrice {
			MaxPrice = variant.Price
		}

		// Слайс размеров
		SizeVal := variant.Value
		SizeVal = strings.ReplaceAll(SizeVal, ",", ".")
		color.Size = append(color.Size, bases.Size{
			Val:      SizeVal,
			IsExit:   variant.InStock,
			DataCode: strconv.Itoa(variant.ItemNumber),
		})
	}
	color.ColorCode = extractColors(pg.Result.Color)
	color.ColorCode = bases.Name2Slug(color.ColorCode)
	color.ColorCode = strings.TrimSpace(color.ColorCode)
	color.Price = MaxPrice
	return color
}

type ColorPriceExit struct {
	Size   string
	Color  string
	IsExit bool
	Price  float64
}

// Перевести структуру товара сайта донора в слайс цветов
func Touch2ColorPriceExit(pg ProductStruct) (varients []ColorPriceExit) {

	varients = make([]ColorPriceExit, 0, len(pg.Result.AllVariants))
	// Цикл по всем вариантам и формирование ценовой политике
	for _, variant := range pg.Result.AllVariants {

		SizeVal := variant.Value
		SizeVal = strings.ReplaceAll(SizeVal, ",", ".")

		color := extractColors(pg.Result.Color)
		color = bases.Name2Slug(color)
		color = strings.TrimSpace(color)

		varients = append(varients, ColorPriceExit{
			Size:   SizeVal,
			IsExit: variant.InStock,
			Color:  color,
			Price:  variant.Price,
		})

	}

	return varients
}
