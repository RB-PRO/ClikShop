package hm_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/hm"
)

// func TestAvailabilityProduct(t *testing.T) {
// 	av1, err1 := hm.AvailabilityProduct("1157823001")
// 	if err1 != nil {
// 		t.Error(err1)
// 	}
// 	fmt.Println(av1)
// 	av2, err2 := hm.AvailabilityProduct("1157823")
// 	if err2 != nil {
// 		t.Error(err2)
// 	}
// 	fmt.Println(av2)
// }

func TestAvailabilityProduct(t *testing.T) {
	Line, ErrorLine := hm.Lines("/tr_tr/kadin/urune-gore-satin-al/elbise/_jcr_content/main/productlisting.display.json", 1)
	if ErrorLine != nil {
		t.Error(ErrorLine)
	}

	Prods := hm.Line2Product2(Line, []bases.Cat{}, "woman")
	LinkROI := 0
	// fmt.Println("Link:", Prods[LinkROI].Link)

	var ErrorParseProduct error
	Prods[LinkROI], ErrorParseProduct = hm.VariableProduct2(Prods[LinkROI])
	if ErrorParseProduct != nil {
		t.Error(ErrorParseProduct)
	}

	var ErrAvailabilityProduct error
	Prods[LinkROI], ErrAvailabilityProduct = hm.AvailabilityProduct(Prods[LinkROI])
	if ErrAvailabilityProduct != nil {
		t.Error(ErrAvailabilityProduct)
	}

	fmt.Println(bases.ProdStr(Prods[LinkROI]))
}

func TestAavailabilityMap(t *testing.T) {
	url := "https://www2.hm.com/tr_tr/productpage.1183407001.html"
	mapAv, ErrAvailabilityProduct := hm.AavailabilityMap(url)
	if ErrAvailabilityProduct != nil {
		t.Error(ErrAvailabilityProduct)
	}
	fmt.Println(mapAv)
}

func TestAvalimity(t *testing.T) {
	var prod bases.Product2
	prod.Link = "https://www2.hm.com/tr_tr/productpage.1170211001.html"
	prod.Article = "1170211001"
	prod.Item = make([]bases.ColorItem, 1)
	prod.Item[0].Link = "https://www2.hm.com/tr_tr/productpage.1170211001.html"

	prod, ErrorParseProduct := hm.VariableProduct2(prod)
	if ErrorParseProduct != nil {
		t.Error("Parsing: VariableProduct2:", ErrorParseProduct)
	}

	prod, ErrAvailabilityProduct := hm.AvailabilityProduct(prod)
	if ErrAvailabilityProduct != nil {
		t.Error(ErrAvailabilityProduct)
	}

	fmt.Println(bases.ProdStr(prod))
}
