package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"ClikShop/common/bases"
)

func PrintSizes() {
	FolderInput := "internal/settings/ty/"
	// FolderOutput := "internal/settings/SS2/"

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
	for _, file := range files {
		// read file
		data, err := os.ReadFile(FolderInput + file.Name())
		if err != nil {
			panic(err)
		}

		// output file
		// filenameReplace := strings.ReplaceAll(file.Name(), ".json", "")
		// filenameReplaces := strings.Split(filenameReplace, "_")
		// filenameReplace = filenameReplaces[2]
		// filenameReplace = strings.ReplaceAll(filenameReplace, ".json", "")
		// FilePatch := fmt.Sprintf(FolderOutput+"zara_%d_%s", i, filenameReplace)
		// FilePatch = strings.ReplaceAll(FilePatch, ".json.json", ".json")
		// fmt.Println(i, FilePatch)

		//
		var varient bases.Variety2
		err = json.Unmarshal(data, &varient)
		if err != nil {
			panic(err)
		}

		for j := range varient.Product {
			// for jj := range varient.Product[j].Item {
			// 	for jjj := range varient.Product[j].Item[jj].Size {
			// 		mapingSize[varient.Product[j].Item[jj].Size[jjj].Val]++
			// 	}
			// }
			mapingSize[varient.Product[j].Manufacturer]++
		}

		// for j := range varient.Product {
		// 	var ErrorTranstate error
		// 	varient.Product[j], ErrorTranstate = Adding.YandexColorRus(varient.Product[j])
		// 	if ErrorTranstate != nil {
		// 		Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
		// 		varient.Product[j], _ = Adding.YandexColorRus(varient.Product[j])
		// 	}
		// 	fmt.Println(i, j)
		// }
		// varient.SaveJson(FilePatch)
	}
	for k := range mapingSize {
		fmt.Println(k)
		appendfile(k)
	}
}

func UpSizesEditProduct() {
	FolderInput := "internal/settings/zara_output2/"
	FolderOutput := "internal/settings/zara_output3/"

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
			for jj := range varient.Product[j].Item {
				for jjj := range varient.Product[j].Item[jj].Size {
					size := varient.Product[j].Item[jj].Size[jjj].Val
					size = strings.ToUpper(size)
					mapingSize[size]++
				}
			}
		}
		for j := range varient.Product {
			// var ErrorTranstate error
			// varient.Product[j], ErrorTranstate = Adding.YandexColorRus(varient.Product[j])
			// if ErrorTranstate != nil {
			// 	Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
			// 	varient.Product[j], _ = Adding.YandexColorRus(varient.Product[j])
			// }
			for jj := range varient.Product[j].Item {

				for jjj := range varient.Product[j].Item[jj].Size {
					size := varient.Product[j].Item[jj].Size[jjj].Val
					size = strings.ToUpper(size)
					varient.Product[j].Item[jj].Size[jjj].Val = size
				}
			}
			fmt.Println(i, j)
		}
		varient.SaveJson(FilePatch)
	}
	for k := range mapingSize {
		fmt.Println(k)
	}
}

func appendfile(data string) {
	f, err := os.OpenFile("tysizes.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(data + "\n"); err != nil {
		panic(err)
	}
}
