package louisvuitton

import (
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func SaveXLSX(FilePathName string, Products []Product) error {
	f := excelize.NewFile()

	// Создаём новый лист
	SheetName := "main"
	if _, err := f.NewSheet(SheetName); err != nil {
		return err
	}
	if err := f.DeleteSheet("Sheet1"); err != nil {
		return err
	}

	// Шапка
	f.SetCellValue(SheetName, "A1", "Brand")
	f.SetCellValue(SheetName, "B1", "SKU")
	f.SetCellValue(SheetName, "C1", "Mark")
	f.SetCellValue(SheetName, "D1", "Category")
	f.SetCellValue(SheetName, "E1", "Title")
	f.SetCellValue(SheetName, "F1", "Description")
	f.SetCellValue(SheetName, "G1", "Photo")
	f.SetCellValue(SheetName, "H1", "Price")
	f.SetCellValue(SheetName, "I1", "Height")
	f.SetCellValue(SheetName, "J1", "Width")
	f.SetCellValue(SheetName, "K1", "Length")
	f.SetCellValue(SheetName, "L1", "SEO title")
	f.SetCellValue(SheetName, "M1", "SEO description")
	f.SetCellValue(SheetName, "N1", "External ID")
	f.SetCellValue(SheetName, "O1", "Parent UID")
	f.SetCellValue(SheetName, "P1", "Editions")
	f.SetCellValue(SheetName, "Q1", "PriceRus")
	f.SetCellValue(SheetName, "R1", "PriceFr")

	// Товары:
	var LineCout int = 2
	for _, Product := range Products {
		f.SetCellValue(SheetName, "A"+strconv.Itoa(LineCout), Product.Brand) //A
		f.SetCellValue(SheetName, "B"+strconv.Itoa(LineCout), Product.SKU)   //B
		// C
		f.SetCellValue(SheetName, "D"+strconv.Itoa(LineCout), strings.Join(Product.Category, ";")) // D
		f.SetCellValue(SheetName, "E"+strconv.Itoa(LineCout), Product.Title)                       // E
		f.SetCellValue(SheetName, "F"+strconv.Itoa(LineCout), Product.Description)                 // F
		// G
		// H
		f.SetCellValue(SheetName, "I"+strconv.Itoa(LineCout), Product.Height) // I
		f.SetCellValue(SheetName, "J"+strconv.Itoa(LineCout), Product.Width)  // J
		f.SetCellValue(SheetName, "K"+strconv.Itoa(LineCout), Product.Length) // K
		// L
		// M
		f.SetCellValue(SheetName, "N"+strconv.Itoa(LineCout), Product.External_ID) // N
		// O
		// P
		LineCout++

		// Вариации товаров:
		for _, Variation := range Product.Variations {
			// A
			f.SetCellValue(SheetName, "B"+strconv.Itoa(LineCout), Variation.SKU) // B
			if Variation.Mark {
				f.SetCellValue(SheetName, "C"+strconv.Itoa(LineCout), "Новинка") // C
			}
			// D
			// E
			// F
			f.SetCellValue(SheetName, "G"+strconv.Itoa(LineCout), strings.Join(Variation.Photo, ";")) // G
			f.SetCellValue(SheetName, "H"+strconv.Itoa(LineCout), Variation.PriceDub)                 // H
			// I
			// J
			// K
			// L
			// M
			f.SetCellValue(SheetName, "N"+strconv.Itoa(LineCout), Variation.External_ID) // N
			f.SetCellValue(SheetName, "O"+strconv.Itoa(LineCout), Variation.Parent_UID)  // O
			f.SetCellValue(SheetName, "P"+strconv.Itoa(LineCout), Variation.Editions)    // P
			f.SetCellValue(SheetName, "Q"+strconv.Itoa(LineCout), Variation.PriceRus)    // P
			f.SetCellValue(SheetName, "R"+strconv.Itoa(LineCout), Variation.PriceFr)     // P
			LineCout++
		}
	}

	// Сохраняем документ
	if err := f.SaveAs(FilePathName); err != nil {
		return err
	}

	// Закрываем документ
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
