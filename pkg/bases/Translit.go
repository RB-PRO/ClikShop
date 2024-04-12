package bases

import (
	"github.com/mozillazg/go-unidecode"
)

// Сделать транслит на Английский
func Translit(str string) string {
	return unidecode.Unidecode(str)
}
