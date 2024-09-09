package zaratr

import (
	"ClikShop/common/bases"
)

const URL string = "https://www.zara.com"

func Touch2Product2(tou Touch) bases.Product2 {
	var Prod bases.Product2
	Prod.Item = make([]bases.ColorItem, 0)

	// Артикул
	Prod.Article = tou.Product.Detail.DisplayReference

	// Название товара
	Prod.Name = tou.Product.Name

	// Ссылка на товар
	Prod.Link = URL + tou.ClientAppConfig.OriginalURL

	// Гендер
	if len(tou.BreadCrumbs) != 0 {
		_, Prod.GenderLabel, _ = bases.GenderBook(tou.BreadCrumbs[0].Keyword, "")
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
		// sizes := make([]string, 0)
		for _, size := range color.Sizes {
			// sizes = append(sizes, size.Name)
			Prod.Size = append(Prod.Size, size.Name) // Все размеры товара
		}

		// Массив размеров, который будет использован для
		Sizes := []bases.Size{}
		for _, SizeItem := range color.Sizes {

			// Проверяем наличие данного размера товара
			var IsLive bool
			if SizeItem.Availability == "in_stock" || SizeItem.Availability == "low_on_stock" {
				IsLive = true
			}

			AddSize := bases.Size{
				Val:    SizeItem.Name,
				IsExit: IsLive,
			}
			Sizes = append(Sizes, AddSize)
		}

		// Создаём массив размеров
		images := make([]string, 0)
		for _, media := range color.Xmedia {
			ImageLink := `https://static.zara.net/photos//` + media.Path + `/w/916/` + media.Name + `.jpg?ts=` + media.Timestamp
			images = append(images, ImageLink)
		}

		// fmt.Println("Coast :=", float64(color.Price)/100)
		// Prod.Item[strings.ToLower(color.Name)] = bases.ProdParam{}
		Prod.Item = append(Prod.Item, bases.ColorItem{
			ColorEng:  color.Name,
			ColorCode: bases.Name2Slug(color.Name),
			Size:      Sizes, // Тут нужно добавить истинные размеры.
			Image:     images,
			Price:     (float64(color.Price) / 100),
		})
	}

	Prod.Size = bases.RemoveDuplicateStr(Prod.Size)

	return Prod
}
