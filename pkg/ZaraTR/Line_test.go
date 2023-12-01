package zaratr_test

import (
	"testing"

	zaratr "github.com/RB-PRO/ClikShop/pkg/ZaraTR"
)

func TestLoadLine(t *testing.T) {
	lin, ErrorCat := zaratr.LoadLine("2215112")
	if ErrorCat != nil {
		t.Error(ErrorCat)
	}
	if len(lin.ProductGroups) == 0 {
		t.Error("Неправльный ответ")
	}
}
