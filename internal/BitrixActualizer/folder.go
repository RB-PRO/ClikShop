package actualizer

import "os"

// Создать папку
// а если она уже создана, то удалить её нахуй и пересоздать
func MakeDir(Path string) error {
	ErrMkdirAll := os.MkdirAll(Path, 0777)
	if ErrMkdirAll != nil {
		return ErrMkdirAll
	}
	return nil
}
