package hm_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/hm"
)

func TestVariableProduct2(t *testing.T) {
	Line, ErrorLine := hm.Lines("/tr_tr/kadin/urune-gore-satin-al/elbise/_jcr_content/main/productlisting.display.json", 1)
	if ErrorLine != nil {
		t.Error(ErrorLine)
	}

	Prods := hm.Line2Product2(Line, []bases.Cat{}, "woman")
	LinkROI := 0
	// fmt.Println("Link:", Prods[LinkROI].Link)
	NewProds, ErrorParseProduct := hm.VariableProduct2(Prods[LinkROI])
	if ErrorParseProduct != nil {
		t.Error(ErrorParseProduct)
	}

	fmt.Println(bases.ProdStr(NewProds))
}
