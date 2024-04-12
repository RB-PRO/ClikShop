package bases

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	case "boy":
		return "Мальчики", "boy", true
	case "girl":
		return "Девочки", "girl", true
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
	str = strings.ReplaceAll(str, "--", "-")
	str = strings.TrimSpace(str)
	// В идеале добавить регулярные выражения для отсеивания лишнего
	return str
}

// Перевести ссылку в название
// Например:
// `urune-gore-satin-al`
// в
// `Urune Gore Satin Al`
func Slug2Name(str string) string {
	str = strings.ReplaceAll(str, "-", " ")
	str = cases.Title(language.Und, cases.NoLower).String(str)
	return str
}

// Вернуть строку в виде продукта в читаемом виде
func ProdStr(Prod Product2) (str string) {
	str += fmt.Sprintf("Название товара: '%v'\n", Prod.Name)
	str += fmt.Sprintf("Полное Название товара: '%v'\n", Prod.FullName)
	str += fmt.Sprintf("- Картинки: %s\n", "\n-"+strings.Join(Prod.Img, "\n-"))

	str += fmt.Sprintf("Ссылка на товар: '%v'\n", Prod.Link)
	str += fmt.Sprintf("Артикул: '%v'\n", Prod.Article)
	str += fmt.Sprintf("Производитель: '%v'\n", Prod.Manufacturer)
	str += fmt.Sprintf("Гендер: '%v'\n", Prod.GenderLabel)
	str += fmt.Sprintf("Все Размеры: '%v'\n", Prod.Size)
	str += fmt.Sprintf("Описание Eng: '%v'\n", Prod.Description.Eng)
	str += fmt.Sprintf("Описание Rus: '%v'\n", Prod.Description.Rus)
	for IndexSize, Sizen := range Prod.Item {
		str += fmt.Sprintf("- %d Вариация с цветом: '%s'\n", IndexSize+1, Sizen.ColorEng)
		str += fmt.Sprintf("--- Код цвета: %s\n", Sizen.ColorCode)
		str += fmt.Sprintf("--- Цена: %v\n", Sizen.Price)
		str += fmt.Sprintf("--- Ссылка: %v\n", Sizen.Link)
		str += fmt.Sprintf("--- Размеры: %v\n", Sizen.Size)
		str += fmt.Sprintf("--- Картинки: %s\n", "\n----"+strings.Join(Sizen.Image, "\n----"))
	}
	return str
}

// Добавить общий список всех размеров в товаре на основании тех размеров, которые содержатся в вариациях данного товара
func EditProdSize(Prod Product2) []string {
	// SizesAll := make([]string, 0, 20)
	var SizesAll []string

	for i := range Prod.Item {
		for j := range Prod.Item[i].Size {
			SizesAll = append(SizesAll, Prod.Item[i].Size[j].Val)
		}
	}
	SizesAll = RemoveDuplicateStr(SizesAll)

	return SizesAll
}

// Оставить в строке только символы и пробелы заменить на нижнее подчёркивание
func KeepLettersAndSpaces(str string) (result string) {
	str = strings.ToLower(str)
	for _, char := range str {
		if unicode.IsLetter(char) {
			result += string(char)
		}
		if unicode.IsSpace(char) {
			result += "_"
		}
	}
	return result
}

// "Мягкий" выход из программы
func ExitSoft() {
	fmt.Println("Press 'q' to quit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" {
			break
		} else {
			fmt.Println("Press 'q' to quit")
		}
	}
}

// Редактирование цены по товарам
func EditCoast(prod Product2, usd float64, walrus float64, delivery int) Product2 {
	for indexKey := range prod.Item {
		// Корректируем данные
		// Курс доллара * цена в долларах * наценка + цена доставки
		price := usd*prod.Item[indexKey].Price*walrus + float64(delivery)
		price = EditDecadense(price)
		prod.Item[indexKey].Price = price
	}
	return prod
}

// Редактирование цены в большую сторону
//
// # Округляем цену в большую сторону по десяткам
//
// Если цена была 5225.77, то станет 5230
func EditDecadense(coast float64) float64 {
	return math.Round(coast/10.0) * 10.0
}

// Создать список картинок для создания main-картинки.
//
//	bases.Product2.Img
func EditIMG(prod Product2) (img []string) {
	if len(prod.Img) == 0 {
		for _, item := range prod.Item {
			img = append(img, item.Image...)
			// for _, image := range item.Image {
			// 	img = append(img, image)
			// }
		}
	}
	return img
}

// Обработать товар, где есть вариации с одинаковым цветом
// Если цвета одинаковые, то добавляет в конце -1 -2 -3 и тд
func EditDoubleColors(prod Product2) Product2 {

	// Создаём мапу со слайсами вариаций товаров
	mapColorItems := make(map[string][]ColorItem)
	for _, ct := range prod.Item {
		mapColorItems[ct.ColorCode] = append(mapColorItems[ct.ColorCode], ct)
	}
	if len(mapColorItems) == len(prod.Item) {
		return prod
	}

	// формирование товаров
	prod.Item = make([]ColorItem, 0, len(prod.Item))
	for _, ct := range mapColorItems {
		if len(ct) == 1 {
			prod.Item = append(prod.Item, ct[0])
			continue
		}
		for i, ct := range ct {
			ct.ColorCode = fmt.Sprintf("%s-%d", ct.ColorCode, i+1)
			ct.ColorEng = fmt.Sprintf("%s-%d", ct.ColorEng, i+1)
			prod.Item = append(prod.Item, ct)
		}
	}

	return prod
}

// Редактирование размер для NOSIZE
func EditOneSize(prod Product2) Product2 {
	for indexKey := range prod.Item {

		if len(prod.Item[indexKey].Size) == 1 {
			if prod.Item[indexKey].Size[0].Val == "" {
				prod.Item[indexKey].Size[0].Val = "NOSIZE"
			}
		}

	}
	return prod
}
