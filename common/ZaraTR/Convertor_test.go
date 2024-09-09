package zaratr_test

import (
	"ClikShop/common/config"
	"fmt"
	"testing"

	zaratr "ClikShop/common/ZaraTR"
)

func TestTouch2Product2(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	zaraService := zaratr.New(cfg)
	if err != nil {
		t.Error(err)
	}

	tou, TouError := zaraService.LoadTouch("ribbed-strappy-vest-top-p03253306")
	if TouError != nil {
		t.Error(TouError)
	}
	prod := zaratr.Touch2Product2(tou)

	fmt.Printf("%+v", prod)
}
