package hmapp

import (
	"fmt"
	"log"
	"strconv"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/hm"
	"github.com/cheggaaa/pb"
	"github.com/playwright-community/playwright-go"
)

func Parsing() {

	err := playwright.Install()
	if err != nil {
		panic(err)
	}

	// Получить слайс категорий
	Categorys, ErrorCategorys := hm.Categorys()
	if ErrorCategorys != nil {
		panic(ErrorCategorys)
	}
	log.Println("wcprod.New: Получен слайс категорий")

	// Получить ядро парсинга эмулятора
	var products []bases.Product2
	core, ErrNewParsingCard := hm.NewParsingCard()
	if ErrNewParsingCard != nil {
		panic(ErrNewParsingCard)
	}
	defer core.Stop()
	log.Println("wcprod.New: Получено ядро парсинга эмулятора")

	// Парсинг всех товаров
	BarCategory := pb.StartNew(len(Categorys))
	BarCategory.Prefix("Парсинг категорий")
	for _, categ := range Categorys {

		// Получить ссылку на все товары json
		LineUrl, ErrLineUrl := core.LineUrl(categ.Link)
		if ErrLineUrl != nil {
			continue
		}
		if LineUrl == "" {
			panic("LineUrl: Nil output")
		}

		// Получить к-во товаров в категории
		cout, ErrorCount := hm.LinesCount(LineUrl)
		if ErrorCount != nil {
			continue
		}

		// Получить все товары
		line, ErrLine := hm.Lines(LineUrl, cout)
		if ErrLine != nil {
			continue
		}

		// Перевести полученый ответ от сервера в слайс Product2 и добавить в него соответствующие данные по каждому товару в зависимоти от категории,
		// а именно: Гендер, Каталог.
		AddProducts := hm.Line2Product2(line, categ.Cat, categ.GendetTag)

		// Добавляем полученный слайс с товарами в общий слайс товаров
		products = append(products, AddProducts...)
		BarCategory.Increment()
	}
	BarCategory.Finish()

	// Сохранить товары в файл XLSX
	varSa := bases.Variety2{Product: products}
	varSa.SaveXlsxCsvs("H&M_Parsing_1")
	log.Println("SaveXlsxCsvs: Сохраняю результат парсинга")

	// Парсинг товаров

	// Парсинг по подслайсами с размером size
	size := 500
	var SubSlice_j, cout int
	for SubSlice_i := 0; SubSlice_i < len(products); SubSlice_i += size {
		SubSlice_j += size
		if SubSlice_j > len(products) {
			SubSlice_j = len(products)
		}
		// do what do you want to with the sub-slice, here just printing the sub-slices
		fmt.Println()

		// Подслайс. Работаем именно с подслайсами, чтобы не перегружать оперативку
		SubSlice := products[SubSlice_i:SubSlice_j]

		BarProducts := pb.StartNew(len(products))
		BarProducts.Prefix(strconv.Itoa(cout))
		for i := range SubSlice {

			// Парсинг всех подпродуктов
			AddingProduct := SubSlice[i]
			for j := range AddingProduct.Item {
				// log.Printf("Парсинг вариаций (%d/%d) товара (%d/%d): %s\n", j+1, len(AddingProduct.Item), i+1, len(products), AddingProduct.Item[j].Link)
				core.VariableProduct3(&AddingProduct, j)
			}

			// Добавить все размеры в товар из всех вариаций товара
			AddingProduct.Size = bases.EditProdSize(AddingProduct)

			BarProducts.Increment()
		}
		BarProducts.Finish()
		cout++
		bases.Variety2{Product: SubSlice}.SaveXlsxCsvs(fmt.Sprintf("H&M_SubSlice_%d_%d-%d", cout, SubSlice_i, SubSlice_i+size))

		// varSa.SaveXlsxCsvs("H&M_Parsing_1")
	}
}
