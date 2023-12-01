package zaratr_test

import (
	"fmt"
	"testing"

	zaratr "github.com/RB-PRO/ClikShop/pkg/ZaraTR"
)

func TestLoadTouch(t *testing.T) {
	// https://www.zara.com/tr/en/linen-blend-longline-bomber-jacket-p03574371.html?ajax=true
	// ribbed-strappy-vest-top-p03253306
	tou, ErrorCat := zaratr.LoadTouch("metallic-block-heel-sandals-p13344110")
	if ErrorCat != nil {
		t.Error(ErrorCat)
	}
	if tou.Product.Name == "" {
		t.Error("Неправльный ответ")
	}
	TouProduct := zaratr.Touch2Product2(tou)
	fmt.Printf("%+v\n\n%v\n\n", TouProduct, TouProduct.Item[0].Price)
	fmt.Println(TouProduct.Item[0].Image)
}
