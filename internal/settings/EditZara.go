package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"

	zaratr "ClikShop/common/ZaraTR"
	"ClikShop/common/bases"
	"ClikShop/common/transrb"
	"ClikShop/common/wcprod"
)

func EditZara() {
	// Создать оьбъект переводчика
	Adding, ErrNewTranslate := wcprod.NewTranslate()
	if ErrNewTranslate != nil {
		panic(ErrNewTranslate)
	}
	files, err := ioutil.ReadDir("internal/settings/jsonzara/")
	if err != nil {
		log.Fatal(err)
	}

	mapFileNameCat := make(map[string][]bases.Cat)
	CatArr, _ := zaratr.CatCycle2() // Получить все категории
	for i, cat := range CatArr.Items {
		line, ErrorLine := zaratr.LoadLine(fmt.Sprintf("%v", cat.RedirectCategoryID))
		if ErrorLine != nil {
			fmt.Println(ErrorLine)
		}
		ProductsLine := make([]zaratr.CommercialComponents, 0)
		for _, ProductGroups := range line.ProductGroups {
			for _, Elements := range ProductGroups.Elements {
				for _, CommercialComponents := range Elements.CommercialComponents {
					// if cout >= 10 { // Максимум 10 товаров в категории
					// 	break
					// }
					CommercialComponents.Cat = cat.Cat
					CommercialComponents.Gender = cat.Gender
					ProductsLine = append(ProductsLine, CommercialComponents)
				}
			}
		}
		if len(ProductsLine) == 0 {
			continue
		}
		FileName := ProductsLine[0].Cat[len(ProductsLine[0].Cat)-1].Slug

		fmt.Println(i, FileName, cat.Cat)

		mapFileNameCat[FileName] = cat.Cat
	}

	for i, file := range files {
		// fmt.Println(file.Name())

		// read file
		data, err := os.ReadFile("internal/settings/jsonzara/" + file.Name())
		if err != nil {
			panic(err)
		}

		filenameReplace := strings.ReplaceAll(file.Name(), ".json", "")

		filenameReplaces := strings.Split(filenameReplace, "_")
		filenameReplace = filenameReplaces[2]

		fmt.Println(i, filenameReplace, mapFileNameCat[filenameReplace])

		var varient bases.Variety2
		err = json.Unmarshal(data, &varient)
		if err != nil {
			panic(err)
		}

		// varient.Product[0].Cat = mapFileNameCat[filenameReplace]
		ProdTranslateCat, _ := Adding.YandexCat(mapFileNameCat[filenameReplace])
		for j := range varient.Product {
			varient.Product[j].Cat = ProdTranslateCat // mapFileNameCat[filenameReplace]
		}

		varient.SaveJson("internal/settings/out/" + file.Name())
	}
}

func EditZara2() {
	// Создать оьбъект переводчика
	Adding, ErrNewTranslate := wcprod.NewTranslate()
	if ErrNewTranslate != nil {
		panic(ErrNewTranslate)
	}
	fmt.Println(Adding)
	files, err := ioutil.ReadDir("internal/settings/jsonzara2/")
	if err != nil {
		log.Fatal(err)
	}

	for i, file := range files {
		// fmt.Println(file.Name())

		filenameReplace := strings.ReplaceAll(file.Name(), ".json", "")

		filenameReplaces := strings.Split(filenameReplace, "_")
		filenameReplace = filenameReplaces[2]

		// FilePatch := "internal/settings/out/" + filenameReplace + ".json"
		filenameReplace = strings.ReplaceAll(filenameReplace, ".json", "")
		FilePatch := fmt.Sprintf("internal/settings/out/zara_%d_%s", i, filenameReplace)
		fmt.Println(i, FilePatch)

		// read file
		data, err := os.ReadFile("internal/settings/jsonzara2/" + file.Name())
		if err != nil {
			panic(err)
		}

		var varient bases.Variety2
		err = json.Unmarshal(data, &varient)
		if err != nil {
			panic(err)
		}

		for j := range varient.Product {
			// varient.Product[j].Cat = ProdTranslateCat // mapFileNameCat[filenameReplace]
			// Adding.YandexDeskription(varient.Product[j].Description.Eng)

			DescriptRus, ErrorTranstate := Adding.YandexDeskription(varient.Product[j].Description.Eng)
			if ErrorTranstate != nil {
				Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
				DescriptRus, _ = Adding.YandexDeskription(varient.Product[j].Description.Eng)
			}

			fmt.Println(i, j, DescriptRus)

			varient.Product[j].Description.Rus = DescriptRus
		}

		varient.SaveJson(FilePatch)
	}
}

func EditZaraColorRus() {
	// Создать оьбъект переводчика
	Adding, ErrNewTranslate := wcprod.NewTranslate()
	if ErrNewTranslate != nil {
		panic(ErrNewTranslate)
	}
	fmt.Println(Adding)
	files, err := ioutil.ReadDir("internal/settings/zara_output/")
	if err != nil {
		log.Fatal(err)
	}

	for i, file := range files {
		// fmt.Println(file.Name())

		filenameReplace := strings.ReplaceAll(file.Name(), ".json", "")

		filenameReplaces := strings.Split(filenameReplace, "_")
		filenameReplace = filenameReplaces[2]

		// FilePatch := "internal/settings/out/" + filenameReplace + ".json"
		filenameReplace = strings.ReplaceAll(filenameReplace, ".json", "")
		FilePatch := fmt.Sprintf("internal/settings/zara_output2/zara_%d_%s.json", i, filenameReplace)
		fmt.Println(i, FilePatch)

		// read file
		data, err := os.ReadFile("internal/settings/zara_output/" + file.Name())
		if err != nil {
			panic(err)
		}

		var varient bases.Variety2
		err = json.Unmarshal(data, &varient)
		if err != nil {
			panic(err)
		}

		for j := range varient.Product {
			var ErrorTranstate error
			// varient.Product[j], ErrorTranstate = Adding.YandexColorRus(varient.Product[j])
			// if ErrorTranstate != nil {
			// 	Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
			// 	varient.Product[j], _ = Adding.YandexColorRus(varient.Product[j])
			// }
			if !IsRussian(varient.Product[j].Name) {
				varient.Product[j].Name, ErrorTranstate = Adding.YandexDeskription(varient.Product[j].Name)
				if ErrorTranstate != nil {
					Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
					varient.Product[j].Name, _ = Adding.YandexDeskription(varient.Product[j].Name)
				}
				fmt.Println(i, j)
			}
		}

		varient.SaveJson(FilePatch)
	}
}

// Если эта фраза на русском языке
func IsRussian(str string) bool {
	var russianCount, englishCount int
	for _, char := range str {
		if unicode.Is(unicode.Latin, char) {
			englishCount++
		}
		if unicode.Is(unicode.Cyrillic, char) {
			russianCount++
		}
	}
	return russianCount > englishCount
}

// Если эта фраза на русском языке
func RussianSymb(str string) (russianCount int) {
	for _, char := range str {
		if unicode.Is(unicode.Cyrillic, char) {
			russianCount++
		}
	}
	return russianCount
}
