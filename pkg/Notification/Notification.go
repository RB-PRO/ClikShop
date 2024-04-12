package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

type Notification struct {
	TG      *telegram.Telegram
	subject string // Тайтл, который объявляется при объявлении нотификации
	name    string // Название программы
}

type JsonStartFile struct {
	Token   string `json:"Token"`
	ChatID  string `json:"ChatID"`
	Subject string `json:"Subject"`
	Name    string `json:"Name"`
}

func NewNotification(FileName string) (*Notification, error) {
	Config, ErrorDataLoad := LoadConfig(FileName)
	if ErrorDataLoad != nil {
		return nil, fmt.Errorf("LoadConfig: %w", ErrorDataLoad)
	}

	// Create a telegram service. Ignoring error for demo simplicity.
	telegramService, ErrorServece := telegram.New(Config.Token)
	if ErrorServece != nil {
		return nil, fmt.Errorf("telegram.New: %w", ErrorServece)
	}

	// Переводить ChatID из string в int64
	ChatID_int, ErrParseInt := strconv.ParseInt(Config.ChatID, 10, 64)
	if ErrParseInt != nil {
		return nil, fmt.Errorf("trconv.ParseInt: %w", ErrParseInt)
	}

	// Добавить ID,куда будут посылаться уведомления
	telegramService.AddReceivers(ChatID_int)

	// Tell our notifier to use the telegram service. You can repeat the above process
	// for as many services as you like and just tell the notifier to use them.
	// Inspired by http middlewares used in higher level libraries.
	notify.UseServices(telegramService)

	// Send a test message.
	ErrorTelegramSend := notify.Send(
		context.Background(),
		Config.Subject,
		Config.Name+": "+"Начинаю работу",
	)
	if ErrorTelegramSend != nil {
		return nil, fmt.Errorf("notify.Send: %w", ErrorTelegramSend)
	}

	return &Notification{TG: telegramService, subject: Config.Subject, name: Config.Name}, nil
}

func (notif *Notification) Sends(message string) error {
	return notif.TG.Send(
		context.Background(),
		notif.subject,
		message,
	)
}

// Загрузить данные из файла
func LoadConfig(filename string) (config JsonStartFile, ErrorFIle error) {
	// Открыть файл
	jsonFile, ErrorFIle := os.Open(filename)
	if ErrorFIle != nil {
		return config, ErrorFIle
	}
	defer jsonFile.Close()

	// Прочитать файл и получить массив byte
	jsonData, ErrorFIle := io.ReadAll(jsonFile)
	if ErrorFIle != nil {
		return config, ErrorFIle
	}

	// Распарсить массив byte в структуру
	if ErrorFIle := json.Unmarshal(jsonData, &config); ErrorFIle != nil {
		return config, ErrorFIle
	}
	return config, ErrorFIle
}
