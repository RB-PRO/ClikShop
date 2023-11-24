// Кастомный пакет логгирования в файл
package gol

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Gol struct {
	logInfo    *log.Logger
	logWarning *log.Logger
	logError   *log.Logger
}

// Создать объект логгирования
func NewGol(ZeroDirection string) (*Gol, error) {
	if _, err := os.Stat(ZeroDirection); os.IsNotExist(err) {
		return nil, fmt.Errorf("folder '%s' does not exist", ZeroDirection)
	}
	FileName := ZeroDirection + time.Now().Format("2006-01-02_15-04") + ".log"
	flags := log.LstdFlags | log.Lshortfile
	FileInfo, _ := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	FileWarinig, _ := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	FileError, _ := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	LogInfo := log.New(FileInfo, "Info:\t", flags)
	LogWarning := log.New(FileWarinig, "Warn:\t", flags)
	LogError := log.New(FileError, "Err:\t", flags)

	return &Gol{
		logInfo:    LogInfo,
		logWarning: LogWarning,
		logError:   LogError,
	}, nil
}

// Логировать состояние
func (l *Gol) Info(value ...interface{}) {
	l.logInfo.Println(value...)
}

// Логгировать предупреждение
func (l *Gol) Warn(value ...interface{}) {
	l.logWarning.Println(value...)
}

// Логгировать ошибку
func (l *Gol) Err(value ...interface{}) {
	l.logError.Println(value...)
}
