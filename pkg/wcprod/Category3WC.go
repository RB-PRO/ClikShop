package wcprod

import (
	"fmt"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
)

// Добавть категорию товара на WC из массива категорий
// и вернуть ID новой категории
func (woo *WcAdd) AddCategoryWC(cat bases.Cat) (NewId int, ErrorAddCat error) {
	for _, Category := range cat {
		// Поиск ID категории
		//	Идём дальше по категории товара
		// добавляем товар на WC, получаем его ID и добавляем его в наше дерево категорий
		CatObj, FindError := woo.FindCat3_WithParam(NewId, Category.Name, Category.Slug)
		if FindError != nil { // Если товар не добавлен
			var OldParent int = NewId
			NewId, ErrorAddCat = woo.UserWC.AddCat_WC(woocommerce.MeCat{ParentID: NewId, Name: Category.Name, Slug: Category.Slug})
			// Обработка ошибки тут должна быть
			woo.AddCategory3(NewId, &Category3Base{Parent: OldParent, Name: Category.Name, Slug: Category.Slug})

			woo.PrintCat3()
			fmt.Println()
		} else {
			NewId = CatObj.Parent
		}
	}
	return NewId, nil
}
