package hm_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/hm"
)

func TestParseCategory(t *testing.T) {
	c, ErrorC := hm.Categorys()
	if ErrorC != nil {
		t.Error(ErrorC)
	}
	fmt.Println(len(c))
	for i, val := range c {
		fmt.Println(i, val)
	}
}

func TestPullOutCat(t *testing.T) {
	cat, ErrCat := hm.PullOutCat("https://www2.hm.com/tr_tr/home/urune-gore-satin-al/dekorasyon.html")
	if ErrCat != nil {
		t.Error(ErrCat)
	}
	if reflect.DeepEqual(cat, []bases.Cat{{Name: "Home", Slug: "home"}, {Name: "Urune Gore Satin Al", Slug: "urune-gore-satin-al"}, {Name: "Dekorasyon", Slug: "dekorasyon"}}) {
		fmt.Println("Output cat:", cat)
		t.Error("Несопостовимые параметры")
	}
}
