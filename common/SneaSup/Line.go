package sneaksup

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"ClikShop/common/bases"
)

// Получить список товаров в результате выполнения ф-й Line
func Lines(link string) (ProductSneakSup []Product, Err error) {

	// Цикл по всем страницам
	for i := 1; ; i++ {

		// Загружаем данные
		line, ErrLinePost := LinePost(link, i)
		if ErrLinePost != nil {
			return nil, ErrLinePost
		}

		// Формируем слайс товаров из структур SneakSup
		ProductSneakSup = append(ProductSneakSup, line.Products...)

		// Выход по причине того, что товаров больше нет
		if line.Pager.PageIndex+1 >= line.Pager.TotalPages {
			break
		}
	}
	return ProductSneakSup, Err
}

// Загрузить список товаров
func LinePost(link string, pagenumber int) (Line LineStruct, ErrLine error) {

	url, errMakeLink := linkTranstore(link, pagenumber)
	if errMakeLink != nil {
		return Line, errMakeLink
	}

	// fmt.Println("URL", url)

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

	// Выполнить запорос
	Response, ErrDo := client.Do(req)
	if ErrDo != nil {
		return LineStruct{}, ErrDo
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return LineStruct{}, ErrorReadAll
	}

	// os.WriteFile("LinePost.txt", BodyPage, 0666)

	// Распарсить полученный json в структуру
	ErrorUnmarshal := json.Unmarshal(BodyPage, &Line)
	if ErrorUnmarshal != nil {
		return LineStruct{}, ErrorUnmarshal
	}

	// Заполнение фактической ссылки на товар
	for iLine := range Line.Products {
		Line.Products[iLine].MYLINK = url
	}

	return Line, ErrNewRequest
}

// Преобразовать ссылку на парсинга line
func linkTranstore(link string, pagenumber int) (string, error) {

	// Парсим ссылку в формат отдачи json
	u, ErrParse := url.Parse(link)
	if ErrParse != nil {
		return "", ErrParse
	}

	// Добавляем аттрибусы, которые соответствуют запросу, который отдаёт json
	q := u.Query()
	q.Set("paginationType", "20")
	q.Set("orderby", "0")
	q.Set("pagenumber", strconv.Itoa(pagenumber))
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// Перевести из line в структуру product2
func Line2Product(SSprods []Product, Category SScat) (prods []bases.Product2) {
	// for _, LineProd := range SSprods {
	// 	var prod bases.Product2

	// 	prod.Name = LineProd.Name
	// 	prod.Link = URL + LineProd.URL
	// 	prod.Article = LineProd.Sku
	// 	prod.FullName = LineProd.DefaultPictureModel.AlternateText
	// 	prod.Manufacturer = LineProd.ManufacturerName

	// 	// for _, Sibligs := range prod.Sibligs {

	// 	// }

	// 	prods = append(prods, prod)
	// }

	// Мапа всех товаров с ключом - Название товара
	// Зачем же так сделано? В структуре товаров LineStruct нет внятного порядочного разнесения по цветам товаров
	// и в истоге в конечном запросе Line мы можем только получить данные по по всем товарам по всем вариациям.
	// Создаём такую мапу и все товары у нас будут распределены по названиями,
	// а в слайсе товаров будут только исходные товары с источника, которые нужно будет представить в новом исходном виде товаров
	SSprodMap := make(map[string][]Product)
	for _, SSprod := range SSprods {
		SSprodMap[SSprod.Name] = append(SSprodMap[SSprod.Name], SSprod)
	}

	// //
	// fmt.Println(len(SSprodMap["Nike W Air Max 270"]))
	// fmt.Println(`len(SSprodMap["Nike W Air Max 270"])`, len(SSprodMap["Nike W Air Max 270"]))
	// for range SSprodMap["Nike W Air Max 270"] {
	// 	for _, SSprod := range SSprodMap["Nike W Air Max 270"] {
	// 		for _, SpecModels := range SSprod.SpecificationAttributeModels { // Цикл по аттрибутам подтоваров
	// 			if SpecModels.SpecificationAttributeID == 3 { // Если это аттрибут цвета
	// 				fmt.Println("---", SpecModels.Name, SpecModels.OptionErpCode)
	// 			}
	// 		}
	// 	}
	// }

	// Создаём новую структуру данных и Добавляем данные по товарам
	for NameProduct, SSproducts := range SSprodMap { // Цикл по мапе
		var pr bases.Product2

		pr.Name = NameProduct                                         // Название товара
		pr.Manufacturer = SSproducts[0].ManufacturerName              // Гендер
		pr.Cat = Category.Cat                                         // Категории
		pr.Link = "https://www.sneaksup.com" + SSproducts[0].URL      // Ссылка на товар
		pr.FullName = SSproducts[0].DefaultPictureModel.AlternateText // Полное название

		SKUs := strings.Split(SSproducts[0].Sku, "_")
		if len(SKUs) == 2 {
			pr.Article = SKUs[0]
		} else {
			pr.Article = SSproducts[0].Sku
		}

		// fmt.Println("len(SSproducts)", len(SSproducts))
		for _, SSprod := range SSproducts { // Цикл по товарам с одинаковыми именами
			var ColorItem bases.ColorItem

			ColorItem.Link = SSprod.MYLINK // Ссылка на line товара

			// Цвет
			for _, SpecModels := range SSprod.SpecificationAttributeModels { // Цикл по аттрибутам подтоваров
				if SpecModels.SpecificationAttributeID == 3 { // Если это аттрибут цвета
					ColorItem.ColorEng = SpecModels.Name
					ColorItem.ColorCode = SpecModels.OptionErpCode
				}
				if SpecModels.SpecificationAttributeID == 9 { // Если это аттрибут гендера
					pr.GenderLabel = "unisex"
					if SpecModels.OptionErpCode == "kadin" {
						pr.GenderLabel = "woman"
					}
					if SpecModels.OptionErpCode == "erkek" {
						pr.GenderLabel = "man"
					}
					if strings.Contains(SpecModels.OptionErpCode, "erkek") && len(SpecModels.OptionErpCode) > 5 {
						pr.GenderLabel = "boy"
					}
					if strings.Contains(SpecModels.OptionErpCode, "kadin") && len(SpecModels.OptionErpCode) > 5 {
						pr.GenderLabel = "girl"
					}
					if strings.Contains(SpecModels.OptionErpCode, "uni") && len(SpecModels.OptionErpCode) > 5 {
						pr.GenderLabel = "unisex"
					}
				}
			}

			// Цена В исходном товаре
			ColorItem.Price = SSprod.ProductPrice.PriceValue

			// Картинки в исходном товаре
			for _, Picture := range SSprod.Pictures {
				ColorItem.Image = append(ColorItem.Image, Picture.OriginalImageURL)
			}

			// Размеры товаров
			for _, ProductAttribute := range SSprod.ProductAttributes {
				if ProductAttribute.ProductID == SSprod.ID { // Если аттрибус относится именно к этому товару
					for _, ProductAttributeValue := range ProductAttribute.ProductAttributeValues {
						ColorItem.Size = append(ColorItem.Size, bases.Size{
							Val:    ProductAttributeValue.Name,
							IsExit: ProductAttributeValue.InStock,
						})
					}
				}
			}

			// Дефолтная картинка
			pr.Img = append(pr.Img, SSprod.DefaultPictureModel.OriginalImageURL)

			// Сохраняем вариацию
			if ColorItem.ColorCode != "" {
				pr.Item = append(pr.Item, ColorItem)
			}
		}
		if len(pr.Item) != 0 {
			prods = append(prods, pr) // Суммирование товаров в общий слайс
		}
	}

	return prods
}
