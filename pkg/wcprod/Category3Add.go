package wcprod

import (
	"errors"

	"github.com/RB-PRO/ClikShop/pkg/woocommerce"
)

// ## Добавить категорию в дерево категорий.
//
// Для функции необходимо получить ID категории новой и ссылку на элемент новой категории
var ErrorIsFind error = errors.New("AddCategory3: Category in entry")

func (woo *WcAdd) AddCategory3(NewId int, NewCat3 *Category3Base) error {
	if NewCat3.Parent == 0 { // Если это базовая категория с ID 0
		woo.Cat3[0].Cat3[NewId] = NewCat3
		woo.Cat3[0].Cat3[NewId].Cat3 = make(map[int]*Category3Base)
		return nil
	}

	//fmt.Println(NewCat3.Parent, NewCat3.Name, NewCat3.Slug)
	//FindMap, ErrorFindCat := woo.FindCat3_WithParam(NewCat3.Parent, NewCat3.Name, NewCat3.Slug)
	FindMap, ErrorFindCat := woo.FindCat3(NewCat3.Parent)
	if ErrorFindCat != nil {
		return ErrorIsFind
	}
	FindMap.Cat3[NewId] = NewCat3
	FindMap.Cat3[NewId].Cat3 = make(map[int]*Category3Base)
	return nil
	//if FindMap.Cat3 == nil {
	//FindMap.Cat3[NewId] = new(Category3Base)
	//}

}

// Преобразовать результат категорий в элемент дерева категорий
func Plc2cat3(PlcCat woocommerce.ProductListCategory) *Category3Base {
	return &Category3Base{
		Name:   PlcCat.Name,
		Slug:   PlcCat.Slug,
		Parent: PlcCat.Parent,
		Cat3:   make(map[int]*Category3Base),
	}
}
