package zaratr

import (
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
)

type CategoryArray struct {
	Items    []Item
	TecalCat []bases.Cat
}

// https://www.zara.com/tr/en/categories?ajax=true
func CatCycle2() (*CategoryArray, error) {

	// Загружаем все категории
	Cat, ErrorCategory := LoadCategory()
	if ErrorCategory != nil {
		return nil, ErrorCategory
	}

	c := new(CategoryArray) // Выделить память в исходную структуру данных

	// Цикл по всем подкатегориям.Они должны быть man, woman, kid
	// CatBasesHome := []bases.Cat{{Name: "Zara", Slug: "zara"}} // Исходная нулевая категория
	CatBasesHome := make([]bases.Cat, 1, 5)
	CatBasesHome[0] = bases.Cat{Name: "Zara", Slug: "zara"}
	for _, CatsCategory1 := range Cat.Categories {
		if CatsCategory1.Name == "WOMAN" || CatsCategory1.Name == "MAN" || CatsCategory1.Name == "KIDS" {
			newslise := make([]bases.Cat, 1, 5)
			copy(newslise, CatBasesHome)
			CatBasesHome1 := append(newslise, bases.Cat{Name: CatsCategory1.Name, Slug: CatsCategory1.Seo.Keyword})

			// Цикл по всем подкатегориям из категории полов
			for _, CatsCategory2 := range CatsCategory1.Subcategories {
				if CatsCategory2.Name == "PERFUMES" ||
					CatsCategory2.Name == "BEAUTY" ||
					CatsCategory2.Name == "HOME" ||
					CatsCategory2.Name == "HOME KIDS" ||
					CatsCategory2.Name == "GIFT CARD" ||
					CatsCategory2.Name == "JOIN LIFE" ||
					CatsCategory2.Name == "ACCESSORIES | JEWELRY" ||
					CatsCategory2.Name == "BAGS" ||
					CatsCategory2.Name == "+ Info" ||
					CatsCategory2.Name == "-" ||
					CatsCategory2.Name == "__" ||
					strings.ToLower(CatsCategory2.Name) == "see all" ||
					strings.ToLower(CatsCategory2.Name) == "view all" ||
					strings.Contains(strings.ToLower(CatsCategory2.Name), "accessories") ||
					strings.Contains(strings.ToLower(CatsCategory2.Name), "beauty") ||
					strings.Contains(strings.ToLower(CatsCategory2.Name), "metallic touch") ||
					strings.Contains(strings.ToLower(CatsCategory2.Name), "special edition") ||
					strings.Contains(strings.ToLower(CatsCategory2.Name), "zara athleticz") ||
					strings.Contains(CatsCategory2.Name, "DIVIDER_MENU") {
					continue
				}
				// fmt.Printf("'%+v'\n", CatsCategory2.Name)
				// if CatsCategory2.Name == "SALE" || CatsCategory2.Name == "NEW COLLECTION" { //
				newslise := make([]bases.Cat, 2, 5)
				copy(newslise, CatBasesHome1)
				CatBasesHome2 := append(newslise, bases.Cat{Name: CatsCategory2.Name, Slug: CatsCategory2.Seo.Keyword})

				for _, CatsCategory3 := range CatsCategory2.Subcategories {
					if CatsCategory3.Name == "HOME" ||
						CatsCategory3.Name == "HOME KIDS" ||
						CatsCategory3.Name == "PERFUMES" ||
						CatsCategory3.Name == "GIFT CARD" ||
						CatsCategory3.Name == "JOIN LIFE" ||
						CatsCategory3.Name == "BEAUTY" ||
						CatsCategory3.Name == "ACCESSORIES | JEWELRY" ||
						CatsCategory3.Name == "BAGS" ||
						CatsCategory3.Name == "-" ||
						strings.ToLower(CatsCategory3.Name) == "see all" ||
						strings.ToLower(CatsCategory3.Name) == "view all" ||
						strings.Contains(strings.ToLower(CatsCategory3.Name), "accessories") ||
						strings.Contains(strings.ToLower(CatsCategory3.Name), "beauty") ||
						strings.Contains(strings.ToLower(CatsCategory3.Name), "metallic touch") ||
						strings.Contains(strings.ToLower(CatsCategory3.Name), "special edition") ||
						strings.Contains(strings.ToLower(CatsCategory3.Name), "zara athleticz") ||
						strings.Contains(CatsCategory2.Name, "DIVIDER_MENU") {
						continue
					}
					// fmt.Printf("'%+v'\n", CatsCategory3.Name)

					// if strings.Contains(CatsCategory3.Name, "JEANS") ||
					// 	strings.Contains(CatsCategory3.Name, "T-SHIRTS") ||
					// 	strings.Contains(CatsCategory3.Name, "SHIRTS") ||
					// 	strings.Contains(CatsCategory3.Name, "SHOES") ||
					// 	strings.Contains(CatsCategory3.Name, "GIRL") ||
					// 	strings.Contains(CatsCategory3.Name, "BOY") ||
					// 	strings.Contains(CatsCategory3.Name, "BOY") ||
					// 	strings.Contains(CatsCategory3.Name, "BOY") { // футболки

					newslise := make([]bases.Cat, 3, 5)
					copy(newslise, CatBasesHome2)
					CatBasesHome3 := append(newslise, bases.Cat{Name: CatsCategory3.Name, Slug: CatsCategory3.Seo.Keyword})
					// fmt.Println(CatBasesHome3)

					// Дальше нужно определиться, что это
					// если это дети, то там нужно ещё цикл по подкатегориям,
					// где уже лежат исходные категории товаров
					var Gender string // Гендр товара
					Gender = bases.Name2Slug(CatsCategory1.Name)

					if Gender == "man" || Gender == "woman" {
						RedirectCategoryID := CatsCategory3.RedirectCategoryID
						if RedirectCategoryID == 0 {
							RedirectCategoryID, _ = CatsCategory3.ID.Int()
						}
						c.Items = append(c.Items, Item{RedirectCategoryID: RedirectCategoryID, Cat: CatBasesHome3, Gender: Gender})
					}
					if Gender == "kids" {
						Gender = "unisex"
						for _, CatsCategory4 := range CatsCategory3.Subcategories {
							if CatsCategory4.Name == "GIFT CARD" ||
								CatsCategory4.Name == "PERFUMES | COSMETICS" ||
								CatsCategory4.Name == "HOME" ||
								CatsCategory4.Name == "HOME KIDS" ||
								CatsCategory4.Name == "PERFUMES" ||
								CatsCategory4.Name == "JOIN LIFE" ||
								CatsCategory4.Name == "BEAUTY" ||
								CatsCategory4.Name == "ACCESSORIES | JEWELRY" ||
								CatsCategory4.Name == "BAGS" ||
								strings.ToLower(CatsCategory4.Name) == "see all" ||
								strings.ToLower(CatsCategory4.Name) == "view all" ||
								strings.Contains(strings.ToLower(CatsCategory4.Name), "accessories") ||
								strings.Contains(strings.ToLower(CatsCategory4.Name), "beauty") ||
								strings.Contains(strings.ToLower(CatsCategory4.Name), "metallic touch") ||
								strings.Contains(strings.ToLower(CatsCategory4.Name), "special edition") ||
								strings.Contains(strings.ToLower(CatsCategory4.Name), "zara athleticz") ||
								strings.Contains(CatsCategory2.Name, "DIVIDER_MENU") {
								continue
							}
							// fmt.Printf("'%+v'\n", CatsCategory4.Name)
							// if strings.Contains(CatsCategory4.Name, "JEANS") ||
							// 	strings.Contains(CatsCategory4.Name, "T-SHIRTS") ||
							// 	strings.Contains(CatsCategory4.Name, "SHIRTS") ||
							// 	strings.Contains(CatsCategory4.Name, "SHOES") ||
							// 	strings.Contains(CatsCategory4.Name, "JEANS") {

							newslise := make([]bases.Cat, 4, 5)
							copy(newslise, CatBasesHome3)
							CatBasesHome4 := append(newslise, bases.Cat{Name: CatsCategory4.Name, Slug: CatsCategory4.Seo.Keyword})
							if strings.Contains(CatsCategory3.Name, "BOY") {
								Gender = "boy"
							}
							if strings.Contains(CatsCategory3.Name, "GIRL") {
								Gender = "girl"
							}

							RedirectCategoryID := CatsCategory4.RedirectCategoryID
							if RedirectCategoryID == 0 {
								RedirectCategoryID, _ = CatsCategory4.ID.Int()
							}
							c.Items = append(c.Items, Item{RedirectCategoryID: RedirectCategoryID, Cat: CatBasesHome4, Gender: Gender})
							// }
						}
					}
					// }
				}
				// }
			}

		}
	}

	return c, nil
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
