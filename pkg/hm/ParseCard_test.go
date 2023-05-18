package hm_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/hm"
)

func TestCard(t *testing.T) {
	Prod, ErrorParseProduct := hm.Product("https://www2.hm.com/tr_tr/productpage.1156720001.html")
	if ErrorParseProduct != nil {
		t.Error(ErrorParseProduct)
	}

	fmt.Println(bases.ProdStr(Prod))
}
