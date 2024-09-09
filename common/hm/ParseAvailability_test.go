package hm_test

import (
	"ClikShop/common/config"
	"fmt"
	"github.com/pkg/errors"
	"testing"

	"ClikShop/common/bases"
	"ClikShop/common/hm"
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
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	hmService := hm.New(cfg)

	// https://www2.hm.com/tr_tr/kadin/urune-gore-satin-al/elbise/_jcr_content/main/productlisting.display.json?page-size=1
	line, ErrorLine := hmService.Lines("/tr_tr/kadin/urune-gore-satin-al/elbise/_jcr_content/main/productlisting.display.json", 1)
	if ErrorLine != nil {
		t.Error(ErrorLine)
		return
	}

	if len(line.Products) == 0 {
		t.Error(errors.New("line is empty"))
	}

	prods := hm.Line2Product2(line, []bases.Cat{}, "woman")
	LinkROI := 0
	// fmt.Println("Link:", Prods[LinkROI].Link)

	prods[LinkROI], err = hmService.VariableProduct2(prods[LinkROI])
	if err != nil {
		t.Error(err)
		return
	}

	prods[LinkROI], err = hmService.AvailabilityProduct(prods[LinkROI])
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(bases.ProdStr(prods[LinkROI]))
}

func TestAavailabilityMap(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	hmService := hm.New(cfg)

	url := "https://www2.hm.com/tr_tr/productpage.1183407001.html"
	mapAv, err := hmService.AavailabilityMap(url)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(mapAv)
}

func TestAvalimity(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	hmService := hm.New(cfg)

	var prod bases.Product2
	prod.Link = "https://www2.hm.com/tr_tr/productpage.1170211001.html"
	prod.Article = "1170211001"
	prod.Item = make([]bases.ColorItem, 1)
	prod.Item[0].Link = "https://www2.hm.com/tr_tr/productpage.1170211001.html"

	prod, err = hmService.VariableProduct2(prod)
	if err != nil {
		t.Error("Parsing: VariableProduct2:", err)
	}

	prod, err = hmService.AvailabilityProduct(prod)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(bases.ProdStr(prod))
}
