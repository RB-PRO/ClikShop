package hm_test

import (
	"ClikShop/common/bases"
	"ClikShop/common/config"
	"ClikShop/common/hm"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// https://www2.hm.com/tr_tr/kadin/urune-gore-satin-al/ust/_jcr_content/main/productlisting.display.json?page-size=1
func TestLinesAll(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	hmService := hm.New(cfg)

	Prod, err := hmService.LinesAll("/tr_tr/cocuk/urune-gore-satin-al/hepsini-incele/_jcr_content/main/productlisting_acbd.display.json")
	if err != nil {
		t.Error(err)
	}
	if len(Prod.Products) == 0 {
		t.Error("LinesAll: Ноль товаров получено")
	}
}

func TestLineUrl(t *testing.T) {
	core, err := hm.NewParsingCard()
	if err != nil {
		t.Error(err)
	}
	url, err := core.LineUrl("/tr_tr/kadin/urune-gore-satin-al/ust.html")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hm.URL + url)
	if url == "" {
		t.Errorf("LineUrl: url is nil")
	}
}

func TestLineUrl2(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	hmService := hm.New(cfg)

	// https://www2.hm.com/tr_tr/kadin/urune-gore-satin-al/elbise.html
	// https://www2.hm.com/tr_tr/kadin/urune-gore-satin-al/denim.html
	categoryURL, err := hmService.LineUrl2("/tr_tr/kadin/urune-gore-satin-al/elbise.html", []bases.Cat{{Name: "Name", Slug: "name"}})
	if err != nil {
		t.Error(err)
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(categoryURL); err != nil {
		panic(err)
	}

	enc.SetIndent("", "  ")
	if err := enc.Encode(categoryURL); err != nil {
		panic(err)
	}

}
