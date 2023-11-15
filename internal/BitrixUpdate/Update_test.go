package bitrixupdate

import (
	"fmt"
	"testing"

	notification "github.com/RB-PRO/SanctionedClothing/pkg/Notification"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
)

func TestHM(t *testing.T) {
	bx := NewTestBX()

	// Получить подробнее о товаре
	// https://213.226.124.16/bitrix/admin/iblock_element_edit.php?IBLOCK_ID=15&type=aspro_lite_catalog&lang=ru&ID=73352&find_section_section=450&WF=Y
	ProductsDetail, ErrProduct := bx.Product([]string{"73352"})
	if ErrProduct != nil {
		t.Error("bitrix: Update: HM: %w", ErrProduct)
	}
	if len(ProductsDetail.Products) == 0 {
		t.Error("bitrix: Update: HM: lenResp = 0")
	}

	// Смотрим ссылку  для определения источника того, откуда пришёл товар
	variationReq, ErrUpdate := bx.UpdateHandM(ProductsDetail)
	if ErrUpdate != nil {
		t.Error("bitrix: Update: HM: %w", ErrUpdate)
	}

	PrintnVariationReq(variationReq)
}

func NewTestBX() *BitrixUser {
	bx := NewBitrixUser()
	Nots, ErrNotification := notification.NewNotification("..\\..\\notification.json")
	if ErrNotification != nil {
		panic(ErrNotification)
	}
	bx.Nots = Nots

	// Загружаем цены
	_, ErrCoasts := bx.Coasts()
	if ErrCoasts != nil {
		panic(ErrCoasts)
	}

	return bx
}

// Печать запросника
func PrintnVariationReq(variationReq []Variation_Request) {
	fmt.Printf("\nvariationReq\n")
	for i := range variationReq {
		fmt.Printf("%+v\n", variationReq[i])
	}
}

func TestUpdates(t *testing.T) {
	ProductID := "61043"
	// Приложение Битрикс
	bx := NewBitrixUser()
	Nots, ErrNotification := notification.NewNotification("..\\..\\notification.json")
	if ErrNotification != nil {
		panic(ErrNotification)
	}
	cb, ErrorCB := cbbank.New() // Цены ЦБ для получение актуального курса
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	bx.cb = cb
	bx.Nots = Nots
	fmt.Printf("Курс: 1₤ = %.2f₽\n", cb.Data.Valute.Try.Value/10)
	_, ErrCoasts := bx.Coasts() // Загружаем цены
	if ErrCoasts != nil {
		panic(ErrCoasts)
	}
	fmt.Println(bx.MapCoast)
	ErrUpdateProduct := bx.UpdateProduct(ProductID) // Обновляем данные по товару
	if ErrUpdateProduct != nil {
		t.Error("bitrix: UpdateProduct", ProductID, ":", ErrUpdateProduct)
	}
}
