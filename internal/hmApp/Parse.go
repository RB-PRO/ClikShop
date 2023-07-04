package hmapp

import (
	"fmt"
	"log"
	"strconv"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/hm"
	"github.com/cheggaaa/pb"
)

func Parsing() {

	// err := playwright.Install()
	// if err != nil {
	// 	panic(err)
	// }

	// Получить слайс категорий
	Categorys, ErrorCategorys := hm.Categorys()
	if ErrorCategorys != nil {
		panic(ErrorCategorys)
	}
	log.Println("wcprod.New: Получен слайс категорий")

	// Получить ядро парсинга эмулятора
	var products []bases.Product2

	// Парсинг всех товаров
	BarCategory := pb.StartNew(len(Categorys))
	BarCategory.Prefix("Парсинг категорий")
	for _, categ := range Categorys {

		// Получить ссылку на все товары json
		LineUrl, ErrLineUrl := hm.LineUrl2(categ.Link)
		if ErrLineUrl != nil {
			panic(ErrLineUrl)
		}
		if LineUrl == "" {
			panic("LineUrl: Nil output")
		}

		// Получить к-во товаров в категории
		cout, ErrorCount := hm.LinesCount(LineUrl)
		if ErrorCount != nil {
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
	}
	BarCategory.Finish()

	// Сохранить товары в файл XLSX
	varSa := bases.Variety2{Product: products}
	varSa.SaveXlsxCsvs("tmp/" + "H&M_ALL")
	log.Println("SaveXlsxCsvs: Сохраняю результат парсинга")

	// Парсинг по подслайсами с размером size
	size := 1000
	BarProducts := pb.StartNew(len(products))
	var SubSlice_j, cout int
	for SubSlice_i := 0; SubSlice_i < len(products); SubSlice_i += size {
		SubSlice_j += size
		if SubSlice_j > len(products) {
			SubSlice_j = len(products)
		}
		// do what do you want to with the sub-slice, here just printing the sub-slices

		// Подслайс. Работаем именно с подслайсами, чтобы не перегружать оперативку
		SubSlice := products[SubSlice_i:SubSlice_j]

		BarProducts.Prefix(strconv.Itoa(cout))
		for i := range SubSlice {

			// Парсинг всех подпродуктов
			AddingProduct := SubSlice[i]

			var ErrorParseProduct error
			AddingProduct, ErrorParseProduct = hm.VariableProduct2(AddingProduct)
			if ErrorParseProduct != nil {
				panic(ErrorParseProduct)
			}

			var ErrAvailabilityProduct error
			AddingProduct, ErrAvailabilityProduct = hm.AvailabilityProduct(AddingProduct)
			if ErrAvailabilityProduct != nil {
				panic(ErrAvailabilityProduct)
			}

			// Добавить все размеры в товар из всех вариаций товара
			AddingProduct.Size = bases.EditProdSize(AddingProduct)

			BarProducts.Increment()
		}
		cout++
		bases.Variety2{Product: SubSlice}.SaveXlsxCsvs(fmt.Sprintf("tmp/H&M_SubSlice_%d_%d-%d", cout, SubSlice_i, SubSlice_i+size))
	}
	BarProducts.Finish()
}
