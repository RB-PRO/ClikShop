package massimoduttiapp

import (
	"fmt"
	"log"
	"strconv"

	massimodutti "ClikShop/common/MassimoDutti"
	"ClikShop/common/bases"
	"ClikShop/common/cbbank"
	"ClikShop/common/transrb"
	"ClikShop/common/wcprod"
	"github.com/cheggaaa/pb"
)

func Parsing() {

	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	log.Println("Курс лиры", cb.Data.Valute.Try.Value/10)

	// Создать оьбъект переводчика
	Translate, ErrNewTranslate := wcprod.NewTranslate()
	if ErrNewTranslate != nil {
		panic(ErrNewTranslate)
	}

	// Получить все категории
	categ, ErrCateg := massimodutti.Category()
	if ErrCateg != nil {
		panic(ErrCateg)
	}

	// Сформировать Слайс категорий из входного результа ответа по всем категориям с сайта
	categs := massimodutti.CategoryBasesForming(categ)

	// Цикл по всем товарам
	// Формируем слайсы с ID товаров и их категории
	var Products []bases.Product2
	BarCategory := pb.StartNew(len(categs))
	BarCategory.Prefix("Парсинг всех артикулов")
	for _, CategoryForSKU := range categs {
		// fmt.Println(i, "/", len(categs))

		// Получаем спимок ID товаров
		prods, ErrSKUs := massimodutti.SKUs(CategoryForSKU.ID)
		if ErrSKUs != nil {
			panic(ErrSKUs)
		}

		// Формируем подслайсы для каждой категории
		sku_slice := make([]int, 0, len(prods.ProductIds))
		for _, idSKU := range prods.ProductIds {
			sku_slice = append(sku_slice, idSKU)
		}

		// Получаем данные по артикулам(id)
		line, ErrLines := massimodutti.Lines(sku_slice)
		if ErrLines != nil {
			panic(ErrLines)
		}

		// Создаём внутренний слайс товаров
		AddingProducts := massimodutti.Line2Product2(line, CategoryForSKU.Cat)

		// Добавляем к итоговому слайсу товаров
		Products = append(Products, AddingProducts...)

		BarCategory.Increment()
	}
	BarCategory.Finish()

	fmt.Println("Всего товаров -", len(Products))

	// ***************************************
	// Парсинг по подслайсами с размером size
	size := 300
	BarProducts := pb.StartNew(len(Products))
	var SubSlice_j, cout int
	for SubSlice_i := 0; SubSlice_i < len(Products); SubSlice_i += size {
		SubSlice_j += size
		if SubSlice_j > len(Products) {
			SubSlice_j = len(Products)
		}

		// Подслайс. Работаем именно с подслайсами, чтобы не перегружать оперативку
		SubSlice := Products[SubSlice_i:SubSlice_j]
		BarProducts.Prefix(strconv.Itoa(cout))
		for i := range SubSlice {
			// Парсинг всех подпродуктов
			AddingProduct := SubSlice[i]

			ID, _ := strconv.Atoi(AddingProduct.Article)
			touch, ErrToucher := massimodutti.Toucher(ID)
			if ErrToucher != nil {
				fmt.Println(Products)
			}
			AddingProduct = massimodutti.Touch2Product2(AddingProduct, touch)

			// Name := AddingProduct.Name
			// Перевести товар
			var ErrorTranstate error
			AddingProduct, ErrorTranstate = Translate.YandexTranslatePart(AddingProduct)
			if ErrorTranstate != nil {
				Translate.Tr, _ = transrb.New(Translate.Tr.FolderID, Translate.Tr.OAuthToken)
				AddingProduct, _ = Translate.YandexTranslatePart(AddingProduct)
			}
			// AddingProduct.Name = Name

			// Добавить все размеры в товар из всех вариаций товара
			AddingProduct.Size = bases.EditProdSize(AddingProduct)

			AddingProduct.Img = bases.EditIMG(AddingProduct)

			// Редактирование цены
			AddingProduct = bases.EditCoast(AddingProduct, cb.Data.Valute.Try.Value/10, 1.3, 500)

			SubSlice[i] = AddingProduct

			BarProducts.Increment()
		}
		cout++
		// bases.Variety2{Product: SubSlice}.SaveXlsxCsvs(fmt.Sprintf("tmp/H&M_SubSlice_%d_%d-%d", cout, SubSlice_i, SubSlice_i+size))
		bases.Variety2{Product: SubSlice}.SaveJson(fmt.Sprintf("tmp/MD/MD_%d_%d-%d", cout, SubSlice_i, SubSlice_i+size))
	}
	BarProducts.Finish()
	bases.ExitSoft()
}
