package wcprod

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	gt "github.com/bas24/googletranslatefree"
)

func ProductTranslate(prod bases.Product2) bases.Product2 {

	prod.Description.Eng = strings.ReplaceAll(prod.Description.Eng, "\t", "")
	prod.Description.Eng = strings.ReplaceAll(prod.Description.Eng, "#", "")
	prod.Description.Rus, _ = gt.Translate(prod.Description.Eng, "en", "ru")
	prod.Name, _ = gt.Translate(prod.Name, "en", "ru")
	prod.FullName, _ = gt.Translate(prod.FullName, "en", "ru")
	prod.FullName = strings.ReplaceAll(prod.FullName, "Артикул:", "")

	// Категории
	prod.Cat[1].Name, _ = gt.Translate(prod.Cat[1].Name, "en", "ru")
	prod.Cat[2].Name, _ = gt.Translate(prod.Cat[2].Name, "en", "ru")
	prod.Cat[3].Name, _ = gt.Translate(prod.Cat[3].Name, "en", "ru")

	// for indexKey := range prod.Item {
	// 	// Если есть мапа с таким-же ключом, то копируем во вторичную переменную значение этой мапы по ключу
	// 	if entry, ok := prod.Item[indexKey]; ok {

	// 		// Корректируем данные
	// 		// Курс доллара * цена в долларах * наценка + цена доставки
	// 		entry.ColorEng, _ = gt.Translate(entry.ColorEng, "en", "ru")

	// 		// Обновляем данные
	// 		prod.Item[indexKey] = entry
	// 	}
	// }

	//tr := translate.New("trnsl.1.1.20170505T201046Z.765061fd7d327f2f.c80d8b95dd956de79d7f9537011fcd3cc802e6e2")
	//tr := translate.New("trnsl.1.1.20191023T124920Z.63524b1f3817bdc2.1719c9be2a2e95a9ce652519943ee104fb9e0a56")
	//tr := translate.New("trnsl.1.1.20190120T184305Z.c3a652a65ff5dac8.3a47d3f48cf9619b3a0d89ad5296f28c220f85ad")

	/*
		response, err := tr.GetLangs("en")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(response.Langs)
			fmt.Println(response.Dirs)
		}

		translation, err := tr.Translate("ru", prod.Description.Eng)
		if err != nil {
			fmt.Println(err)
		} else {
			prod.Description.Rus = translation.Result()
		}
		translation, err = tr.Translate("ru", prod.Name)
		if err != nil {
			fmt.Println(err)
		} else {
			prod.Name = translation.Result()
		}
		translation, err = tr.Translate("ru", prod.FullName)
		if err != nil {
			fmt.Println(err)
		} else {
			prod.FullName = translation.Result()
		}
	*/

	return prod
}

func (woo *WcAdd) YandexCat(InputCat []bases.Cat) ([]bases.Cat, error) {
	// Категории
	var cats []string
	for i := 1; i < len(InputCat); i++ {
		cats = append(cats, InputCat[i].Name)
	}
	TranslateCats, ErorTranslateCat := woo.Tr.Trans(cats)
	if ErorTranslateCat != nil {
		return []bases.Cat{}, ErorTranslateCat
	}
	for ind := 1; ind < len(InputCat); ind++ {
		InputCat[ind].Name = TranslateCats[ind-1]
	}
	return InputCat, nil
}
func (woo *WcAdd) YandexDeskription(InputSrt string) (string, error) {
	// Категории
	// Описание
	InputSrt = strings.ReplaceAll(InputSrt, "\t", "")
	InputSrt = strings.ReplaceAll(InputSrt, "#", "")

	// Первая буква заглавная
	r := []rune(InputSrt)
	if len(r) != 0 {
		r[0] = unicode.ToUpper(r[0])
		InputSrt = string(r)
	}

	TranslateNames, ErorTranslate := woo.Tr.Trans([]string{InputSrt})
	if ErorTranslate != nil {
		return "", ErorTranslate
	}
	if len(TranslateNames) == 1 {
		InputSrt = TranslateNames[0]
	}

	return InputSrt, nil
}

// перевести цвета вариаций
func (woo *WcAdd) YandexColorRus(prod bases.Product2) (bases.Product2, error) {

	ColorRusSlice := make([]string, 0, len(prod.Item))
	ColorRus := make(map[string]string)
	var cout int
	for _, item := range prod.Item {
		if item.ColorRus == "" {
			ColorRusSlice = append(ColorRusSlice, item.ColorEng)
			cout++
		}
	}
	if cout == 0 {
		return prod, nil
	}

	TranslateColorRusSlice, ErorTranslate := woo.Tr.Trans(ColorRusSlice)
	if ErorTranslate != nil {
		return prod, ErorTranslate
	}

	for i := range ColorRusSlice {
		ColorRus[ColorRusSlice[i]] = TranslateColorRusSlice[i]
	}
	for i := range prod.Item {
		if val, ok := ColorRus[prod.Item[i].ColorEng]; ok {
			prod.Item[i].ColorRus = val
		}
	}
	fmt.Println("Перевожу")
	return prod, nil
}

func (woo *WcAdd) YandexTranslate(prod bases.Product2) (bases.Product2, error) {
	// Описание
	prod.Description.Eng = strings.ReplaceAll(prod.Description.Eng, "\t", "")
	prod.Description.Eng = strings.ReplaceAll(prod.Description.Eng, "#", "")

	// Первая буква заглавная
	r := []rune(prod.Description.Eng)
	if len(r) != 0 {
		r[0] = unicode.ToUpper(r[0])
		prod.Description.Eng = string(r)
	}

	// Краткое описание
	prod.FullName = strings.ReplaceAll(prod.FullName, "SKU:", "")

	// Переводим имя
	TranslateNames, ErorTranslate := woo.Tr.Trans([]string{prod.Description.Eng, prod.Name, prod.FullName})
	if ErorTranslate != nil {
		return prod, ErorTranslate
	}
	if len(TranslateNames) == 3 {
		prod.Description.Rus, prod.Name, prod.FullName = TranslateNames[0], TranslateNames[1], TranslateNames[2]
	}
	// fmt.Printf("TranslateNames %d %+#v", len(TranslateNames), TranslateNames)

	// Вариации
	var item []string
	for _, it := range prod.Item {
		item = append(item, it.ColorEng)
	}
	TranslateItem, ErorTranslateItem := woo.Tr.Trans(item)
	if ErorTranslateItem != nil {
		return prod, ErorTranslateItem
	}

	for KeyColor := range prod.Item {
		prod.Item[KeyColor].ColorRus = TranslateItem[KeyColor]
		// for KeySize := range prod.Item[KeyColor].Item {
		// 	prod.Item[KeyColor].Item[KeySize].ColorEng = TranslateItem[KeyColor]
		// }
	}

	// Вторичное описание
	var Spec []string
	for i, v := range prod.Specifications {
		Spec = append(Spec, i, v)
	}
	TransSpec, ErorTransSpec := woo.Tr.Trans(Spec)
	if ErorTransSpec != nil {
		return prod, ErorTransSpec
	}
	NewSpec := make(map[string]string)
	for i := 0; i < len(TransSpec); i += 2 {
		NewSpec[TransSpec[i]] = TransSpec[i+1]
	}
	prod.Specifications = NewSpec

	return prod, nil
}

// Перевод определённых частей товара, а именно:
//   - Название
//   - Описание
//   - Цвета
func (woo *WcAdd) YandexTranslatePart(prod bases.Product2) (bases.Product2, error) {
	// Описание
	prod.Description.Eng = strings.ReplaceAll(prod.Description.Eng, "\t", "")
	prod.Description.Eng = strings.ReplaceAll(prod.Description.Eng, "#", "")

	// Переводим имя
	TranslateNames, ErorTranslate := woo.Tr.Trans([]string{prod.Description.Eng, prod.Name})
	if ErorTranslate != nil {
		return prod, ErorTranslate
	}
	if len(TranslateNames) == 2 {
		prod.Description.Rus, prod.Name = TranslateNames[0], TranslateNames[1]
	}

	// Вариации
	var item []string
	for _, it := range prod.Item {
		item = append(item, it.ColorEng)
	}
	TranslateItem, ErorTranslateItem := woo.Tr.Trans(item)
	if ErorTranslateItem != nil {
		return prod, ErorTranslateItem
	}

	for KeyColor := range prod.Item {
		prod.Item[KeyColor].ColorRus = TranslateItem[KeyColor]
	}

	return prod, nil
}
