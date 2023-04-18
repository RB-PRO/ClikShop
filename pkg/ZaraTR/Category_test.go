package zaratr_test

import (
	"testing"

	zaratr "github.com/RB-PRO/SanctionedClothing/pkg/ZaraTR"
)

func TestLoadCategory(t *testing.T) {
	cat, ErrorCat := zaratr.LoadCategory()
	if ErrorCat != nil {
		t.Error(ErrorCat)
	}
	if cat.Categories[0].Name != "WOMAN" {
		t.Error("Неправльный ответ")
	}
}
