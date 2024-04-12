package wcprod_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/wcprod"
)

func TestTranslate(t *testing.T) {

	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		t.Error(errorInitWcAdd)
	}
	Variety2 := varietBasesVariety2()

	for _, prod := range Variety2.Product {
		TranslateProd, ErrorTransalte := Adding.YandexTranslate(prod)
		if ErrorTransalte != nil {
			t.Error(ErrorTransalte)
		}
		fmt.Printf("%+v", TranslateProd)
	}
}
