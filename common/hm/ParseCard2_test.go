package hm_test

import (
	"ClikShop/common/config"
	"fmt"
	"testing"

	"ClikShop/common/bases"
	"ClikShop/common/hm"
)

func TestVariableProduct2(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	hmService := hm.New(cfg)

	Line, ErrorLine := hmService.Lines("/tr_tr/kadin/yeni-urunler/giysi/_jcr_content/main/productlisting.display.json", 20)
	if ErrorLine != nil {
		t.Error(ErrorLine)
	}

	Prods := hm.Line2Product2(Line, []bases.Cat{}, "woman")
	LinkROI := 12
	fmt.Println("Link:", Prods[LinkROI].Link)
	fmt.Println("len:", len(Prods))

	NewProds, ErrorParseProduct := hmService.VariableProduct2(Prods[LinkROI])
	if ErrorParseProduct != nil {
		t.Error(ErrorParseProduct)
	}
	NewProds, ErrAvailabilityProduct := hmService.AvailabilityProduct(NewProds)
	if ErrAvailabilityProduct != nil {
		t.Error(ErrAvailabilityProduct)
	}

	fmt.Println(bases.ProdStr(NewProds))
}

func TestVariableDescription2(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	hmService := hm.New(cfg)

	line, err := hmService.Lines("/tr_tr/kadin/urune-gore-satin-al/elbise/_jcr_content/main/productlisting.display.json", 1)
	if err != nil {
		t.Error(err)
	}

	Prods := hm.Line2Product2(line, []bases.Cat{}, "woman")
	LinkROI := 0
	fmt.Println("Link:", Prods[LinkROI].Link)
	newProds, err := hmService.VariableDescription2(Prods[LinkROI])
	if err != nil {
		t.Error(err)
	}

	fmt.Println(newProds.Description.Eng)
	fmt.Println(newProds.Specifications)
}
