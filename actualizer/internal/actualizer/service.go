package actualizer

import (
	md "ClikShop/common/MassimoDutti"
	zara "ClikShop/common/ZaraTR"
	"ClikShop/common/apibitrix"
	"ClikShop/common/cbbank"
	"ClikShop/common/config"
	"ClikShop/common/gol"
	"ClikShop/common/hm"
	tg "ClikShop/common/tginformer"
	"ClikShop/common/transrb"
	"fmt"
)

type Service struct {
	BitrixService *apibitrix.Service
	BankService   *cbbank.Service
	Translate     *transrb.Translate
	TG            *tg.Service
	Gol           *gol.Gol
	SKU           map[string]bool
	mapCoast      map[string]apibitrix.CoastMap

	hmService   *hm.Service
	mdService   *md.Service
	zaraService *zara.Service
	//ssService   *sneaksup.Service
	//tyService   *trendyol.Service
}

func New(cfg config.Config) (*Service, error) {
	translateService, err := transrb.New(cfg.Translator.FolderId, cfg.Translator.OauthToken)
	if err != nil {
		return nil, fmt.Errorf("translate service: %v", err)
	}

	tgService, err := tg.New(cfg.Telegram.Token, cfg.Telegram.ChatID)
	if err != nil {
		return nil, fmt.Errorf("central bank service: %v", err)
	}

	return &Service{
		BitrixService: apibitrix.New(),
		TG:            tgService,
		BankService:   cbbank.New(),
		Gol:           gol.NewGol("actualizer"),
		Translate:     translateService,
		hmService:     hm.New(cfg),
		mdService:     md.New(cfg),
		zaraService:   zara.New(cfg),
	}, nil
}