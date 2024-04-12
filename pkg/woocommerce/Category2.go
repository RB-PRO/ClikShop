package woocommerce

import (
	"errors"

	"github.com/RB-PRO/ClikShop/pkg/bases"
)

func (user *User) AddCat2(plc *Categorys, cat []bases.Cat) (NewAddParentId int, errorAdd error) {
	for itteration, NewCat := range cat {
		var FindParent int
		if itteration == 0 {
			FindParent = 0
		}

		//Find, errorFInd := plc.FindSlugName(NewCat.Name, NewCat.Slug)                 // Поиск по имени и Ссылке
		Find, errorFInd := plc.FindSlugNameParent(NewCat.Name, NewCat.Slug, FindParent) //  Поиск по имени и Ссылке и родительской категории
		if errorFInd != nil {                                                           // Если найдено
			// Добавляем товар
			if itteration == 0 {
				NewAddParentId = 0
			}

			NewAddParentId, errorAdd = user.localAdd(plc, NewAddParentId, NewCat.Name, NewCat.Slug)

		} else {
			NewAddParentId = Find.ID
		}
		FindParent = NewAddParentId // Присваение ID поиска по родительской категории
	}
	return NewAddParentId, nil
}

func (user *User) localAdd(plc *Categorys, parentId int, NewName string, NewSlug string) (NewId int, errorAdd error) {

	//fmt.Println("AddCat_WC", parentId, NewName, NewSlug)
	NewId, errorAdd = user.AddCat_WC(MeCat{
		ParentID:    parentId,
		Name:        NewName,
		Slug:        NewSlug,
		Description: "Автокатегория",
	})
	if errorAdd != nil {
		return 0, errorAdd
	}

	plc.Category = append(plc.Category, ProductListCategory{
		ID:     NewId,
		Name:   NewName,
		Slug:   NewSlug,
		Parent: parentId,
	})
	return NewId, nil
}

// Поиск в массиве категорий по Slug
func (plc Categorys) FindSlug(Slug string) (ProductListCategory, error) {
	for index := range plc.Category {
		if plc.Category[index].Slug == Slug {
			return plc.Category[index], nil
		}
	}
	return ProductListCategory{}, errors.New("plc.FindSlug: Не нашёл категории с таким SLUG = " + Slug)
}

// Поиск в массиве категорий по Name
func (plc Categorys) FindName(Name string) (ProductListCategory, error) {
	for index := range plc.Category {
		if plc.Category[index].Name == Name {
			return plc.Category[index], nil
		}
	}
	return ProductListCategory{}, errors.New("plc.FindSlug: Не нашёл категории с таким Name = " + Name)
}

// Поиск в массиве категорий по Slug + Name
func (plc Categorys) FindSlugName(Name, Slug string) (ProductListCategory, error) {
	for index := range plc.Category {
		if plc.Category[index].Name == Name && plc.Category[index].Slug == Slug {
			return plc.Category[index], nil
		}
	}
	return ProductListCategory{}, errors.New("plc.FindSlug: Не нашёл категории с таким Name = " + Name + ", с таким SLUG = " + Slug)
}

// Поиск в массиве категорий по Slug + Name
func (plc Categorys) FindSlugNameParent(Name, Slug string, Parent int) (ProductListCategory, error) {
	for index := range plc.Category {
		if plc.Category[index].Name == Name && plc.Category[index].Slug == Slug && plc.Category[index].Parent == Parent {
			return plc.Category[index], nil
		}
	}
	return ProductListCategory{}, errors.New("plc.FindSlug: Не нашёл категории с таким Name = " + Name + ", с таким SLUG = " + Slug)
}
