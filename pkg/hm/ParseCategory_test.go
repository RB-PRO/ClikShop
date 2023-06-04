package hm_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/hm"
)

func TestParseCategory(t *testing.T) {
	c, ErrorC := hm.Categorys()
	if ErrorC != nil {
		t.Error(ErrorC)
	}
	fmt.Println(len(c))
}
