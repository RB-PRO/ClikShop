package tg

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Service struct {
	tg     *tgbotapi.BotAPI
	chatID int64
}

func New(Token string, chatID int64) (TG *Service, Err error) {
	bot, Err := tgbotapi.NewBotAPI(Token)
	if Err != nil {
		return nil, fmt.Errorf("telegram: auth bot: %v", Err)
	}
	// bot.Debug = true

	return &Service{tg: bot, chatID: chatID}, nil
}
