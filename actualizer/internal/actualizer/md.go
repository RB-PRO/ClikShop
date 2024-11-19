package actualizer

import (
	"ClikShop/common/gol"
	"fmt"
	"strconv"

	md "ClikShop/common/MassimoDutti"
	"ClikShop/common/bases"
	"github.com/cheggaaa/pb"
)

// Структура HM для парсинга
type MD struct {
	service *md.Service
	*gol.Gol
}

func NewMD(service *md.Service) *MD {
	return &MD{
		service: service,
		Gol:     gol.NewGol("HM"),
	}
}

// Парсинг данных и сохранение их в файлы
//
//	Заменить во всех файлах нужно символы '\u0026' на '&'
func (s *MD) Scraper() (string, error) {
	folder := "tmp/md"
	ReMakeDir(folder)

	// Получить все категории
	categ, ErrCateg := s.service.Category()
	if ErrCateg != nil {
		panic(ErrCateg)
	}

	// Сформировать Слайс категорий из входного результа ответа по всем категориям с сайта
	categs := md.CategoryBasesForming(categ)

	// Цикл по всем товарам
	// Формируем слайсы с ID товаров и их категории
	var index int = 1
	categs = categs[1:]
	for icateg, CategoryForSKU := range categs {

		// Получаем спимок ID това
		prods, ErrSKUs := s.service.SKUs(CategoryForSKU.ID)
		if ErrSKUs != nil {
			panic(ErrSKUs)
		}

		// Получаем данные по артикулам(id)
		line, ErrLines := s.service.Lines(prods.ProductIds)
		if ErrLines != nil {
			//panic(ErrLines)
			fmt.Println(ErrLines)
		}

		// Создаём внутренний слайс товаров
		Products := md.Line2Product2(line, CategoryForSKU.Cat)

		BarProducts := pb.StartNew(len(Products))
		BarProducts.Prefix(fmt.Sprintf("[%d/%d]", icateg+1, len(categs)))
		for i := range Products {

			ID, _ := strconv.Atoi(Products[i].Article)
			touch, ErrToucher := s.service.Toucher(ID)
			if ErrToucher != nil {
				fmt.Println(Products)
			}
			Products[i] = md.Touch2Product2(Products[i], touch)

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
