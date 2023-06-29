package wcprod

// Загрузить картинку на сервер imgbb на 10 минуток,  пока товар загрузится на сервер ClikShop

import (
	"fmt"
	"net/http"

	"github.com/RB-PRO/SanctionedClothing/pkg/imgbb"
)

// Загрузить файл на сервис imgbb
func (woo *WcAdd) UploadFile(FileName string) (string, error) {

	// Переводим картинку в формат base64
	base64, ErrBase64 := imgbb.PicToBase64(FileName)
	if ErrBase64 != nil {
		return "", ErrBase64
	}

	// Загрузка картинки на сервер imgbb
	Response, ErrorUpload := woo.Imgbb.Upload(base64, "RB_PRO", 260000)
	if ErrorUpload != nil {
		return "", ErrorUpload
	}

	// Проверка на отрицательный ответ
	if Response.Status != http.StatusOK {
		return "", fmt.Errorf("RefrashLink: Response Status is %v", Response.Status)
	}

	return Response.Data.Image.URL, nil
}
