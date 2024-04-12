package sneaksup

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/bases"
)

func TestCategory(t *testing.T) {
	cc := Category()
	// fmt.Println(cc)

	prods := []bases.Product2{}
	for _, val := range cc {
		prods = append(prods, bases.Product2{Cat: val.Cat, Link: val.Link})
	}
	Variety2 := bases.Variety2{Product: prods}
	Variety2.SaveXlsxCsvs("sneaksup")
}

func TestLine(t *testing.T) {

}

func TestLinkTranstore(t *testing.T) {

	var tests = []struct {
		inp string
		out string
	}{
		{"https://www.sneaksup.com/kadin-ayakkabi-sneaker",
			"https://www.sneaksup.com/kadin-ayakkabi-sneaker?paginationType=20&orderby=0&pagenumber=1"},
		{"https://www.sneaksup.com/kadin-ayakkabi-basketbol-ayakkabisi?pagenumber=1",
			"https://www.sneaksup.com/kadin-ayakkabi-basketbol-ayakkabisi?pagenumber=1&paginationType=20&orderby=0&pagenumber=1"},
		{"https://www.sneaksup.com/cocuk?specs=01-03-yas-cocuk-erkek,01-03-yas-cocuk-kiz,01-03-yas-cocuk-uni",
			"https://www.sneaksup.com/cocuk?specs=01-03-yas-cocuk-erkek,01-03-yas-cocuk-kiz,01-03-yas-cocuk-uni&paginationType=20&orderby=0&pagenumber=1"},
	}

	for _, e := range tests {
		res, err := linkTranstore(e.inp, 1)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(res)
		// if res != e.out {
		// 	t.Errorf("Find(%s) = %s, expected %s", e.inp, res, e.out)
		// }
	}
}

func TestLines(t *testing.T) {
	// Для проверки - https://www.sneaksup.com/kadin-giyim-esofman-alti?paginationType=20&orderby=0&pagenumber=1
	line, ErrLines := Lines("https://www.sneaksup.com/kadin-giyim-esofman-alti")
	if ErrLines != nil {
		t.Error(ErrLines)
	}
	fmt.Println("Cout Products:", len(line))
}

func TestProducts(t *testing.T) {
	// Для проверки - https://www.sneaksup.com/kadin-giyim-sweatshirt?paginationType=20&orderby=0&pagenumber=1
	url := "https://www.sneaksup.com/kadin-giyim-sweatshirt"
	res, _ := linkTranstore(url, 1)
	linePost, ErrLinePost := LinePost(res, 1)
	if ErrLinePost != nil {
		t.Error(ErrLinePost)
	}
	Prods := Line2Product(linePost.Products, SScat{})

	//

	fmt.Println(bases.ProdStr(Prods[0]))
}

func TestDescription(t *testing.T) {
	Description, ErrDesc := Description("https://www.sneaksup.com/nike-gamma-force-dx9176-103")
	if ErrDesc != nil {
		t.Error(ErrDesc)
	}
	fmt.Printf("'%s'\n", Description)
}

func TestAavailability(t *testing.T) {
	colors, ErrAavailability := Aavailability("https://www.sneaksup.com/jordan-w-j-brkln-flc-pant-dq4478-010")
	if ErrAavailability != nil {
		t.Error(ErrAavailability)
	}
	for _, color := range colors {
		fmt.Printf("%+v\n", color)
	}
}
