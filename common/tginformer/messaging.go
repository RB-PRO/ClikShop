package tg

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Отправить сообщение
func (s *Service) Message(data string) error {
	msg := tgbotapi.NewMessage(s.chatID, data)
	_, err := s.tg.Send(msg)
	return err
}

// Отправить сообщение
func (s *Service) MessagePhoto(data string, images []string) error {
	// msg := tgbotapi.NewMessage(tg.config.ChatID, data)
	// _, err := tg.TG.Send(msg)
	imagesInterf := make([]interface{}, len(images))
	for i := range images {
		photo := tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL(images[i]))
		if i == 0 {
			photo.Caption = data
		}
		imagesInterf[i] = photo
	}

	size := 9
	var jSubSlice int
	first := true
	for iSubSlice := 0; iSubSlice < len(imagesInterf); iSubSlice += size {
		jSubSlice += size
		if jSubSlice > len(imagesInterf) {
			jSubSlice = len(imagesInterf)
		}

		// Подслайс. Работаем именно с подслайсами, чтобы не перегружать оперативку
		SubSlice := imagesInterf[iSubSlice:jSubSlice]

		if !first {
			time.Sleep(1 * time.Second)
		}

		// Все фотографии для отправки в ообщении
		mediaGroup := tgbotapi.NewMediaGroup(s.chatID, SubSlice)
		_, err := s.tg.Send(mediaGroup)
		if err != nil {
			fmt.Printf("send error: %+#v\n", err)
		}

		first = false
	}

	return nil
}
