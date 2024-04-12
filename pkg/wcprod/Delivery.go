package wcprod

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// Загрузить список, где в каждой строке:
//   - наименование подкатегории
//   - цена доставки
func XlsxDelivery() (map[string]int, error) {
	FileName := "Delivery.xlsx"
	SheetName := "main"
	f, err := excelize.OpenFile(FileName)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Не могу открыть файл " + FileName)
	}
	defer f.Close() // Закрыть файл в конце выполнения функции

	// Получить все строки в SheetName
	rows, err := f.GetRows(SheetName)
	if err != nil {
		return nil, errors.New("Не могу получить строки в файле " + FileName + " на листе " + SheetName)
	}

	Delivery := make(map[string]int)
	for _, row := range rows {
		ratio, errorStrConv := strconv.Atoi(row[1])
		if errorStrConv != nil {
			return nil, errors.New("Не могу распарсить данные(string>int) " + row[1] + " в файле " + FileName + " на листе " + SheetName)
		}
		Delivery[row[0]] = ratio
	}

	return Delivery, nil
}
