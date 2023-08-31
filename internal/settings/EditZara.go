package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	zaratr "github.com/RB-PRO/SanctionedClothing/pkg/ZaraTR"
	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/transrb"
	"github.com/RB-PRO/SanctionedClothing/pkg/wcprod"
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
			Adding.YandexDeskription(varient.Product[j].Description.Eng)

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
