package wcprod

import (
	"errors"
	"fmt"

	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
)

// Структура категории
type Category3Base struct {
	Name   string
	Slug   string
	Parent int                    // Родитель
	Cat3   map[int]*Category3Base // Мапа потомков
}

// Сформировать дерево категорий
func (woo *WcAdd) FormMapCat3() error {
	plc := woo.Plc
	fmt.Println("Длина входного массива категорий:", len(plc.Category))
	for _, ValPlc := range plc.Category {
		fmt.Println(ValPlc.ID, ValPlc.Name, ValPlc.Parent)
	}

	//LenPlc := len(plc.Category)
	var LenCat3 int = 0
	Cat := make(map[int]*Category3Base)
	Cat[0] = &Category3Base{Parent: 0}

	// Идём по массиву
	for len(plc.Category) != LenCat3 {
		fmt.Println("->", len(plc.Category), LenCat3, len(plc.Category) != LenCat3)
		for IndexPlc, PlcVal := range plc.Category {
			if !PlcVal.IsAdd3 { // Если товар не добавлен в категории
				ErrorAdd := woo.AddCategory3(&Cat, PlcVal.ID, plc2cat3(PlcVal)) // Добавить товар
				if ErrorAdd != nil {
					return ErrorAdd
				}

				plc.Category[IndexPlc].IsAdd3 = true // Поставить чек, то товар добавлен
				LenCat3++
			}
		}
	}

	woo.Cat3 = Cat

	return nil
}

// Преобразовать результат категорий в своё дерево категорий
func plc2cat3(PlcCat woocommerce.ProductListCategory) Category3Base {
	return Category3Base{
		Name:   PlcCat.Name,
		Slug:   PlcCat.Slug,
		Parent: PlcCat.Parent,
	}
}

// Добавить категорию в товар
func (woo *WcAdd) AddCategory3(cat *map[int]*Category3Base, NewID int, NewCat3 Category3Base) error {

	// Если это базовая категория с ID 0
	FindMap, ErrorFindCat := woo.FindCat3(NewID)
	if errors.Is(ErrorFindCat, ErrorNotFoundCatId) {
		return ErrorFindCat
	}
	if errors.Is(ErrorFindCat, ErrorNotFoundCatfindCategory3Base) {
		return ErrorFindCat
	}

	FindMap.Cat3[NewID] = &NewCat3

	return nil
}

/////////////////////////////////////////////////////////////////////////////////

var ErrorNotFoundCatId error = errors.New("FindCat3: Not found Category3Base in Cat3")

// Поиск по ID в дереве категорий
func (woo *WcAdd) FindCat3(FindId int) (*Category3Base, error) {
	fmt.Println(" woo.Cat3", woo.Cat3)
	for Id, CatBase := range woo.Cat3 {
		if Id == FindId { // Если это этот ID
			return CatBase, nil
		}

		fmt.Println("CatBase", CatBase.Cat3)
		FindCat, ErrorFind := findCategory3Base(CatBase, FindId)
		if ErrorFind != nil {
			return FindCat, nil
		}
	}
	return nil, ErrorNotFoundCatId
}

var ErrorNotFoundCatfindCategory3Base error = errors.New("findCategory3Base: not found Category3Base in Cat3")

// Поиск именно в ячейке
func findCategory3Base(cat *Category3Base, FindId int) (*Category3Base, error) {
	fmt.Println("cat.Cat3", cat.Cat3)
	for Id, CatBase := range cat.Cat3 {
		fmt.Println("ID", Id)
		if Id == FindId { // Если это этот ID
			return CatBase, nil
		}
		//if len(CatBase.Cat3) != 0 {
		if CatBase.Cat3 != nil {
			FindCat, ErrorFindCat := findCategory3Base(CatBase, FindId)
			if ErrorFindCat != nil { // Если не нашёл
				return FindCat, nil
			}
		}
	}
	return nil, ErrorNotFoundCatfindCategory3Base
}

/////////////////////////////////////////////////////////////////////////////////

// ## Печать категорий товаров.
//
// Пример:
// `0
// - 1 Test1 test1
// -- 2 Test2 test2
// --- 3 Test3 test3
// ---- 4 Test4 test4
// -- 22 Test22 test22
// --- 33 Test33 test33
// ---- 44 Test44 test44`
func (woo *WcAdd) PrintCat3() {
	var prefix string = "-"

	for Id, CatBase := range woo.Cat3 {
		fmt.Println(Id)

		printCategory3Base(CatBase, prefix)
	}
}

// Обход всех потомков и вывод на экран
func printCategory3Base(cat *Category3Base, prefix string) {
	for Id, CatBase := range cat.Cat3 {
		fmt.Println(prefix, Id, CatBase.Name, CatBase.Slug)
		printCategory3Base(CatBase, prefix+"-")
	}
}
