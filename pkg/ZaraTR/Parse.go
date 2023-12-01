package zaratr

import (
	"fmt"
	"log"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/cheggaaa/pb"
)

func Parsing() bases.Variety2 {

	// Категории
	CatArr, _ := CatCycle2() // Получить все категории
	log.Println("Всего", len(CatArr.Items), "категорий")

	// Все товары
	ProductsLine := make([]CommercialComponents, 0)
	bar := pb.StartNew(len(CatArr.Items))
	for _, cat := range CatArr.Items {
		line, ErrorLine := LoadLine(fmt.Sprintf("%v", cat.RedirectCategoryID))
		if ErrorLine != nil {
			fmt.Println(ErrorLine)
		}
		bar.Increment()

		var cout int
		for _, ProductGroups := range line.ProductGroups {
			for _, Elements := range ProductGroups.Elements {
				for _, CommercialComponents := range Elements.CommercialComponents {
					// if cout >= 10 { // Максимум 10 товаров в категории
					// 	break
					// }
					CommercialComponents.Cat = cat.Cat
					CommercialComponents.Gender = cat.Gender
					ProductsLine = append(ProductsLine, CommercialComponents)
					cout++
				}
			}
		}
	}
	bar.Finish()
	log.Println("Всего", len(ProductsLine), "товара(ов)")

	// парсим товары
	var Variety bases.Variety2
	bar2 := pb.StartNew(len(ProductsLine))
	for _, prod := range ProductsLine {
		// log.Printf("(%d/%d) Парсинг товара: %v", i+1, len(ProductsLine), fmt.Sprintf(TouchURL, prod.Seo.Keyword+"-p"+prod.Seo.SeoProductID))
		touch, _ := LoadTouch(prod.Seo.Keyword + "-p" + prod.Seo.SeoProductID) // Выполняем запрос
		Prod2 := Touch2Product2(touch)                                         // АПереводим в структуру Product2
		Prod2.Cat = prod.Cat                                                   // Обновляем категории
		Prod2.GenderLabel = prod.Gender                                        // Обнволяем гендер

		Variety.Product = append(Variety.Product, Prod2)
		bar2.Increment()
	}
	bar2.Finish()

	return Variety
}
