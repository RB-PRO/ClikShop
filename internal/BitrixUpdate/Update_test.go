package bitrixupdate

import (
	"fmt"
	"testing"

	"ClikShop/common/apibitrix"
	"ClikShop/common/cbbank"
)

// Печать запросника
func PrintnVariationReq(variationReq []apibitrix.Variation_Request) {
	fmt.Printf("\nvariationReq\n")
	for i := range variationReq {
		fmt.Printf("%+v\n", variationReq[i])
	}
}
func TestREject(t *testing.T) {
	fmt.Println(naaktstring("123"))
	fmt.Println(naaktstring("123ыфв -ф 2131ё--"))
	fmt.Println(naaktstring("123asd"))
	fmt.Println(naaktstring("123asdzxc"))
}

func TestUpdates(t *testing.T) {
	ProductID := "494730"

	// Приложение Битрикс
	// bx := NewBitrixUser()
	BitrixUser, ErrBX := apibitrix.NewBitrixUser()
	if ErrBX != nil {
		t.Error(ErrBX)
	}
	bx := BitrixUpdator{BX: BitrixUser}

	cb, ErrorCB := cbbank.New() // Цены ЦБ для получение актуального курса
	if ErrorCB != nil {
		t.Error(ErrorCB)
	}
	bx.BX.CB = cb
	fmt.Printf("Курс: 1₤ = %.2f₽\n", cb.Data.Valute.Try.Value/10)
	_, ErrCoasts := bx.BX.Coasts() // Загружаем цены
	if ErrCoasts != nil {
		t.Error(ErrCoasts)
	}
	fmt.Println(bx.BX.MapCoast)
	ErrUpdateProduct := bx.UpdateProduct(ProductID) // Обновляем данные по товару
	if ErrUpdateProduct != nil {
		t.Error("bitrix: UpdateProduct", ProductID, ":", ErrUpdateProduct)
	}
}
