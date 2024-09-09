package notification

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

type Service struct {
	name string
	TG   *telegram.Telegram
}

type Config struct {
	Name   string `json:"name"`
	ChatID string `json:"chat_id"`
	Token  string `json:"token"`
}

func New(cfg Config) (*Service, error) {

	// Create a telegram service. Ignoring error for demo simplicity.
	telegramService, err := telegram.New(cfg.Token)
	if err != nil {
		return nil, fmt.Errorf("telegram service: %v", err)
	}

	ChatID, err := strconv.ParseInt(cfg.ChatID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("notification service: %v", err)
	}

	// Добавить ID,куда будут посылаться уведомления
	telegramService.AddReceivers(ChatID)

	// Tell our notifier to use the telegram service. You can repeat the above process
	// for as many services as you like and just tell the notifier to use them.
	// Inspired by http middlewares used in higher level libraries.
	notify.UseServices(telegramService)

	// Send a test message
	if err := notify.Send(
		context.Background(),
		cfg.Name,
		"Start service",
	); err != nil {
		return nil, fmt.Errorf("notify.Send: %w", err)
	}

	return &Service{
		name: cfg.Name,
		TG:   telegramService,
	}, nil
}

func (s *Service) Sends(message string) error {
	return s.TG.Send(
		context.Background(),
		s.name,
		message,
	)
}
