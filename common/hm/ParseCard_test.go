package hm_test

import (
	"ClikShop/common/config"
	"fmt"
	"testing"

	"ClikShop/common/hm"
)

// TODO: Delete
// It is cut out, like everything connected with playwright
func TestVariableProduct3(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	hmService := hm.New(cfg)

	Line, ErrorLine := hmService.Lines("/tr_tr/kadin/urune-gore-satin-al/elbise/_jcr_content/main/productlisting.display.json", 1)
	if ErrorLine != nil {
		t.Error(ErrorLine)
	}
	_ = Line

	//Prods := hm.Line2Product2(Line, []bases.Cat{}, "woman")
	//core, _ := hm.NewParsingCard()
	//LinkROI := 0
	//fmt.Println("Link:", Prods[LinkROI].Link)
	//for CoutColor := range Prods[LinkROI].Item {
	//	ErrorParseProduct := core.VariableProduct3(&Prods[LinkROI], CoutColor)
	//	if ErrorParseProduct != nil {
	//		t.Error(ErrorParseProduct)
	//	}
	//}
	//
	//fmt.Println(bases.ProdStr(Prods[LinkROI]))
}

func TestStrFromSKU(t *testing.T) {
	fmt.Println(hm.StrFromSKU("1163274001002"))
}
