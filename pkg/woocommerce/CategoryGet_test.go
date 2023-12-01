package woocommerce_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/woocommerce"
)

func TestProductsCategories(t *testing.T) {
	consumer_key, _ := dataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := dataFile("secret_key")     // Секретный код пользователя
	// Авторизация
	userWC, _ := woocommerce.New(consumer_key, secret_key)

	// Получить список всех категорий
	cats, errCat := userWC.ProductsCategories()
	if errCat != nil {
		t.Error(errCat)
	}
	if len(cats.Category) == 0 {
		t.Error("Но полученных категорий")
	}
	fmt.Println("Всего найдено", len(cats.Category), "записей.")
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
