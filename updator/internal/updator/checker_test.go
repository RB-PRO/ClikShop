package updator

import (
	"ClikShop/common/config"
	tg "ClikShop/common/tginformer"
	"log"
	"testing"
	"time"
)

func TestChecker(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		log.Fatalln(err)
	}

	s, err := New(cfg)
	if err != nil {
		t.Error(err)
	}

	tgService, err := tg.New(cfg.Telegram.Token, cfg.Telegram.ChatID)
	if err != nil {
		t.Error(err)
	}

	checker, err := tgService.NewChecker("START", tg.LoadNumericFromConfig(cfg))
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Millisecond * 300)

	_ = checker.Run(cfg, s.CheckParseFromLink)
}
