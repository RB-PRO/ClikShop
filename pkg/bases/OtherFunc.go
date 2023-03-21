package bases

import (
	"io"
	"os"
	"strings"
)

// Удалить дубликаты в слайсе
func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// Перевести /sweaters/CKvXARDQ1wHiAgIBAg.zso в sweaters
func FormingColorEng(input string) (output string) {
	input = strings.ReplaceAll(input, " ", "-")
	input = strings.ReplaceAll(input, "'", "")
	input = strings.ReplaceAll(input, "/", "_")
	output = strings.ToLower(input)
	return output
}

// Словарь, который используется для Name в GenderLabel
// и
// роидетльской категории. Например Женщины/woman
//
//	Функция принимает Woman[или]woman, а отдаёт Женщины
func GenderBook(name, slug string) (string, string, bool) {
	nameLower := strings.ToLower(name) // Сделать нижний шрифт
	switch nameLower {
	case "women":
		return "Женщины", "women", true
	case "man":
		return "Мужчины", "man", true
	case "men":
		return "Мужчины", "man", true
	case "kid":
		return "Дети", "kid", true
	default:
		return "Унисекс", "unisex", false
	}
}

// Получение значение из файла
func DataFile(filename string) (string, error) {
	// Открыть файл
	fileToken, errorToken := os.Open(filename)
	if errorToken != nil {
		return "", errorToken
	}

	// Прочитать значение файла
	data := make([]byte, 64)
	n, err := fileToken.Read(data)
	if err == io.EOF { // если конец файла
		return "", errorToken
	}
	fileToken.Close() // Закрытие файла

	return string(data[:n]), nil
}
