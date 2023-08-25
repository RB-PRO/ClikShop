package sneaksup

import "testing"

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
		if res != e.out {
			t.Errorf("Find(%s) = %s, expected %s",
				e.inp, res, e.out)
		}
	}
}
