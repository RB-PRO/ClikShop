package zaraapp

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	zaratr "github.com/RB-PRO/SanctionedClothing/pkg/ZaraTR"
	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
	"github.com/RB-PRO/SanctionedClothing/pkg/wcprod"
	"github.com/cheggaaa/pb"
)

func Parse() {
	varient := zaratr.Parsing()
	varient.SaveXlsxCsvs("Zara")
}

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
				varient.Product[i] = EditCoast(varient.Product[i], cb.Data.Valute.Try.Value/10, walrus, ActualDelivery)
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
	// "Мягкий" выход из программы
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" {
			break
		} else {
			fmt.Println("Press 'q' to quit")
		}
	}
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
