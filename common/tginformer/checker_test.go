package tg_test

import (
	"ClikShop/common/bases"
	"ClikShop/common/config"
	tg "ClikShop/common/tginformer"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestChecker(t *testing.T) {
	cfg, err := config.ParseConfig("../../config.json")
	if err != nil {
		log.Fatalln(err)
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

	parseFunc := func(cfg config.Config, link string) (string, error) {

		linkSplit := strings.Split(link, "_")
		if len(linkSplit) != 2 {
			return "", errors.New("error link format " + link)
		}
		number, err := strconv.Atoi(linkSplit[1])
		if err != nil {
			return "", errors.New("error link format " + link)
		}
		switch {
		case strings.Contains(link, bases.TagMD):
			parseLink := cfg.Updater.PingLinks.MassimoDutti[number]
			return parseLink, nil
		case strings.Contains(link, bases.TagHM):
			parseLink := cfg.Updater.PingLinks.HM[number]
			return parseLink, nil
		case strings.Contains(link, bases.TagZara):
			parseLink := cfg.Updater.PingLinks.Zara[number]
			return parseLink, nil
		case strings.Contains(link, bases.TagSS):
			parseLink := cfg.Updater.PingLinks.SneakSup[number]
			return parseLink, nil
		case strings.Contains(link, bases.TagTY):
			parseLink := cfg.Updater.PingLinks.Trendyol[number]
			return parseLink, nil
		default:
			return "", fmt.Errorf("не знаю, какую логику применить к тегу %s", link)
		}
	}
	_ = checker.Run(cfg, parseFunc)
}
