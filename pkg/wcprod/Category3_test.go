package wcprod_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/wcprod"
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

func TestFindCat3(t *testing.T) { // Поиск товара по ID
	// Создаём экземпляр загрузчика данных
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

	FindProd, ErrorFind := Adding.FindCat3(33)
	if ErrorFind != nil {
		t.Error(ErrorFind)
	}
	fmt.Println("FindCat3: Нашёл товар", FindProd.Name, FindProd.Slug, FindProd.Parent)
}

func TestPrintCat3(t *testing.T) { // Печать категории
	// Создаём экземпляр загрузчика данных
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

	Adding.PrintCat3()
}
