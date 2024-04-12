package zaraapp

import (
	"fmt"
	"log"
	"strconv"

	zaratr "github.com/RB-PRO/ClikShop/pkg/ZaraTR"
	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/cbbank"
	"github.com/RB-PRO/ClikShop/pkg/transrb"
	"github.com/RB-PRO/ClikShop/pkg/wcprod"
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
