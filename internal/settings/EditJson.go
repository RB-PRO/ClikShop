// Пакет для всяких микропрограмм серверных
package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"ClikShop/common/bases"
)

// go build -o settings cmd/main/bank.go
// settings MD_SubSlice_3_2000-3000.json  MD_SubSlice_5_4000-5000.json  MD_SubSlice_7_6000-7000.json MD_SubSlice_2_1000-2000.json  MD_SubSlice_4_3000-4000.json  MD_SubSlice_6_5000-6000.json  MD_SubSlice_8_7000-8000.json
func EditJson1() {
	if len(os.Args) == 1 {
		log.Fatal("Подайте на вход список файлов")
	}
	files := os.Args[1:]
	fmt.Println(files)

	for _, FileName := range files {

		// read file
		data, err := os.ReadFile(FileName)
		if err != nil {
			panic(err)
		}

		var varient bases.Variety2

		// unmarshall it
		err = json.Unmarshal(data, &varient)
		if err != nil {
			panic(err)
		}

		for i := range varient.Product {
			var ImageMain []string
			// if len(varient.Product[i].Item) != 0 {
			// 	if len(varient.Product[i].Item[0].Image) != 0 {
			// 		ImageMain = varient.Product[i].Item[0].Image[0]
			// 	}
			// }
			for _, item := range varient.Product[i].Item {
				for _, image := range item.Image {
					ImageMain = append(ImageMain, image)
				}

			}
			varient.Product[i].Img = ImageMain
		}
		varient.SaveJson("out/" + FileName)
	}
}

func EditJson() {
	files, err := ioutil.ReadDir("internal/settings/json/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())

		// read file
		data, err := os.ReadFile("internal/settings/json/" + file.Name())
		if err != nil {
			panic(err)
		}

		var varient bases.Variety2

		// unmarshall it
		err = json.Unmarshal(data, &varient)
		if err != nil {
			panic(err)
		}
		varient.SaveJson("internal/settings/out/" + file.Name())
	}
}
