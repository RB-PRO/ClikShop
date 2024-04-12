package hmapp_test

import (
	"fmt"
	"strings"
	"testing"

	hmapp "github.com/RB-PRO/ClikShop/internal/hmApp"
	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/cbbank"
	"github.com/RB-PRO/ClikShop/pkg/hm"
	"github.com/RB-PRO/ClikShop/pkg/imgbb"
	"github.com/RB-PRO/ClikShop/pkg/transrb"
	"github.com/RB-PRO/ClikShop/pkg/wcprod"
	ikurl "github.com/imagekit-developer/imagekit-go/url"
)

// go test -v -run ^TestParseAddProduct$ github.com/RB-PRO/ClikShop/internal/hmApp
func TestParseAddProduct(t *testing.T) {
	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		t.Error(ErrorCB)
	}
	fmt.Println("Курс лиры к рублю -", cb.Data.Valute.Try.Value/10)

	// Загружаем товары на WC
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		t.Error(errorInitWcAdd)
	}

	// Получить слайс категорий
	Categorys, ErrorCategorys := hm.Categorys()
	if ErrorCategorys != nil {
		t.Error(ErrorCategorys)
	}
	fmt.Println(len(Categorys))

	// Получить ядро парсинга эмулятора
	var products []bases.Product2
	core, ErrNewParsingCard := hm.NewParsingCard()
	if ErrNewParsingCard != nil {
		t.Error(ErrNewParsingCard)
	}

	// Тут якобы начало цикла
	categ := Categorys[0]

	// Получить ссылку на все товары json
	LineUrl, ErrLineUrl := core.LineUrl(categ.Link)
	if ErrLineUrl != nil {
		t.Error(ErrLineUrl)
	}

	// Получить к-во товаров в категории
	cout, ErrorCount := hm.LinesCount(LineUrl)
	if ErrorCount != nil {
		t.Error(ErrorCount)
	}

	// Получить все товары
	line, ErrLine := hm.Lines(LineUrl, cout)
	if ErrLine != nil {
		t.Error(ErrLine)
	}
	// Перевести полученый ответ от сервера в слайс Product2 и добавить в него соответствующие данные по каждому товару в зависимоти от категории,
	// а именно: Гендер, Каталог.
	AddProducts := hm.Line2Product2(line, categ.Cat, categ.GendetTag)

	// Добавляем полученный слайс с товарами в общий слайс товаров
	products = append(products, AddProducts[8])

	///////////////////////////////////////
	delivery := 500 // Доставка
	walrus := 1.3   // Моржа
	var ErrorTranstate error
	AddingProduct := products[0]
	for j := range AddingProduct.Item {
		core.VariableProduct3(&AddingProduct, j)
	}
	// Проверка загружен ли этот товар или нет
	if _, ok := Adding.AllProdSKU[products[0].Article]; !ok {

		AddingProduct.Size = bases.EditProdSize(AddingProduct)
		for i_image := range AddingProduct.Item {
			var ImagesNewLinks []string
			for j_image := range AddingProduct.Item[i_image].Image {
				var OutputImage string = AddingProduct.Item[i_image].Image[j_image]
				OutputImage = strings.ReplaceAll(OutputImage, "&call=url[file:/product/fullscreen]", "&call=url%5Bfile%3A%2Fproduct%2Ffullscreen%5D")
				OutputImage = strings.ReplaceAll(OutputImage, "&call=url[file:/product/fullscreen", "&call=url%5Bfile%3A%2Fproduct%2Ffullscreen%5D")

				//
				// Загрузить локальную картинку
				FilePathWebp := "tmp/local.webp"
				FilePathJpg := strings.ReplaceAll(FilePathWebp, ".webp", ".jpg")
				ErrDownload := imgbb.DownloadFile(FilePathWebp, OutputImage)
				if ErrDownload != nil {
					fmt.Println("ErrDownload:", ErrDownload)
					continue
				}
				if strings.Contains(OutputImage, "set=format%5Bwebp%5D") {
					ErrWebp2Jpg := bases.Webp2Jpg(FilePathWebp, FilePathJpg)
					if ErrWebp2Jpg != nil {
						fmt.Println("ErrWebp2Jpg:", ErrWebp2Jpg)
						continue
					}
				}
				// var ErrorRefrash error
				// OutputImage, ErrorRefrash = Adding.UploadFile(FilePathJpg)
				// if ErrorRefrash != nil {
				// 	fmt.Println("ErrorRefrash:", ErrorRefrash)
				// 	continue
				// }

				// ctx := context.TODO()
				// meta, rerer := Adding.IK.Metadata.FromUrl(ctx, OutputImage)
				// if rerer != nil {
				// 	fmt.Println("Adding.IK.Url:", rerer)
				// 	continue
				// }
				// OutputImage = meta

				ImagesNewLinks = append(ImagesNewLinks, OutputImage)
			}
			AddingProduct.Item[i_image].Image = ImagesNewLinks
		}
		// AddingProduct.Item[0].Image = []string{AddingProduct.Item[0].Image[0]}
		// AddingProduct.Item[1].Image = []string{AddingProduct.Item[0].Image[0]}

		// Редактирование цены
		AddingProduct = hmapp.EditCoast(AddingProduct, cb.Data.Valute.Try.Value/10, walrus, delivery)

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
			t.Error(errorAddProductWC)
		} else {
			products[0].Upload = true
		}
		Adding.AllProdSKU[AddingProduct.Article] = true
	} else {
		fmt.Println("Этот товар уже существует")
	}
}

//

// go test -v -run ^TestParseAddProductAllCategorys$ github.com/RB-PRO/ClikShop/internal/hmApp
func TestParseAddProductAllCategorys(t *testing.T) {
	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		t.Error(ErrorCB)
	}
	fmt.Println("Курс лиры к рублю -", cb.Data.Valute.Try.Value/10)

	// Загружаем товары на WC
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		t.Error(errorInitWcAdd)
	}

	// Получить слайс категорий
	Categorys, ErrorCategorys := hm.Categorys()
	if ErrorCategorys != nil {
		t.Error(ErrorCategorys)
	}

	// Получить ядро парсинга эмулятора
	var products []bases.Product2
	core, ErrNewParsingCard := hm.NewParsingCard()
	if ErrNewParsingCard != nil {
		t.Error(ErrNewParsingCard)
	}

	// Тут якобы начало цикла
	// categ := Categorys[0]

	for _, categ := range Categorys {

		// Получить ссылку на все товары json
		LineUrl, ErrLineUrl := core.LineUrl(categ.Link)
		if ErrLineUrl != nil {
			t.Error(ErrLineUrl)
		}

		// // Получить к-во товаров в категории
		// cout, ErrorCount := hm.LinesCount(LineUrl)
		// if ErrorCount != nil {
		// 	t.Error(ErrorCount)
		// }

		// Получить все товары
		line, ErrLine := hm.Lines(LineUrl, 10)
		if ErrLine != nil {
			t.Error(ErrLine)
		}
		// Перевести полученый ответ от сервера в слайс Product2 и добавить в него соответствующие данные по каждому товару в зависимоти от категории,
		// а именно: Гендер, Каталог.
		AddProducts := hm.Line2Product2(line, categ.Cat, categ.GendetTag)

		// Добавляем полученный слайс с товарами в общий слайс товаров
		products = append(products, AddProducts[8])

		///////////////////////////////////////
		delivery := 500 // Доставка
		walrus := 1.3   // Моржа
		var ErrorTranstate error
		AddingProduct := products[0]
		for j := range AddingProduct.Item {
			core.VariableProduct3(&AddingProduct, j)
		}
		// Проверка загружен ли этот товар или нет
		if _, ok := Adding.AllProdSKU[products[0].Article]; !ok {

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
							Path: FilePathJpg,
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
							Path: FilePathJpg,
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
			AddingProduct = hmapp.EditCoast(AddingProduct, cb.Data.Valute.Try.Value/10, walrus, delivery)

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
				t.Error(errorAddProductWC)
			} else {
				products[0].Upload = true
			}
			Adding.AllProdSKU[AddingProduct.Article] = true
		} else {
			fmt.Println("Этот товар уже существует")
		}
	}

}
