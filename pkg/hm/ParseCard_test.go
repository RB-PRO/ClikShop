package hm_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/hm"
)

func TestVariableProduct3(t *testing.T) {
	Line, ErrorLine := hm.Lines("/tr_tr/kadin/urune-gore-satin-al/elbise/_jcr_content/main/productlisting.display.json", 1)
	if ErrorLine != nil {
		t.Error(ErrorLine)
	}

	Prods := hm.Line2Product2(Line, []bases.Cat{}, "woman")
	core, _ := hm.NewParsingCard()
	LinkROI := 0
	fmt.Println("Link:", Prods[LinkROI].Link)
	for CoutColor := range Prods[LinkROI].Item {
		ErrorParseProduct := core.VariableProduct3(&Prods[LinkROI], CoutColor)
		if ErrorParseProduct != nil {
			t.Error(ErrorParseProduct)
		}
	}

	fmt.Println(bases.ProdStr(Prods[LinkROI]))
}

func TestStrFromSKU(t *testing.T) {
	fmt.Println(hm.StrFromSKU("1163274001002"))
}
