package zaratr

import (
	"strings"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

type CategoryArray struct {
	Items    []Item
	TecalCat []bases.Cat
}

// https://www.zara.com/tr/en/categories?ajax=true
func CatCycle() *CategoryArray {

	// Загружаем все категории
	Cat, ErrorCategory := LoadCategory()
	if ErrorCategory != nil {
		panic(ErrorCategory)
	}

	c := new(CategoryArray)   // Выделить память в исходную структуру данных
	c.Items = make([]Item, 0) // Массив конечных категорий

	// Цикл по всем подкатегориям
	for _, catCaters := range Cat.Categories {
		if catCaters.Name == "WOMAN" || catCaters.Name == "MAN" || catCaters.Name == "KIDS" {

			// c.TecalCat = append(c.TecalCat, bases.Cat{
			// 	Name: catCaters.Name,
			// 	Slug: strings.ToLower(catCaters.Name),
			// })
			c.TecalCat = []bases.Cat{{
				Name: catCaters.Name,
				Slug: strings.ToLower(catCaters.Name),
			}}
			c.cycle(&catCaters)
		}
	}

	return c
}

// Перебрать все категории циклом и на выходе получить массив категорий
func (c *CategoryArray) cycle(catCaters *Subcategories) {
	for _, cat := range catCaters.Subcategories {
		if filter(cat.Name) {
			continue
		}

		// Добавляем в текущий аргумент массива
		c.TecalCat = append(c.TecalCat, bases.Cat{
			Name: cat.Name,
			Slug: strings.ToLower(cat.Name),
		})

		if len(cat.Subcategories) == 0 {

			// На это решение я учил 2.5 часа.
			// Прошу ознакомиться со [статьёй], где описывается про копирование слайсов и почему это надо/не надо делать.
			//
			// [статьёй]: https://gosamples.dev/copy-slice/
			cat.Item.Cat = append(cat.Item.Cat, c.TecalCat...) // Актуализировать массив категорий

			c.Items = append(c.Items, cat.Item) // Добавить в массив категорий

			c.TecalCat = c.TecalCat[:len(c.TecalCat)-1] // Удалить последний элемент массива

		} else {
			c.cycle(&cat)
		}
	}
	c.TecalCat = c.TecalCat[:len(c.TecalCat)-1] // Удалить последний элемент массива
}

// Фильтр категорий.
// То имя категории, которое есть в каталоге, не учитываем
func filter(str string) bool {
	switch strings.ToUpper(str) {
	case "ALL PRODUCTS":
		return true
	case "VIEW ALL":
		return true
	case "ZARA SRPLS":
		return true
	case "BAGS":
		return true
	case "ACCESSORIES":
		return true
	case "BIKINIS | SWIMSUITS":
		return true
	case "PERFUMES":
		return true
	case "LINGERIE":
		return true
	case "BEAUTY":
		return true
	case "SPECIAL EDITION":
		return true
	case "GIFT CARD":
		return true
	case "E-GIFT CARD":
		return true
	case "BAGS | BACKPACKS":
		return true
	case "+ INFO":
		return true
	case "COMPANY":
		return true
	case "HOME KIDS":
		return true
	case "HOME":
		return true
	default:
		return false
	}
}
