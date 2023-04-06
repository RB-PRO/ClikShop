package wcprod

import "fmt"

/////////////////////////////////////////////////////////////////////////////////
////////////////////////////Функции печати категорий/////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

// ## Печать категорий товаров.
//
// Пример:
//
//   - 1 Test1 test1
//     -- 2 Test2 test2
//     --- 3 Test3 test3
//     ---- 4 Test4 test4
//     -- 22 Test22 test22
//     --- 33 Test33 test33
//     ---- 44 Test44 test44
func (woo *WcAdd) PrintCat3() {
	var prefix string = "-"
	for _, CatBase := range woo.Cat3 {
		printCategory3Base(CatBase, prefix)
	}
}

// Обход всех потомков и вывод на экран
func printCategory3Base(cat *Category3Base, prefix string) {
	for Id, CatBase := range cat.Cat3 {
		fmt.Println(prefix, Id, CatBase.Name, CatBase.Slug)
		printCategory3Base(CatBase, prefix+"-")
	}
}
