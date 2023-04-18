package zaratr_test

import (
	"testing"

	zaratr "github.com/RB-PRO/SanctionedClothing/pkg/ZaraTR"
)

func TestLoadTouch(t *testing.T) {
	tou, ErrorCat := zaratr.LoadTouch("ribbed-strappy-vest-top-p03253306")
	if ErrorCat != nil {
		t.Error(ErrorCat)
	}
	if tou.Product.Name == "" {
		t.Error("Неправльный ответ")
	}
}
