package actualizer

import (
	"fmt"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/hm"
	"github.com/cheggaaa/pb"
)

// Структура HM для парсинга
type HM struct {
	*bitrixActualizer
}

func NewHM(bx *bitrixActualizer) *HM {
	return &HM{bx}
}

// Парсинг данных и сохранение их в файлы
//
//	Заменить во всех файлах нужно символы '\u0026' на '&'
func (bx *HM) screper() (string, error) {
	folder := "hm"
	ReMakeDir(folder)

	// Получить слайс категорий
	Categorys, ErrorCategorys := hm.Categorys()
	if ErrorCategorys != nil {
		panic(ErrorCategorys)
	}
	bx.GLOG.Info("hm.Categorys: Получен слайс категорий")

	var count int
	for icateg, categ := range Categorys {
		// if icateg < 28 {
		// 	continue
		// }

		// Получить ссылку на все товары json
		LineUrl, ErrLineUrl := hm.LineUrl2(categ.Link)
		if ErrLineUrl != nil {
			bx.GLOG.Err(ErrLineUrl)
			panic(ErrLineUrl)
		}
		if LineUrl == "" {
			// bx.GLOG.Err("LineUrl: Nil output")
			panic("LineUrl: Nil output")
		}

		// Получить к-во товаров в категории
		cout, ErrorCount := hm.LinesCount(LineUrl)
		if ErrorCount != nil {
			// bx.GLOG.Warn("LinesCount:", hm.URL+categ.Link, hm.URL+LineUrl)
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

		BarProducts := pb.StartNew(len(SubSlice))
		BarProducts.Prefix(fmt.Sprintf("[%d/%d]", icateg+1, len(Categorys)))
		for i := range SubSlice {
			// Парсинг всех подпродуктов
			AddingProduct := SubSlice[i]
			// gol.Info("Parsing: ", AddingProduct.Link)

			// Размеры и картинки
			var ErrorParseProduct error
			AddingProduct, ErrorParseProduct = hm.VariableProduct2(AddingProduct)
			if ErrorParseProduct != nil {
				bx.GLOG.Err("VariableProduct2:", ErrorParseProduct)
				continue
				// panic(ErrorParseProduct)
			}

			// Данные по рамерам
			var ErrAvailabilityProduct error
			AddingProduct, ErrAvailabilityProduct = hm.AvailabilityProduct(AddingProduct)
			if ErrAvailabilityProduct != nil {
				bx.GLOG.Err("AvailabilityProduct:", ErrAvailabilityProduct)
				continue
				// panic(ErrAvailabilityProduct)
			}

			// Описание товара
			var ErrVariableDescription2 error
			AddingProduct, ErrVariableDescription2 = hm.VariableDescription2(AddingProduct)
			if ErrVariableDescription2 != nil {
				bx.GLOG.Err("VariableDescription2:", ErrVariableDescription2)
				continue
				// panic(ErrVariableDescription2)
			}

			// Добавить все размеры в товар из всех вариаций товара
			AddingProduct.Size = bases.EditProdSize(AddingProduct)
			AddingProduct.Img = bases.EditIMG(AddingProduct)
			AddingProduct.Cat = ProdTranslateCat

			SubSlice[i] = AddingProduct

			BarProducts.Increment()
			count++
		}
		bases.Variety2{Product: SubSlice}.SaveJson(fmt.Sprintf("%s/hm_%d_%s",
			folder, icateg, categ.Cat[len(categ.Cat)-1].Slug))
		BarProducts.Finish()
	}
	return folder, nil
}
