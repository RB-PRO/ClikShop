package massimodutti

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"ClikShop/common/bases"
)

// Получить сведения по товару, путём выполнения [запроса].
//
// На вход нужно подать id товара, о котором необходимо получить информацию.
//
// [запроса]: https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/27185784/detail?languageId=-1&appId=1
func (s *Service) Toucher(id int) (Touch, error) {
	c, err := s.NewServiceCollector()
	if err != nil {
		return Touch{}, errors.Wrap(err, "create service collector: ")
	}

	url := fmt.Sprintf("https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/%v/detail?languageId=-1&appId=1", id)

	headers := http.Header{}
	headers.Add("authority", "www.massimodutti.com")
	headers.Add("accept", "application/json, text/plain, */*")
	headers.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	headers.Add("referer", "https://www.massimodutti.com/")
	headers.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	headers.Add("sec-ch-ua-mobile", "?0")
	headers.Add("sec-ch-ua-platform", "\"Linux\"")
	headers.Add("sec-fetch-dest", "empty")
	headers.Add("sec-fetch-mode", "cors")
	headers.Add("sec-fetch-site", "same-origin")
	headers.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36")

	var response Touch
	c.OnResponse(func(r *colly.Response) {
		if err := json.Unmarshal(r.Body, &response); err != nil {
			log.Println("ERROR:500:", err)
			return
		}
	})

	return response, c.Request(http.MethodGet, url, nil, nil, headers)
}

// Перевести результаты парсинга в стандартный для нас формат
func Touch2Product2(Product bases.Product2, touch Touch) bases.Product2 {

	if len(touch.BundleProductSummaries) == 0 {
		return bases.Product2{}
	}

	// Название товара
	Product.Name = strings.TrimSpace(touch.Name)

	// Описание товара
	Description := make([]string, 0)
	for _, attr := range touch.Attributes {
		if attr.Type == "DESCRIPTION" {
			Description = append(Description, attr.Value)
		}
	}
	Product.Description.Eng = strings.Join(Description, "\n")

	// Производитель
	Product.Manufacturer = "Massimo Dutti"

	// Пол
	Product.GenderLabel, _, _ = bases.GenderBook(touch.SectionNameEN, bases.Name2Slug(touch.SectionNameEN))

	// Вариации данного товара
	var Item []bases.ColorItem

	// fmt.Println("len(touch.Detail.Colors)", len(touch.Detail.Colors))
	// fmt.Printf("%+v\n\n", touch.Detail)

	IdToUrl := make(map[string]string)
	if len(touch.Detail.Colors) != 0 {
		for _, Colors := range touch.Detail.Colors {
			IdToUrl[Colors.ID] = Colors.Image.URL
		}
	} else {
		for _, Colors := range touch.BundleProductSummaries[0].Detail.Colors {
			IdToUrl[Colors.ID] = Colors.Image.URL
		}
	}

	// Картинки новой версии
	images2 := make(map[string][]string, 0)
	if len(touch.Detail.Colors) != 0 {
		for _, Xmedia := range touch.Detail.Xmedia {
			// fmt.Println("Xmedia.ColorCode", Xmedia.ColorCode)
			for _, XmediaLocations := range Xmedia.XmediaLocations {
				if XmediaLocations.Set == 0 {
					for _, Locations := range XmediaLocations.Locations {

						switch Locations.Location.(type) {
						case int:
							if Locations.Location == 1 {
								for _, imageCode := range Locations.MediaLocations {
									LinkStrs := strings.Split(imageCode, "_")
									if len(LinkStrs) == 4 {
										images2[Xmedia.ColorCode] = append(images2[Xmedia.ColorCode], fmt.Sprintf("https://static.massimodutti.net/3/photos/%s_%s_%s_16.jpg", IdToUrl[Xmedia.ColorCode], LinkStrs[1], LinkStrs[2]))
									}
								}
							}
						case string:
							fmt.Println("Locations.Location", Locations.Location)
						}
					}
				}
			}
		}
	} else {
		for _, Xmedia := range touch.BundleProductSummaries[0].Detail.Xmedia {
			//fmt.Println("Xmedia.ColorCode", Xmedia.ColorCode)
			for _, XmediaLocations := range Xmedia.XmediaLocations {
				if XmediaLocations.Set == 0 {
					for _, Locations := range XmediaLocations.Locations {

						switch Locations.Location.(type) {
						case int:
							if Locations.Location == 1 {
								for _, imageCode := range Locations.MediaLocations {
									LinkStrs := strings.Split(imageCode, "_")
									if len(LinkStrs) == 4 {
										images2[Xmedia.ColorCode] = append(images2[Xmedia.ColorCode], fmt.Sprintf("https://static.massimodutti.net/3/photos/%s_%s_%s_16.jpg", IdToUrl[Xmedia.ColorCode], LinkStrs[1], LinkStrs[2]))
									}
								}
							}
						case string:
							fmt.Println("Locations.Location", Locations.Location)
						}

						// if Locations.Location == 1 {
						// 	for _, imageCode := range Locations.MediaLocations {
						// 		LinkStrs := strings.Split(imageCode, "_")
						// 		if len(LinkStrs) == 4 {
						// 			images2[Xmedia.ColorCode] = append(images2[Xmedia.ColorCode], fmt.Sprintf("https://static.massimodutti.net/3/photos/%s_%s_%s_16.jpg", IdToUrl[Xmedia.ColorCode], LinkStrs[1], LinkStrs[2]))
						// 		}
						// 	}
						// }
					}
				}
			}
		}
	}

	// for iImg, Img := range images2 {
	// 	fmt.Println("-", iImg, "--", Img)
	// }

	if len(touch.Detail.Colors) != 0 {
		for _, colorSet := range touch.Detail.Colors {
			// fmt.Println("colorSet.ID", colorSet.ID)
			// Ищем и вставляем цену
			Price := 0.0
			if len(colorSet.Sizes) > 0 {
				if n, err := strconv.ParseFloat(colorSet.Sizes[0].Price, 64); err == nil {
					Price = n / 100
				}
			}

			// Формирование слайса размеров
			var sizes []bases.Size
			for _, ValSize := range colorSet.Sizes {
				AddSize := bases.Size{Val: ValSize.Name, IsExit: ValSize.IsBuyable}
				if ValSize.VisibilityValue != "SHOW" {
					AddSize.IsExit = false
				}
				sizes = append(sizes, AddSize)
			}

			// Картинки
			images := make([]string, 0)
			ColorCode := colorSet.ID // Получаем ID цвета
			for _, XmediaColor := range touch.Detail.Xmedia {
				if XmediaColor.ColorCode == ColorCode { // Если это тот самый необходимый цвет
					if len(XmediaColor.XmediaItems) > 0 {
						for _, MediasColor := range XmediaColor.XmediaItems[0].Medias {
							images = append(images, fmt.Sprintf("https://static.massimodutti.net/3/photos/%s/%s16.jpg", XmediaColor.Path, MediasColor.IDMedia))
							// fmt.Println("IMAGES", fmt.Sprintf("https://static.massimodutti.net/3/photos/%s/%s16.jpg", XmediaColor.Path, MediasColor.IDMedia))
						}
					}
				}
			}
			images = bases.RemoveDuplicateStr(images)

			// Формируем структур цветов и размеров
			Item = append(Item, bases.ColorItem{
				ColorEng:  colorSet.Name,
				ColorCode: bases.Name2Slug(colorSet.Name),
				Link:      fmt.Sprintf("https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/%v/detail?languageId=-1&appId=1", Product.Article),
				Price:     Price,
				Size:      sizes,
				Image:     images,
				// Image:     images2[colorSet.ID],
			})
		}
		Product.Article = touch.Detail.DisplayReference // Артикул(id)
	} else {
		for _, colorSet := range touch.BundleProductSummaries[0].Detail.Colors {
			// fmt.Println("colorSet.ID", colorSet.ID)
			// Ищем и вставляем цену
			Price := 0.0
			if len(colorSet.Sizes) > 0 {
				if n, err := strconv.ParseFloat(colorSet.Sizes[0].Price, 64); err == nil {
					Price = n / 100
				}
			}

			// Формирование слайса размеров
			var sizes []bases.Size
			for _, ValSize := range colorSet.Sizes {
				AddSize := bases.Size{Val: ValSize.Name, IsExit: ValSize.IsBuyable}
				if ValSize.VisibilityValue != "SHOW" {
					AddSize.IsExit = false
				}
				sizes = append(sizes, AddSize)
			}

			// Картинки
			images := make([]string, 0)
			ColorCode := colorSet.ID // Получаем ID цвета
			for _, XmediaColor := range touch.BundleProductSummaries[0].Detail.Xmedia {
				if XmediaColor.ColorCode == ColorCode { // Если это тот самый необходимый цвет
					if len(XmediaColor.XmediaItems) > 0 {
						for _, MediasColor := range XmediaColor.XmediaItems[0].Medias {
							images = append(images, fmt.Sprintf("https://static.massimodutti.net/3/photos/%s/%s16.jpg", XmediaColor.Path, MediasColor.IDMedia))
						}
					}
				}
			}
			images = bases.RemoveDuplicateStr(images)

			// Формируем структур цветов и размеров
			Item = append(Item, bases.ColorItem{
				ColorEng:  colorSet.Name,
				ColorCode: bases.Name2Slug(colorSet.Name),
				Link:      fmt.Sprintf("https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/%v/detail?languageId=-1&appId=1", Product.Article),
				Price:     Price,
				Size:      sizes,
				Image:     images,
				// Image:     images2[colorSet.ID],
			})
		}
		Product.Article = touch.BundleProductSummaries[0].Detail.DisplayReference // Артикул(id)
	}
	//fmt.Println("for _, Xmedia := range touch.Detail.Xmedia {")
	// TEST

	Product.Item = Item
	Product.Link = fmt.Sprintf("https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/%v/detail?languageId=-1&appId=1", touch.ID)

	return Product
}
