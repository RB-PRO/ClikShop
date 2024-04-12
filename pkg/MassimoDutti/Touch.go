package massimodutti

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
)

// Получить сведения по товару, путём выполнения [запроса].
//
// # На вход нужно подать id товара, о котором необходимо получить информацию.
//
// [запроса]: https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/27185784/detail?languageId=-1&appId=1
func Toucher(id int) (touch Touch, ErrCategory error) {

	url := fmt.Sprintf("https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/%v/detail?languageId=-1&appId=1", id)

	// fmt.Println("url:", url)

	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return touch, ErrNewRequest
	}
	req.Header.Add("authority", "www.massimodutti.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("referer", "https://www.massimodutti.com/")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36")

	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return touch, ErrDo
	}
	defer res.Body.Close()

	// if res.StatusCode == http.StatusOK {
	// 	ErrNewDecoder := json.NewDecoder(res.Body).Decode(&touch)
	// 	if ErrNewDecoder != nil {
	// 		return touch, ErrNewDecoder
	// 	}
	// } else {
	// 	return touch, errors.New("Toucher: http.Status is not ok")
	// }
	ErrNewDecoder := json.NewDecoder(res.Body).Decode(&touch)
	if ErrNewDecoder != nil {
		return touch, ErrNewDecoder
	}

	return touch, nil
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
