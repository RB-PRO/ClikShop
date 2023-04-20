package cbbank_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
)

func TestUSD(t *testing.T) {
	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		t.Error(ErrorCB)
	}
	fmt.Println("Курс доллара", cb.Data.Valute.Usd.Value)
}
