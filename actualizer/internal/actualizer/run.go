package actualizer

import (
	"fmt"
	"github.com/pkg/errors"
)

func (s *Service) Run() {
	s.run()
}

func (s *Service) run() {
	coastsMap, err := s.BitrixService.Coasts()
	if err != nil {
		panic(errors.Wrap(err, "bitrix: Coasts: Не смог создать запрос: %w"))
	}
	s.mapCoast = coastsMap

	rate, err := s.BankService.Lira()
	if err != nil {
		panic(err)
	}
	_ = s.TG.Message(fmt.Sprintf("Курс: 1₤ = %.2f₽", rate))
	_ = s.TG.Message(fmt.Sprintf("Ценовая политика %v", coastsMap))

	// Загружаем цены
	_, err = s.BitrixService.Coasts()
	if err != nil {
		panic(err)
	}

	// Загрузить все артикулы
	s.SKU, err = s.BitrixService.SKU()
	if err != nil {
		panic(err)
	}
	_ = s.TG.Message(fmt.Sprintf("Получил %d артикулов из Bitrix", len(s.SKU)))

	shops := []Shop{NewHM(s.hmService), NewMD(s.mdService), NewZARA(s.zaraService), NewTY(), NewSS()}
	//shops := []Shop{NewZARA(s.zaraService), NewTY(), NewSS()}
	for _, shop := range shops {

		// Парсинг товаров
		folder, err := shop.Scraper()
		if err != nil {
			s.Gol.Err(fmt.Sprintf("shop.screper(): %s: %v", folder, err))
			continue
		}

		// Вычитание товаров
		if err := s.Sub(folder); err != nil {
			s.Gol.Err(fmt.Sprintf("%v: bx.Sub: %v", folder, err))
			return
		}

		// Удаление дубликатов
		if err := s.DeleteRepeated(folder); err != nil {
			s.Gol.Err(fmt.Sprintf("%v: bx.DeleteRepeated: %v", folder, err))
			return
		}

		// Перевод
		if err := s.Trans(folder); err != nil {
			s.Gol.Err(fmt.Sprintf("%v: bx.Trans: %v", folder, err))
			return
		}

		// Публикация товара
		if err := s.Push(folder); err != nil {
			s.Gol.Err(fmt.Sprintf("%v: bx.ErrPush: %v", folder, err))
			return
		}

	}
}
