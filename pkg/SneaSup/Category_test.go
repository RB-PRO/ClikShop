package sneaksup

import (
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

func TestCategory(t *testing.T) {
	cc := Category()
	// fmt.Println(cc)

	prods := []bases.Product2{}
	for _, val := range cc {
		prods = append(prods, bases.Product2{Cat: val.cat, Link: val.link})
	}
	Variety2 := bases.Variety2{Product: prods}
	Variety2.SaveXlsxCsvs("sneaksup")
}
