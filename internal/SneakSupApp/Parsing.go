package sneaksupapp

import (
	"fmt"
	"strings"

	sneasup "ClikShop/common/SneaSup"
	"ClikShop/common/bases"
	"ClikShop/common/cbbank"
	"ClikShop/common/transrb"
	"ClikShop/common/wcprod"
	"github.com/cheggaaa/pb"
)

func Parsing() {

	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	fmt.Println("Курс лиры", cb.Data.Valute.Try.Value/10)

	// Создать оьбъект переводчика
	Translate, ErrNewTranslate := wcprod.NewTranslate()
	if ErrNewTranslate != nil {
		panic(ErrNewTranslate)
	}

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

		// Products = Products[:2]

		// Редактирование товара, согласно требованиям
		for iProd := range Products {
			BarCategory.Prefix(fmt.Sprintf("%s [%d/%d]", Namecategory, iProd, len(Products)))
			for iItem := range Products[iProd].Item {
				for iSize := range Products[iProd].Item[iItem].Size {
					Products[iProd].Item[iItem].Size[iSize].Val = strings.TrimSpace(Products[iProd].Item[iItem].Size[iSize].Val)
					Products[iProd].Item[iItem].Size[iSize].Val = strings.ReplaceAll(Products[iProd].Item[iItem].Size[iSize].Val, "Yaş", "лет")
					Products[iProd].Item[iItem].Size[iSize].Val = strings.ReplaceAll(Products[iProd].Item[iItem].Size[iSize].Val, "YAŞ", "лет")
				}
			}
			Products[iProd].Size = bases.EditProdSize(Products[iProd])
			Products[iProd] = bases.EditCoast(Products[iProd], cb.Data.Valute.Try.Value/10, 1.4, 600)
			Products[iProd].Description.Eng, _ = sneasup.Description(Products[iProd].Link)
			Products[iProd].Img = bases.EditIMG(Products[iProd])
			img := make([]string, 0)
			for iItem := range Products[iProd].Item {
				img = append(img, Products[iProd].Item[iItem].Image...)
			}
			Products[iProd].Img = img

			Name := Products[iProd].Name
			var ErrorTranstate error // Перевести товар
			Products[iProd], ErrorTranstate = Translate.YandexTranslatePart(Products[iProd])
			if ErrorTranstate != nil {
				fmt.Println(ErrorTranstate)
				Translate.Tr, _ = transrb.New(Translate.Tr.FolderID, Translate.Tr.OAuthToken)
				Products[iProd], _ = Translate.YandexTranslatePart(Products[iProd])
			}
			Products[iProd].Name = Name
			// break
		}

		// Сохранить в файл
		bases.Variety2{Product: Products}.SaveJson(fmt.Sprintf("tmp/SS/SS_%d_%s", iCategory, Namecategory))
		BarCategory.Increment()
		// break
	}
	fmt.Println("Всего товаров:", CoutProduct)
	BarCategory.Finish()

}
