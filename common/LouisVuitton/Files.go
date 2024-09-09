package louisvuitton

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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
	dr := &Direction{zeroPath: zeroPath}
	dr.MakeDir(zeroPath) // psd add error
	return dr
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
func (dr *Direction) SavePhotos(links []string, Path string) ([]string, error) {
	NewPhotos := make([]string, len(links))
	for ilink, link := range links {
		var FileName string
		linkClasters := strings.Split(link, "_")
		if len(linkClasters) > 2 {
			FileName = strings.Join(linkClasters[len(linkClasters)-2:], "_")
		}

		if FileName == "" {
			FileName = strconv.Itoa(ilink) + ".png"
		}

		FilePath := fmt.Sprintf("%s%s%s", dr.zeroPath, Path, FileName)
		// fmt.Println("FilePath", FilePath, link)

		// Сохраняем фото
		ErrSavePhoto := dr.SavePhoto(link, FilePath)
		if ErrSavePhoto != nil {
			log.Println(ErrSavePhoto)
			continue
		}
		NewPhotos[ilink] = dr.zeroPath + Path + FileName

		time.Sleep(time.Second)
	}
	return NewPhotos, nil
}

func (dr *Direction) SavePhoto(link string, FilePath string) error {
	client := &http.Client{}
	// client.Timeout = time.Minute
	req, ErrNewRequest := http.NewRequest(http.MethodGet, link, nil)
	if ErrNewRequest != nil {
		return ErrNewRequest
	}

	req.Header.Add("authority", "ru.louisvuitton.com")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "none")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.4.971 Yowser/2.5 Safari/537.36")

	// Выполнить запрос
	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return ErrDo
	}
	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("SavePhoto: wrong status code: %d", res.StatusCode)
	}

	//open a file for writing
	file, ErrCreate := os.Create(FilePath)
	if ErrCreate != nil {
		return ErrCreate
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, ErrCopy := io.Copy(file, res.Body)
	if ErrCopy != nil {
		return ErrCopy
	}
	return nil
}
