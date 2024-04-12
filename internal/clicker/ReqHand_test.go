package clicker

import (
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/wcprod"
)

// go test -v -run ^TestHands$ github.com/RB-PRO/ClikShop/internal/Clicker
func TestHands(t *testing.T) {
	Adding, errorInitWcAdd := wcprod.New2() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		t.Error(errorInitWcAdd)
	}

	hands, ErrorHand := Hands(Adding)
	if ErrorHand != nil {
		t.Error(ErrorHand)
	}

	if len(hands) != 0 {
		t.Error("Zero hands")
	}

}
