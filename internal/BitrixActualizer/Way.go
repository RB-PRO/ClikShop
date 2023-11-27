package actualizer

// Файл для полного описания пути товара:
//	- Вычитание
//	- Перевод
//	- Публикация в Bitrix

import (
	"fmt"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/transrb"
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
	bx.BX.Nots.Sends(fmt.Sprintf("#actualizer\nВ бренде %s всего спарсили %d товаров, а после вычитания осталось %d",
		Folder, ProductCoutOrig, ProductCoutSub))
	return nil
}

// Перевод товаров
//
// Передаём путь к папке с файлами с товарами и эта функция обработает все товары из папки,
// удалив дубли и сохранит новые товары в новой папке `*_sub`
func (bx *bitrixActualizer) Trans(Folder string) error {
	ProductTranslating := 0
	folderSub := Folder + "_sub"
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
			// Перевести товар
			TranslateProd, ErrorTranstate := bx.TR.TranslateProduct2(Variety.Product[i])
			if ErrorTranstate != nil {
				bx.TR, _ = transrb.New(bx.TR.FolderID, bx.TR.OAuthToken)
				Variety.Product[i], _ = bx.TR.TranslateProduct2(Variety.Product[i])
			} else {
				Variety.Product[i] = TranslateProd
			}
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
	for ifile, file := range files {
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
			ProductPushAll++
			barProduct.Increment()
		}
		barProduct.Finish()
		ErrProduct.SaveJson2(FolderErr + "/" + file)
	}
	bx.BX.Nots.Sends(fmt.Sprintf("#actualizer\nТоваров для загрузки в бренде %s - %d\nУспешно загружено %d из %d(%.2f) товаров",
		Folder, ProductPushAll, ProductPushDone, ProductPushAll, float64(ProductPushDone/ProductPushAll)))
	return nil
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
