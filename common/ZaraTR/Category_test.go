package zaratr_test

import (
	"ClikShop/common/config"
	"fmt"
	"testing"

	zaratr "ClikShop/common/ZaraTR"
)

func TestLoadCategory(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	zaraService := zaratr.New(cfg)
	if err != nil {
		t.Error(err)
	}

	cat, ErrorCat := zaraService.LoadCategory()
	if ErrorCat != nil {
		t.Error(ErrorCat)
	}
	if cat.Categories[0].Name != "WOMAN" {
		t.Error("Неправльный ответ")
	}
	fmt.Println("Всего категорий", len(cat.Categories))
}

func TestCatCycle(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	zaraService := zaratr.New(cfg)
	if err != nil {
		t.Error(err)
	}

	cycCat := zaraService.CatCycle()
	// fmt.Println(cycCat)
	for _, cat := range cycCat.Items {
		fmt.Println(cat.Name)
	}
	fmt.Println(len(cycCat.Items))
}
func TestCatCycle2(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	zaraService := zaratr.New(cfg)
	if err != nil {
		t.Error(err)
	}

	cycCat, err := zaraService.CatCycle2()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(cycCat.Items))
	for i, v := range cycCat.Items {
		fmt.Println(i, v.Gender, v.RedirectCategoryID, v.Cat)
	}
}
