package zaraapp

import (
	"bufio"
	"fmt"
	"os"

	zaratr "github.com/RB-PRO/SanctionedClothing/pkg/ZaraTR"
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

	Variety := zaratr.Parsing()

	Variety.SaveXlsxCsvs("Zara")

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

// Получить категории и создать из них структуру категорий
func LoadCat3(Adding *wcprod.WcAdd, Category zaratr.Category) error {
	return nil
}
