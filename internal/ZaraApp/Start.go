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
)

// Начать парсить и одновременно загружать товары
func Start() {

	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	fmt.Println("Курс лиры", cb.Data.Valute.Try.Value)

	// Загружаем товары на WC //
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		log.Fatalln(errorInitWcAdd)
	}

	varient := zaratr.Parsing()

	varient.SaveXlsxCsvs("Zara")

	// Загружаем товары
	delivery := 500 // Доставка
	walrus := 1.3   // Моржа
	// 3500
	for i := 2500; i < len(varient.Product)-2; i++ {
		fmt.Printf("Start: Загружаю товар (%d/%d)", i, len(varient.Product)-2)
		if !varient.Product[i].Upload {
			// Формирование адекватной цены доставки из файла
			ActualDelivery := Adding.EditDelivery(varient.Product[i].Cat, delivery)
			varient.Product[i] = EditCoast(varient.Product[i], cb.Data.Valute.Try.Value/10, walrus, ActualDelivery)
			//errorAddProductWC := Adding.AddProduct(wcprod.ProductTranslate(varient.Product[i])) //.AddAttr()
			varient.Product[i], _ = Adding.YandexTranslate(varient.Product[i])
			errorAddProductWC := Adding.AddProduct(varient.Product[i]) //.AddAttr()
			if errorAddProductWC != nil {
				varient.Product[i].Upload = true
			}
		}
	}
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
