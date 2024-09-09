package pm6_test

import (
	"fmt"
	"strconv"
	"testing"

	"ClikShop/common/bases"
	"ClikShop/common/pm6"
)

func TestAllPages(t *testing.T) {
	pmm, ErrorPMM := pm6.NewPM()
	if ErrorPMM != nil {
		t.Error(ErrorPMM)
	}
	pagesInt := pmm.AllPages("/null/.zso?s=isNew/desc/goLiveDate/desc/recentSalesStyle/desc/")
	fmt.Println("Всего страниц:", pagesInt)
	if pagesInt == 0 {
		t.Error("Неправильное к-во товаров. По ссылке https://www.6pm.com/null/.zso?s=isNew/desc/goLiveDate/desc/recentSalesStyle/desc/ Должно быть \"[не ноль]\", а получено " + "\"" + strconv.Itoa(pagesInt) + "\"")
	}
}
func TestParsePageWithVarienty(t *testing.T) {
	pmm, _ := pm6.NewPM()
	linkPages := "/null/.zso?s=isNew/desc/goLiveDate/desc/recentSalesStyle/desc/" // Ссылка на страницу товаров
	pagesInt := 2                                                                 // Получить сколько всего страниц товаров есть

	var varient bases.Variety2                                 // Массив базы данных товаров
	varient = pmm.ParsePageWithVarienty(varient, linkPages, 0) // Парсим первую страницу товаров
	for i := 1; i <= pagesInt; i++ {                           // Цикл по всем страницам товаров
		// Сортируем товары и записываем их в готовую базу данных varient
		varient = pmm.ParsePageWithVarienty(varient, linkPages, i) // Парсим первую страницу товаров
	}
	PrintVarient(varient) // Печать
}

func PrintVarient(varient bases.Variety2) {
	fmt.Println("len", len(varient.Product))
	for index, value := range varient.Product {
		strs := ""
		for key := range value.Item {
			strs += value.Item[key].ColorCode + ", "
		}
		fmt.Println(index, ":", ">"+value.Name+"<", "color:", strs)
	}
}
