package zaratr

import "fmt"

func Parsing() {

	// Загрузить категории
	Cat, ErrorCategory := LoadCategory()
	if ErrorCategory != nil {
		panic(ErrorCategory)
	}

	//

	Items := CatCycle(Cat) // Наполнить цикл
	fmt.Println(len(Items))
}

// Перебрать все категории циклом и на выходе получить массив категорий
func CatCycle(Cat Category) []Item {
	Items := make([]Item, 0) // Массив конечных категорий

	for _, catCaters := range Cat.Categories {
		Items = append(Items, cycle(&catCaters))
	}

	return Items
}

func cycle(subs *Subcategories) []Item {
	Items := make([]Item, 0) // Массив конечных категорий
	for _, cat := range subs.Subcategories {

		// Если это конечная папка, то добавляем категорию
		if len(cat.Subcategories) == 0 {
			Items = append(Items, &cat.Item)
			return nil
		}

		// Цикл по потомку
		cycle(subs, Items)
	}
}
