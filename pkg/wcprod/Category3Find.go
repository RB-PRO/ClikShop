package wcprod

import (
	"errors"
)

/////////////////////////////////////////////////////////////////////////////////
/////////////////////////Функции поиска категегории//////////////////////////////
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

var ErrorNotFoundCatId_WC error = errors.New("FindCFindCat3_WithParamat3: Not found Category3Base in Cat3")

func (woo *WcAdd) FindCat3_WithParam(Parent int, Name string, Slug string) (*Category3Base, int, error) {
	// fmt.Println(" woo.Cat3", woo.Cat3)
	for Id, CatBase := range woo.Cat3 {
		//fmt.Println("Id,CatBase.Name,CatBase.Slug", Id, CatBase.Name, CatBase.Slug)
		if CatBase.Parent == Parent && CatBase.Name == Name && CatBase.Slug == Slug { // Если это этот ID
			return CatBase, Id, nil
		}
		//fmt.Println("CatBase", CatBase.Cat3)
		FindCat, Id, ErrorFind := findCategory3Base_WithParam(CatBase, Parent, Name, Slug)
		//fmt.Println("CatBase!!!", FindCat, ErrorFind)
		if ErrorFind == nil {
			return FindCat, Id, nil
		}
	}
	return nil, 0, ErrorNotFoundCatId
}

var ErrorNotFoundCatfindCategory3Base_WC error = errors.New("findCategory3Base_WithParam: not found Category3Base in Cat3")

// Поиск именно в ячейке
func findCategory3Base_WithParam(cat *Category3Base, Parent int, Name string, Slug string) (*Category3Base, int, error) {
	// fmt.Println("cat.Cat3", cat.Cat3)
	for Id, CatBase := range cat.Cat3 {
		//fmt.Println("Id,CatBase.Name,CatBase.Slug", Id, CatBase.Name, CatBase.Slug, "-", Parent, Name, Slug)
		if CatBase.Parent == Parent && CatBase.Name == Name && CatBase.Slug == Slug { // Если это этот ID
			return CatBase, Id, nil
		}
		//if len(CatBase.Cat3) != 0 {
		if CatBase.Cat3 != nil && len(CatBase.Cat3) != 0 {
			FindCat, Id, ErrorFindCat := findCategory3Base_WithParam(CatBase, Parent, Name, Slug)
			//fmt.Println("---FindCat", FindCat, ErrorFindCat, " -- ", FindCat.Name, FindCat.Slug, FindCat.Parent)
			if ErrorFindCat == nil { // Если не нашёл
				return FindCat, Id, nil
			}
		}
		//fmt.Println("Exit")
	}
	return &Category3Base{}, 0, ErrorNotFoundCatfindCategory3Base
}
