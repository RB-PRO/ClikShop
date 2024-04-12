package massimodutti

import (
	"fmt"
	"strconv"
	"testing"
)

func TestParsing2(t *testing.T) {

	// Получить все категории
	categ, ErrCateg := Category()
	if ErrCateg != nil {
		panic(ErrCateg)
	}
	// Сформировать Слайс категорий из входного результа ответа по всем категориям с сайта
	categs := CategoryBasesForming(categ)

	// Получаем спимок ID товаров
	prods, ErrSKUs := SKUs(categs[0].ID)
	if ErrSKUs != nil {
		panic(ErrSKUs)
	}

	// Получаем данные по артикулам(id)
	line, ErrLines := Lines(prods.ProductIds)
	if ErrLines != nil {
		panic(ErrLines)
	}

	// Создаём внутренний слайс товаров
	Products := Line2Product2(line, categs[0].Cat)

	ID, _ := strconv.Atoi(Products[0].Article)
	touch, ErrToucher := Toucher(ID)
	if ErrToucher != nil {
		fmt.Println(Products)
	}
	ProductTouch := Touch2Product2(Products[0], touch)

	fmt.Printf("$+$ %+v\n", ProductTouch)
}
