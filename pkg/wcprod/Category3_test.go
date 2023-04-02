package wcprod_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/wcprod"
)

func TestFormMapCat3(t *testing.T) {
	// Создаём экземпляр загрузчика данных
	Adding, errorInitWcAdd := wcprod.New()
	if errorInitWcAdd != nil {
		t.Error(errorInitWcAdd)
	}

	// Создать категорию товаров
	ErrorFormCat := Adding.FormMapCat3()
	if ErrorFormCat != nil {
		t.Error(ErrorFormCat)
	}

	Adding.PrintCat3()

	fmt.Printf("%+v", Adding.Cat3)

}
