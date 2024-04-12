package zaratr_test

import (
	"fmt"
	"testing"

	zaratr "github.com/RB-PRO/ClikShop/pkg/ZaraTR"
)

func TestLoadCategory(t *testing.T) {
	cat, ErrorCat := zaratr.LoadCategory()
	if ErrorCat != nil {
		t.Error(ErrorCat)
	}
	if cat.Categories[0].Name != "WOMAN" {
		t.Error("Неправльный ответ")
	}
	fmt.Println("Всего категорий", len(cat.Categories))
}

func TestCatCycle(t *testing.T) {
	cycCat := zaratr.CatCycle()
	// fmt.Println(cycCat)
	for _, cat := range cycCat.Items {
		fmt.Println(cat.Name)
	}
	fmt.Println(len(cycCat.Items))
}
func TestCatCycle2(t *testing.T) {
	cycCat, err := zaratr.CatCycle2()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(cycCat.Items))
	for i, v := range cycCat.Items {
		fmt.Println(i, v.Gender, v.RedirectCategoryID, v.Cat)
	}
}
