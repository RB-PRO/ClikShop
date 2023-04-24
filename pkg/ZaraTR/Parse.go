package zaratr

import (
	"fmt"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/cheggaaa/pb"
)

func Parsing() bases.Variety2 {

	// Категории
	CatArr := CatCycle() // Получить все категории
	fmt.Println("Всего", len(CatArr.Items), "категорий")

	// Все товары
	ProductsLine := make([]CommercialComponents, 0)
	bar := pb.StartNew(len(CatArr.Items))
	for _, cat := range CatArr.Items {
		line, ErrorLine := LoadLine(fmt.Sprintf("%v", cat.ID.value))
		if ErrorLine != nil {
			fmt.Println(ErrorLine)
		}
		bar.Increment()

		if len(line.ProductGroups) != 0 {
			if len(line.ProductGroups) != 0 {
				if len(line.ProductGroups[0].Elements) != 0 {
					for ind := range line.ProductGroups[0].Elements[0].CommercialComponents { // Циклом обновляем категории
						if line.ProductGroups[0].Elements[0].CommercialComponents[ind].Type == "Product" { // Если это сам товар
							line.ProductGroups[0].Elements[0].CommercialComponents[ind].Cat = cat.Cat
						}
					}
					ProductsLine = append(ProductsLine, line.ProductGroups[0].Elements[0].CommercialComponents...)
				}
			}
		}
	}
	bar.Finish()
	fmt.Println("Всего", len(ProductsLine), "товара(ов)")

	// парсим товары
	var Variety bases.Variety2
	bar2 := pb.StartNew(len(ProductsLine))
	for _, prod := range ProductsLine {
		touch, _ := LoadTouch(prod.Seo.Keyword + "-p" + prod.Seo.SeoProductID)
		Prod2 := Touch2Product2(touch)
		Prod2.Cat = prod.Cat // Обновляем категнории

		Variety.Product = append(Variety.Product, Prod2)
		bar2.Increment()
	}
	bar2.Finish()

	return Variety
}
