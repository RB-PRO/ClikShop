package bases

import (
	"fmt"
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
	input = strings.ReplaceAll(input, "/", "-")
	input = strings.ReplaceAll(input, "--", "-")
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
	case "woman":
		return "Женщины", "women", true
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

// Перевести цвет в ссылку для цвета
//
// Названите в ярлык
func Name2Slug(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "-")
	str = strings.ReplaceAll(str, "/", "-")
	str = strings.ReplaceAll(str, "--", "-")
	// В идеале добавить регулярные выражения для отсеивания лишнего
	return str
}

// Вернуть строку в виде продукта в читаемом виде
func ProdStr(Prod Product2) (str string) {
	str += fmt.Sprintf("Название товара: '%v'\n", Prod.Name)
	str += fmt.Sprintf("Полное Название товара: '%v'\n", Prod.FullName)

	str += fmt.Sprintf("Ссылка на товар: '%v'\n", Prod.Link)
	str += fmt.Sprintf("Артикул: '%v'\n", Prod.Article)
	str += fmt.Sprintf("Производитель: '%v'\n", Prod.Manufacturer)
	str += fmt.Sprintf("Гендер: '%v'\n", Prod.GenderLabel)
	str += fmt.Sprintf("Все Размеры: '%v'\n", Prod.Size)
	str += fmt.Sprintf("Описание: '%v'\n", Prod.Description.Eng)
	for IndexSize, Sizen := range Prod.Item {
		str += fmt.Sprintf("- %d Вариация с цветом: '%s'\n", IndexSize, Sizen.ColorEng)
		str += fmt.Sprintf("--- Код цвета: %s\n", Sizen.ColorCode)
		str += fmt.Sprintf("--- Цена: %v\n", Sizen.Price)
		str += fmt.Sprintf("--- Размеры: %v\n", Sizen.Size)
		str += fmt.Sprintf("--- Картинки: %s\n", Sizen.Image)
	}

	return str
}
