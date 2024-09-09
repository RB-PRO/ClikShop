package bases

import (
	"image/jpeg"
	"os"

	webps "github.com/nickalie/go-webpbin"
)

func Webp2Jpg(InputFileName, OutputFileName string) error {
	// Открываем файл с изображением в формате WebP
	webpFile, err := os.Open(InputFileName)
	if err != nil {
		return err
	}
	defer webpFile.Close()

	// Декодируем изображение в формате WebP
	webpImage, err := webps.Decode(webpFile)
	if err != nil {
		return err
	}

	// Создаем новый файл для сохранения изображения в формате JPEG
	jpegFile, err := os.Create(OutputFileName)
	if err != nil {
		return err
	}
	defer jpegFile.Close()

	// Кодируем изображение в формате JPEG и сохраняем его в файл
	if err := jpeg.Encode(jpegFile, webpImage, &jpeg.Options{Quality: 80}); err != nil {
		return err
	}

	return nil
}
