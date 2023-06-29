package hmapp

import (
	"fmt"
	"math"
	"strings"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
	"github.com/RB-PRO/SanctionedClothing/pkg/hm"
	"github.com/RB-PRO/SanctionedClothing/pkg/imgbb"
	"github.com/RB-PRO/SanctionedClothing/pkg/transrb"
	"github.com/RB-PRO/SanctionedClothing/pkg/wcprod"
	"github.com/cheggaaa/pb"
	ikurl "github.com/imagekit-developer/imagekit-go/url"
)

// Начать парсить и одновременно загружать товары
func Start() {

	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	fmt.Println("Курс лиры к рублю -", cb.Data.Valute.Try.Value/10)

	// Загружаем товары на WC
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		panic(errorInitWcAdd)
	}

	// Получить слайс категорий
	Categorys, ErrorCategorys := hm.Categorys()
	if ErrorCategorys != nil {
		panic(ErrorCategorys)
	}

	// Получить ядро парсинга эмулятора
	var products []bases.Product2
	core, ErrNewParsingCard := hm.NewParsingCard()
	if ErrNewParsingCard != nil {
		panic(ErrNewParsingCard)
	}

	// Парсинг всех товаров
	BarCategory := pb.StartNew(len(Categorys))
	BarCategory.Prefix("Парсинг категорий")
	for _, categ := range Categorys {
		// Получить ссылку на все товары json
		LineUrl, ErrLineUrl := core.LineUrl(categ.Link)
		if ErrLineUrl != nil {
			continue
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
	varSa.SaveXlsxCsvs("H&M_Products")

	///////////////////////////////////////

	// Парсинг товаров и последующая загрузка на CkikShop
	delivery := 500 // Доставка
	walrus := 1.3   // Моржа
	var ErrorTranstate error
	for i := range products {
		fmt.Printf("Парсинг вариаций товаров (%d/%d)", i+1, len(products))
		AddingProduct := products[i]
		for j := range AddingProduct.Item {
			core.VariableProduct3(&AddingProduct, j)
		}

		// Проверка загружен ли этот товар или нет
		if _, ok := Adding.AllProdSKU[products[i].Article]; !ok {

			AddingProduct.Size = bases.EditProdSize(AddingProduct)

			// Редактирование цены
			AddingProduct = EditCoast(products[i], cb.Data.Valute.Try.Value/10, walrus, delivery)

			// Перевести название, цвета и описание с турецкого на Русский
			AddingProduct, ErrorTranstate = Adding.YandexTranslate(AddingProduct)
			if ErrorTranstate != nil {
				Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
				AddingProduct, _ = Adding.YandexTranslate(AddingProduct)
			}

			// Добавить товар на сайт ClikShop
			errorAddProductWC := Adding.AddProduct(AddingProduct)
			if errorAddProductWC == nil {
				products[i].Upload = true
			}
			Adding.AllProdSKU[AddingProduct.Article] = true
		}
	}

	// // Загружаем товары
	// delivery := 500 // Доставка
	// walrus := 1.3   // Моржа
	// for i := 0; i < len(varient.Product)-2; i++ {
	// 	fmt.Printf("Start: Загружаю товар (%d/%d)", i, len(varient.Product)-2)
	// 	if !varient.Product[i].Upload {
	// 		if _, ok := Adding.AllProdSKU[varient.Product[i].Article]; !ok {
	// 			// Формирование адекватной цены доставки из файла
	// 			ActualDelivery := Adding.EditDelivery(varient.Product[i].Cat, delivery)
	// 			varient.Product[i] = EditCoast(varient.Product[i], cb.Data.Valute.Try.Value/10, walrus, ActualDelivery)
	// 			//errorAddProductWC := Adding.AddProduct(wcprod.ProductTranslate(varient.Product[i])) //.AddAttr()
	// 			var ErrorTranstate error
	// 			varient.Product[i], ErrorTranstate = Adding.YandexTranslate(varient.Product[i])
	// 			if ErrorTranstate != nil {
	// 				Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
	// 				varient.Product[i], _ = Adding.YandexTranslate(varient.Product[i])
	// 			}
	// 			errorAddProductWC := Adding.AddProduct(varient.Product[i]) //.AddAttr()
	// 			if errorAddProductWC != nil {
	// 				varient.Product[i].Upload = true
	// 			}
	// 			Adding.AllProdSKU[varient.Product[i].Article] = true
	// 		}
	// 	}
	// }

	bases.ExitSoft() // "Мягкий" выход из программы
}

// Редактирование цены по товарам
func EditCoast(prod bases.Product2, usd float64, walrus float64, delivery int) bases.Product2 {
	for indexKey := range prod.Item {
		// Корректируем данные
		// Курс доллара * цена в долларах * наценка + цена доставки
		price := usd*prod.Item[indexKey].Price*walrus + float64(delivery)
		price = EditDecadense(price)
		prod.Item[indexKey].Price = price
	}
	return prod
}

// Редактирование цены в большую сторону
//
// # Округляем цену в большую сторону по десяткам
//
// Если цена была 5225.77, то станет 5230
func EditDecadense(coast float64) float64 {
	return math.Round(coast/10.0) * 10.0
}

// Отличтка слайса от пустых элементов
func ClearSlise(strs []string) (output []string) {
	for i := range strs {
		if strs[i] != "" {
			output = append(output, strs[i])
		}
	}
	return output
}
func Start2() {
	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	fmt.Println("Курс лиры к рублю -", cb.Data.Valute.Try.Value/10)

	// Загружаем товары на WC
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		panic(errorInitWcAdd)
	}

	// Получить слайс категорий
	Categorys, ErrorCategorys := hm.Categorys()
	if ErrorCategorys != nil {
		panic(ErrorCategorys)
	}

	// Получить ядро парсинга эмулятора
	// var products []bases.Product2
	core, ErrNewParsingCard := hm.NewParsingCard()
	if ErrNewParsingCard != nil {
		panic(ErrNewParsingCard)
	}

	// Тут якобы начало цикла
	// categ := Categorys[0]

	for _, categ := range Categorys {

		// Получить ссылку на все товары json
		LineUrl, ErrLineUrl := core.LineUrl(categ.Link)
		if ErrLineUrl != nil {
			panic(ErrLineUrl)
		}

		// // Получить к-во товаров в категории
		// cout, ErrorCount := hm.LinesCount(LineUrl)
		// if ErrorCount != nil {
		// 	panic(ErrorCount)
		// }

		// Получить все товары
		line, ErrLine := hm.Lines(LineUrl, 10)
		if ErrLine != nil {
			panic(ErrLine)
		}
		// Перевести полученый ответ от сервера в слайс Product2 и добавить в него соответствующие данные по каждому товару в зависимоти от категории,
		// а именно: Гендер, Каталог.
		AddProducts := hm.Line2Product2(line, categ.Cat, categ.GendetTag)

		///////////////////////////////////////
		delivery := 500 // Доставка
		walrus := 1.3   // Моржа
		var ErrorTranstate error
		AddingProduct := AddProducts[0]
		for j := range AddingProduct.Item {
			core.VariableProduct3(&AddingProduct, j)
		}
		// Проверка загружен ли этот товар или нет
		if _, ok := Adding.AllProdSKU[AddingProduct.Article]; !ok {

			AddingProduct.Size = bases.EditProdSize(AddingProduct)
			for i_image := range AddingProduct.Item {
				var ImagesNewLinks []string
				for j_image := range AddingProduct.Item[i_image].Image {
					var OutputImage string = AddingProduct.Item[i_image].Image[j_image]
					OutputImage = strings.ReplaceAll(OutputImage, "&call=url[file:/product/fullscreen]", "&call=url%5Bfile%3A%2Fproduct%2Ffullscreen%5D")
					OutputImage = strings.ReplaceAll(OutputImage, "&call=url[file:/product/fullscreen", "&call=url%5Bfile%3A%2Fproduct%2Ffullscreen%5D")

					//
					if strings.Contains(OutputImage, "set=format%5Bwebp%5D") {
						// Загрузить локальную картинку
						FilePathWebp := "tmp/local.webp"
						FilePathJpg := strings.ReplaceAll(FilePathWebp, ".webp", ".jpg")
						ErrDownload := imgbb.DownloadFile(FilePathWebp, OutputImage)
						if ErrDownload != nil {
							fmt.Println("ErrDownload:", ErrDownload)
							continue
						}

						ErrWebp2Jpg := bases.Webp2Jpg(FilePathWebp, FilePathJpg)
						if ErrWebp2Jpg != nil {
							fmt.Println("ErrWebp2Jpg:", ErrWebp2Jpg)
							continue
						}

						// var ErrorRefrash error
						// OutputImage, ErrorRefrash = Adding.UploadFile(FilePathJpg)
						// if ErrorRefrash != nil {
						// 	fmt.Println("ErrorRefrash:", ErrorRefrash)
						// 	continue
						// }

						url, ErrIKurl := Adding.IK.Url(ikurl.UrlParam{
							Src: OutputImage,
						})
						if ErrIKurl != nil {
							fmt.Println("Adding.IK.Url:", ErrIKurl)
							continue
						}
						fmt.Println("IK:", url, OutputImage)
						OutputImage = url
					} else {
						FilePathJpg := "tmp/local.jpg"
						ErrDownload := imgbb.DownloadFile(FilePathJpg, OutputImage)
						if ErrDownload != nil {
							fmt.Println("ErrDownload:", ErrDownload)
							continue
						}

						// var ErrorRefrash error
						// OutputImage, ErrorRefrash = Adding.UploadFile(FilePathJpg)
						// if ErrorRefrash != nil {
						// 	fmt.Println("ErrorRefrash:", ErrorRefrash)
						// 	continue
						// }

						url, ErrIKurl := Adding.IK.Url(ikurl.UrlParam{
							Src: OutputImage,
						})
						if ErrIKurl != nil {
							fmt.Println("Adding.IK.Url:", ErrIKurl)
							continue
						}
						fmt.Println("IK:", url, OutputImage)
						OutputImage = url

					}
					ImagesNewLinks = append(ImagesNewLinks, OutputImage)
				}
				AddingProduct.Item[i_image].Image = ImagesNewLinks
			}
			// AddingProduct.Item[0].Image = []string{AddingProduct.Item[0].Image[0]}
			// AddingProduct.Item[1].Image = []string{AddingProduct.Item[0].Image[0]}

			// Редактирование цены
			AddingProduct = EditCoast(AddingProduct, cb.Data.Valute.Try.Value/10, walrus, delivery)

			// Перевести название, цвета и описание с турецкого на Русский
			AddingProduct, ErrorTranstate = Adding.YandexTranslatePart(AddingProduct)
			if ErrorTranstate != nil {
				Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
				AddingProduct, _ = Adding.YandexTranslatePart(AddingProduct)
			}

			fmt.Println(bases.ProdStr(AddingProduct))

			// Добавить товар на сайт ClikShop
			errorAddProductWC := Adding.AddProduct(AddingProduct)
			if errorAddProductWC != nil {
				panic(errorAddProductWC)
			} else {
				AddingProduct.Upload = true
			}
			Adding.AllProdSKU[AddingProduct.Article] = true
		} else {
			fmt.Println("Этот товар уже существует")
		}
	}

}
