package notification

import (
	"context"
	"strconv"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

type Notification struct {
	TG      *telegram.Telegram
	subject string // Тайтл, который объявляется при объявлении нотификации
	name    string // Название программы
}

func NewNotification(token, ChatID, subject, name string) (Notification, error) {

	// Create a telegram service. Ignoring error for demo simplicity.
	telegramService, ErrorServece := telegram.New(token)
	if ErrorServece != nil {
		return Notification{}, ErrorServece
	}

	// Переводить ChatID из string в int64
	ChatID_int, ErrParseInt := strconv.ParseInt(ChatID, 10, 64)
	if ErrParseInt != nil {
		return Notification{}, ErrParseInt
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
		subject,
		name+": "+"Начинаю работу",
	)
	if ErrorTelegramSend != nil {
		return Notification{}, ErrorTelegramSend
	}

	return Notification{TG: telegramService, subject: subject, name: name}, nil
}

func (notif Notification) Sends(message string) error {
	return notif.TG.Send(
		context.Background(),
		notif.subject,
		message,
	)
}
