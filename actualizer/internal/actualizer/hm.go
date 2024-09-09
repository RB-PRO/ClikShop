package actualizer

import (
	"ClikShop/common/gol"
	"fmt"
	"strings"

	"ClikShop/common/bases"
	"ClikShop/common/hm"
	"github.com/cheggaaa/pb"
)

// Структура HM для парсинга
type HM struct {
	service *hm.Service
	*gol.Gol
}

func NewHM(serviceHM *hm.Service) *HM {
	return &HM{
		service: serviceHM,
		Gol:     gol.NewGol("HM"),
	}
}

// Парсинг данных и сохранение их в файлы
//
//	Заменить во всех файлах нужно символы '\u0026' на '&'
func (s *HM) Scraper() (string, error) {
	folder := "hm"
	ReMakeDir(folder)

	// Получить слайс категорий
	Categorys, err := s.service.Categorys()
	if err != nil {
		panic(err)
	}
	s.Info("hm.Categorys: Получен слайс категорий")

	for icateg, categ := range Categorys {
		// if icateg < 28 {
		// 	continue
		// }

		if strings.Contains(categ.Link, "view-all") {
			continue
		}

		// Получить ссылку на все товары json
		categoryURLs, err := s.service.LineUrl2(categ.Link, categ.Cat)
		if err != nil {
			s.Err(err)
			panic(err)
		}

		for _, catURL := range categoryURLs {
			if catURL.URL == "" {
				s.Gol.Err("LineUrl: Nil output")
				//panic("LineUrl: Nil output")
				continue
			}

			// Получить к-во товаров в категории
			//count, err := s.service.LinesCount(catURL.URL)
			//if err != nil {
			//	s.Gol.Warn("LinesCount:", hm.URL+categ.Link, hm.URL+catURL.URL)
			//}
			//if count == 0 {
			//	continue
			//}

			// Получить все товары
			// TODO: тут ваще непонятно 5к оставил
			line, _ := s.service.Lines(catURL.URL, 5000)

			// Перевести полученный ответ от сервера в слайс Product2 и добавить в него соответствующие данные по каждому товару
			// в зависимости от категории, а именно: Гендер, Каталог.
			SubSlice := hm.Line2Product2(line, categ.Cat, categ.GendetTag)

			// Переведённая категория
			if len(SubSlice) == 0 {
				continue
			}
			ProdTranslateCat := SubSlice[0].Cat

			BarProducts := pb.StartNew(len(SubSlice))
			BarProducts.Prefix(fmt.Sprintf("[%d/%d]", icateg+1, len(Categorys)))
			for i := range SubSlice {
				// Парсинг всех под-продуктов
				AddingProduct := SubSlice[i]

				// Размеры и картинки
				AddingProduct, err = s.service.VariableProduct2(AddingProduct)
				if err != nil {
					s.Err("VariableProduct2:", err)
					continue
				}

				// Данные по рамерам
				AddingProduct, err = s.service.AvailabilityProduct(AddingProduct)
				if err != nil {
					s.Err("AvailabilityProduct:", err)
					continue
				}

				// Описание товара
				AddingProduct, err = s.service.VariableDescription2(AddingProduct)
				if err != nil {
					s.Err("VariableDescription2:", err)
					continue
				}

				// Добавить все размеры в товар из всех вариаций товара
				AddingProduct.Size = bases.EditProdSize(AddingProduct)
				AddingProduct.Img = bases.EditIMG(AddingProduct)
				AddingProduct.Cat = ProdTranslateCat

				SubSlice[i] = AddingProduct

				BarProducts.Increment()
			}
			_ = bases.Variety2{Product: SubSlice}.SaveJson(fmt.Sprintf("%s/hm_%d_%s",
				folder, icateg, categ.Cat[len(categ.Cat)-1].Slug))
			BarProducts.Finish()
		}

	}
	return folder, nil
}
