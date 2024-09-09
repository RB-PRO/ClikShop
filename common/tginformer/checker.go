package tg

import (
	"ClikShop/common/bases"
	"ClikShop/common/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

type Checker struct {
	MessageID int
	*Service
}

func (s *Service) NewChecker(startMessage string, key tgbotapi.InlineKeyboardMarkup) (*Checker, error) {
	msg := tgbotapi.NewMessage(s.chatID, startMessage)
	msg.ReplyMarkup = key

	responseMsg, err := s.tg.Send(msg)
	if err != nil {
		return nil, err
	}

	return &Checker{MessageID: responseMsg.MessageID, Service: s}, nil
}

func (s *Checker) Run(cfg config.Config, parseFunc func(cfg config.Config, link string) (string, error)) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in run checker service", r)
		}
	}()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := s.tg.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := s.tg.Request(callback); err != nil {
				panic(err)
			}

			msg := tgbotapi.MessageConfig{}
			parseMsg, err := parseFunc(cfg, update.CallbackQuery.Data)
			if err != nil {
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "ERROR parse!\n"+update.CallbackQuery.Data+"\nErr: "+err.Error())
			} else {
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, parseMsg)
			}

			// And finally, send a message containing the data received.
			if _, err := s.tg.Send(msg); err != nil {
				panic(err)
			}
		}
	}
	return nil
}

func LoadNumericFromConfig(config config.Config) (key tgbotapi.InlineKeyboardMarkup) {
	col := make([][]tgbotapi.InlineKeyboardButton, 5)

	createRowFunc := func(links []string, InlineKeyboardButton string) []tgbotapi.InlineKeyboardButton {
		rowKeyboardButton := make([]tgbotapi.InlineKeyboardButton, 0, len(links))
		for i := range links {
			InlineKeyboardButtonSuffix := InlineKeyboardButton + "_" + strconv.Itoa(i)
			rowKeyboardButton = append(rowKeyboardButton, tgbotapi.NewInlineKeyboardButtonData(InlineKeyboardButtonSuffix, InlineKeyboardButtonSuffix))
		}
		return rowKeyboardButton
	}

	col[0] = createRowFunc(config.Updater.PingLinks.MassimoDutti, bases.TagMD)
	col[1] = createRowFunc(config.Updater.PingLinks.HM, bases.TagHM)
	col[2] = createRowFunc(config.Updater.PingLinks.Zara, bases.TagZara)
	col[3] = createRowFunc(config.Updater.PingLinks.SneakSup, bases.TagSS)
	col[4] = createRowFunc(config.Updater.PingLinks.Trendyol, bases.TagTY)

	return tgbotapi.NewInlineKeyboardMarkup(col...)
}
