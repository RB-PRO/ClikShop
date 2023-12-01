package hm_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/hm"
)

func TestVariableProduct2(t *testing.T) {
	Line, ErrorLine := hm.Lines("/tr_tr/kadin/yeni-urunler/giysi/_jcr_content/main/productlisting.display.json", 20)
	if ErrorLine != nil {
		t.Error(ErrorLine)
	}

	Prods := hm.Line2Product2(Line, []bases.Cat{}, "woman")
	LinkROI := 12
	fmt.Println("Link:", Prods[LinkROI].Link)
	fmt.Println("len:", len(Prods))
	NewProds, ErrorParseProduct := hm.VariableProduct2(Prods[LinkROI])
	if ErrorParseProduct != nil {
		t.Error(ErrorParseProduct)
	}
	NewProds, ErrAvailabilityProduct := hm.AvailabilityProduct(NewProds)
	if ErrAvailabilityProduct != nil {
		t.Error(ErrAvailabilityProduct)
	}

	fmt.Println(bases.ProdStr(NewProds))
}

func TestVariableDescription2(t *testing.T) {
	Line, ErrorLine := hm.Lines("/tr_tr/kadin/urune-gore-satin-al/elbise/_jcr_content/main/productlisting.display.json", 1)
	if ErrorLine != nil {
		t.Error(ErrorLine)
	}

	Prods := hm.Line2Product2(Line, []bases.Cat{}, "woman")
	LinkROI := 0
	fmt.Println("Link:", Prods[LinkROI].Link)
	NewProds, ErrorParseProduct := hm.VariableDescription2(Prods[LinkROI])
	if ErrorParseProduct != nil {
		t.Error(ErrorParseProduct)
	}

	fmt.Println(NewProds.Description.Eng)
	fmt.Println(NewProds.Specifications)
}
