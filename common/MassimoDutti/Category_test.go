package massimodutti_test

import (
	massimodutti "ClikShop/common/MassimoDutti"
	"ClikShop/common/config"
	"fmt"
	"testing"
)

func TestCategory(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	mdService, err := massimodutti.New(cfg)
	if err != nil {
		t.Error(err)
	}

	categs, ErrCateg := mdService.Category()
	if ErrCateg != nil {
		t.Error(ErrCateg)
	}
	for i, category := range categs.Categories {
		for j, subcategory := range category.Subcategories {
			// Проверка на тип категории.
			// Заметили, что именно с этим типом категория считается валидной
			if subcategory.Type == "22" {
				fmt.Printf("%d - %s, %d - %s\n", i, category.Name, j, subcategory.Name)
			}
		}
	}
}
