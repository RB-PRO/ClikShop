package zaratr_test

import (
	"ClikShop/common/config"
	"testing"

	zaratr "ClikShop/common/ZaraTR"
)

func TestLoadLine(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	zaraService := zaratr.New(cfg)
	if err != nil {
		t.Error(err)
	}

	lin, ErrorCat := zaraService.LoadLine("2215112")
	if ErrorCat != nil {
		t.Error(ErrorCat)
	}
	if len(lin.ProductGroups) == 0 {
		t.Error("incorrect product group")
	}
}
