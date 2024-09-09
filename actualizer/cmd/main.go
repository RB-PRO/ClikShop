package main

import (
	"ClikShop/actualizer/internal/actualizer"
	"ClikShop/common/config"
	"log"
)

func main() {
	cfg, err := config.ParseConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	actualizerService, err := actualizer.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	actualizerService.Run()

	//go func() {
	//	for {
	//		if err := updateService.Run(); err != nil {
	//			log.Fatalln(err)
	//		}
	//		break
	//	}
	//}()

}
