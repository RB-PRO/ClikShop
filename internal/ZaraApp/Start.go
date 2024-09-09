package zaraapp

import (
	"log"

	zaratr "ClikShop/common/ZaraTR"
	"ClikShop/common/bases"
	"ClikShop/common/cbbank"
	"ClikShop/common/wcprod"
	"github.com/cheggaaa/pb"
)

// Начать парсить и одновременно загружать товары
func Start() {

	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	log.Println("cbbank: Курс лиры", cb.Data.Valute.Try.Value/10)

	// Загружаем товары на WC //
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		log.Fatalln(errorInitWcAdd)
	}

	// Пропарсить все товары на заре
	varient := zaratr.Parsing()

	// Сохранить данные в файл xlsx
	varient.SaveXlsxCsvs("Zara")
	log.Println("Парсинг: Сохраняю", len(varient.Product), "товаров")

	// Загружаем товары
	delivery := 500 // Доставка
	walrus := 1.3   // Моржа
	bar := pb.StartNew(len(varient.Product) - 2)
	for i := 0; i < len(varient.Product)-2; i++ {
		log.Printf("Start: Загружаю товар (%d/%d): %s\n", i+1, len(varient.Product)-2, varient.Product[i].Link)
		if !varient.Product[i].Upload {
			if _, ok := Adding.AllProdSKU[varient.Product[i].Article]; !ok {
				// Формирование адекватной цены доставки из файла
				ActualDelivery := Adding.EditDelivery(varient.Product[i].Cat, delivery)
				varient.Product[i] = bases.EditCoast(varient.Product[i], cb.Data.Valute.Try.Value/10, walrus, ActualDelivery)
				// var ErrorTranstate error
				// varient.Product[i], ErrorTranstate = Adding.YandexTranslate(varient.Product[i])
				// if ErrorTranstate != nil {
				// 	Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
				// 	varient.Product[i], _ = Adding.YandexTranslate(varient.Product[i])
				// }
				errorAddProductWC := Adding.AddProduct(varient.Product[i]) //.AddAttr()
				if errorAddProductWC != nil {
					varient.Product[i].Upload = true
				}
				Adding.AllProdSKU[varient.Product[i].Article] = true
			}
		}
		bar.Increment()
	}
	bar.Finish()

	bases.ExitSoft() // "Мягкий" выход из программы
}
