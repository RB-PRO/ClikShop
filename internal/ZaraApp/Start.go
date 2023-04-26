package zaraapp

import (
	"bufio"
	"fmt"
	"log"
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

	// Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	// if errorInitWcAdd != nil {
	// 	log.Fatalln(errorInitWcAdd)
	// }

	varient := zaratr.Parsing()

	varient.SaveXlsxCsvs("Zara")

	// Загружаем товары на WC //
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		log.Fatalln(errorInitWcAdd)
	}
	// Загружаем товары
	delivery := 100 // Доставка
	walrus := 1.5   // Моржа
	for i := 0; i < len(varient.Product)-2; i++ {
		if !varient.Product[i].Upload {
			// Формирование адекватной цены доставки из файла
			ActualDelivery := Adding.EditDelivery(varient.Product[i].Cat, delivery)
			varient.Product[i] = EditCoast(varient.Product[i], cb.Data.Valute.Try.Value, walrus, ActualDelivery)
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
		// Если есть мапа с таким-же ключом, то копируем во вторичную переменную значение этой мапы по ключу
		if entry, ok := prod.Item[indexKey]; ok {

			// Корректируем данные
			// Курс доллара * цена в долларах * наценка + цена доставки
			entry.Price = usd*entry.Price*walrus + float64(delivery)

			// Обновляем данные
			prod.Item[indexKey] = entry
		}
	}
	return prod
}
