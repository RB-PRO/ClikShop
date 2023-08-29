package zaraapp

import (
	"fmt"
	"log"
	"strconv"

	zaratr "github.com/RB-PRO/SanctionedClothing/pkg/ZaraTR"
	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
	"github.com/RB-PRO/SanctionedClothing/pkg/gol"
	"github.com/RB-PRO/SanctionedClothing/pkg/transrb"
	"github.com/RB-PRO/SanctionedClothing/pkg/wcprod"
	"github.com/cheggaaa/pb"
)

func Parse() {
	varient := zaratr.Parsing()
	varient.SaveXlsxCsvs("Zara")
}
func Parsing3() {
	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	log.Println("Курс лиры", cb.Data.Valute.Try.Value/10)

	// Создать оьбъект переводчика
	Translate, ErrNewTranslate := wcprod.NewTranslate()
	if ErrNewTranslate != nil {
		panic(ErrNewTranslate)
	}

	// Парсинг
	varient := zaratr.Parsing()
	varient.SaveJson("tmp/ZARA")

	// ***************************************
	// Парсинг по подслайсами с размером size
	size := 300
	BarProducts := pb.StartNew(len(varient.Product))
	var SubSlice_j, cout int
	for SubSlice_i := 0; SubSlice_i < len(varient.Product); SubSlice_i += size {
		SubSlice_j += size
		if SubSlice_j > len(varient.Product) {
			SubSlice_j = len(varient.Product)
		}

		// Подслайс. Работаем именно с подслайсами, чтобы не перегружать оперативку
		SubSlice := varient.Product[SubSlice_i:SubSlice_j]
		BarProducts.Prefix(strconv.Itoa(cout))
		for i := range SubSlice {
			// Парсинг всех подпродуктов
			AddingProduct := SubSlice[i]

			// Редактируем товар
			AddingProduct = bases.EditCoast(AddingProduct, cb.Data.Valute.Try.Value/10, 1.3, 500)
			AddingProduct.Size = bases.EditProdSize(AddingProduct)
			AddingProduct.Img = bases.EditIMG(AddingProduct)

			// Перевести товар
			var ErrorTranstate error
			AddingProduct, ErrorTranstate = Translate.YandexTranslatePart(AddingProduct)
			if ErrorTranstate != nil {
				Translate.Tr, _ = transrb.New(Translate.Tr.FolderID, Translate.Tr.OAuthToken)
				AddingProduct, _ = Translate.YandexTranslatePart(AddingProduct)
			}

			SubSlice[i] = AddingProduct

			BarProducts.Increment()
		}
		cout++
		// bases.Variety2{Product: SubSlice}.SaveXlsxCsvs(fmt.Sprintf("tmp/H&M_SubSlice_%d_%d-%d", cout, SubSlice_i, SubSlice_i+size))
		bases.Variety2{Product: SubSlice}.SaveJson(fmt.Sprintf("tmp/ZARA_SubSlice_%d_%d-%d", cout, SubSlice_i, SubSlice_i+size))
	}
	BarProducts.Finish()
	bases.ExitSoft()
}

func Parsing() {
	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}

	glog := gol.NewGol()
	log.Println("Курс лиры", cb.Data.Valute.Try.Value/10)
	// Категории
	CatArr, _ := zaratr.CatCycle2() // Получить все категории
	log.Println("Всего", len(CatArr.Items), "категорий")

	// Все товары

	allID := make(map[string]bool)
	var cout int
	for i, cat := range CatArr.Items {
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
			Prod2.Cat = prod.Cat                                                          // Обновляем категории
			Prod2.GenderLabel = prod.Gender                                               // Обнволяем гендер

			// Редактируем товар
			Prod2 = bases.EditCoast(Prod2, cb.Data.Valute.Try.Value/10, 1.3, 500)
			Prod2.Size = bases.EditProdSize(Prod2)
			Prod2.Img = bases.EditIMG(Prod2)

			Variety.Product = append(Variety.Product, Prod2)
			cout++
		}
		if len(Variety.Product) > 0 {
			Variety.SaveXlsx(fmt.Sprintf("tmp/ZARA/xlsx/Zara_%d_%v", i, Variety.Product[0].Cat[len(Variety.Product[0].Cat)-1].Name))
			Variety.SaveJson(fmt.Sprintf("tmp/ZARA/json/Zara_%d_%v", i, Variety.Product[0].Cat[len(Variety.Product[0].Cat)-1].Name))
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
