package actualizer

import (
	"fmt"
	"log"

	zaratr "github.com/RB-PRO/ClikShop/pkg/ZaraTR"
	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/cheggaaa/pb"
)

// Структура HM для парсинга
type ZARA struct {
	*bitrixActualizer
}

func NewZARA(bx *bitrixActualizer) *ZARA {
	return &ZARA{bx}
}

// Парсинг данных и сохранение их в файлы
//
//	Заменить во всех файлах нужно символы '\u0026' на '&'
func (bx *ZARA) screper() (string, error) {
	folder := "zara"
	ReMakeDir(folder)

	// Категории
	CatArr, _ := zaratr.CatCycle2() // Получить все категории
	log.Println("Всего", len(CatArr.Items), "категорий")

	// Все товары
	allID := make(map[string]bool)
	var index int
	for i, cat := range CatArr.Items {
		// if i < 129 {
		// 	continue
		// }
		line, ErrorLine := zaratr.LoadLine(fmt.Sprintf("%v", cat.RedirectCategoryID))
		if ErrorLine != nil {
			fmt.Println(ErrorLine)
			bx.GLOG.Err(fmt.Sprintf("Парсера ZARA: i=%d, Неудачная загрузка по ссылке: https://www.zara.com/tr/en/category/%d/products Ошибка: %v",
				i, cat.RedirectCategoryID, ErrorLine))
		}
		// bar.Increment()

		ProductsLine := make([]zaratr.CommercialComponents, 0)
		for _, ProductGroups := range line.ProductGroups {
			for _, Elements := range ProductGroups.Elements {
				for _, CommercialComponents := range Elements.CommercialComponents {
					// if cout >= 10 { // Максимум 10 товаров в категории
					// 	break
					// }
					CommercialComponents.Cat = cat.Cat
					CommercialComponents.Gender = cat.Gender
					ProductsLine = append(ProductsLine, CommercialComponents)
				}
			}
		}

		ProductsLine = RemoveDuplicateProductsLine(ProductsLine)

		// Парсим товары
		var Variety bases.Variety2

		if len(ProductsLine) == 0 {
			bx.GLOG.Warn(fmt.Sprintf("Парсера ZARA: i=%d, длина Line = 0", i))
			continue
		}

		// Переведённая категория
		FileName := ProductsLine[0].Cat[len(ProductsLine[0].Cat)-1].Slug
		ProdTranslateCat := ProductsLine[0].Cat

		barProduct := pb.StartNew(len(ProductsLine))
		barProduct.Prefix(fmt.Sprintf("[%d/%d]", i+1, len(CatArr.Items)))
		for _, prod := range ProductsLine {
			barProduct.Increment()
			if _, valueok := allID[prod.Reference]; !valueok {
				allID[prod.Reference] = true
			} else {
				bx.GLOG.Warn("Парсера ZARA: Дубль:", fmt.Sprintf(zaratr.TouchURL, prod.Seo.Keyword+"-p"+prod.Seo.SeoProductID))
				continue
			}

			bx.GLOG.Info("Парсера ZARA: LoadTouch:", fmt.Sprintf(zaratr.TouchURL, prod.Seo.Keyword+"-p"+prod.Seo.SeoProductID))

			if prod.Name == "LOOK" || prod.Name == "" {
				continue
			}
			touch, _ := zaratr.LoadTouch(prod.Seo.Keyword + "-p" + prod.Seo.SeoProductID) // Выполняем запрос
			Prod2 := zaratr.Touch2Product2(touch)                                         // АПереводим в структуру Product2
			Prod2.Cat = ProdTranslateCat                                                  //prod.Cat // Обновляем категории
			Prod2.GenderLabel = prod.Gender                                               // Обнволяем гендер

			// Редактируем товар
			Prod2.Size = bases.EditProdSize(Prod2)
			Prod2.Img = bases.EditIMG(Prod2)

			Variety.Product = append(Variety.Product, Prod2)
		}

		Variety.SaveJson(fmt.Sprintf("%s/zara_%d_%v",
			folder, index, FileName))
		barProduct.Finish()
		index++
	}
	return folder, nil
}

// Удалить дубликаты в товарах ProductsLine
func RemoveDuplicateProductsLine(ProductsLine []zaratr.CommercialComponents) []zaratr.CommercialComponents {
	allKeys := make(map[string]bool)
	list := []zaratr.CommercialComponents{}
	for _, item := range ProductsLine {
		if _, value := allKeys[item.ID.String()]; !value {
			allKeys[item.ID.String()] = true
			list = append(list, item)
		}
	}
	return list
}
