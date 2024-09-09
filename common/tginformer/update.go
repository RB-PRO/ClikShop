package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UpdateMassage struct {
	MessageID int
	*Service
}

func (s *Service) NewUpdMsg(Message string) (*UpdateMassage, error) {
	msg := tgbotapi.NewMessage(s.chatID, Message)
	responseMsg, err := s.tg.Send(msg)
	if err != nil {
		return nil, err
	}
	return &UpdateMassage{MessageID: responseMsg.MessageID, Service: s}, nil
}

func (upd *UpdateMassage) Update(Message string) error {
	UpdateMSG := tgbotapi.NewEditMessageText(upd.chatID, upd.MessageID, Message)
	_, err := upd.tg.Send(UpdateMSG)
	return err
}
