package wcprod_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/wcprod"
	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
)

func TestFormMapCat3(t *testing.T) {
	// Создаём экземпляр загрузчика данных
	Adding, errorInitWcAdd := wcprod.New()
	if errorInitWcAdd != nil {
		t.Error(errorInitWcAdd)
	}

	// Создать категорию товаров
	ErrorFormCat := Adding.FormMapCat3()
	if ErrorFormCat != nil {
		t.Error(ErrorFormCat)
	}

	Adding.PrintCat3()

	fmt.Printf("%+v", Adding.Cat3)

}

func TestAddCategory3(t *testing.T) {
	fmt.Println("TestAddCategory3:")
	Adding := newCat3() // Создаём экземпляр загрузчика данных
	Adding.PrintCat3()

	NewCateg := wcprod.Plc2cat3(woocommerce.ProductListCategory{
		ID:     55,
		Name:   "Test55",
		Slug:   "test55",
		Parent: 0,
	})

	if ErrorAdd := Adding.AddCategory3(55, NewCateg); ErrorAdd != nil {
		t.Error(ErrorAdd)
	}
	Adding.PrintCat3()
}

func TestFindCat3(t *testing.T) { // Поиск товара по ID
	Adding := newCat3() // Создаём экземпляр загрузчика данных

	FindProd, ErrorFind := Adding.FindCat3(0)
	if ErrorFind != nil {
		t.Error(ErrorFind)
	}
	fmt.Println("FindCat3: Нашёл товар", FindProd.Name, FindProd.Slug, FindProd.Parent)
}

func TestPrintCat3(t *testing.T) { // Печать категории
	Adding := newCat3() // Создаём экземпляр загрузчика данных

	Adding.PrintCat3()
}

func newCat3() *wcprod.WcAdd {
	Adding := new(wcprod.WcAdd)
	Adding.Cat3 = make(map[int]*wcprod.Category3Base)

	Adding.Cat3[0] = &wcprod.Category3Base{
		Parent: 0,
		Cat3:   map[int]*wcprod.Category3Base{},
	}
	Adding.Cat3[0].Cat3[1] = &wcprod.Category3Base{
		Name:   "Test1",
		Slug:   "test1",
		Parent: 0,
		Cat3:   map[int]*wcprod.Category3Base{},
	}
	Adding.Cat3[0].Cat3[1].Cat3[2] = &wcprod.Category3Base{
		Name:   "Test2",
		Slug:   "test2",
		Parent: 1,
		Cat3:   map[int]*wcprod.Category3Base{},
	}
	Adding.Cat3[0].Cat3[1].Cat3[2].Cat3[3] = &wcprod.Category3Base{
		Name:   "Test3",
		Slug:   "test3",
		Parent: 2,
		Cat3:   map[int]*wcprod.Category3Base{},
	}
	Adding.Cat3[0].Cat3[1].Cat3[2].Cat3[3].Cat3[4] = &wcprod.Category3Base{
		Name:   "Test4",
		Slug:   "test4",
		Parent: 3,
		Cat3:   map[int]*wcprod.Category3Base{},
	}

	Adding.Cat3[0].Cat3[1].Cat3[22] = &wcprod.Category3Base{
		Name:   "Test22",
		Slug:   "test22",
		Parent: 1,
		Cat3:   map[int]*wcprod.Category3Base{},
	}
	Adding.Cat3[0].Cat3[1].Cat3[22].Cat3[33] = &wcprod.Category3Base{
		Name:   "Test33",
		Slug:   "test33",
		Parent: 22,
		Cat3:   map[int]*wcprod.Category3Base{},
	}
	Adding.Cat3[0].Cat3[1].Cat3[22].Cat3[33].Cat3[44] = &wcprod.Category3Base{
		Name:   "Test44",
		Slug:   "test44",
		Parent: 33,
		Cat3:   map[int]*wcprod.Category3Base{},
	}
	return Adding
}
