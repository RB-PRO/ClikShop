package bases

import (
	"encoding/json"
	"os"
)

func (variety Variety2) SaveJson(filename string) error {
	f, ErrCreateFile := os.Create(filename + ".json")
	if ErrCreateFile != nil {
		return ErrCreateFile
	}
	as_json, ErrMarshalIndent := json.MarshalIndent(variety, "", "\t")
	if ErrMarshalIndent != nil {
		return ErrMarshalIndent
	}
	f.Write(as_json)
	f.Close()
	return nil
}
