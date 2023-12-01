package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
)

// ////////////////////////////
func EditSS_FilesOfSize3() {
	FolderInput := "internal/settings/SS2/"
	FolderOutput := "internal/settings/SS3/"

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

	// vievall := "view-all"

	filesName := make(map[string]bool)
	prodsmap := make(map[string]bool)
	var ALLprods []bases.Product2
	for _, file := range files {
		filenameReplace := file.Name() // output file
		filenameReplace = strings.ReplaceAll(filenameReplace, "2.json2.json", "")
		filenameReplace = strings.ReplaceAll(filenameReplace, ".json2.json", "")
		filenameReplace = strings.ReplaceAll(filenameReplace, ".json", "")
		if ok, _ := filesName[filenameReplace]; ok {
			continue
		}
		filesName[filenameReplace] = true

		//FilePatch := fmt.Sprintf(FolderOutput+"%s", filenameReplace)
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

		var prods []bases.Product2
		for _, prod := range varient.Product {
			// if len(prod.Article) == 10 {
			// 	prod.Article = prod.Article[:7]
			// }
			// prod.Img = bases.EditIMG(prod)

			if _, ok := prodsmap[prod.Article]; !ok {
				prodsmap[prod.Article] = true

				if prod.Description.Rus == "" {
					prod.Description.Rus = "."
				}

				for jj := range prod.Cat {
					if prod.Cat[jj].Slug == "" {
						prod.Cat[jj].Slug = bases.Name2Slug(prod.Cat[jj].Name)
						prod.Cat[jj].Slug = bases.Translit(prod.Cat[jj].Slug)
					}
				}

				tecItemColor := make(map[string]int)
				for iItem, colorItem := range prod.Item {
					if _, ok := tecItemColor[colorItem.ColorRus]; !ok {
						tecItemColor[colorItem.ColorRus]++
					} else {
						prod.Item[iItem].ColorRus += strconv.Itoa(tecItemColor[colorItem.ColorRus])
						prod.Item[iItem].ColorCode += strconv.Itoa(tecItemColor[colorItem.ColorRus])
						prod.Item[iItem].ColorEng += strconv.Itoa(tecItemColor[colorItem.ColorRus])
						tecItemColor[colorItem.ColorRus]++
					}
				}

				prods = append(prods, prod)
				ALLprods = append(ALLprods, prod)
			}
		}

		size := 30
		var SubSlice_j int
		for SubSlice_i := 0; SubSlice_i < len(prods); SubSlice_i += size {
			SubSlice_j += size
			if SubSlice_j > len(prods) {
				SubSlice_j = len(prods)
			}

			// Подслайс. Работаем именно с подслайсами, чтобы не перегружать оперативку
			SubSlice := prods[SubSlice_i:SubSlice_j]

			fmt.Printf(FolderOutput+"%s_%d-%d.json qweqwe\n", strings.ReplaceAll(filenameReplace, "internal/settings/SS2/", ""), SubSlice_i, SubSlice_j)
			FilePatch := fmt.Sprintf(FolderOutput+"%s_%d-%d", filenameReplace, SubSlice_i, SubSlice_j)
			bases.Variety2{SubSlice}.SaveJson(FilePatch)
		}

		//varient.SaveJson(FilePatch)
		// break
	}
	fmt.Println(">>>", len(ALLprods))
	bases.Variety2{ALLprods}.SaveJson(FolderOutput + "ALL")

}
