package wcprod

import (
	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/woocommerce"
)

// Добавть категорию товара на WC из массива категорий
// и вернуть ID новой категории
// Передаём в эту функцию bases.Cat и получаем на выходе ID категории.
func (woo *WcAdd) AddCategoryWC(cat []bases.Cat) (NewId int, ErrorAddCat error) {
	for _, Category := range cat {
		// Поиск ID категории
		//	Идём дальше по категории товара
		// добавляем товар на WC, получаем его ID и добавляем его в наше дерево категорий
		//fmt.Println("ищу", NewId, Category.Name)
		_, Id, FindError := woo.FindCat3_WithParam(NewId, Category.Name, Category.Slug)
		if FindError != nil { // Если товар не добавлен
			var OldParent int = NewId
			//fmt.Println("Создаю категорию", OldParent)
			NewId, ErrorAddCat = woo.UserWC.AddCat_WC(woocommerce.MeCat{ParentID: NewId, Name: Category.Name, Slug: Category.Slug})
			if ErrorAddCat != nil { // Обработка ошибки тут должна быть
				return 0, ErrorAddCat
			}

			woo.AddCategory3(NewId, &Category3Base{Parent: OldParent, Name: Category.Name, Slug: Category.Slug})

		} else {
			NewId = Id
		}
	}
	return NewId, nil
}
