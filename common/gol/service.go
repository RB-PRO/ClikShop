package gol

import (
	"fmt"
	"log"
	"os"
)

type Gol struct {
	name       string
	logInfo    *log.Logger
	logWarning *log.Logger
	logError   *log.Logger
}

// NewGol create new logger service
func NewGol(name string) *Gol {
	LogInfo := log.New(os.Stdout, fmt.Sprintf("[INFO/%s]\t", name), log.Ldate|log.Ltime)
	LogWarning := log.New(os.Stdout, fmt.Sprintf("[WARN/%s]\t", name), log.Ldate|log.Ltime|log.Lshortfile)
	LogError := log.New(os.Stdout, fmt.Sprintf("[ERR/%s]\t", name), log.Ldate|log.Ltime|log.Lshortfile)

	return &Gol{
		logInfo:    LogInfo,
		logWarning: LogWarning,
		logError:   LogError,
	}
}
