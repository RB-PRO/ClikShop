package wcprod

// # Структура дерева категорий с ключом - ID ключа мапы подкатегории
//
//	Дерево категорий - сложная структура данных, которая содержит информацию о категории и может сожержать ссылки на дочерние категории
//
// В данной структуре категорий имеется ряд особенностей. Например базовая категория имеет ID 0 и создаётся изначально.
type Category3Base struct {
	// Содержимое категории:
	Name string // Навазние категории
	Slug string // Ярлык(ссылка)

	// Системная информация
	Parent int                    // Родитель
	Cat3   map[int]*Category3Base // Мапа потомков
}

// ### Сформировать дерево категорий из массива входных данных.
//
// Пример массива входных данных:
//
//	[]ArrData{
//		ID     int    // id категории
//		Name   string // Название категории
//		Parent int    // id родительской категории
//	}
func (woo *WcAdd) FormMapCat3() error {
	plc := woo.Plc.Category // Входной массив данных
	// for index, value := range plc {
	// fmt.Println(index, value.ID, value.Name, value.Parent)
	// }
	//fmt.Println("len(plc)", len(plc))

	var LenCat3 int = 1       // Переменная-счётчик, которая сравнивается с общим к-вом входных данных из исходного массива
	for len(plc) >= LenCat3 { // Пока не добавили все товары в дерево категорий
		for IndexPlc, PlcVal := range plc { // Цикл по всем полям входной структуры
			//fmt.Println(">>>", len(plc), LenCat3, IndexPlc, plc[IndexPlc].IsAdd3, plc[IndexPlc].Name)
			if !PlcVal.IsAdd3 { // Если товар не добавлен в категории

				// Добавить товар. Передаём ID нового товара и элемент массива входного, который и планируем добавить
				// В элементе входного массива обязательно должна быть переменная Parent, которая содержит ID родителя
				// Если получаем ошибку в результате добавления товара, то продолжаем иттерацию по входному массиву
				//fmt.Println("ADD", PlcVal.ID, PlcVal.Name, PlcVal.Parent)
				if ErrorAdd := woo.AddCategory3(PlcVal.ID, Plc2cat3(PlcVal)); ErrorAdd != nil {
					continue
				}

				// Фиксируем во входной структуре, что товар добавлен в категорию товаров.
				plc[IndexPlc].IsAdd3 = true
				LenCat3++

			}
		}
	}
	return nil
}

/*
>>> 8 4 5 false Test3
ADD 3686 Test3 3685
>>> 8 4 6 true TestBase
- 3683 Test0 test0
- 3709 TestBase testbase
- 15 Товары без категории misc
*/
