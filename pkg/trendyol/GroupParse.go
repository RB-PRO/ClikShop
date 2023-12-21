package trendyol

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
)

const URL string = "https://www.trendyol.com"
const group_URL string = "https://public.trendyol.com/discovery-web-websfxproductgroups-santral/api/v1/product-groups/%d"

func ParseGroup(ProductGroupID int) (pg GroupStruct, Err error) {
	url := fmt.Sprintf(group_URL, ProductGroupID) // Рабочая ссылка для парсинга
	// fmt.Println("Lines:", url)
	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return GroupStruct{}, fmt.Errorf("http.NewRequest: %v", ErrNewRequest)
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36")

	// Выполнить запорос
	Response, ErrDo := client.Do(req)
	if ErrDo != nil {
		return GroupStruct{}, fmt.Errorf("client.Do: %v", ErrDo)
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return GroupStruct{}, fmt.Errorf("io.ReadAll: %v", ErrorReadAll)
	}

	// Распарсить полученный json в структуру
	ErrorUnmarshal := json.Unmarshal(BodyPage, &pg)
	if ErrorUnmarshal != nil {
		return GroupStruct{}, fmt.Errorf("json.Unmarshal: %v", ErrorUnmarshal)
	}

	return pg, nil
}

// Получить товар по его ID группы
func Product(ProductGroupID int) (prod bases.Product2, Err error) {
	pg, ErrGroup := ParseGroup(ProductGroupID)
	if ErrGroup != nil {
		return prod, fmt.Errorf("ParseGroup: %v", ErrGroup)
	}

	if len(pg.Result.SlicingAttributes) == 0 {
		return productOwner(ProductGroupID)

		// url := fmt.Sprintf(group_URL, ProductGroupID)
		// fmt.Println("Ну тут одиночный товар. Запускаю второй алгоритм")
		// return prod, fmt.Errorf("в товаре %d не найдены slicingAttributes товаров: %v", ProductGroupID, url)
	}

	if len(pg.Result.SlicingAttributes[0].Attributes) == 0 {
		url := fmt.Sprintf(group_URL, ProductGroupID)
		return prod, fmt.Errorf("в товаре %d не найдены attributes товаров: %v", ProductGroupID, url)
	}

	// Производитель
	prod.Manufacturer = pg.Result.SlicingAttributes[0].Brand.Name

	// Ссылка на товар
	prod.Link = fmt.Sprintf(group_URL, ProductGroupID)

	for iDonProd, DonProd := range pg.Result.SlicingAttributes[0].Attributes {
		var ProductID int
		if len(DonProd.Contents) == 1 {
			ProductID = DonProd.Contents[0].ID
		} else {
			continue
		}

		// Парсинг вариации
		pd, ErrProd := ParseProduct(ProductID)
		if ErrProd != nil {
			return prod, fmt.Errorf("ParseProduct: %v", ErrProd)
		}

		// Тк инфа по товару лежит только в самой вариации,
		// то будем брать инфу с первого товара
		if iDonProd == 1 {
			prod.Name = DonProd.Contents[0].Name // Название товара

			// Описание товара
			for _, ds := range pd.Result.Description {
				prod.Description.Eng = prod.Description.Eng + ds.Text + "\n"
			}
			re := regexp.MustCompile("[[^]]*]")
			prod.Description.Eng = re.ReplaceAllString(prod.Description.Eng, "")

			// Гендер товара
			// fmt.Println("pd.Result.Gender.ID ", pd.Result.Gender.ID, pd.Result.Gender.Name)
			switch pd.Result.Gender.ID {
			case 1:
				prod.GenderLabel = "women"
			case 2:
				prod.GenderLabel = "man"
			case 3:
				prod.GenderLabel = "unisex"
			}

			//  Артикул
			prod.Article = pd.Result.ProductCode

			// Категория
			CategsStrs := strings.Split(pd.Result.Category.Hierarchy, "/")
			prod.Cat = make([]bases.Cat, 0, len(CategsStrs)+1)
			prod.Cat = append(prod.Cat, bases.Cat{
				Name: "trendyol",
				Slug: bases.Name2Slug("trendyol"),
			})
			for _, categ := range CategsStrs {
				prod.Cat = append(prod.Cat, bases.Cat{
					Name: categ,
					Slug: bases.Name2Slug(categ),
				})
			}

		}

		// Фотографии вариации товаров
		images := make([]string, 0, len(pd.Result.Images))
		for _, img := range pd.Result.Images {
			images = append(images, "https://cdn.dsmcdn.com/mnresize/1200/1800"+img)
		}

		// Вариации товаров
		color := Touch2ColorItem(pd)
		prod.Item = append(prod.Item, bases.ColorItem{
			ColorEng:  DonProd.Name,
			ColorCode: DonProd.BeautifiedName,
			Size:      color.Size,
			Price:     pd.Result.Price.OriginalPrice.Value,
			// Link: URL + DonProd.Contents[0].URL,
			Link:  fmt.Sprintf(product_URL, DonProd.Contents[0].ID),
			Image: images,
		})
	}

	return prod, Err
}

// Парсинг с ProductGroupID для обычного варианта
func productOwner(ProductID int) (prod bases.Product2, Err error) {

	// Парсинг вариации
	pd, ErrProd := ParseProduct(ProductID)
	if ErrProd != nil {
		return prod, fmt.Errorf("ParseProduct: %v", ErrProd)
	}

	prod.Name = pd.Result.Name // Название товара

	// Описание товара
	for _, ds := range pd.Result.Description {
		prod.Description.Eng = prod.Description.Eng + ds.Text + "\n"
	}
	re := regexp.MustCompile("[[^]]*]")
	prod.Description.Eng = re.ReplaceAllString(prod.Description.Eng, "")

	// Гендер товара
	// fmt.Println("pd.Result.Gender.ID ", pd.Result.Gender.ID, pd.Result.Gender.Name)
	switch pd.Result.Gender.ID {
	case 1:
		prod.GenderLabel = "women"
	case 2:
		prod.GenderLabel = "man"
	case 3:
		prod.GenderLabel = "unisex"
	}

	//  Артикул
	prod.Article = pd.Result.ProductCode

	// Категория
	CategsStrs := strings.Split(pd.Result.Category.Hierarchy, "/")
	prod.Cat = make([]bases.Cat, 0, len(CategsStrs)+1)
	prod.Cat = append(prod.Cat, bases.Cat{
		Name: "trendyol",
		Slug: bases.Name2Slug("trendyol"),
	})
	for _, categ := range CategsStrs {
		prod.Cat = append(prod.Cat, bases.Cat{
			Name: categ,
			Slug: bases.Name2Slug(categ),
		})
	}

	// Фотографии вариации товаров
	images := make([]string, 0, len(pd.Result.Images))
	for _, img := range pd.Result.Images {
		images = append(images, "https://cdn.dsmcdn.com/mnresize/1200/1800"+img)
	}

	// Вариации товаров
	color := Touch2ColorItem(pd)
	prod.Item = append(prod.Item, bases.ColorItem{
		ColorEng:  pd.Result.Color,
		ColorCode: bases.Name2Slug(pd.Result.Color),
		Size:      color.Size,
		Price:     pd.Result.Price.OriginalPrice.Value,
		// Link: URL + DonProd.Contents[0].URL,
		Link:  fmt.Sprintf(product_URL, ProductID),
		Image: images,
	})

	return prod, nil
}
