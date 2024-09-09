package hmapp

import (
	"fmt"
	"strconv"

	"ClikShop/common/bases"
	"ClikShop/common/gol"
	"ClikShop/common/hm"
	"ClikShop/common/transrb"
	"ClikShop/common/wcprod"
	"github.com/cheggaaa/pb"
)

func Parse() {

	gol, _ := gol.NewGol("logs/")

	// Создать оьбъект переводчика
	Adding, ErrNewTranslate := wcprod.NewTranslate()
	if ErrNewTranslate != nil {
		panic(ErrNewTranslate)
	}

	// Получить слайс категорий
	Categorys, ErrorCategorys := hm.Categorys()
	if ErrorCategorys != nil {
		panic(ErrorCategorys)
	}
	gol.Info("wcprod.New: Получен слайс категорий")

	// Получить ядро парсинга эмулятора
	var products []bases.Product2

	// Парсинг всех товаров
	BarCategory := pb.StartNew(len(Categorys))
	BarCategory.Prefix("Парсинг категорий")
	for _, categ := range Categorys {

		// Получить ссылку на все товары json
		LineUrl, ErrLineUrl := hm.LineUrl2(categ.Link)
		if ErrLineUrl != nil {
			gol.Err(ErrLineUrl)
			panic(ErrLineUrl)
		}
		if LineUrl == "" {
			gol.Err("LineUrl: Nil output")
			panic("LineUrl: Nil output")
		}

		// Получить к-во товаров в категории
		cout, ErrorCount := hm.LinesCount(LineUrl)
		if ErrorCount != nil {
			gol.Warn("LinesCount:", hm.URL+categ.Link, hm.URL+LineUrl)
			fmt.Println("LinesCount:", hm.URL+categ.Link, hm.URL+LineUrl)
		}

		// Получить все товары
		line, _ := hm.Lines(LineUrl, cout)

		// Перевести полученый ответ от сервера в слайс Product2 и добавить в него соответствующие данные по каждому товару
		// в зависимоти от категории, а именно: Гендер, Каталог.
		AddProducts := hm.Line2Product2(line, categ.Cat, categ.GendetTag)

		// Добавляем полученный слайс с товарами в общий слайс товаров
		products = append(products, AddProducts...)
		BarCategory.Increment()
		// break
	}
	BarCategory.Finish()
	gol.Info("Line: Done")

	// Сохранить товары в файл XLSX
	varSa := bases.Variety2{Product: products}
	varSa.SaveXlsxCsvs("tmp/" + "HM_ALL")
	gol.Info("SaveXlsxCsvs: Сохраняю результат line")
	// products = products[:10]

	// Парсинг по подслайсами с размером size
	size := 1000
	BarProducts := pb.StartNew(len(products))
	var SubSlice_j, cout int
	for SubSlice_i := 0; SubSlice_i < len(products); SubSlice_i += size {
		SubSlice_j += size
		if SubSlice_j > len(products) {
			SubSlice_j = len(products)
		}

		// Подслайс. Работаем именно с подслайсами, чтобы не перегружать оперативку
		SubSlice := products[SubSlice_i:SubSlice_j]

		BarProducts.Prefix(strconv.Itoa(cout))
		for i := range SubSlice {
			// Парсинг всех подпродуктов
			AddingProduct := SubSlice[i]
			// gol.Info("Parsing: ", AddingProduct.Link)

			// Размеры и картинки
			var ErrorParseProduct error
			AddingProduct, ErrorParseProduct = hm.VariableProduct2(AddingProduct)
			if ErrorParseProduct != nil {
				gol.Err("Parsing: VariableProduct2:", ErrorParseProduct)
				panic(ErrorParseProduct)
			}

			// Данные по рамерам
			var ErrAvailabilityProduct error
			AddingProduct, ErrAvailabilityProduct = hm.AvailabilityProduct(AddingProduct)
			if ErrAvailabilityProduct != nil {
				gol.Err("Parsing: AvailabilityProduct:", ErrAvailabilityProduct)
				panic(ErrAvailabilityProduct)
			}

			// Описание товара
			var ErrVariableDescription2 error
			AddingProduct, ErrVariableDescription2 = hm.VariableDescription2(AddingProduct)
			if ErrVariableDescription2 != nil {
				gol.Err("Parsing: VariableDescription2:", ErrVariableDescription2)
				panic(ErrVariableDescription2)
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

			SubSlice[i] = AddingProduct

			gol.Info("Parsing: Done")
			BarProducts.Increment()
		}
		cout++
		// bases.Variety2{Product: SubSlice}.SaveXlsxCsvs(fmt.Sprintf("tmp/HM_SubSlice_%d_%d-%d", cout, SubSlice_i, SubSlice_i+size))
		bases.Variety2{Product: SubSlice}.SaveJson(fmt.Sprintf("tmp/HM_SubSlice_%d_%d-%d", cout, SubSlice_i, SubSlice_i+size))
	}
	BarProducts.Finish()
}
