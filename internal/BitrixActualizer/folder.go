package actualizer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/RB-PRO/ClikShop/pkg/bases"
)

// Создать папку
// а если она уже создана, то удалить её нахуй и пересоздать
func MakeDir(Path string) error {
	ErrMkdirAll := os.MkdirAll(Path, os.ModePerm)
	if ErrMkdirAll != nil {
		return ErrMkdirAll
	}
	return nil
}

// Пересоздать папку
func ReMakeDir(Path string) error {

	// Абсолютный путь до папки. Если его нет, то удаляем всё
	absFolderPath, _ := filepath.Abs(Path)

	// Если папка существует - удаляем
	if _, err := os.Stat(absFolderPath); err == nil {
		os.RemoveAll(Path)
	}

	// Создание пути
	ErrMkdirAll := os.MkdirAll(Path, 0777)
	if ErrMkdirAll != nil {
		return ErrMkdirAll
	}
	return nil
}

// Получить все файлы из папки
func FolderFiles(Folder string) (files []string, Err error) {
	entries, Err := os.ReadDir(Folder)
	if Err != nil {
		return nil, fmt.Errorf("os.ReadDir: %v", Err)
	}
	for _, e := range entries {
		files = append(files, e.Name())
	}
	return files, nil
}

func ProdFile(folder, filename string) (Variety bases.Variety2, Err error) {

	// Считать файл
	data, Err := os.ReadFile(folder + filename)
	if Err != nil {
		return bases.Variety2{}, fmt.Errorf("os.ReadFile: %v", Err)
	}

	// Распарсить
	Err = json.Unmarshal(data, &Variety)
	if Err != nil {
		return bases.Variety2{}, fmt.Errorf("os.ReadFile: %v", Err)
	}

	return Variety, nil
}
