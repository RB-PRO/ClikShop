package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
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
