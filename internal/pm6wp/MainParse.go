package pm6wp

import (
	"fmt"
	"log"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/cbbank"
	"github.com/RB-PRO/ClikShop/pkg/pm6"
	"github.com/RB-PRO/ClikShop/pkg/wcprod"
	//"github.com/RB-PRO/ClikShop/pkg/wcprod"
)

// Функция, которая занимается вызовом функций парсинга и загрузки товаров
// Входные параметры:
// PageStart int  // Стартовая страница
// walrus float64 // Моржа
// delivery int   // Стоймость доставки
func Work(PageStart int, walrus float64, delivery int) {
	pmm, ErrorPMM := pm6.NewPM()
	if ErrorPMM != nil {
		log.Fatalln(ErrorPMM)
	}

	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		log.Fatalln(ErrorCB)
	}
	usd := cb.Data.Valute.Usd.Value

	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		log.Fatalln(errorInitWcAdd)
	}

	linkPages := "/null/.zso?s=brandNameFacetLC/asc/productName/asc/" // Ссылка на страницу товаров

	linkPages = "/hats/COfWARCJ1wHiAgIBAg.zso?s=brandNameFacetLC/asc/productName/asc/"
	linkPages = "/women-clothing/CKvXAcABAeICAgEY.zso?s=brandNameFacetLC/asc/productName/asc/"
	PageEnd := pmm.AllPages(linkPages) // Получить сколько всего страниц товаров есть
	fmt.Println("[pm6wp]: Всего страниц", PageEnd)
	PageEnd = 391                                                      // До этого мы парсим
	var varient bases.Variety2                                         // Массив базы данных товаров
	varient = pmm.ParsePageWithVarienty(varient, linkPages, PageStart) // Парсим первую страницу товаров
	for i := PageStart + 1; i <= PageEnd; i++ {                        // Цикл по всем страницам товаров
		fmt.Println("[pmwp]: Парсинг страниц", i, "/", PageEnd)

		// Сортируем товары и записываем их в готовую базу данных varient
		varient = pmm.ParsePageWithVarienty(varient, linkPages, i) // Парсим первую страницу товаров

		for j := 0; j < len(varient.Product); j++ {
			fmt.Println(">>", j, "/", len(varient.Product))
			if varient.Product[j].Manufacturer == "" {
				for key := range varient.Product[j].Item {
					//fmt.Println("parse", varient.Product[j].Item[key].Link)
					pmm.ParseProduct(&varient.Product[j], varient.Product[j].Item[key].Link)
				}
			}
		}

		// Загружаем товары
		for i := 0; i < len(varient.Product)-2; i++ {
			if !varient.Product[i].Upload {
				// Формирование адекватной цены доставки из файла
				ActualDelivery := Adding.EditDelivery(varient.Product[i].Cat, delivery)
				varient.Product[i] = EditCoast(varient.Product[i], usd, walrus, ActualDelivery)
				//errorAddProductWC := Adding.AddProduct(wcprod.ProductTranslate(varient.Product[i])) //.AddAttr()
				errorAddProductWC := Adding.AddProduct(varient.Product[i]) //.AddAttr()
				if errorAddProductWC != nil {
					varient.Product[i].Upload = true
				}
			}
		}
	}

	varient.SaveXlsxCsvs("TEST")
}

// Редактирование цены по товарам
func EditCoast(prod bases.Product2, usd float64, walrus float64, delivery int) bases.Product2 {
	for indexKey := range prod.Item {
		// Корректируем данные
		// Курс доллара * цена в долларах * наценка + цена доставки
		prod.Item[indexKey].Price = usd*prod.Item[indexKey].Price*walrus + float64(delivery)
	}
	return prod
}
