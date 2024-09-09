package louisvuitton

import (
	"fmt"
	"testing"
)

func coreTest() *Core {
	return NewCore("607e3016889f431fb8020693311016c9", "60bbcdcD722D411B88cBb72C8246a22F")
}

func TestCore(t *testing.T) {
	core := NewCore("", "")
	core.UpdateCore()
}

func TestCategory(t *testing.T) {
	categs, Err := Category()
	if Err != nil {
		t.Error(Err)
	}

	for i, v := range categs {
		fmt.Println(i, v.CategoryTag, v.Path)
	}
	fmt.Println("len(categs)", len(categs))
}

func TestPages(t *testing.T) {
	core := coreTest()
	Products := core.Pages("t1z0ff7q")
	fmt.Println("Всего товаров", len(Products))
}

func TestPage(t *testing.T) {
	core := coreTest()
	page, errpage := core.PageSingle("t1z0ff7q", 0)
	if errpage != nil {
		t.Error(errpage)
	}
	if page.NbPages == 0 {
		t.Error("PageSingle: Не получить получить ответ или распарсить его")
	}
	for _, v := range page.Hits {
		fmt.Println(v.ProductID)
	}
}

func TestTouch(t *testing.T) {
	core := coreTest()
	touch, ErrTouch := core.Toucher("nvprod3900007v")
	if ErrTouch != nil {
		t.Error(ErrTouch)
	}
	if touch.Name == "" {
		t.Error("Toucher: Не получить получить ответ или распарсить его")
	}
}

func TestLink2Cat(t *testing.T) {

	var tests = []struct {
		input   string
		output  string
		output2 string
	}{
		{"/rus-ru/art-of-living/books-and-stationery/hard-cover-books/_/N-t1134j9w", "LV/art-of-living/books-and-stationery/hard-cover-books/", "t1134j9w"},
		{"/rus-ru/women/handbags/_/N-tfr7qdp", "LV/women/handbags/", "tfr7qdp"},
		{"/rus-ru/stories/gifting", "LV/stories/gifting/", ""},
		{"/rus-ru/magazine", "LV/magazine/", ""},
	}

	for _, e := range tests {
		res := Link2Cat(e.input)
		if res.Path != e.output {
			t.Errorf("Link2Cat(%s) = %s, expected %s",
				e.input, res, e.output)
		}
		if res.CategoryTag != e.output2 {
			t.Errorf("Link2Cat(%s) = %s, expected %s",
				e.input, res, e.output2)
		}
	}
}

func TestSaveXLSX(t *testing.T) {
	core := coreTest()
	touch, ErrTouch := core.Toucher("001054")
	if ErrTouch != nil {
		t.Error(ErrTouch)
	}
	if touch.Name == "" {
		t.Error("Toucher: Не получить получить ответ или распарсить его")
	}

	Prod := TouchResponse2Product(touch)
	Prods := []Product{Prod}

	ErrSave := SaveXLSX("lv.xlsx", Prods)
	if ErrSave != nil {
		t.Error(ErrSave)
	}
}
