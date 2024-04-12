package wcprod_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/wcprod"
	"github.com/RB-PRO/ClikShop/pkg/woocommerce"
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

	fmt.Printf("%+v", Adding.Cat3[0].Cat3)

}

func TestAddCategory3(t *testing.T) {
	fmt.Println("TestAddCategory3:")
	// 1
	AddingNull := new(wcprod.WcAdd)
	AddingNull.Cat3 = make(map[int]*wcprod.Category3Base)
	AddingNull.Cat3[0] = &wcprod.Category3Base{}
	AddingNull.Cat3[0].Cat3 = make(map[int]*wcprod.Category3Base)
	//AddingNull.PrintCat3()
	NewCategNull := wcprod.Plc2cat3(woocommerce.ProductListCategory{
		ID:     55,
		Name:   "Test55",
		Slug:   "test55",
		Parent: 0,
	})
	if ErrorAdd := AddingNull.AddCategory3(55, NewCategNull); ErrorAdd != nil {
		t.Error(ErrorAdd)
	}
	AddingNull.PrintCat3()
	fmt.Println()

	// 2
	fmt.Println("TestAddCategory3:")
	Adding := newCat3() // Создаём экземпляр загрузчика данных
	Adding.PrintCat3()

	NewCateg := wcprod.Plc2cat3(woocommerce.ProductListCategory{
		ID:     55,
		Name:   "Test55",
		Slug:   "test55",
		Parent: 1,
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
func TestFindCat3_WithParam(t *testing.T) { // Поиск товара по ID
	fmt.Println("FindCat3_WithParam: 1")
	Adding := newCat3() // Создаём экземпляр загрузчика данных

	FindProd, _, ErrorFind := Adding.FindCat3_WithParam(2, "Test3", "test3")
	if ErrorFind != nil {
		t.Error(ErrorFind)
	}
	Adding.PrintCat3()
	fmt.Println("FindCat3_WithParam: Нашёл товар", FindProd.Name, FindProd.Slug, FindProd.Parent)
	fmt.Println()

	fmt.Println("FindCat3_WithParam: 2")
	Adding2 := newCat3() // Создаём экземпляр загрузчика данных

	FindProd2, _, ErrorFind2 := Adding2.FindCat3_WithParam(1, "Test2", "test2")
	if ErrorFind2 != nil {
		t.Error(ErrorFind2)
	}
	Adding.PrintCat3()
	fmt.Println("FindCat3_WithParam: Нашёл товар", FindProd2.Name, FindProd2.Slug, FindProd2.Parent)
	fmt.Println()
}

func TestPrintCat3(t *testing.T) { // Печать категории
	Adding := newCat3() // Создаём экземпляр загрузчика данных

	Adding.PrintCat3()
}

func newCat3() *wcprod.WcAdd {

	Adding := new(wcprod.WcAdd)
	Adding.Cat3 = make(map[int]*wcprod.Category3Base)
	Adding.Cat3[0] = &wcprod.Category3Base{}
	Adding.Cat3[0].Cat3 = make(map[int]*wcprod.Category3Base)

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

func TestAddCategoryWC(t *testing.T) {
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

	AddCat := make([]bases.Cat, 4)
	AddCat[0].Name = "Test0"
	AddCat[0].Slug = "test0"
	AddCat[1].Name = "Test1"
	AddCat[1].Slug = "test1"
	AddCat[2].Name = "Test222"
	AddCat[2].Slug = "test222"
	AddCat[3].Name = "Test333"
	AddCat[3].Slug = "test333"

	NewId, ErrorAdd := Adding.AddCategoryWC(AddCat)
	if ErrorAdd != nil {
		t.Error(ErrorAdd)
	}
	fmt.Println("Новый ID добавленной категории", NewId)
	Adding.PrintCat3()
}
