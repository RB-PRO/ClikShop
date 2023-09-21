package sneaksupapp

import (
	"fmt"
	"strings"

	sneasup "github.com/RB-PRO/SanctionedClothing/pkg/SneaSup"
	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
	"github.com/cheggaaa/pb"
)

func Parsing() {

	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	fmt.Println("Курс лиры", cb.Data.Valute.Try.Value/10)

	// // Создать оьбъект переводчика
	// Translate, ErrNewTranslate := wcprod.NewTranslate()
	// if ErrNewTranslate != nil {
	// 	panic(ErrNewTranslate)
	// }

	SneakSupCategory := sneasup.Category()
	fmt.Println("Всего категорий:", len(SneakSupCategory))

	var CoutProduct int

	BarCategory := pb.StartNew(len(SneakSupCategory))
	for iCategory, Category := range SneakSupCategory {
		Namecategory := strings.ReplaceAll(Category.Link, "https://www.sneaksup.com/", "") // Название категории
		BarCategory.Prefix(Namecategory)

		// Загружаем line по категории
		line, ErrLines := sneasup.Lines(Category.Link)
		if ErrLines != nil {
			fmt.Println(ErrLines)
		}

		// Переводим в свою структур товаров
		Products := sneasup.Line2Product(line, Category)

		CoutProduct += len(line)

		// Редактирование товара, согласно требованиям
		for iProd := range Products {
			Products[iProd].Size = bases.EditProdSize(Products[iProd])
			Products[iProd] = bases.EditCoast(Products[iProd], cb.Data.Valute.Try.Value/10, 1.3, 500)
			Products[iProd].Description.Eng, _ = sneasup.Description(Products[iProd].Link)
			// var ErrorTranstate error // Перевести товар
			// Products[iProd], ErrorTranstate = Translate.YandexTranslate(Products[iProd])
			// if ErrorTranstate != nil {
			// 	Translate.Tr, _ = transrb.New(Translate.Tr.FolderID, Translate.Tr.OAuthToken)
			// 	Products[iProd], _ = Translate.YandexTranslate(Products[iProd])
			// }
		}

		// Сохранить в файл
		bases.Variety2{Product: Products}.SaveJson(fmt.Sprintf("tmp/SS/SS_%d_%s", iCategory, Namecategory))
		BarCategory.Increment()
	}
	BarCategory.Finish()

}
