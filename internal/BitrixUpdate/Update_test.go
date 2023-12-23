package bitrixupdate

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/apibitrix"
	"github.com/RB-PRO/ClikShop/pkg/cbbank"
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
	// https://www2.hm.com/tr_tr/productpage.1115995001.html
	ProductID := "420286"

	// Приложение Битрикс
	// bx := NewBitrixUser()
	BitrixUser, ErrBX := apibitrix.NewBitrixUser()
	if ErrBX != nil {
		panic(ErrBX)
	}
	bx := BitrixUpdator{BX: BitrixUser}

	cb, ErrorCB := cbbank.New() // Цены ЦБ для получение актуального курса
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	bx.BX.CB = cb
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
