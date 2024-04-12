package actualizer

// Файл для полного описания пути товара:
//	- Вычитание
//	- Перевод
//	- Публикация в Bitrix

import (
	"fmt"
	"strings"
	"time"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/transrb"
	"github.com/cheggaaa/pb"
)

// Вычитание товаров
//
// Передаём путь к папке с файлами с товарами и эта функция обработает все товары из папки,
// удалив дубли и сохранит новые товары в новой папке `*_sub`
func (bx *bitrixActualizer) Sub(Folder string) error {
	ProductCoutOrig, ProductCoutSub := 0, 0
	folderSub := Folder + "_sub"
	MakeDir(folderSub) // Создаём папку вычитания

	// Получаем файлы из папки источника со спарсенными товарами
	files, ErrFolderFiles := FolderFiles(Folder + "/")
	if ErrFolderFiles != nil {
		return fmt.Errorf("FolderFiles: %v", ErrFolderFiles)
	}
	for _, file := range files {
		// Читаем файл с товарами
		VarietyOrig, ErrProdFile := ProdFile(Folder+"/", file)
		if ErrProdFile != nil {
			bx.GLOG.Warn(fmt.Sprintf("%s%s: Не удалось прочитать файл: %v",
				Folder, file, ErrProdFile))
			continue
		}
		// Вычитаем товары
		VarietySub := bx.subtraction(VarietyOrig)
		ProductCoutOrig += len(VarietyOrig.Product)
		ProductCoutSub += len(VarietySub.Product)

		// Сохраняем вычитанные товары в файл
		VarietySub.SaveJson2(fmt.Sprintf("%s%s", folderSub+"/", file))
	}
	bx.BX.Nots.Sends(fmt.Sprintf("#actualizer\nВ бренде %s всего спарсили %d товаров, а после вычитания осталось %d, это %.2f%%",
		Folder, ProductCoutOrig, ProductCoutSub, 100.0*float64(ProductCoutSub)/float64(ProductCoutOrig)))
	return nil
}

// Удалить дубликаты
//
// Передаём путь к папке с файлами с товарами и эта функция обработает все товары из папки,
// удалив дубли и сохранит новые товары в новой папке `*_sub`
func (bx *bitrixActualizer) DeleteRepeated(Folder string) error {
	folderSub := Folder + "_sub"
	folderDr := Folder + "_dr"
	MakeDir(folderDr + "/") // Создаём папку вычитания

	// Получаем файлы из папки источника со спарсенными товарами
	files, ErrFolderFiles := FolderFiles(folderSub + "/")
	if ErrFolderFiles != nil {
		return fmt.Errorf("FolderFiles: %v", ErrFolderFiles)
	}

	var IgnoredProducts, TotalProducts int
	sku := make(map[string]bool)
	for _, file := range files {
		// Читаем файл с товарами
		Variety, ErrProdFile := ProdFile(folderSub+"/", file)
		if ErrProdFile != nil {
			bx.GLOG.Warn(fmt.Sprintf("%s%s: Не удалось прочитать файл: %v",
				Folder, file, ErrProdFile))
			continue
		}

		var SavedVar bases.Variety2

		// Обработка товаров
		for i := range Variety.Product {
			if _, ok := sku[Variety.Product[i].Article]; ok { // Если был уже такой SKU
				bx.GLOG.Warn(fmt.Sprintf("DeleteRepeated: %s: Товар %s с артикулом %s уже был, поэтому его игнорирую",
					file, Variety.Product[i].Name, Variety.Product[i].Article))
				IgnoredProducts++
			} else { // Если такого SKU не было
				SavedVar.Product = append(SavedVar.Product, Variety.Product[i])
				sku[Variety.Product[i].Article] = true
			}
			TotalProducts++
		}

		// Сохраняем вычитанные товары в файл
		SavedVar.SaveJson2(folderDr + "/" + file)
	}
	bx.BX.Nots.Sends(fmt.Sprintf("#actualizer\nУдаление дублей в папке '%s': Убрал %d товаров из %d, это %.2f%%",
		folderDr, IgnoredProducts, TotalProducts, 100.0*float64(IgnoredProducts)/float64(TotalProducts)))
	return nil
}

// Перевод товаров
//
// Передаём путь к папке с файлами с товарами и эта функция обработает все товары из папки,
// удалив дубли и сохранит новые товары в новой папке `*_sub`
func (bx *bitrixActualizer) Trans(Folder string) error {
	ProductTranslating := 0
	folderSub := Folder + "_dr"
	folderTr := Folder + "_tr"
	MakeDir(folderTr + "/") // Создаём папку вычитания

	// Получаем файлы из папки источника со спарсенными товарами
	files, ErrFolderFiles := FolderFiles(folderSub + "/")
	if ErrFolderFiles != nil {
		return fmt.Errorf("FolderFiles: %v", ErrFolderFiles)
	}
	for ifile, file := range files {
		// Читаем файл с товарами
		Variety, ErrProdFile := ProdFile(folderSub+"/", file)
		if ErrProdFile != nil {
			bx.GLOG.Warn(fmt.Sprintf("%s%s: Не удалось прочитать файл: %v",
				Folder, file, ErrProdFile))
			continue
		}

		// Перевод товаров
		barProduct := pb.StartNew(len(Variety.Product))
		barProduct.Prefix(fmt.Sprintf("Перевод [%d/%d]", ifile+1, len(files)))
		for i := range Variety.Product {
			Name := Variety.Product[i].Name

			var ErrCoast error
			Variety.Product[i], ErrCoast = bx.Coast(Variety.Product[i])
			if ErrCoast != nil {
				bx.GLOG.Err(fmt.Sprintf("[%d/%d]: %s: Товар '%s' %v",
					ifile, len(files), file, Variety.Product[i].Name, ErrCoast))
				continue
			}

			// Перевести товар
			TranslateProd, ErrorTranstate := bx.TR.TranslateProduct2(Variety.Product[i])
			if ErrorTranstate != nil {
				// fmt.Println("\nErrorTranstate ", ErrorTranstate, "\n ")
				bx.TR, _ = transrb.New(bx.TR.FolderID, bx.TR.OAuthToken)
				Variety.Product[i], _ = bx.TR.TranslateProduct2(Variety.Product[i])
			} else {
				Variety.Product[i] = TranslateProd
			}

			// Проверка того, что этот товар из СС(не связано с рейх-канцелярией)
			if strings.Contains(Variety.Product[i].Link, "sneaksup.com") {
				Variety.Product[i].Name = Name
			}

			if strings.Contains(Variety.Product[i].Link, "trendyol") {
				time.Sleep(time.Second)
				transName, ErrorTranstate2 := bx.TR.TransENG(Name)
				// fmt.Println("transName", transName, "ErrorTranstate2", ErrorTranstate2, "\n ")
				if ErrorTranstate2 != nil {
					// fmt.Println("\nErrorTranstate2 ", ErrorTranstate2, "\n ")
					bx.TR, _ = transrb.New(bx.TR.FolderID, bx.TR.OAuthToken)
					Variety.Product[i].Name, _ = bx.TR.TransENG(Name)
				} else {
					Variety.Product[i].Name = transName
				}

			}

			// Автоперевод размеров по годам.
			for Item_i := range Variety.Product[i].Item {
				for Size_i := range Variety.Product[i].Item[Item_i].Size {
					Variety.Product[i].Item[Item_i].Size[Size_i].Val = strings.ReplaceAll(
						Variety.Product[i].Item[Item_i].Size[Size_i].Val,
						"Yaş", "лет",
					)
				}
			}
			for Size_i := range Variety.Product[i].Size {
				Variety.Product[i].Size[Size_i] = strings.ReplaceAll(
					Variety.Product[i].Size[Size_i],
					"Yaş", "лет",
				)
			}

			// "Yaş"

			barProduct.Increment()
		}
		barProduct.Finish()

		ProductTranslating += len(Variety.Product)

		// Сохраняем вычитанные товары в файл
		Variety.SaveJson2(folderTr + "/" + file)
	}
	bx.BX.Nots.Sends(fmt.Sprintf("#actualizer\nВ бренде %s перевели %d товара(ов)",
		Folder, ProductTranslating))
	return nil
}

// Редактирование цены для товара
func (bx *bitrixActualizer) Coast(prod bases.Product2) (bases.Product2, error) {
	var Walrus float64
	var Delivery int
	if cm, ok := bx.BX.MapCoast[prod.Manufacturer]; ok {
		Walrus = cm.Walrus
		Delivery = cm.Delivery
	} else {
		cm = bx.BX.MapCoast["ss"]
		Walrus = cm.Walrus
		Delivery = cm.Delivery
		// return prod, fmt.Errorf("для производителя '%s' не найдены конфигурации цены, все цены: '%+v'",
		// 	prod.Manufacturer, bx.BX.MapCoast)
	}
	for i := range prod.Item {
		prod.Item[i].Price = bases.EditDecadense(Walrus*prod.Item[i].Price + float64(Delivery))
	}
	return prod, nil
}

// Полуть мапу всех артикулов товаров из магазина в Bitrix
func (bx *bitrixActualizer) subtraction(a bases.Variety2) (deductible bases.Variety2) {
	for i := range a.Product {
		if _, ok := bx.SKU[a.Product[i].Article]; !ok {
			deductible.Product = append(deductible.Product, a.Product[i])
		}
	}
	return deductible
}

////////////////////////////////

// Загрузка товаров в Bitrix
//
// Эта функция загрузит все товары из папки и создат `*_err` папку для ошибочных товаров
func (bx *bitrixActualizer) Push(Folder string) error {
	FolderProds := Folder + "_tr"
	FolderErr := Folder + "_err"
	MakeDir(FolderErr + "/") // Создаём папку вычитания
	ProductPushDone, ProductPushFalse, ProductPushAll := 0, 0, 0
	// Получаем файлы в папке
	files, ErrFolderFiles := FolderFiles(FolderProds + "/")
	if ErrFolderFiles != nil {
		return fmt.Errorf("FolderFiles: %v", ErrFolderFiles)
	}
	sku := make(map[string]bool)
	// for ifile, file := range files {
	for ifile := 0; ifile < len(files); ifile++ {
		file := files[ifile]
		var ErrProduct bases.Variety2
		// Читаем файл с товарами
		VarietyOrig, ErrProdFile := ProdFile(FolderProds+"/", file)
		if ErrProdFile != nil {
			bx.GLOG.Warn(fmt.Sprintf("%s%s: Не удалось прочитать файл: %v",
				Folder, file, ErrProdFile))
			// fmt.Println(ErrProdFile)
			continue
		}
		barProduct := pb.StartNew(len(VarietyOrig.Product))
		barProduct.Prefix(fmt.Sprintf("Публикация [%d/%d]", ifile+1, len(files)))
		for i := range VarietyOrig.Product {
			if _, ok := sku[VarietyOrig.Product[i].Article]; !ok {
				ID, ErrAddProd := bx.BX.AddProduct(VarietyOrig.Product[i])
				if ErrAddProd != nil {
					bx.BX.Log.Err(fmt.Sprintf("Ошибка при загрузке товара %s, Ошибка: %v",
						VarietyOrig.Product[i].Name, ErrAddProd))
					ErrProduct.Product = append(ErrProduct.Product, VarietyOrig.Product[i])
					ProductPushFalse++
				} else {
					bx.BX.Log.Info(fmt.Sprintf("Загрузил товар %s, bx - https://213.226.124.16/bitrix/admin/iblock_element_edit.php?IBLOCK_ID=15&type=aspro_lite_catalog&lang=ru&ID=%d&find_section_section=0&WF=Y , donor - %s",
						VarietyOrig.Product[i].Name, ID, VarietyOrig.Product[i].Link))
					ProductPushDone++
				}
				sku[VarietyOrig.Product[i].Article] = true
				ProductPushAll++
			} else { // Товар такой товар уже возникал при загрузке
				bx.BX.Log.Err(fmt.Sprintf("Товар '%s' с артикулом %s уже был при загрузке. Файл %s",
					VarietyOrig.Product[i].Name, VarietyOrig.Product[i].Article, file))
			}
			barProduct.Increment()
		}
		barProduct.Finish()
		ErrProduct.SaveJson2(FolderErr + "/" + file)
	}
	bx.BX.Nots.Sends(fmt.Sprintf("#actualizer\nТоваров для загрузки в бренде %s - %d\nУспешно загружено %d из %d(%.2f%%) товаров",
		Folder, ProductPushAll, ProductPushDone, ProductPushAll, 100.0*float64(ProductPushDone)/float64(ProductPushAll)))
	return nil
}
