package massimodutti

import (
	"ClikShop/common/bases"
)

// Структура категорий, которая исползуется для обозначения всех товаров по его древу категорий
type CategoryBasesForm struct {
	ID  int         // ID категории
	Cat []bases.Cat // Массив категорий
}

// Сформировать слайс категорий CategoryBasesForm для дальнейшего парсинга по ним
func CategoryBasesForming(catt Categories) (c []CategoryBasesForm) {
	for _, category := range catt.Categories {
		for _, subcategory := range category.Subcategories {
			//fmt.Printf("%d //%d - %s, %d - %s\n", subcategory.Type, i, category.Name, j, subcategory.Name)

			if subcategory.Name != "COLLECTION" {
				continue
			}
			for _, subcategoryTwo := range subcategory.Subcategories {
				c = append(c, CategoryBasesForm{
					ID: subcategoryTwo.ID,
					Cat: []bases.Cat{
						{Name: "Massimo Dutti", Slug: bases.Name2Slug("Massimo Dutti")},
						{Name: category.Name, Slug: bases.Name2Slug(category.Name)},
						{Name: subcategory.Name, Slug: bases.Name2Slug(subcategory.Name)},
						{Name: subcategoryTwo.Name, Slug: bases.Name2Slug(subcategoryTwo.Name)},
					},
				})
			}
		}
	}
	return c
}
