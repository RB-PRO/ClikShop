package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"ClikShop/common/bases"
	"ClikShop/common/hm"
	"ClikShop/common/transrb"
	"ClikShop/common/wcprod"
	"github.com/cheggaaa/pb"
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

// ////////////////////////////
// комплексная работа со всеми файлами hm
func EditHM_FilesOfSize() {
	FolderInput := "internal/settings/hm_output5/"
	FolderOutput := "internal/settings/hm_output6/"

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

	vievall := "view-all"

	sku := make(map[string]int)
	filesName := make(map[string]bool)
	for ifile, file := range files {
		if ifile >= 51 && ifile <= 56 {
			continue
		}
		if strings.Contains(file.Name(), vievall) {
			continue
		}
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

			}
			// varient.Product[j], _ = hm.AvailabilityProduct(varient.Product[j])
			sku[varient.Product[j].Article]++

			Bar.Increment()
		}
		Bar.Finish()
		varient.SaveJson(FilePatch)
	}
}

// ////////////////////////////
// комплексная работа со всеми файлами hm
func EditHM_FilesOfSize2() {
	FolderInput := "internal/settings/md_output4/"
	FolderOutput := "internal/settings/md_output5/"

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

	// vievall := "view-all"

	filesName := make(map[string]bool)
	for _, file := range files {
		// if ifile >= 51 && ifile <= 56 {
		// 	continue
		// }
		// if strings.Contains(file.Name(), vievall) {
		// 	continue
		// }
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
			// varient.Product[j].Img = bases.EditIMG(varient.Product[j])
			// var prod bases.Product2 = varient.Product[j]

			// prod, _ = hm.VariableProduct2(prod)

			// prod, _ = hm.AvailabilityProduct(prod)

			// mapcoloraval := make(map[string][]bases.Size)
			// for _, v := range prod.Item {
			// 	// fmt.Println(v, v.Size)
			// 	mapcoloraval[v.ColorEng] = v.Size
			// }
			// // fmt.Println(mapcoloraval)
			// for i := range varient.Product[j].Item {
			// 	// prod.Item[i].Size = mapcoloraval[prod.Item[i].ColorEng]
			// 	copy(prod.Item[i].Size, mapcoloraval[prod.Item[i].ColorEng])
			// }
			// varient.Product[j] = prod

			// }
			// varient.Product[j], _ = hm.AvailabilityProduct(varient.Product[j])

			Name, ErrorTranstate := Adding.YandexDeskription(varient.Product[j].Name)
			if ErrorTranstate != nil {
				fmt.Println(ErrorTranstate)
				Adding.Tr, _ = transrb.New(Adding.Tr.FolderID, Adding.Tr.OAuthToken)
				Name, _ = Adding.YandexDeskription(varient.Product[j].Name)
			}
			varient.Product[j].Name = Name

			Bar.Increment()
			// if j == 10 {
			// 	break
			// }
		}
		Bar.Finish()
		varient.SaveJson(FilePatch)
		// break
	}
}

// ////////////////////////////
func EditHM_FilesOfSize3() {
	FolderInput := "internal/settings/hm_output10/"
	FolderOutput := "internal/settings/hm_output11/"

	files, err := ioutil.ReadDir(FolderInput)
	if err != nil {
		log.Fatal(err)
	}

	// vievall := "view-all"

	filesName := make(map[string]bool)
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

		fmt.Println(file.Name())
		// Bar := pb.StartNew(len(varient.Product))
		for j := range varient.Product {
			gender := varient.Product[j].GenderLabel
			varient.Product[j].Cat[1].Name = gender
			varient.Product[j].Cat[1].Slug = gender

			var description string
			for k, v := range varient.Product[j].Specifications {
				description += "\n" + k + ": " + v
			}
			varient.Product[j].Description.Rus += description
			// varient.Product[j].Img = bases.EditIMG(varient.Product[j])
			// var prod bases.Product2 = varient.Product[j]

			// prod, _ = hm.VariableProduct2(prod)

			// prod, _ = hm.AvailabilityProduct(prod)

			// mapcoloraval := make(map[string][]bases.Size)
			// for _, v := range prod.Item {
			// 	// fmt.Println(v, v.Size)
			// 	mapcoloraval[v.ColorEng] = v.Size
			// }
			// // fmt.Println(mapcoloraval)
			// for i := range varient.Product[j].Item {
			// 	// prod.Item[i].Size = mapcoloraval[prod.Item[i].ColorEng]
			// 	copy(prod.Item[i].Size, mapcoloraval[prod.Item[i].ColorEng])
			// }
			// varient.Product[j] = prod

			// }
			// varient.Product[j], _ = hm.AvailabilityProduct(varient.Product[j])

			// Bar.Increment()
		}
		// Bar.Finish()
		varient.SaveJson(FilePatch)
		// break
	}
}
