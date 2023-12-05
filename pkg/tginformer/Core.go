package tg

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	*tgbotapi.BotAPI
	config config
}

type config struct {
	Token  string `json:"Token"`
	ChatID int64  `json:"ChatID"`
}

func NewTelegram(ConfigFileName string) (TG *Telegram, Err error) {

	// Загрузка конфига
	cf, Err := LoadConfig(ConfigFileName)
	if Err != nil {
		return nil, fmt.Errorf("загрузка конфигурационного файла %s: %v",
			ConfigFileName, Err)
	}

	// Авторизация бота
	bot, Err := tgbotapi.NewBotAPI(cf.Token)
	if Err != nil {
		return nil, fmt.Errorf("авторизация бота: %v", Err)
	}
	// bot.Debug = true

	return &Telegram{BotAPI: bot, config: cf}, nil
}
