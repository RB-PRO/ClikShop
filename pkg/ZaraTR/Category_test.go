package zaratr_test

import (
	"fmt"
	"testing"

	zaratr "github.com/RB-PRO/SanctionedClothing/pkg/ZaraTR"
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
