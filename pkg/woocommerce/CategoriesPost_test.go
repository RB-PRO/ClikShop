package woocommerce

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/bases"
)

// Тестирование добавления товара и обновления категории
func TestAddCat(t *testing.T) {
	consumer_key, _ := dataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := dataFile("secret_key")     // Секретный код пользователя
	// Авторизация
	userWC, _ := New(consumer_key, secret_key)
	// [{Name:Men Slug:men} {Name:Clothing Slug:clothing} {Name:Hoodies & Sweatshirts Slug:hoodies-sweatshirts} {Name:'47 NHL Slug:47-nhl}]

	categ := make([]bases.Cat, 4)
	categ[0].Name = "Men"
	categ[0].Slug = "men"
	categ[1].Name = "Clothing"
	categ[1].Slug = "clothing"
	categ[2].Name = "Hoodies & Sweatshirts"
	categ[2].Slug = "hoodies-sweatshirts"
	categ[3].Name = "'47 NHL"
	categ[3].Slug = "47-nhl"

	// Получить дерева категорий
	plc, errPLC := userWC.ProductsCategories()
	if errPLC != nil {
		t.Error(errPLC)
	}
	fmt.Println("Сделал массив товаров plc, длиной", len(plc.Category))
	NodeCategoryes := NewCategoryes()                   // Создать категории
	NodeCategoryes, _ = plc.FormingNode(NodeCategoryes) // Обновить в соответствии с массивом данных
	NodeCategoryes.PrintInorderName("-")                // Напечатать дерево категорий
	fmt.Println("Сделал структуру товаров NodeCategoryes")
	// Создать категории для товаров и получить её ID
	idCat1, errorAddCat := userWC.AddCat(NodeCategoryes, categ)
	if errorAddCat != nil {
		t.Error(errorAddCat)
	}
	fmt.Println("NodeCategoryes.Len()", NodeCategoryes.Len())
	// Создать категории для товаров и получить её ID
	idCat2, errorAddCat2 := userWC.AddCat(NodeCategoryes, categ)
	if errorAddCat2 != nil {
		t.Error(errorAddCat2)
	}
	fmt.Println("NodeCategoryes.Len()", NodeCategoryes.Len())
	if idCat1 != idCat2 {
		t.Error("Сперва сделал категорию, потом попытался добавить идентичную и получил другой ID.\nidCat1 -", idCat1, "\nidCat2 -", idCat2)
	}

}

// Тестируем добавление категории
func TestAddCat_WC(t *testing.T) {

	consumer_key, _ := dataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := dataFile("secret_key")     // Секретный код пользователя

	// Авторизация
	userWC, _ := New(consumer_key, secret_key)
	cat := MeCat{
		Name: "Кофты",
		Slug: "kofta",
	}
	ParentID, ParentError := userWC.AddCat_WC(cat)
	t.Log(ParentID)
	if ParentError != nil {
		t.Error(ParentError)
	}

	cat = MeCat{
		Name:     "Кофты_подкатегория",
		Slug:     "kofta_child",
		ParentID: ParentID,
	}
	t.Log(userWC.AddCat_WC(cat))

	// Должно вывестись два числа.
	// Первое - ID категории, второе - ID подкатегории
}

func TestFormingNode(t *testing.T) {

	consumer_key, _ := dataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := dataFile("secret_key")     // Секретный код пользователя
	// Авторизация
	userWC, _ := New(consumer_key, secret_key)

	// Получить дерево категорий
	plc, errPLC := userWC.ProductsCategories()
	if errPLC != nil {
		t.Error(errPLC)
	}

	NodeCategoryes := NewCategoryes() // Создать категории
	var errorNodeCategoryes error
	NodeCategoryes, errorNodeCategoryes = plc.FormingNode(NodeCategoryes) // Обновить в соответствии с массивом данных
	if errorNodeCategoryes != nil {
		t.Error(errorNodeCategoryes)
	}

	NodeCategoryes.PrintInorderName("-") // Напечатать дерево категорий

}

// Получение значение из файла
func dataFile(filename string) (string, error) {
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
