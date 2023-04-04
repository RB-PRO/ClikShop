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

	var LenCat3 int = 1

	// Идём по массиву
	for len(plc.Category) != LenCat3 {
		for IndexPlc, PlcVal := range plc.Category {
			// fmt.Println(IndexPlc, PlcVal.Parent, PlcVal.ID, PlcVal.Name, PlcVal.Slug, "---", len(plc.Category), LenCat3)

			if !PlcVal.IsAdd3 { // Если товар не добавлен в категории

				AddCat := Plc2cat3(PlcVal)
				fmt.Println(AddCat, &AddCat)

				ErrorAdd := woo.AddCategory3(PlcVal.ID, AddCat) // Добавить товар
				if ErrorAdd != nil {
					continue
					// 	fmt.Println(ErrorAdd)
					//return ErrorAdd
				}
				plc.Category[IndexPlc].IsAdd3 = true // Поставить чек, то товар добавлен
				LenCat3++
			}
		}
	}

	return nil
}

// Преобразовать результат категорий в своё дерево категорий
func Plc2cat3(PlcCat woocommerce.ProductListCategory) *Category3Base {
	return &Category3Base{
		Name:   PlcCat.Name,
		Slug:   PlcCat.Slug,
		Parent: PlcCat.Parent,
		Cat3:   make(map[int]*Category3Base),
	}
}

// Добавить категорию в товар
func (woo *WcAdd) AddCategory3(NewId int, NewCat3 *Category3Base) error {
	if NewCat3.Parent == 0 {
		woo.Cat3[0].Cat3 = make(map[int]*Category3Base)
		woo.Cat3[0].Cat3[NewId] = NewCat3
		return nil
	}

	// Если это базовая категория с ID 0
	FindMap, ErrorFindCat := woo.FindCat3(NewCat3.Parent)
	if ErrorFindCat != nil {
		return ErrorFindCat
	}
	FindMap.Cat3[NewId] = NewCat3
	FindMap.Cat3[NewId].Cat3 = make(map[int]*Category3Base)

	return nil
}

/////////////////////////////////////////////////////////////////////////////////

var ErrorNotFoundCatId error = errors.New("FindCat3: Not found Category3Base in Cat3")

// Поиск по ID в дереве категорий
func (woo *WcAdd) FindCat3(FindId int) (*Category3Base, error) {
	//fmt.Println(" woo.Cat3", woo.Cat3)
	for Id, CatBase := range woo.Cat3 {
		if Id == FindId { // Если это этот ID
			return CatBase, nil
		}
		//fmt.Println("CatBase", CatBase.Cat3)
		FindCat, ErrorFind := findCategory3Base(CatBase, FindId)
		//fmt.Println("CatBase!!!", FindCat, ErrorFind)
		if ErrorFind == nil {
			return FindCat, nil
		}
	}
	return nil, ErrorNotFoundCatId
}

var ErrorNotFoundCatfindCategory3Base error = errors.New("findCategory3Base: not found Category3Base in Cat3")

// Поиск именно в ячейке
func findCategory3Base(cat *Category3Base, FindId int) (*Category3Base, error) {
	//fmt.Println("cat.Cat3", cat.Cat3)
	for Id, CatBase := range cat.Cat3 {
		//fmt.Println("ID", Id == FindId, Id, FindId)
		if Id == FindId { // Если это этот ID
			return CatBase, nil
		}
		//if len(CatBase.Cat3) != 0 {
		if CatBase.Cat3 != nil && len(CatBase.Cat3) != 0 {
			FindCat, ErrorFindCat := findCategory3Base(CatBase, FindId)
			//fmt.Println("---FindCat", FindCat, ErrorFindCat, " -- ", FindCat.Name, FindCat.Slug, FindCat.Parent)
			if ErrorFindCat == nil { // Если не нашёл
				return FindCat, nil
			}
		}
		//fmt.Println("Exit")
	}
	return &Category3Base{}, ErrorNotFoundCatfindCategory3Base
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
	for _, CatBase := range woo.Cat3 {
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
