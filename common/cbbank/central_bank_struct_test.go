package cbbank_test

import (
	"fmt"
	"testing"

	"ClikShop/common/cbbank"
)

func TestUSD(t *testing.T) {
	cb := cbbank.New()
	lira, err := cb.Lira()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Курс лиры", lira)
}
