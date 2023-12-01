package wcprod

// Загрузить картинку на сервер imgbb на 10 минуток,  пока товар загрузится на сервер ClikShop

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/RB-PRO/ClikShop/pkg/imgbb"
)

// Загрузить файл на сервис imgbb
func (woo *WcAdd) UploadFile(FileName string) (string, error) {

	// Переводим картинку в формат base64
	base64, ErrBase64 := imgbb.PicToBase64(FileName)
	if ErrBase64 != nil {
		return "", ErrBase64
	}

	// Загрузка картинки на сервер imgbb
	Response, ErrorUpload := woo.Imgbb.Upload(base64, "RB_PRO", 6000)
	if ErrorUpload != nil {
		return "", ErrorUpload
	}

	// Проверка на отрицательный ответ
	if Response.Status != http.StatusOK {
		return "", fmt.Errorf("RefrashLink: Response Status is %v", Response.Status)
	}

	return Response.Data.Image.URL, nil
}

// Convert Picture in base64
func PicToBase64(filename string) (string, error) {
	imgFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()

	fInfo, fInfoError := imgFile.Stat()
	if fInfoError != nil {
		return "", fInfoError
	}
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	_, ReadError := fReader.Read(buf)
	if ReadError != nil {
		return "", ReadError
	}

	return base64.StdEncoding.EncodeToString(buf), nil
}
