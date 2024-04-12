package massimodutti

import (
	"fmt"
	"testing"
)

func TestCategory(t *testing.T) {
	categs, ErrCateg := Category()
	if ErrCateg != nil {
		t.Error(ErrCateg)
	}
	for i, i_val := range categs.Categories {
		for j, j_val := range i_val.Subcategories {
			// Проверка на тип категории.
			//Заметили, что именно с этим типом категория считается валидной
			if j_val.Type == "22" {
				fmt.Printf("%d - %s, %d - %s\n", i, i_val.Name, j, j_val.Name)
			}
		}
	}
}
