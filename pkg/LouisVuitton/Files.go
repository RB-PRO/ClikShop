package louisvuitton

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Фунционал файл отвечает за работу директорией, создание, отслеживание, скачивание данных по позициям

// Структура управления файловой системы парсера
//
// Под капотом директория пути до самой папки с фотографиями
// и методы, которые отвечают за сохранение данных
type Direction struct {
	zeroPath string // Путь до папки, в которую будут сохраняться данные
}

// Созждать обхект работы с файловой системой
func NewDir(zeroPath string) *Direction {
	return &Direction{zeroPath: zeroPath}
}

// Пересоздать папку
func (dr *Direction) MakeDir(Path string) error {

	// Абсолютный путь до папки. Если его нет, то удаляем всё
	absFolderPath, _ := filepath.Abs(dr.zeroPath + Path)

	// Если папка существует - удаляем
	if _, err := os.Stat(absFolderPath); err == nil {
		os.RemoveAll(dr.zeroPath + Path)
	}

	// Создание пути
	ErrMkdirAll := os.MkdirAll(dr.zeroPath+Path, 0777)
	if ErrMkdirAll != nil {
		return ErrMkdirAll
	}
	return nil
}

// Сохранить фотографии в определённую папочку
func (dr *Direction) SavePhotos(links []string, Path string) error {
	for ilink, link := range links {
		// don't worry about errors
		response, ErrGet := http.Get(link)
		if ErrGet != nil {
			return ErrGet
		}
		defer response.Body.Close()

		//open a file for writing
		file, ErrCreate := os.Create(fmt.Sprintf("%s/%d.jpg", Path, ilink))
		if ErrCreate != nil {
			return ErrCreate
		}
		defer file.Close()

		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, ErrCopy := io.Copy(file, response.Body)
		if ErrCopy != nil {
			return ErrCopy
		}
	}
	return nil
}
