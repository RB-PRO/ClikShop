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
