package louisvuitton

import (
	"fmt"
	"testing"
)

func TestFolder(t *testing.T) {
	dir := NewDir("")
	errMakedir := dir.MakeDir("REAL/REAL")
	if errMakedir != nil {
		t.Error(errMakedir)
	}
}

func TestSavePhotos(t *testing.T) {
	core := coreTest()
	touch, ErrTouch := core.Toucher("001054")
	if ErrTouch != nil {
		t.Error(ErrTouch)
	}
	if touch.Name == "" {
		t.Error("Toucher: Не получить получить ответ или распарсить его")
	}

	Prod := TouchResponse2Product(touch)

	dir := NewDir("LV4/")

	dir.SavePhotos([]string{"https://ru.louisvuitton.com/images/is/image/lv/1/PP_VP_L/louis-vuitton-сумка-keepall-50-с-плечевым-ремнём-lv-aerogram-багаж--M22609_PM1_Worn view.png"}, "")

	PathToProduct := Prod.SKU + "/"
	dir.MakeDir(PathToProduct)
	for j := range Prod.Variations {
		PathToVariation := PathToProduct + Prod.Variations[j].SKU + "/"
		dir.MakeDir(PathToVariation)
		dir.SavePhotos(Prod.Variations[j].Photo, PathToVariation)
	}
	// dir.SavePhotos()
}

func TestPageAndSaveImages(t *testing.T) {
	core := coreTest()
	page, errpage := core.PageSingle("t1z0ff7q", 0)
	if errpage != nil {
		t.Error(errpage)
	}
	if page.NbPages == 0 {
		t.Error("PageSingle: Не получить получить ответ или распарсить его")
	}
	// for _, v := range page.Hits { fmt.Println(v.ProductID) }

	fmt.Println("page.Hits[0].ProductID", page.Hits[0].ProductID)

	touch, ErrToucher := core.Toucher(page.Hits[0].ProductID)
	if ErrToucher != nil {
		t.Error(ErrToucher)
	}
	Prod := TouchResponse2Product(touch)

	fmt.Println("Все фотографии:")
	for i := range Prod.Variations {
		for j := range Prod.Variations[i].Photo {
			fmt.Println(i, j, Prod.Variations[i].Photo[j])
		}
	}

	// Объект для работы с папками
	dir := NewDir("LV/")

	// Загрузить одно фото
	fmt.Println("Загружаю фото:")
	fmt.Println(Prod.Variations[0].Photo[0])
	ErrSavePhoto := dir.SavePhoto(Prod.Variations[0].Photo[0], "LV/test.png")
	if ErrSavePhoto != nil {
		t.Error(ErrSavePhoto)
	}

	// Загрузить все фотографии по товару
	PathToProduct := Prod.SKU + "/"
	dir.MakeDir(PathToProduct)
	for j := range Prod.Variations {
		PathToVariation := PathToProduct + Prod.Variations[j].SKU + "/"
		dir.MakeDir(PathToVariation)
		Prod.Variations[j].Photo, _ = dir.SavePhotos(Prod.Variations[j].Photo, PathToVariation)

	}

}
