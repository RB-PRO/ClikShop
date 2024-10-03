package updator

import (
	"ClikShop/common/apibitrix"
	"ClikShop/common/bases"
	"ClikShop/common/config"
	"fmt"
	"log"
	"sync"
	"testing"
)

// Печать запросника
func PrintnVariationReq(variationReq []apibitrix.Variation_Request) {
	fmt.Printf("\nvariationReq\n")
	for i := range variationReq {
		fmt.Printf("%+v\n", variationReq[i])
	}
}

func TestREject(t *testing.T) {
	fmt.Println(naakt("123").String())
	fmt.Println(naakt("123ыфв -ф 2131ё--").String())
	fmt.Println(naakt("123asd").String())
	fmt.Println(naakt("123asdzxc").String())
}

func TestUpdates(t *testing.T) {

	cfg, err := config.ParseConfig("../../../config.json")
	if err != nil {
		log.Fatalln(err)
	}

	service, err := New(cfg)
	if err != nil {
		t.Error(err)
	}

	coastsMap, err := service.BitrixService.Coasts()
	if err != nil {
		t.Error(err)
	}

	exchangeRateLira, err := service.BankService.Lira()
	if err != nil {
		t.Error(err)
	}

	sku, err := service.BitrixService.SKU()
	if err != nil {
		t.Error(err)
	}
	_ = sku

	priceFunc := func(brand string, price float64) float64 {
		return bases.EditDecadense(
			exchangeRateLira*
				price*
				coastsMap[brand].Walrus +
				float64(coastsMap[brand].Delivery),
		)
	}

	ProductID := "1026427"

	var wg sync.WaitGroup
	wg.Add(1)
	go service.updating(
		[]string{ProductID},
		1,
		priceFunc,
		&wg,
	)
	wg.Wait()

}
