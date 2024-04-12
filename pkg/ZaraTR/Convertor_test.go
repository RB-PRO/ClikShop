package zaratr_test

import (
	"fmt"
	"testing"

	zaratr "github.com/RB-PRO/ClikShop/pkg/ZaraTR"
)

func TestTouch2Product2(t *testing.T) {
	tou, TouError := zaratr.LoadTouch("ribbed-strappy-vest-top-p03253306")
	if TouError != nil {
		t.Error(TouError)
	}
	prod := zaratr.Touch2Product2(tou)

	fmt.Printf("%+v", prod)
}
