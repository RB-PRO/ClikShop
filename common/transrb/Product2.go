package transrb

import (
	"fmt"
	"strings"
	"unicode"

	"ClikShop/common/bases"
)

func (tr *Translate) TranslateProduct2(prod bases.Product2) (bases.Product2, error) {

	TranslateProduct := prod

	// Описание
	TranslateProduct.Description.Eng = strings.ReplaceAll(TranslateProduct.Description.Eng, "\t", "")
	TranslateProduct.Description.Eng = strings.ReplaceAll(TranslateProduct.Description.Eng, "#", "")

	// Первая буква заглавная
	r := []rune(strings.ToLower(TranslateProduct.Name))
	if len(r) != 0 {
		r[0] = unicode.ToUpper(r[0])
		TranslateProduct.Name = string(r)
	}
	// Краткое описание
	TranslateProduct.FullName = strings.ReplaceAll(TranslateProduct.FullName, "SKU:", "")

	// Название товара
	TranslateName, ErorTranslate := tr.Trans([]string{TranslateProduct.Name})
	if ErorTranslate != nil {
		return TranslateProduct, ErorTranslate
	}
	if len(TranslateName) != 1 {
		return bases.Product2{}, fmt.Errorf("name: len(Name) != 1")
	}
	TranslateProduct.Name = TranslateName[0]

	// Длинное название товара
	if len(TranslateProduct.FullName) != 0 {
		TranslateFullName, ErorTranslate := tr.Trans([]string{TranslateProduct.FullName})
		if ErorTranslate != nil {
			return TranslateProduct, ErorTranslate
		}
		if len(TranslateFullName) != 1 {
			return bases.Product2{}, fmt.Errorf("fullName: len(FullName) != 1")
		}
		TranslateProduct.FullName = TranslateFullName[0]
	}

	// Описание товара
	if len(TranslateProduct.Description.Eng) != 0 {
		TranslateDescription, ErorTranslate := tr.Trans([]string{TranslateProduct.Description.Eng})
		if ErorTranslate != nil {
			return TranslateProduct, ErorTranslate
		}
		if len(TranslateDescription) != 1 {
			return bases.Product2{}, fmt.Errorf("description: len(FullName) != 1")
		}
		TranslateProduct.Description.Rus = TranslateDescription[0]
	}

	// Вариации
	var item []string
	for _, it := range prod.Item {
		item = append(item, it.ColorEng)
	}
	TranslateItem, ErorTranslateItem := tr.Trans(item)
	if ErorTranslateItem != nil {
		return prod, ErorTranslateItem
	}
	for KeyColor := range prod.Item {
		prod.Item[KeyColor].ColorRus = TranslateItem[KeyColor]
	}

	// // Вторичное описание
	// var Spec []string
	// for i, v := range prod.Specifications {
	// 	Spec = append(Spec, i, v)
	// }
	// TransSpec, ErorTransSpec := tr.Trans(Spec)
	// if ErorTransSpec != nil {
	// 	return prod, ErorTransSpec
	// }
	// NewSpec := make(map[string]string)
	// for i := 0; i < len(TransSpec); i += 2 {
	// 	NewSpec[TransSpec[i]] = TransSpec[i+1]
	// }
	// prod.Specifications = NewSpec

	return TranslateProduct, nil
}
