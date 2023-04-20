package zaratr

import "fmt"

func Parsing() {

	//

	Items := CatCycle() // Наполнить цикл
	fmt.Println(len(Items))
}

func CatCycle() []Item {
	Cat, ErrorCategory := LoadCategory()
	if ErrorCategory != nil {
		panic(ErrorCategory)
	}
	Items := make([]Item, 0) // Массив конечных категорий

	for _, catCaters := range Cat.Categories {
		cycle(&catCaters, &Items)
	}

	return Items
}

func cycle(catCaters *Subcategories, Items *[]Item) {
	for _, c := range catCaters.Subcategories {
		if len(c.Subcategories) == 0 {

		}
	}
}

/*
// Перебрать все категории циклом и на выходе получить массив категорий
func CatCycle() []Item {
	// Загрузить категории
	Cat, ErrorCategory := LoadCategory()
	if ErrorCategory != nil {
		panic(ErrorCategory)
	}

	Items := make([]Item, 0) // Массив конечных категорий

	for _, catCaters := range Cat.Categories {
		Items = append(Items, cycle(&catCaters)...)
	}

	return Items
}

func cycle(subs *Subcategories) []Item {
	Items := make([]Item, 0) // Массив конечных категорий
	for _, cat := range subs.Subcategories {

		// Если это конечная папка, то добавляем категорию
		if len(cat.Subcategories) == 0 {
			Items = append(Items)
			return nil
		}

		// Цикл по потомку
		cycle(subs)
	}

	return Items
}
*/
