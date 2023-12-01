package actualizer

import (
	"fmt"
	"strings"

	sneasup "github.com/RB-PRO/SanctionedClothing/pkg/SneaSup"
	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/cheggaaa/pb"
)

func (bx *bitrixActualizer) ss(folder string) {

	MakeDir(folder)

	SneakSupCategory := sneasup.Category()

	for iCategory, Category := range SneakSupCategory {
		Namecategory := strings.ReplaceAll(Category.Link, "https://www.sneaksup.com/", "") // Название категории

		// Загружаем line по категории
		line, ErrLines := sneasup.Lines(Category.Link)
		if ErrLines != nil {
			fmt.Println(ErrLines)
		}

		// Переводим в свою структур товаров
		Products := sneasup.Line2Product(line, Category)

		// Редактирование товара, согласно требованиям
		BarProducts := pb.StartNew(len(Products))
		BarProducts.Prefix(fmt.Sprintf("[%d/%d]", iCategory+1, len(SneakSupCategory)))
		for iProd := range Products {
			// BarCategory.Prefix(fmt.Sprintf("%s [%d/%d]", Namecategory, iProd, len(Products)))
			for iItem := range Products[iProd].Item {
				for iSize := range Products[iProd].Item[iItem].Size {
					Products[iProd].Item[iItem].Size[iSize].Val = strings.TrimSpace(Products[iProd].Item[iItem].Size[iSize].Val)
					Products[iProd].Item[iItem].Size[iSize].Val = strings.ReplaceAll(Products[iProd].Item[iItem].Size[iSize].Val, "Yaş", "лет")
					Products[iProd].Item[iItem].Size[iSize].Val = strings.ReplaceAll(Products[iProd].Item[iItem].Size[iSize].Val, "YAŞ", "лет")
				}
			}
			Products[iProd].Size = bases.EditProdSize(Products[iProd])
			Products[iProd].Description.Eng, _ = sneasup.Description(Products[iProd].Link)
			Products[iProd].Img = bases.EditIMG(Products[iProd])
			img := make([]string, 0)
			for iItem := range Products[iProd].Item {
				img = append(img, Products[iProd].Item[iItem].Image...)
			}
			Products[iProd].Img = img
			BarProducts.Increment()
		}
		BarProducts.Finish()

		// Сохранить в файл
		bases.Variety2{Product: Products}.SaveJson(fmt.Sprintf("%s/ss_%d_%s",
			folder, iCategory, Namecategory))

	}
}
