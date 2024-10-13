package updator

import (
	"ClikShop/common/bases"
	"ClikShop/common/config"
	tg "ClikShop/common/tginformer"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"sync"
	"time"
)

var iterationNumber = 1

func (s *Service) Run() error {
	defer func() {
		if r := recover(); r != nil {
			s.Gol.Errf("Recovered in: %v", r)
		}
	}()

	coastsMap, err := s.BitrixService.Coasts()
	if err != nil {
		panic(errors.Wrap(err, "bitrix: Coasts: Не смог создать запрос: %w"))
	}

	exchangeRateLira, err := s.BankService.Lira()
	if err != nil {
		panic(errors.Wrap(err, "error getting rate lira: %w"))
	}

	sku, err := s.BitrixService.SKU()
	if err != nil {
		panic(errors.Wrap(err, "error getting rate lira: %w"))
	}
	_ = sku

	productsID, err := s.BitrixService.Products()
	if err != nil {
		panic(errors.Wrap(err, "error getting rate lira: %w"))
	}

	cfg, err := config.ParseConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	checker, err := s.TG.NewChecker(fmt.Sprintf("Запуск обновлятора.\nНомер №%d", iterationNumber), tg.LoadNumericFromConfig(cfg))
	if err != nil {
		panic(errors.Wrap(err, "error start checker: %w"))
	}
	go checker.Run(cfg, s.CheckParseFromLink)

	priceFunc := func(brand string, price float64) float64 {
		return bases.EditDecadense(
			exchangeRateLira*price*coastsMap[brand].Walrus +
				float64(coastsMap[brand].Delivery),
		)
	}

	countMicroUpdater := 3

	// mashup products
	rand.Shuffle(len(productsID), func(i, j int) { productsID[i], productsID[j] = productsID[j], productsID[i] })

	var wg sync.WaitGroup

	partSize := len(productsID) / countMicroUpdater
	for i := 0; i < countMicroUpdater; i++ {
		start := i * partSize
		end := start + partSize
		if i == countMicroUpdater-1 {
			end = len(productsID)
		}

		subProductsID := productsID[start:end]
		wg.Add(len(subProductsID))
		go s.updating(subProductsID, i, priceFunc, &wg)
	}
	wg.Wait()
	iterationNumber++

	return nil
}

func (s *Service) updating(productsID []string, number int, priceFunc func(brand string, price float64) float64, wg *sync.WaitGroup) {
	prefix := fmt.Sprintf("#updator\nМикроОбновлятор №%d\n", number+1)
	timeStart := time.Now()
	tgUpdateMessage, err := s.TG.NewUpdMsg(prefix + "Начинаю обновлять товары!\nВремя: " + timeStart.Format(time.DateTime))
	if err != nil {
		panic(err)
	}

	var MicroGoodUpdate int
	for iProdID, ProdID := range productsID {
		if iProdID%100 == 0 || iProdID == len(productsID) {
			_ = tgUpdateMessage.Update(
				fmt.Sprintf(" - Обновил %d товаров из %d, это %.2f%%\nПроцент успеха: %.2f%%\nНачал в %s",
					iProdID, len(productsID), 100.0*float64(iProdID)/float64(len(productsID)), 100.0*float64(MicroGoodUpdate)/float64(iProdID), timeStart.Format(time.DateTime)),
			)
		}

		if err := s.updateProduct(ProdID, priceFunc); err != nil {
			s.Gol.Warn(fmt.Sprintf("МикроОбновлятор №%d: обновление товара %s: %s", number, ProdID, err))
		} else {
			MicroGoodUpdate++
			//DoneUpdateProduct++
		}
		wg.Done()
	}
}
