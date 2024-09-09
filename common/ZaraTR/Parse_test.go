package zaratr

import (
	"ClikShop/common/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"ClikShop/common/bases"
)

func TestCatCycle(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	zaraService := New(cfg)
	if err != nil {
		t.Error(err)
	}

	CatArr := zaraService.CatCycle() // Наполнить цикл
	fmt.Println(len(CatArr.Items))

}
func TestParsing(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	zaraService := New(cfg)
	if err != nil {
		t.Error(err)
	}

	// go test -timeout 12000s -run ^TestParsing$ ClikShop/common/ZaraTR
	zaraService.Parsing()
}
func TestParseTouch2Product2(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	zaraService := New(cfg)
	if err != nil {
		t.Error(err)
	}

	product, err := zaraService.LoadFantomTouch("running-trainers-with-chunky-soles-p12341120")
	if err != nil {
		t.Error(err)
	}

	// fmt.Printf("%+v", Product)
	fmt.Println("Товар:")
	// fmt.Printf("%+v\n", Product.Item)
	fmt.Printf("%+v\n", product.Name)
}

// Комплексный тест парсинга
func TestComplexParse(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	zaraService := New(cfg)
	if err != nil {
		t.Error(err)
	}

	// Категории
	CatArr := zaraService.CatCycle() // Получить все категории
	fmt.Println("Всего", len(CatArr.Items), "категорий")
	// for _, cat := range CatArr.Items {
	// 	if cat.ID.value == "2184366" {
	// 		fmt.Println(cat.ID, cat.Name, cat.Cat)
	// 	}
	// }
	var cat Item

	for ind, val := range CatArr.Items {
		if val.ID.value == "2184366" {
			cat = val
			fmt.Printf("%v - cat: %+v\n\n", ind, cat)
		}
	}
	if cat.ID.value == "" {
		t.Error("Не нашёл товар с категорией 2184366")
	}

	// Список всех товаров
	// cat := CatArr.Items[1]
	fmt.Println("ID категории", cat.ID.value)
	fmt.Println("Категория товара:", cat.Cat) // WOMAN > SHIRTS > Satin
	fmt.Printf("Весь товар: %v\n\n", cat.Cat) // WOMAN > SHIRTS > Satin
	line, ErrorLine := zaraService.LoadLine(fmt.Sprintf("%v", cat.ID.value))
	if ErrorLine != nil {
		fmt.Println(ErrorLine)
	}

	/////////////

	ProductsLine := make([]CommercialComponents, 0)
	if len(line.ProductGroups) != 0 {
		if len(line.ProductGroups) != 0 {
			if len(line.ProductGroups[0].Elements) != 0 {
				for ind := range line.ProductGroups[0].Elements[0].CommercialComponents { // Циклом обновляем категории
					if line.ProductGroups[0].Elements[0].CommercialComponents[ind].Type == "Product" { // Если это сам товар
						line.ProductGroups[0].Elements[0].CommercialComponents[ind].Cat = cat.Cat
						ProductsLine = append(ProductsLine, line.ProductGroups[0].Elements[0].CommercialComponents[ind])
					}
				}
			}
		}
	}
	fmt.Println("Всего", len(ProductsLine), "товара(ов)")

	// Сам товар
	prod := ProductsLine[0]
	var Variety bases.Variety2
	fmt.Println("Ссылка на товар", (prod.Seo.Keyword + "-p" + prod.Seo.SeoProductID))
	touch, _ := zaraService.LoadTouch(prod.Seo.Keyword + "-p" + prod.Seo.SeoProductID)
	Prod2 := Touch2Product2(touch)
	Prod2.Cat = prod.Cat // Обновляем категнории

	fmt.Printf("%+v", Prod2)

	Variety.Product = append(Variety.Product, Prod2)
	Variety.SaveXlsx("Zara")
}

func TestTou(t *testing.T) {
	// Specify the URL you want to access
	url := "https://www.zara.com/tr/en/running-trainers-with-chunky-soles-p12341120.html?ajax=true"

	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Create an HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers to mimic a real browser request
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	// Set any additional headers if needed

	// Set cookies if needed
	// req.Header.Set("Cookie", "key1=value1; key2=value2")

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	// fmt.Println(string(body))

	var responseData map[string]interface{}

	// Unmarshal JSON
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Get the "location" data
	location, ok := responseData["location"].(string)
	if !ok {
		fmt.Println("Error: 'location' field not found or not a string")
		return
	}

	// Print the location data
	// fmt.Println("Location:", location)
	detal(location)
}

func detal(location string) {
	// Specify the URL you want to access
	url := location

	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Create an HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers to mimic a real browser request
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	// Set any additional headers if needed

	// Set cookies if needed
	// req.Header.Set("Cookie", "key1=value1; key2=value2")

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}
