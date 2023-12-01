package massimodutti

import (
	"github.com/RB-PRO/ClikShop/pkg/bases"
)

// Структура категорий, которая исползуется для обозначения всех товаров по его древу категорий
type CategoryBasesForm struct {
	ID  int         // ID категории
	Cat []bases.Cat // Массив категорий
}

// Сформировать слайс категорий CategoryBasesForm для дальнейшего парсинга по ним
func CategoryBasesForming(catt Categories) (c []CategoryBasesForm) {
	for _, i_val := range catt.Categories {
		for _, j_val := range i_val.Subcategories {
			// Проверка на тип категории.
			//Заметили, что именно с этим типом категория считается валидной
			if j_val.Type == "22" {
				// fmt.Printf("%d - %s, %d - %s\n", i, i_val.Name, j, j_val.Name)
				c = append(c, CategoryBasesForm{
					ID: j_val.ID,
					Cat: []bases.Cat{
						{Name: "Massimo Dutti", Slug: bases.Name2Slug("Massimo Dutti")},
						{Name: i_val.Name, Slug: bases.Name2Slug(i_val.Name)},
						{Name: j_val.Name, Slug: bases.Name2Slug(j_val.Name)},
					},
				})
			}
		}
	}
	return c
}
