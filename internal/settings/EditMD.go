package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/transrb"
	"github.com/RB-PRO/ClikShop/pkg/wcprod"
	"github.com/cheggaaa/pb"
)

func EditMD() {
	FolderInput := "internal/settings/md_output2/"
	FolderOutput := "internal/settings/md_output3/"

	// Создать оьбъект переводчика
	Adding, ErrNewTranslate := wcprod.NewTranslate()
	if ErrNewTranslate != nil {
		panic(ErrNewTranslate)
	}
	fmt.Println(Adding)

	files, err := ioutil.ReadDir(FolderInput)
	if err != nil {
		log.Fatal(err)
	}

	sku := make(map[string]int)
	filesName := make(map[string]bool)
	products := make([]bases.Product2, 0)
	for _, file := range files {
		filenameReplace := file.Name() // output file
		filenameReplace = strings.ReplaceAll(filenameReplace, "2.json2.json", "")
		filenameReplace = strings.ReplaceAll(filenameReplace, ".json2.json", "")
		filenameReplace = strings.ReplaceAll(filenameReplace, ".json", "")
		if ok, _ := filesName[filenameReplace]; ok {
			continue
		}
		filesName[filenameReplace] = true

		FilePatch := fmt.Sprintf(FolderOutput+"%s", filenameReplace)
		// fmt.Println(i, FilePatch)

		// read file
		data, err := os.ReadFile(FolderInput + file.Name())
		if err != nil {
			panic(err)
		}
		var varient bases.Variety2
		err = json.Unmarshal(data, &varient)
		if err != nil {
			panic(err)
		}

		Bar := pb.StartNew(len(varient.Product))

		for j := range varient.Product {
			Bar.Prefix(file.Name())
			// for jj := range varient.Product[j].Item {
			// 	for jjj := range varient.Product[j].Item[jj].Size {
			// 		mapingSize[varient.Product[j].Item[jj].Size[jjj].Val]++
			// 	}
			// }
			// fmt.Println(file.Name(), "[", j, "/", len(varient.Product), "]")

			if _, ok := sku[varient.Product[j].Article]; !ok {
				//fmt.Println(varient.Product[j].Name, RussianSymb(varient.Product[j].Name), len(varient.Product[j].Name)/2, RussianSymb(varient.Product[j].Name) < len(varient.Product[j].Name)/2)

				var ErrorTranstate error
				Name := varient.Product[j].Name
				if RussianSymb(Name) < len([]rune(Name))/2 {
					Bar.Prefix(file.Name() + " Перевожу")
					varient.Product[j], ErrorTranstate = Adding.YandexTranslate(varient.Product[j])
					if ErrorTranstate != nil {
						Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
						varient.Product[j], _ = Adding.YandexTranslate(varient.Product[j])
					}
				}

				products = append(products, varient.Product[j])

			}
			// varient.Product[j], _ = hm.AvailabilityProduct(varient.Product[j])
			sku[varient.Product[j].Article]++

			Bar.Increment()
		}
		Bar.Finish()
		varient.SaveJson(FilePatch)
	}

	fmt.Println(len(products))
}
