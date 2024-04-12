package wcprod_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/wcprod"
)

func TestXlsxDelivery(t *testing.T) {
	// СОбираем данные из файла
	Delivery, errorFile := wcprod.XlsxDelivery()
	if errorFile != nil {
		t.Error(errorFile)
	}
	if len(Delivery) == 0 {
		t.Error("Длина Мапы равна нулю")
	}

	// Вывод
	for key, value := range Delivery {
		fmt.Println(key, "-", value)
	}
}
