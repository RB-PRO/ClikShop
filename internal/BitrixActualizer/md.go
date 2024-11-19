package actualizer

import (
	"ClikShop/common/gol"
	"fmt"
	"strconv"

	massimodutti "ClikShop/common/MassimoDutti"
	"ClikShop/common/bases"
	"github.com/cheggaaa/pb"
)

// Структура HM для парсинга
type MD struct {
	*gol.Gol
}

func NewMD() *MD {
	return &MD{
		Gol: gol.NewGol("HM"),
	}
}

// Парсинг данных и сохранение их в файлы
//
//	Заменить во всех файлах нужно символы '\u0026' на '&'
func (bx *MD) Scraper() (string, error) {
	folder := "tmp/md"
	ReMakeDir(folder)

	// Получить все категории
	categ, ErrCateg := massimodutti.Category()
	if ErrCateg != nil {
		panic(ErrCateg)
	}

	// Сформировать Слайс категорий из входного результа ответа по всем категориям с сайта
	categs := massimodutti.CategoryBasesForming(categ)

	// Цикл по всем товарам
	// Формируем слайсы с ID товаров и их категории
	var index int = 1
	categs = categs[1:]
	for icateg, CategoryForSKU := range categs {

		// Получаем спимок ID това
		prods, ErrSKUs := massimodutti.SKUs(CategoryForSKU.ID)
		if ErrSKUs != nil {
			panic(ErrSKUs)
		}

		// Получаем данные по артикулам(id)
		line, ErrLines := massimodutti.Lines(prods.ProductIds)
		if ErrLines != nil {
			panic(ErrLines)
		}

		// Создаём внутренний слайс товаров
		Products := massimodutti.Line2Product2(line, CategoryForSKU.Cat)

		BarProducts := pb.StartNew(len(Products))
		BarProducts.Prefix(fmt.Sprintf("[%d/%d]", icateg+1, len(categs)))
		for i := range Products {

			ID, _ := strconv.Atoi(Products[i].Article)
			touch, ErrToucher := massimodutti.Toucher(ID)
			if ErrToucher != nil {
				fmt.Println(Products)
			}
			Products[i] = massimodutti.Touch2Product2(Products[i], touch)

			// Добавить все размеры в товар из всех вариаций товара
			Products[i].Size = bases.EditProdSize(Products[i])

			Products[i].Img = bases.EditIMG(Products[i])

			BarProducts.Increment()
		}
		bases.Variety2{Product: Products}.SaveJson(fmt.Sprintf("%s/md_%d_%d",
			folder, index, CategoryForSKU.ID))
		BarProducts.Finish()
		index++
	}
	return folder, nil
}
