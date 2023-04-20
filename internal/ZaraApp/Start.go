package zaraapp

import (
	"fmt"
	"log"

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
	fmt.Println("Курс доллара", cb.Data.Valute.Usd.Value)

	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		log.Fatalln(errorInitWcAdd)
	}

	Category, ErrorCat := zaratr.LoadCategory()
	if ErrorCat != nil {
		panic(ErrorCat)
	}

	ErrorCat = LoadCat3(Adding, Category)
	if ErrorCat != nil {
		panic(ErrorCat)
	}

}

// Получить категории и создать из них структуру категорий
func LoadCat3(Adding *wcprod.WcAdd, Category zaratr.Category) error {
	return nil
}
