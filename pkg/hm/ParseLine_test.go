package hm_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/hm"
)

// https://www2.hm.com/tr_tr/kadin/urune-gore-satin-al/ust/_jcr_content/main/productlisting.display.json?page-size=1
func TestLinesAll(t *testing.T) {
	Prod, ErrorParseProduct := hm.LinesAll("/tr_tr/cocuk/urune-gore-satin-al/hepsini-incele/_jcr_content/main/productlisting_acbd.display.json")
	if ErrorParseProduct != nil {
		t.Error(ErrorParseProduct)
	}
	if len(Prod.Products) == 0 {
		t.Error("LinesAll: Ноль товаров получено")
	}
}

func TestLineUrl(t *testing.T) {
	core, ErrNewParsingCard := hm.NewParsingCard()
	if ErrNewParsingCard != nil {
		t.Error(ErrNewParsingCard)
	}
	url, ErrLineUrl := core.LineUrl("/tr_tr/kadin/urune-gore-satin-al/ust.html")
	if ErrLineUrl != nil {
		t.Error(ErrLineUrl)
	}
	fmt.Println(hm.URL + url)
	if url == "" {
		t.Errorf("LineUrl: url is nil")
	}
}

func TestLineUrl2(t *testing.T) {
	url, ErrLineUrl := hm.LineUrl2("/tr_tr/kadin/urune-gore-satin-al/ust.html")
	if ErrLineUrl != nil {
		t.Error(ErrLineUrl)
	}
	fmt.Println(hm.URL + url)
	if url == "" {
		t.Errorf("LineUrl: url is nil")
	}
}
