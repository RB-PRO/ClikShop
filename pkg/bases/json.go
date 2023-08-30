package bases

import (
	"bytes"
	"encoding/json"
	"os"
)

func (variety Variety2) SaveJson(filename string) error {
	f, ErrCreateFile := os.Create(filename + ".json")
	if ErrCreateFile != nil {
		return ErrCreateFile
	}
	// as_json, ErrMarshalIndent := json.MarshalIndent(variety, "", "\t")
	as_json, ErrMarshalIndent := MarshalMy(variety)
	if ErrMarshalIndent != nil {
		return ErrMarshalIndent
	}
	f.Write(as_json)
	f.Close()
	return nil
}

func MarshalMy(i interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(i)
	return bytes.TrimRight(buffer.Bytes(), "\n"), err
}
