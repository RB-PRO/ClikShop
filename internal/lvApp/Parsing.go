package lvapp

import (
	"fmt"
	"log"

	lv "github.com/RB-PRO/ClikShop/pkg/LouisVuitton"
	"github.com/cheggaaa/pb"
)

func Parsing() {

	// Получаем список категорий
	Categorys, ErrCat := lv.Category()
	if ErrCat != nil {
		panic(ErrCat)
	}

	// Объект для работы с папками
	dir := lv.NewDir("")
	dir.MakeDir("LV/")

	// Создаём ядро парсинга
	core := lv.NewCore("", "")
	core.UpdateCore()

	fmt.Println("Всего категорий:", len(Categorys))

	var Prods []lv.Product
	for icategory, category := range Categorys {
		if icategory == 29 {

			// Получить список всех товаров
			Products := core.Pages(category.CategoryTag)

			// Парсинг внутри категории
			var ProdsCategory []lv.Product
			BarProducts := pb.StartNew(len(Products))
			BarProducts.Prefix(fmt.Sprintf("[%d/%d] Парсинг", icategory, len(Categorys)))
			for _, product := range Products {
				touch, ErrToucher := core.Toucher(product.ProductID)
				if ErrToucher != nil {
					log.Println("Error Toucher:", product.ProductID)
				}
				ProdsCategory = append(ProdsCategory, lv.TouchResponse2Product(touch))
				BarProducts.Increment()
			}
			BarProducts.Finish()

			// Сличаем цены с Дубаями
			var SKUs []string
			for _, product := range ProdsCategory {
				SKUs = append(SKUs, product.SKU)
			}
			mapPriceAE, ErrPrice := lv.Price(SKUs, "eng-ae", "AE")
			if ErrPrice != nil {
				panic(ErrPrice)
			}
			mapPriceFR, ErrPrice := lv.Price(SKUs, "fra-fr", "FR")
			if ErrPrice != nil {
				panic(ErrPrice)
			}
			for i := range ProdsCategory {
				for j := range ProdsCategory[i].Variations {
					ProdsCategory[i].Variations[j].PriceDub = mapPriceAE[ProdsCategory[i].SKU]
					ProdsCategory[i].Variations[j].PriceFr = mapPriceFR[ProdsCategory[i].SKU]
				}
			}

			// Работа с папками и картинками
			BarPhoto := pb.StartNew(len(Products))
			BarPhoto.Prefix(fmt.Sprintf("[%d/%d] Фото", icategory, len(Categorys)))
			for i := range ProdsCategory {
				PathToProduct := category.Path + ProdsCategory[i].SKU + "/"
				dir.MakeDir(PathToProduct)
				for j := range ProdsCategory[i].Variations {
					PathToVariation := PathToProduct + ProdsCategory[i].Variations[j].SKU + "/"
					dir.MakeDir(PathToVariation)
					ProdsCategory[i].Variations[j].Photo, _ = dir.SavePhotos(ProdsCategory[i].Variations[j].Photo, PathToVariation)
				}
				BarPhoto.Increment()
				if i == 10 {
					break
				}
			}
			BarPhoto.Finish()

			// Сохраняем полученные результаты
			Prods = append(Prods, ProdsCategory...)
			break
		}
	}

	lv.SaveXLSX("lv.xlsx", Prods)
}
