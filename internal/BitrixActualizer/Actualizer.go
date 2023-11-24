package actualizer

import (
	"fmt"

	notification "github.com/RB-PRO/SanctionedClothing/pkg/Notification"
	"github.com/RB-PRO/SanctionedClothing/pkg/apibitrix"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
	"github.com/RB-PRO/SanctionedClothing/pkg/gol"
)

type bitrixActualizer struct {
	BX   *apibitrix.BitrixUser
	GLOG *gol.Gol
}

// Создать актуализатор
func NewActualizer() (*bitrixActualizer, error) {
	bx, ErrNewBitrixUser := apibitrix.NewBitrixUser()
	if ErrNewBitrixUser != nil {
		return nil, fmt.Errorf("apibitrix.NewBitrixUser: %v", ErrNewBitrixUser)
	}
	glog, ErrNewLogs := gol.NewGol("logs/")
	if ErrNewLogs != nil {
		return nil, fmt.Errorf("gol.NewGol: %v", ErrNewLogs)
	}
	return &bitrixActualizer{BX: bx, GLOG: glog}, nil
}

func Start() {
	// Приложение Битрикс
	bx, ErrNewActualizer := NewActualizer()
	if ErrNewActualizer != nil {
		panic(fmt.Errorf("gol.NewGol: %v", ErrNewActualizer))
	}
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

	// Загрузить все артикулы
	sku, ErrSKU := bx.MapSKU()
	if ErrSKU != nil {
		panic(ErrSKU)
	}
	bx.BX.Nots.Sends(fmt.Sprintf("Получил %d артикулов из Bitrix", len(sku)))

	bx.md()
	bx.zara()
	bx.hm()
	bx.ss()
}

// func SaveJson(filename string, data []string) error {
// 	f, ErrCreateFile := os.Create(filename + ".json")
// 	if ErrCreateFile != nil {
// 		return ErrCreateFile
// 	}
// 	// as_json, ErrMarshalIndent := json.MarshalIndent(variety, "", "\t")
// 	as_json, ErrMarshalIndent := MarshalMy(data)
// 	if ErrMarshalIndent != nil {
// 		return ErrMarshalIndent
// 	}
// 	f.Write(as_json)
// 	f.Close()
// 	return nil
// }

// func MarshalMy(i interface{}) ([]byte, error) {
// 	buffer := &bytes.Buffer{}
// 	encoder := json.NewEncoder(buffer)
// 	encoder.SetEscapeHTML(false)
// 	err := encoder.Encode(i)
// 	return bytes.TrimRight(buffer.Bytes(), "\n"), err
// }
