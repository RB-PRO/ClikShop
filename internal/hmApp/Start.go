package hmapp

import (
	"context"
	"fmt"
	"log"
	"math"
	"strings"

	"ClikShop/common/bases"
	"ClikShop/common/cbbank"
	"ClikShop/common/hm"
	"ClikShop/common/imgbb"
	"ClikShop/common/transrb"
	"ClikShop/common/wcprod"
	"github.com/cheggaaa/pb"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/playwright-community/playwright-go"
)

// Начать парсить и одновременно загружать товары
func Start() {
	err := playwright.Install()
	if err != nil {
		panic(err)
	}

	// Создаём объект ядра парсинга, который включает в себя все необходимые функции
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		panic(errorInitWcAdd)
	}
	log.SetOutput(Adding.LogFile)                // set log out put
	log.SetFlags(log.Lshortfile | log.LstdFlags) // optional: log date-time, filename, and line number
	log.Println("wcprod.New: Загрузили ядро")

	// Получение курса валюты
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	log.Println("cbbank: Курс лиры", cb.Data.Valute.Try.Value/10)

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
	varSa.SaveXlsxCsvs("H&M_Products")
	log.Println("Всего товаров", len(products))
	log.Println("SaveXlsxCsvs: Сохраняю результат парсинга")

	///////////////////////////////////////

	// Парсинг товаров и последующая загрузка на CkikShop
	delivery := 500 // Доставка
	walrus := 1.3   // Моржа
	var ErrorTranstate error
	BarProducts := pb.StartNew(len(products))
	BarProducts.Prefix("Парсинг и загрузка товаров")
	for i := range products {
		// Проверка загружен ли этот товар или нет
		if _, ok := Adding.AllProdSKU[products[i].Article]; !ok {

			// Теперь работа будет с переменной AddingProduct, которая содержит добавляемый товар
			AddingProduct := products[i]
			for j := range AddingProduct.Item {
				log.Printf("Парсинг вариаций (%d/%d) товара (%d/%d): %s\n", j+1, len(AddingProduct.Item), i+1, len(products), AddingProduct.Item[j].Link)
				core.VariableProduct3(&AddingProduct, j)
			}

			// Добавить все размеры в товар из всех вариаций товара
			AddingProduct.Size = bases.EditProdSize(AddingProduct)

			// Танцы с бубном над картинками товара
			for i_image := range AddingProduct.Item {
				var ImagesNewLinks []string
				for j_image := range AddingProduct.Item[i_image].Image {
					var OutputImage string = AddingProduct.Item[i_image].Image[j_image]
					var FilePathJpg string
					FilePathWebp := "tmp/local.webp"
					FilePathJpg = strings.ReplaceAll(FilePathWebp, ".webp", ".jpg")
					if strings.Contains(OutputImage, "set=format%5Bwebp%5D") {
						// Загрузить локальную картинку
						ErrDownload := imgbb.DownloadFile(FilePathWebp, OutputImage)
						if ErrDownload != nil {
							log.Println("ErrDownload:", ErrDownload)
							continue
						}

						ErrWebp2Jpg := bases.Webp2Jpg(FilePathWebp, FilePathJpg)
						if ErrWebp2Jpg != nil {
							log.Println("ErrWebp2Jpg:", ErrWebp2Jpg)
							continue
						}

					} else {
						ErrDownload := imgbb.DownloadFile(FilePathJpg, OutputImage)
						if ErrDownload != nil {
							log.Println("ErrDownload:", ErrDownload)
							continue
						}
					}

					base64, ErrBase64 := wcprod.PicToBase64(FilePathJpg)
					if ErrBase64 != nil {
						fmt.Println("ErrBase64", ErrBase64)
						log.Println("PicToBase64:", ErrBase64)
						continue
					}

					// Загрузить картинку
					res, ErrUpload := Adding.IK.Uploader.Upload(context.Background(), base64, uploader.UploadParam{FileName: "RB_PRO.jpg"})
					if ErrUpload != nil {
						fmt.Println("ErrUpload", ErrUpload)
						log.Println("IK.Uploader.Upload:", ErrUpload)
						continue
					}

					OutputImage = res.Data.Url
					// fmt.Println(OutputImage)

					// Загрузить картинку на сервис imgbb
					// var ErrorRefrash error
					// OutputImage, ErrorRefrash = Adding.UploadFile(FilePathJpg)
					// if ErrorRefrash != nil {
					// 	fmt.Println("ErrorRefrash:", ErrorRefrash)
					// 	continue
					// }
					ImagesNewLinks = append(ImagesNewLinks, OutputImage)
				}
				AddingProduct.Item[i_image].Image = ImagesNewLinks
			}

			// Редактирование цены
			AddingProduct = EditCoast(AddingProduct, cb.Data.Valute.Try.Value/10, walrus, delivery)

			// Перевести название, цвета и описание с турецкого на Русский
			AddingProduct, ErrorTranstate = Adding.YandexTranslate(AddingProduct)
			if ErrorTranstate != nil {
				Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
				AddingProduct, _ = Adding.YandexTranslate(AddingProduct)
			}

			// Добавить товар на сайт ClikShop
			errorAddProductWC := Adding.AddProduct(AddingProduct)
			if errorAddProductWC != nil {
				fmt.Println("Adding.AddProduct:", errorAddProductWC)
			} else {
				products[0].Upload = true
			}

			Adding.AllProdSKU[AddingProduct.Article] = true
		} else {
			log.Println("Adding.AddProduct: Товар с артикулом", products[i].Article)
		}
		BarProducts.Increment()
	}
	BarProducts.Finish()

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
