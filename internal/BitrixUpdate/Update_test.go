package bitrixupdate

import (
	"fmt"
	"testing"

	notification "github.com/RB-PRO/SanctionedClothing/pkg/Notification"
	"github.com/RB-PRO/SanctionedClothing/pkg/apibitrix"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
)

// Печать запросника
func PrintnVariationReq(variationReq []apibitrix.Variation_Request) {
	fmt.Printf("\nvariationReq\n")
	for i := range variationReq {
		fmt.Printf("%+v\n", variationReq[i])
	}
}

func TestUpdates(t *testing.T) {
	ProductID := "140542"
	// Приложение Битрикс
	// bx := NewBitrixUser()
	BitrixUser, ErrBX := apibitrix.NewBitrixUser()
	if ErrBX != nil {
		panic(ErrBX)
	}
	bx := bitrixUpdator{BitrixUser}
	Nots, ErrNotification := notification.NewNotification("../../notification.json")
	if ErrNotification != nil {
		panic(ErrNotification)
	}
	cb, ErrorCB := cbbank.New() // Цены ЦБ для получение актуального курса
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	bx.BX.CB = cb
	bx.BX.Nots = Nots
	fmt.Printf("Курс: 1₤ = %.2f₽\n", cb.Data.Valute.Try.Value/10)
	_, ErrCoasts := bx.BX.Coasts() // Загружаем цены
	if ErrCoasts != nil {
		panic(ErrCoasts)
	}
	fmt.Println(bx.BX.MapCoast)
	ErrUpdateProduct := bx.UpdateProduct(ProductID) // Обновляем данные по товару
	if ErrUpdateProduct != nil {
		t.Error("bitrix: UpdateProduct", ProductID, ":", ErrUpdateProduct)
	}
}
