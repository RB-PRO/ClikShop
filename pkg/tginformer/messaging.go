package tg

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Отправить сообщение
func (tg *Telegram) Message(data string) error {
	msg := tgbotapi.NewMessage(tg.config.ChatID, data)
	_, err := tg.Send(msg)
	return err
}

// Отправить сообщение
func (tg *Telegram) MessagePhoto(data string, images []string) error {
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
	var SubSlice_j int
	var first bool = true
	for SubSlice_i := 0; SubSlice_i < len(imagesInterf); SubSlice_i += size {
		SubSlice_j += size
		if SubSlice_j > len(imagesInterf) {
			SubSlice_j = len(imagesInterf)
		}

		// Подслайс. Работаем именно с подслайсами, чтобы не перегружать оперативку
		SubSlice := imagesInterf[SubSlice_i:SubSlice_j]

		// fmt.Println(len(SubSlice), first)

		if !first {
			// Добавить сообщение
			time.Sleep(1 * time.Second)
		}
		// Все фотографии для отправки в ообщении
		mediaGroup := tgbotapi.NewMediaGroup(tg.config.ChatID, SubSlice)
		_, ErrSend := tg.Send(mediaGroup)
		if ErrSend != nil {
			fmt.Printf("send: %v\n", ErrSend)
			// return fmt.Errorf("send: %v", ErrSend)
		}
		// fmt.Printf("%+v\n", message.MediaGroupID)

		first = false
	}

	return nil
}
