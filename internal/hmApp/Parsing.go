package hmapp

import (
	"fmt"
	"log"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/cbbank"
	"github.com/RB-PRO/ClikShop/pkg/gol"
	"github.com/RB-PRO/ClikShop/pkg/hm"
	"github.com/RB-PRO/ClikShop/pkg/transrb"
	"github.com/RB-PRO/ClikShop/pkg/wcprod"
	"github.com/cheggaaa/pb"
)

// Парсинг данных и сохранение их в файлы
//
//	Заменить во всех файлах нужно символы '\u0026' на '&'
func Parsing() {
	glog, _ := gol.NewGol("logs/")

	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	glog.Info("Курс лиры", cb.Data.Valute.Try.Value/10)

	// Создать оьбъект переводчика
	Adding, ErrNewTranslate := wcprod.NewTranslate()
	if ErrNewTranslate != nil {
		panic(ErrNewTranslate)
	}
	glog.Info("Загрузил переводчик", Adding.Tr.OAuthToken)

	// Получить слайс категорий
	Categorys, ErrorCategorys := hm.Categorys()
	if ErrorCategorys != nil {
		panic(ErrorCategorys)
	}
	glog.Info("wcprod.New: Получен слайс категорий")

	var count int
	for icateg, categ := range Categorys {
		if icateg < 28 {
			continue
		}

		// Получить ссылку на все товары json
		LineUrl, ErrLineUrl := hm.LineUrl2(categ.Link)
		if ErrLineUrl != nil {
			glog.Err(ErrLineUrl)
			panic(ErrLineUrl)
		}
		if LineUrl == "" {
			glog.Err("LineUrl: Nil output")
			panic("LineUrl: Nil output")
		}

		// Получить к-во товаров в категории
		cout, ErrorCount := hm.LinesCount(LineUrl)
		if ErrorCount != nil {
			glog.Warn("LinesCount:", hm.URL+categ.Link, hm.URL+LineUrl)
			fmt.Println("LinesCount:", hm.URL+categ.Link, hm.URL+LineUrl)
		}

		// Получить все товары
		line, _ := hm.Lines(LineUrl, cout)

		// Перевести полученый ответ от сервера в слайс Product2 и добавить в него соответствующие данные по каждому товару
		// в зависимоти от категории, а именно: Гендер, Каталог.
		SubSlice := hm.Line2Product2(line, categ.Cat, categ.GendetTag)

		// Переведённая категория
		if len(SubSlice) == 0 {
			continue
		}
		ProdTranslateCat := SubSlice[0].Cat
		ProdTranslateCat, _ = Adding.YandexCat(ProdTranslateCat)

		BarProducts := pb.StartNew(len(SubSlice))
		BarProducts.Prefix(fmt.Sprintf("[%d/%d]", icateg, len(Categorys)))
		for i := range SubSlice {
			// Парсинг всех подпродуктов
			AddingProduct := SubSlice[i]
			// gol.Info("Parsing: ", AddingProduct.Link)

			// Размеры и картинки
			var ErrorParseProduct error
			AddingProduct, ErrorParseProduct = hm.VariableProduct2(AddingProduct)
			if ErrorParseProduct != nil {
				glog.Err("Parsing: VariableProduct2:", ErrorParseProduct)
				continue
				//panic(ErrorParseProduct)
			}

			// Данные по рамерам
			var ErrAvailabilityProduct error
			AddingProduct, ErrAvailabilityProduct = hm.AvailabilityProduct(AddingProduct)
			if ErrAvailabilityProduct != nil {
				glog.Err("Parsing: AvailabilityProduct:", ErrAvailabilityProduct)
				continue
				//panic(ErrAvailabilityProduct)
			}

			// Описание товара
			var ErrVariableDescription2 error
			AddingProduct, ErrVariableDescription2 = hm.VariableDescription2(AddingProduct)
			if ErrVariableDescription2 != nil {
				glog.Err("Parsing: VariableDescription2:", ErrVariableDescription2)
				continue
				//panic(ErrVariableDescription2)
			}

			// Перевести товар
			var ErrorTranstate error
			AddingProduct, ErrorTranstate = Adding.YandexTranslate(AddingProduct)
			if ErrorTranstate != nil {
				Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
				AddingProduct, _ = Adding.YandexTranslate(AddingProduct)
			}

			// Добавить все размеры в товар из всех вариаций товара
			AddingProduct.Size = bases.EditProdSize(AddingProduct)
			AddingProduct = bases.EditCoast(AddingProduct, cb.Data.Valute.Try.Value/10, 1.3, 500)
			AddingProduct.Img = bases.EditIMG(AddingProduct)
			AddingProduct.Cat = ProdTranslateCat

			SubSlice[i] = AddingProduct

			BarProducts.Increment()
			count++
		}
		// bases.Variety2{Product: SubSlice}.SaveXlsxCsvs(fmt.Sprintf("tmp/HM_SubSlice_%d_%d-%d", cout, SubSlice_i, SubSlice_i+size))
		bases.Variety2{Product: SubSlice}.SaveJson(fmt.Sprintf("tmp/HM/json/hm_%d_%s", icateg, categ.Cat[len(categ.Cat)-1].Slug))
		BarProducts.Finish()
	}
	log.Println("Всего", count, "товара(ов)")
}
