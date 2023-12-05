package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Структура для работы с обновлёнными сообщениями
type UpdateMassage struct {
	MessageID int
	*Telegram
}

// Создаём сообщение, которое будет изменяться
func (TG *Telegram) NewUpdMsg(Message string) (*UpdateMassage, error) {
	msg := tgbotapi.NewMessage(TG.config.ChatID, Message)
	RespMsg, err := TG.Send(msg)
	return &UpdateMassage{MessageID: RespMsg.MessageID, Telegram: TG}, err
}

// Обновить сообщение
func (upd *UpdateMassage) Update(Message string) error {
	// fmt.Println(upd.config.ChatID, upd.MessageID, Message)
	UpdateMSG := tgbotapi.NewEditMessageText(upd.config.ChatID, upd.MessageID, Message)
	_, err := upd.Send(UpdateMSG)
	return err
}
