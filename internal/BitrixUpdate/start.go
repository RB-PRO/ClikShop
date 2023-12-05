package bitrixupdate

import (
	"fmt"
	"strings"
	"time"

	notification "github.com/RB-PRO/ClikShop/pkg/Notification"
	"github.com/RB-PRO/ClikShop/pkg/apibitrix"
	"github.com/RB-PRO/ClikShop/pkg/cbbank"
	tg "github.com/RB-PRO/ClikShop/pkg/tginformer"
)

type bitrixUpdator struct {
	BX *apibitrix.BitrixUser
}

func Start() {

	// Приложение Битрикс
	BitrixUser, ErrBX := apibitrix.NewBitrixUser()
	if ErrBX != nil {
		panic(ErrBX)
	}
	bx := bitrixUpdator{BitrixUser}
	Nots, ErrNotification := notification.NewNotification("notification_updator.json")
	if ErrNotification != nil {
		panic(ErrNotification)
	}
	tg, ErrTG := tg.NewTelegram("sender.json")
	if ErrTG != nil {
		panic(ErrTG)
	}

	for {
		TimeStart := time.Now()

		cb, ErrorCB := cbbank.New() // Цены ЦБ для получение актуального курса
		if ErrorCB != nil {
			panic(ErrorCB)
		}
		bx.BX.CB = cb
		bx.BX.Nots = Nots
		// bx.BX.Nots.Sends(fmt.Sprintf("#updator\nКурс: 1₤ = %.2f₽", cb.Data.Valute.Try.Value/10))
		tg.Message(fmt.Sprintf("#updator\nКурс: 1₤ = %.2f₽", cb.Data.Valute.Try.Value/10))

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
		// bx.BX.Nots.Sends(fmt.Sprintf("В Bitrix всего %d товаров.", len(ProductsID)))

		tgUpdate, err := tg.NewUpdMsg("Начинаю обновлять товары!\nВремя: " + TimeStart.Format("15:04 02.01.2006"))
		if err != nil {
			panic(err)
		}

		var goodUpdate int
		// Цикл по всем товарам
		for iProductID, ProductID := range ProductsID {

			// if (iProductID+1)%100 == 0 {
			// 	bx.BX.Nots.Sends(fmt.Sprintf("Обработка товаров: (%d/%d)", iProductID+1, len(ProductsID)))
			// }

			if (iProductID+1)%10 == 0 {
				tgUpdate.Update(fmt.Sprintf("#updator\nОбновил %d товаров из %d, это %.2f%%\nНачал в %s",
					iProductID, len(ProductsID), float64(iProductID+1)/float64(len(ProductsID)), TimeStart.Format("15:04 02.01.2006")))
			}

			// Обновляем данные по товару
			ErrUpdateProduct := bx.UpdateProduct(ProductID)
			if ErrUpdateProduct != nil {
				bx.BX.Log.Warn(fmt.Sprintf("Цикл: UpdateProduct %s: %s", ProductID, ErrUpdateProduct))
			} else {
				goodUpdate++
			}

			// break
		}

		// bx.BX.Nots.Sends(fmt.Sprintf("#updator\nПрошёл цикл обновлятора по %d товарам", len(ProductsID)))
		// tgUpdate.Update(fmt.Sprintf("#updator\nПрошёл цикл обновлятора по %d товарам", len(ProductsID)))

		tgUpdate.Update(fmt.Sprintf("#updator\nУспешно обновлено %d товаров из %d, это %.2f%%\nНачал в %s\nЗакончил в %s\nЭто - %s",
			goodUpdate, len(ProductsID), float64(goodUpdate)/float64(len(ProductsID)), TimeStart.Format("15:04 02.01.2006"), time.Now().Format("15:04 02.01.2006"), time.Now().Sub(TimeStart)))

	}
}

// Свести строку к одному типу
func EditColorName(str string) string {
	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, "_", "")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)
	return str
}
