package zaraapp

import (
	"fmt"
	"log"

	zaratr "github.com/RB-PRO/ClikShop/pkg/ZaraTR"
	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/cbbank"
	"github.com/RB-PRO/ClikShop/pkg/gol"
	"github.com/RB-PRO/ClikShop/pkg/transrb"
	"github.com/RB-PRO/ClikShop/pkg/wcprod"
	"github.com/cheggaaa/pb"
)

func Parsing() {
	glog, _ := gol.NewGol("logs/")

	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	glog.Info("Курс лиры", cb.Data.Valute.Try.Value/10)

	// Создать оьбъект переводчика
	Adding, ErrNewTranslate := wcprod.NewTranslate()
	if ErrNewTranslate != nil {
		panic(ErrNewTranslate)
	}
	glog.Info("Загрузил переводчик", Adding.Tr.OAuthToken)

	// Категории
	CatArr, _ := zaratr.CatCycle2() // Получить все категории
	log.Println("Всего", len(CatArr.Items), "категорий")

	// Все товары
	allID := make(map[string]bool)
	var cout int
	for i, cat := range CatArr.Items {
		if i < 129 {
			continue
		}
		line, ErrorLine := zaratr.LoadLine(fmt.Sprintf("%v", cat.RedirectCategoryID))
		if ErrorLine != nil {
			fmt.Println(ErrorLine)
		}
		// bar.Increment()

		ProductsLine := make([]zaratr.CommercialComponents, 0)
		for _, ProductGroups := range line.ProductGroups {
			for _, Elements := range ProductGroups.Elements {
				for _, CommercialComponents := range Elements.CommercialComponents {
					// if cout >= 10 { // Максимум 10 товаров в категории
					// 	break
					// }
					CommercialComponents.Cat = cat.Cat
					CommercialComponents.Gender = cat.Gender
					ProductsLine = append(ProductsLine, CommercialComponents)
				}
			}
		}

		ProductsLine = RemoveDuplicateProductsLine(ProductsLine)

		// Парсим товары
		var Variety bases.Variety2
		bar2 := pb.StartNew(len(ProductsLine))
		bar2.Prefix(fmt.Sprintf("[%d/%d]", i, len(CatArr.Items)))

		// Переведённая категория
		FileName := ProductsLine[0].Cat[len(ProductsLine[0].Cat)-1].Slug
		ProdTranslateCat := ProductsLine[0].Cat
		ProdTranslateCat, _ = Adding.YandexCat(ProdTranslateCat)

		for _, prod := range ProductsLine {
			bar2.Increment()
			if _, valueok := allID[prod.Reference]; !valueok {
				allID[prod.Reference] = true
			} else {
				glog.Warn("Дубль:", fmt.Sprintf(zaratr.TouchURL, prod.Seo.Keyword+"-p"+prod.Seo.SeoProductID))
				continue
			}

			glog.Info("Парсинг LoadTouch:", fmt.Sprintf(zaratr.TouchURL, prod.Seo.Keyword+"-p"+prod.Seo.SeoProductID))
			touch, _ := zaratr.LoadTouch(prod.Seo.Keyword + "-p" + prod.Seo.SeoProductID) // Выполняем запрос
			Prod2 := zaratr.Touch2Product2(touch)                                         // АПереводим в структуру Product2
			Prod2.Cat = ProdTranslateCat                                                  //prod.Cat // Обновляем категории
			Prod2.GenderLabel = prod.Gender                                               // Обнволяем гендер

			// Перевести товар
			var ErrorTranstate error
			Prod2, ErrorTranstate = Adding.YandexTranslate(Prod2)
			if ErrorTranstate != nil {
				Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
				Prod2, _ = Adding.YandexTranslate(Prod2)
			}

			// Редактируем товар
			Prod2 = bases.EditCoast(Prod2, cb.Data.Valute.Try.Value/10, 1.3, 500)
			Prod2.Size = bases.EditProdSize(Prod2)
			Prod2.Img = bases.EditIMG(Prod2)

			Variety.Product = append(Variety.Product, Prod2)
			cout++
		}
		if len(Variety.Product) > 0 {
			Variety.SaveXlsx(fmt.Sprintf("tmp/ZARA/xlsx/zara_%d_%v", i, FileName)) //Variety.Product[0].Cat[len(Variety.Product[0].Cat)-1].Name))
			Variety.SaveJson(fmt.Sprintf("tmp/ZARA/json/zara_%d_%v", i, FileName)) //Variety.Product[0].Cat[len(Variety.Product[0].Cat)-1].Name))
		}
		bar2.Finish()
	}
	log.Println("Всего", cout, "товара(ов)")
}

// Удалить дубликаты в товарах ProductsLine
func RemoveDuplicateProductsLine(ProductsLine []zaratr.CommercialComponents) []zaratr.CommercialComponents {
	allKeys := make(map[string]bool)
	list := []zaratr.CommercialComponents{}
	for _, item := range ProductsLine {
		if _, value := allKeys[item.ID.String()]; !value {
			allKeys[item.ID.String()] = true
			list = append(list, item)
		}
	}
	return list
}
