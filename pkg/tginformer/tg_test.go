package tg

import (
	"testing"
	"time"
)

func TestSends(t *testing.T) {
	tg, ErrTG := NewTelegram("..\\..\\sender.json")
	if ErrTG != nil {
		t.Error(ErrTG)
	}
	Message := "ТЕСТИК)"
	images := []string{"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_1_1_1.jpg?ts=1700750710738",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_2_1_1.jpg?ts=1700750710565",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_2_2_1.jpg?ts=1700750710661",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_2_3_1.jpg?ts=1700750706867",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_2_4_1.jpg?ts=1700750710661",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_15_17_1.jpg?ts=1700553678715",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_15_18_1.jpg?ts=1700553679139",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_15_19_1.jpg?ts=1700553678759",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_15_20_1.jpg?ts=1700553678565",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_15_21_1.jpg?ts=1700553679003",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_6_3_1.jpg?ts=1700569804690",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_6_4_1.jpg?ts=1700569804725",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_6_5_1.jpg?ts=1700569804070",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_6_8_1.jpg?ts=1700569803147",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_6_9_1.jpg?ts=1700569803815",
		"https://static.zara.net/photos///2024/V/0/1/p/8073/032/800/2/w/916/8073032800_6_10_1.jpg?ts=1700569804402",
	}
	ErrMessage := tg.MessagePhoto(Message, images)
	if ErrMessage != nil {
		t.Error(ErrMessage)
	}
}

func TestUpdate(t *testing.T) {
	tg, ErrTG := NewTelegram("..\\..\\sender.json")
	if ErrTG != nil {
		t.Error(ErrTG)
	}

	upd, err := tg.NewUpdMsg("START")
	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Second)
	upd.Update("1")
	time.Sleep(time.Second)
	upd.Update("2")
	time.Sleep(time.Second)
	upd.Update("3")

}
