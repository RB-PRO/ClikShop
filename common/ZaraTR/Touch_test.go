package zaratr_test

import (
	"ClikShop/common/config"
	"fmt"
	"testing"

	zaratr "ClikShop/common/ZaraTR"
)

func TestLoadTouch(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	zaraService := zaratr.New(cfg)
	if err != nil {
		t.Error(err)
	}

	// https://www.zara.com/tr/en/linen-blend-longline-bomber-jacket-p03574371.html?ajax=true
	// ribbed-strappy-vest-top-p03253306
	tou, err := zaraService.LoadTouch("metallic-block-heel-sandals-p13344110")
	if err != nil {
		t.Error(err)
	}
	if tou.Product.Name == "" {
		t.Error("incorrect product name")
	}
	TouProduct := zaratr.Touch2Product2(tou)
	fmt.Printf("%+v\n\n%v\n\n", TouProduct, TouProduct.Item[0].Price)
	fmt.Println(TouProduct.Item[0].Image)
}
