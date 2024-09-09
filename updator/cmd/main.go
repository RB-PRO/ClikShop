package main

import (
	"ClikShop/common/config"
	"ClikShop/updator/internal/updator"
	"log"
)

func main() {
	cfg, err := config.ParseConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	updateService, err := updator.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	updateService.Run()

	//go func() {
	//	for {
	//		if err := updateService.Run(); err != nil {
	//			log.Fatalln(err)
	//		}
	//		break
	//	}
	//}()

}
