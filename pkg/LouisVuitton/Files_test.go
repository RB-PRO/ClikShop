package louisvuitton

import "testing"

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
