package wcprod_test

// Загрузить картинку на сервер imgbb на 10 минуток,  пока товар загрузится на сервер ClikShop

import (
	"context"
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/wcprod"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
)

func TestUploadFile(t *testing.T) {
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		t.Error(errorInitWcAdd)
	}

	// link, err := Adding.UploadFile(`D:\Desktop\Work\program\go\src\github.com\RB-PRO\ClikShop\pkg\wcprod\test.jpg`)
	// if err != nil {
	// 	t.Error(err)
	// }

	// FilePathJpg := `D:\Desktop\Work\program\go\src\github.com\RB-PRO\ClikShop\pkg\wcprod\test.jpg`

	base64, err := wcprod.PicToBase64("test.jpg")
	if err != nil {
		t.Error(err)
	}
	res, ErrUpload := Adding.IK.Uploader.Upload(context.Background(), base64, uploader.UploadParam{FileName: "test.jpg"})
	if ErrUpload != nil {
		t.Error(ErrUpload)
	}
	OutputImage := res.Data.Url
	fmt.Println(OutputImage)

}
