package bitrixupdate

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/RB-PRO/ClikShop/pkg/apibitrix"
	"github.com/RB-PRO/ClikShop/pkg/cbbank"
	tg "github.com/RB-PRO/ClikShop/pkg/tginformer"
)

type BitrixUpdator struct {
	BX            *apibitrix.BitrixUser
	TG            *tg.Telegram
	ActiveWorkers int // Количество активных воркеров
}

func NewBitrixUpdator() (*BitrixUpdator, error) {

	// Проверка того, что есть аргумент к-ва воркеров
	if len(os.Args) != 2 {
		return nil, fmt.Errorf("please choose arguments of count workers")
	}

	// Получение к-ва воркеров
	var ActiveWorkers int
	if ArgVal, ErrAtoi := strconv.Atoi(os.Args[1]); ErrAtoi != nil {
		return nil, fmt.Errorf("strconv.Atoi: %v", ErrAtoi)
	} else {
		ActiveWorkers = ArgVal
	}

	// Приложение Битрикс
	BitrixUser, ErrBX := apibitrix.NewBitrixUser()
	if ErrBX != nil {
		return nil, fmt.Errorf("apibitrix.NewBitrixUser: %v", ErrBX)
	}

	// Телеграмус
	fileTG := "sender.json"
	tg, ErrTG := tg.NewTelegram(fileTG)
	if ErrTG != nil {
		return nil, fmt.Errorf("tg.NewTelegram: %s: %v", fileTG, ErrTG)
	}
	return &BitrixUpdator{BX: BitrixUser, TG: tg, ActiveWorkers: ActiveWorkers}, nil
}

func Start() {
	bx, ErrBX := NewBitrixUpdator()
	if ErrBX != nil {
		log.Fatalf("NewBitrixUpdator: %v", ErrBX)
	}

	ErrWork := bx.Wath()
	if ErrBX != nil {
		log.Fatalf("Wath: %v", ErrWork)
	}

}

func (bx *BitrixUpdator) Wath() error {
	for {
		TimeStart := time.Now()

		cb, ErrorCB := cbbank.New() // Цены ЦБ для получение актуального курса
		if ErrorCB != nil {
			panic(ErrorCB)
		}
		bx.BX.CB = cb
		// bx.BX.Nots = Nots
		// bx.BX.Nots.Sends(fmt.Sprintf("#updator\nКурс: 1₤ = %.2f₽", cb.Data.Valute.Try.Value/10))
		bx.TG.Message(fmt.Sprintf("#updator\nКурс: 1₤ = %.2f₽", cb.Data.Valute.Try.Value/10))

		// Загружаем цены
		_, ErrCoasts := bx.BX.Coasts()
		if ErrCoasts != nil {
			panic(ErrCoasts)
		}

		// Получаем списки товаров
		ProductsID, ErrProducts := bx.BX.Products()
		if ErrProducts != nil {
			panic(ErrProducts)
		}

		// рандомно мешаем товары
		rand.Shuffle(len(ProductsID),
			func(i, j int) { ProductsID[i], ProductsID[j] = ProductsID[j], ProductsID[i] })
		// bx.BX.Nots.Sends(fmt.Sprintf("В Bitrix всего %d товаров. Перемешал их рандомно.", len(ProductsID)))

		var DoneUpdateProduct int
		// chUpdate := make(chan int, 1)
		var wg sync.WaitGroup
		CoutMictoUpdator := bx.ActiveWorkers // количество микрообновляторов
		// CoutMictoUpdator = 3 // количество микрообновляторов
		wg.Add(CoutMictoUpdator)

		fmt.Println(len(ProductsID))

		partSize := len(ProductsID) / CoutMictoUpdator
		for i := 0; i < CoutMictoUpdator; i++ {
			start := i * partSize
			end := start + partSize
			if i == CoutMictoUpdator-1 {
				end = len(ProductsID)
			}
			fmt.Println(len(ProductsID[start:end]))
			time.Sleep(2 * time.Second)

			// ЗАпускаем поток на обновление
			go func(ProductsID []string, NumberMicro int) {
				Prefix := fmt.Sprintf("#updator\nМикроОбновлятор №%d\n", NumberMicro+1)
				tgUpdate, err := bx.TG.NewUpdMsg(Prefix + "Начинаю обновлять товары!\nВремя: " + TimeStart.Format("15:04 02.01.2006"))
				if err != nil {
					panic(err)
				}
				//////////////////
				// tgUpdate.Update("TEST")
				var MicroGoodUpdate int
				for iProdID, ProdID := range ProductsID {
					if iProdID%100 == 0 || iProdID == len(ProductsID) {
						tgUpdate.Update(fmt.Sprintf(" - Обновил %d товаров из %d, это %.2f%%\nПроцент успеха: %.2f%%\nНачал в %s",
							iProdID, len(ProductsID), 100.0*float64(iProdID)/float64(len(ProductsID)), 100.0*float64(MicroGoodUpdate)/float64(iProdID), TimeStart.Format("15:04 02.01.2006")))
					}

					// Обновляем данные по товару
					ErrUpdateProduct := bx.UpdateProduct(ProdID)
					if ErrUpdateProduct != nil {
						bx.BX.Log.Warn(fmt.Sprintf("МикроОбновлятор №%d: UpdateProduct %s: %s", NumberMicro, ProdID, ErrUpdateProduct))
					} else {
						MicroGoodUpdate++
						DoneUpdateProduct++
					}
				}
				wg.Done()
			}(ProductsID[start:end], i)
		}
		wg.Wait()
		bx.TG.Message(fmt.Sprintf("#updator\nУспешно обновлено %d товаров из %d, это %.2f%%\nНачал в %s\nЗакончил в %s\nЭто - %s",
			DoneUpdateProduct, len(ProductsID), 100.0*float64(DoneUpdateProduct)/float64(len(ProductsID)), TimeStart.Format("15:04 02.01.2006"), time.Now().Format("15:04 02.01.2006"), time.Since(TimeStart).String()))
	}

	// for {
	// 	TimeStart := time.Now()

	// 	cb, ErrorCB := cbbank.New() // Цены ЦБ для получение актуального курса
	// 	if ErrorCB != nil {
	// 		panic(ErrorCB)
	// 	}
	// 	bx.BX.CB = cb
	// 	// bx.BX.Nots = Nots
	// 	// bx.BX.Nots.Sends(fmt.Sprintf("#updator\nКурс: 1₤ = %.2f₽", cb.Data.Valute.Try.Value/10))
	// 	tg.Message(fmt.Sprintf("#updator\nКурс: 1₤ = %.2f₽", cb.Data.Valute.Try.Value/10))

	// 	// Загружаем цены
	// 	_, ErrCoasts := bx.BX.Coasts()
	// 	if ErrCoasts != nil {
	// 		panic(ErrCoasts)
	// 	}

	// 	// Получаем списки товаров
	// 	ProductsID, ErrProducts := bx.BX.Products()
	// 	if ErrProducts != nil {
	// 		panic(ErrProducts)
	// 	}
	// 	// bx.BX.Nots.Sends(fmt.Sprintf("В Bitrix всего %d товаров.", len(ProductsID)))

	// 	tgUpdate, err := tg.NewUpdMsg("Начинаю обновлять товары!\nВремя: " + TimeStart.Format("15:04 02.01.2006"))
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// Цикл по всем товарам
	// 	var goodUpdate int
	// 	for iProductID, ProductID := range ProductsID {

	// 		if (iProductID+1)%10 == 0 {
	// 			tgUpdate.Update(fmt.Sprintf("#updator\nОбновил %d товаров из %d, это %.2f%%\nНачал в %s",
	// 				iProductID+1, len(ProductsID), 100.0*float64(iProductID+1)/float64(len(ProductsID)), TimeStart.Format("15:04 02.01.2006")))
	// 		}

	// 		// Обновляем данные по товару
	// 		ErrUpdateProduct := bx.UpdateProduct(ProductID)
	// 		if ErrUpdateProduct != nil {
	// 			bx.BX.Log.Warn(fmt.Sprintf("Цикл: UpdateProduct %s: %s", ProductID, ErrUpdateProduct))
	// 		} else {
	// 			goodUpdate++
	// 		}

	// 	}

	// 	// bx.BX.Nots.Sends(fmt.Sprintf("#updator\nПрошёл цикл обновлятора по %d товарам", len(ProductsID)))
	// 	// tgUpdate.Update(fmt.Sprintf("#updator\nПрошёл цикл обновлятора по %d товарам", len(ProductsID)))

	// 	tgUpdate.Update(fmt.Sprintf("#updator\nУспешно обновлено %d товаров из %d, это %.2f%%\nНачал в %s\nЗакончил в %s\nЭто - %s",
	// 		goodUpdate, len(ProductsID), 100.0*float64(goodUpdate)/float64(len(ProductsID)), TimeStart.Format("15:04 02.01.2006"), time.Now().Format("15:04 02.01.2006"), time.Since(TimeStart).String()))

	// }
}

// Свести строку к одному типу
func EditColorName(str string) string {
	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, "_", "")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)
	return str
}
