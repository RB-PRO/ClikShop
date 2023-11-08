package bitrixupdate

import (
	"fmt"
	"strings"

	notification "github.com/RB-PRO/SanctionedClothing/pkg/Notification"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
)

func Start() {

	// Приложение Битрикс
	bx := NewBitrixUser()
	Nots, ErrNotification := notification.NewNotification("notification.json")
	if ErrNotification != nil {
		panic(ErrNotification)
	}
	cb, ErrorCB := cbbank.New() // Цены ЦБ для получение актуального курса
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	bx.cb = cb
	bx.Nots = Nots

	bx.Nots.Sends(fmt.Sprintf("Курс: 1₤ = %.2f₽", cb.Data.Valute.Try.Value/10))

	// Загружаем цены
	_, ErrCoasts := bx.Coasts()
	if ErrCoasts != nil {
		panic(ErrCoasts)
	}

	// Получаем списки товаров
	ProductsID, ErrProducts := bx.Products()
	if ErrProducts != nil {
		panic(ErrProducts)
	}
	bx.Nots.Sends(fmt.Sprintf("В Bitrix всего %d товаров.", len(ProductsID)))

	// Цикл по всем товарам
	for _, ProductID := range ProductsID {

		// Обновляем данные по товару
		ErrUpdateProduct := bx.UpdateProduct(ProductID)
		if ErrUpdateProduct != nil {
			bx.log.Warn(fmt.Sprintf("Цикл: UpdateProduct %s: %s", ProductID, ErrUpdateProduct))
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
