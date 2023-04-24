package zaratr

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

func TestCatCycle(t *testing.T) {

	CatArr := CatCycle() // Наполнить цикл
	fmt.Println(len(CatArr.Items))

}
func TestParsing(t *testing.T) {
	// go test -timeout 12000s -run ^TestParsing$ github.com/RB-PRO/SanctionedClothing/pkg/ZaraTR
	Parsing()
}
func TestTouch2Product2(t *testing.T) {

	touch, _ := LoadTouch("jggr-pnt-10-p02837201")
	Product := Touch2Product2(touch)

	fmt.Printf("%+v", Product)
}

// Комплексный тест парсинга
func TestComplexParse(t *testing.T) {

	// Категории
	CatArr := CatCycle() // Получить все категории
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
	line, ErrorLine := LoadLine(fmt.Sprintf("%v", cat.ID.value))
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
	touch, _ := LoadTouch(prod.Seo.Keyword + "-p" + prod.Seo.SeoProductID)
	Prod2 := Touch2Product2(touch)
	Prod2.Cat = prod.Cat // Обновляем категнории

	fmt.Printf("%+v", Prod2)

	Variety.Product = append(Variety.Product, Prod2)
	Variety.SaveXlsx("Zara")
}
