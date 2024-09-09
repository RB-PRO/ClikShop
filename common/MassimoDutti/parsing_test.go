package massimodutti_test

import (
	massimodutti "ClikShop/common/MassimoDutti"
	"ClikShop/common/config"
	"fmt"
	"strconv"
	"testing"
)

func TestParsing2(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		t.Error(err)
	}
	mdService, err := massimodutti.New(cfg)
	if err != nil {
		t.Error(err)
	}

	// Получить все категории
	categ, err := mdService.Category()
	if err != nil {
		panic(err)
	}
	// Сформировать Слайс категорий из входного результа ответа по всем категориям с сайта
	categs := massimodutti.CategoryBasesForming(categ)

	// Получаем спимок ID товаров
	prods, err := mdService.SKUs(categs[0].ID)
	if err != nil {
		panic(err)
	}

	// Получаем данные по артикулам(id)
	line, err := mdService.Lines(prods.ProductIds)
	if err != nil {
		panic(err)
	}

	// Создаём внутренний слайс товаров
	Products := massimodutti.Line2Product2(line, categs[0].Cat)

	ID, _ := strconv.Atoi(Products[0].Article)
	touch, ErrToucher := mdService.Toucher(ID)
	if ErrToucher != nil {
		fmt.Println(Products)
	}
	ProductTouch := massimodutti.Touch2Product2(Products[0], touch)

	fmt.Printf("$+$ %+v\n", ProductTouch)
}
