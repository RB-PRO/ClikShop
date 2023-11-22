package bitrixupdate

import (
	"fmt"
	"strings"

	notification "github.com/RB-PRO/SanctionedClothing/pkg/Notification"
	"github.com/RB-PRO/SanctionedClothing/pkg/apibitrix"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
)

type bitrixUpdator struct {
	BX *apibitrix.BitrixUser
}

func Start() {

	// Приложение Битрикс
	bx := bitrixUpdator{apibitrix.NewBitrixUser()}
	Nots, ErrNotification := notification.NewNotification("notification.json")
	if ErrNotification != nil {
		panic(ErrNotification)
	}
	cb, ErrorCB := cbbank.New() // Цены ЦБ для получение актуального курса
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	bx.BX.CB = cb
	bx.BX.Nots = Nots

	bx.BX.Nots.Sends(fmt.Sprintf("Курс: 1₤ = %.2f₽", cb.Data.Valute.Try.Value/10))

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
	bx.BX.Nots.Sends(fmt.Sprintf("В Bitrix всего %d товаров.", len(ProductsID)))

	// Цикл по всем товарам
	for iProductID, ProductID := range ProductsID {

		if (iProductID+1)%100 == 0 {
			bx.BX.Nots.Sends(fmt.Sprintf("Обработка товаров: (%d/%d)", iProductID+1, len(ProductsID)))
		}

		// Обновляем данные по товару
		ErrUpdateProduct := bx.UpdateProduct(ProductID)
		if ErrUpdateProduct != nil {
			bx.BX.Log.Warn(fmt.Sprintf("Цикл: UpdateProduct %s: %s", ProductID, ErrUpdateProduct))
		}

		// break
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
