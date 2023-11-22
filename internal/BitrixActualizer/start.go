package actualizer

import (
	"fmt"

	notification "github.com/RB-PRO/SanctionedClothing/pkg/Notification"
	"github.com/RB-PRO/SanctionedClothing/pkg/apibitrix"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
)

type bitrixActualizer struct {
	BX *apibitrix.BitrixUser
}

func NewActualizer() *bitrixActualizer {
	return &bitrixActualizer{apibitrix.NewBitrixUser()}
}

func Start() {
	// Приложение Битрикс
	// bx := bitrixActualizer{apibitrix.NewBitrixUser()}
	bx := NewActualizer()
	Nots, ErrNotif := notification.NewNotification("notification.json")
	if ErrNotif != nil {
		panic(ErrNotif)
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
	_, ErrSKU := bx.MapSKU()
	if ErrSKU != nil {
		panic(ErrSKU)
	}

}

// Полуть мапу всех артикулов товаров из магазина в Bitrix
func (bx *bitrixActualizer) MapSKU() (map[string]bool, error) {
	// Получаем список ID всех товаров
	ProdsID, ErrProducts := bx.BX.Products()
	if ErrProducts != nil {
		return nil, ErrProducts
	}
	bx.BX.Nots.Sends(fmt.Sprintf("В Bitrix всего %d товаров.", len(ProdsID)))

	skus := make(map[string]bool)
	size := 500
	var SubSlice_j, cout int
	// Цикл по всем-всем товарам с целью формирования мапы всех артикулов
	for SubSlice_i := 0; SubSlice_i < len(ProdsID); SubSlice_i += size {
		SubSlice_j += size
		if SubSlice_j > len(ProdsID) {
			SubSlice_j = len(ProdsID)
		}
		// fmt.Println(SubSlice_i, SubSlice_j)
		// Получаем список ID всех товаров
		ProdInfo, ErrProdInfo := bx.BX.Product(ProdsID[SubSlice_i:SubSlice_j])
		if ErrProdInfo != nil {
			return nil, ErrProdInfo
		}
		for i := range ProdInfo.Products {
			skus[ProdInfo.Products[i].ID] = true
		}
		cout++
	}
	bx.BX.Nots.Sends(fmt.Sprintf("Получил подробную информацию о %d товарах.", len(skus)))
	// fmt.Println(skus)
	return skus, nil
}
