package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/hm"
)

func EditHM() {
	// Создать оьбъект переводчика
	files, err := ioutil.ReadDir("internal/settings/json/")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(files); i += 2 {
		file := files[i]
		// fmt.Println(file.Name())

		filenameReplace := strings.ReplaceAll(file.Name(), ".json", "")

		filenameReplaces := strings.Split(filenameReplace, "_")
		filenameReplace = filenameReplaces[2]

		// FilePatch := "internal/settings/out/" + filenameReplace + ".json"
		filenameReplace = strings.ReplaceAll(filenameReplace, ".json", "")
		FilePatch := fmt.Sprintf("internal/settings/out/hm_%d_%s", i, filenameReplace)
		fmt.Println(i, FilePatch)

		// read file
		data, err := os.ReadFile("internal/settings/json/" + file.Name())
		if err != nil {
			panic(err)
		}

		var varient bases.Variety2
		err = json.Unmarshal(data, &varient)
		if err != nil {
			panic(err)
		}

		varient.SaveJson(FilePatch)
	}
}

// Редактировать наличие всех товаров
func EditIsExitHM() {
	FolderInput := "internal/settings/hm_output2/"
	FolderOutput := "internal/settings/hm_output3/"

	// // Создать оьбъект переводчика
	// Adding, ErrNewTranslate := wcprod.NewTranslate()
	// if ErrNewTranslate != nil {
	// 	panic(ErrNewTranslate)
	// }
	// fmt.Println(Adding)
	files, err := ioutil.ReadDir(FolderInput)
	if err != nil {
		log.Fatal(err)
	}

	mapingSize := make(map[string]int)
	for i, file := range files {
		// read file
		data, err := os.ReadFile(FolderInput + file.Name())
		if err != nil {
			panic(err)
		}

		// output file
		filenameReplace := strings.ReplaceAll(file.Name(), ".json", "")
		filenameReplaces := strings.Split(filenameReplace, "_")
		filenameReplace = filenameReplaces[2]
		filenameReplace = strings.ReplaceAll(filenameReplace, ".json", "")
		FilePatch := fmt.Sprintf(FolderOutput+"zara_%d_%s", i, filenameReplace)
		FilePatch = strings.ReplaceAll(FilePatch, ".json.json", ".json")
		// fmt.Println(i, FilePatch)

		//
		var varient bases.Variety2
		err = json.Unmarshal(data, &varient)
		if err != nil {
			panic(err)
		}

		for j := range varient.Product {
			NewProds, ErrAvailabilityProduct := hm.AvailabilityProduct(varient.Product[j])
			if ErrAvailabilityProduct != nil {
				panic(ErrAvailabilityProduct)
			}
			fmt.Println(bases.ProdStr(NewProds))
		}

		// for j := range varient.Product {
		// 	for jj := range varient.Product[j].Item {
		// 		for jjj := range varient.Product[j].Item[jj].Size {
		// 			size := varient.Product[j].Item[jj].Size[jjj].Val
		// 			size = strings.ToUpper(size)
		// 			mapingSize[size]++
		// 		}
		// 	}
		// }
		// for j := range varient.Product {
		// 	// var ErrorTranstate error
		// 	// varient.Product[j], ErrorTranstate = Adding.YandexColorRus(varient.Product[j])
		// 	// if ErrorTranstate != nil {
		// 	// 	Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
		// 	// 	varient.Product[j], _ = Adding.YandexColorRus(varient.Product[j])
		// 	// }
		// 	for jj := range varient.Product[j].Item {

		// 		for jjj := range varient.Product[j].Item[jj].Size {
		// 			size := varient.Product[j].Item[jj].Size[jjj].Val
		// 			size = strings.ToUpper(size)
		// 			varient.Product[j].Item[jj].Size[jjj].Val = size
		// 		}
		// 	}
		// 	fmt.Println(i, j)
		// }
		varient.SaveJson(FilePatch)
	}
	for k := range mapingSize {
		fmt.Println(k)
	}
}
