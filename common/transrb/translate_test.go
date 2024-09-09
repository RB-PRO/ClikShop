package transrb_test

import (
	"io"
	"os"
	"testing"

	"ClikShop/common/transrb"
)

func TestTrans(t *testing.T) {

	inputStr := []string{"hello"}
	outputStr := "здравствуйте"

	FolderID, _ := DataFile("FolderID")
	OAuthToken, _ := DataFile("OAuthToken")

	tr, err := transrb.New(FolderID, OAuthToken)
	if err != nil {
		t.Error(err)
	}

	answerTranslate, errorTranslate := tr.Trans(inputStr)
	if errorTranslate != nil {
		t.Error(errorTranslate)
	}
	if len(answerTranslate) == 0 {
		t.Error("Массив вывода равен нулю")
	}
	if outputStr != answerTranslate[0] {
		t.Errorf(`Неверный перевод.
Получено:    "%v"
Должно быть: "%v"`, answerTranslate, outputStr)
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
	data := make([]byte, 512)
	n, err := fileToken.Read(data)
	if err == io.EOF { // если конец файла
		return "", errorToken
	}
	fileToken.Close() // Закрытие файла

	return string(data[:n]), nil
}

func TestTrans2(t *testing.T) {

	inputStr := "Downtown Kadın Siyah/bordo Ceket"
	outputStr := "Downtown Women's Black/burgundy Jacket"

	FolderID, _ := DataFile("..\\..\\FolderID")
	OAuthToken, _ := DataFile("..\\..\\OAuthToken")

	tr, err := transrb.New(FolderID, OAuthToken)
	if err != nil {
		t.Error(err)
	}

	answerTranslate, errorTranslate := tr.TransENG(inputStr)
	if errorTranslate != nil {
		t.Error(errorTranslate)
	}
	if len(answerTranslate) == 0 {
		t.Error("Массив вывода равен нулю")
	}
	if outputStr != answerTranslate {
		t.Errorf(`Неверный перевод.
Получено:    "%v"
Должно быть: "%v"`, answerTranslate, outputStr)
	}

}
