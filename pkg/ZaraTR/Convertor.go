package zaratr

import (
	"strings"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

const URL string = "https://www.zara.com"

func Touch2Product2(tou Touch) bases.Product2 {
	var Prod bases.Product2
	Prod.Item = make(map[string]bases.ProdParam)

	// Артикул
	Prod.Article = tou.Product.Detail.DisplayReference

	// Название товара
	Prod.Name = tou.Product.Name

	// Ссылка на товар
	Prod.Link = URL + tou.ClientAppConfig.OriginalURL

	// Гендер
	if len(tou.BreadCrumbs) != 0 {
		Prod.GenderLabel = tou.BreadCrumbs[0].Keyword
	} else {
		Prod.GenderLabel = "unisex"
	}

	// Производитель
	Prod.Manufacturer = tou.Product.Brand.BrandGroupCode

	// Краткое описание
	Prod.FullName = tou.DocInfo.Description

	// Краткое описание
	Prod.Description.Eng = tou.DocInfo.Description

	// Цикл по цветам
	for _, color := range tou.Product.Detail.Colors {

		// Создаём массив размеров
		sizes := make([]string, 0)
		for _, size := range color.Sizes {
			sizes = append(sizes, size.Name)
			Prod.Size = append(Prod.Size, size.Name) // Все размеры товара
		}

		// Создаём массив размеров
		images := make([]string, 0)
		for _, media := range color.Xmedia {
			ImageLink := `https://static.zara.net/photos//` + media.Path + `/w/916/` + media.Name + `.jpg?ts=` + media.Timestamp
			images = append(images, ImageLink)
		}

		Prod.Item[strings.ToLower(color.Name)] = bases.ProdParam{
			ColorEng: color.Name,
			Size:     sizes,
			Image:    images,
			Price:    float64(color.Price) / 100,
		}
	}

	Prod.Size = bases.RemoveDuplicateStr(Prod.Size)

	return Prod
}
