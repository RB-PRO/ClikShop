package actualizer

import (
	"fmt"

	notification "github.com/RB-PRO/SanctionedClothing/pkg/Notification"
	"github.com/RB-PRO/SanctionedClothing/pkg/apibitrix"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
	"github.com/RB-PRO/SanctionedClothing/pkg/gol"
	"github.com/RB-PRO/SanctionedClothing/pkg/transrb"
	"github.com/RB-PRO/SanctionedClothing/pkg/wcprod"
)

type bitrixActualizer struct {
	BX   *apibitrix.BitrixUser
	SKU  map[string]bool
	GLOG *gol.Gol
	TR   *transrb.Translate
}

// Создать актуализатор
func NewActualizer() (*bitrixActualizer, error) {

	// Логгер
	glog, ErrNewLogs := gol.NewGol("logs/")
	if ErrNewLogs != nil {
		return nil, fmt.Errorf("gol.NewGol: %v", ErrNewLogs)
	}

	// Битрикс-пользователь
	bx, ErrNewBitrixUser := apibitrix.NewBitrixUser()
	if ErrNewBitrixUser != nil {
		return nil, fmt.Errorf("apibitrix.NewBitrixUser: %v", ErrNewBitrixUser)
	}

	// Подгружаем цены для формирования адекватной цены на товары
	_, ErrCoasts := bx.Coasts()
	if ErrCoasts != nil {
		panic(ErrCoasts)
	}
	glog.Info(fmt.Sprintf("Ценовая политика %v", bx.MapCoast))

	// Переводчик
	ConfigRBFileName := "config_rb.json" // Загрузка конфига
	ConfigRB, ErrOpenConfigRB := wcprod.LoadConfig(ConfigRBFileName)
	if ErrOpenConfigRB != nil {
		return nil, fmt.Errorf("wcprod.LoadConfig: New: Read config file '%s' error: %v", ConfigRBFileName, ErrOpenConfigRB.Error())
	}
	tr, ErrTranslate := transrb.New(ConfigRB.FolderID, ConfigRB.OAuthToken)
	if ErrTranslate != nil {
		return nil, fmt.Errorf("transrb.New: %v", ErrTranslate)
	}

	return &bitrixActualizer{BX: bx, GLOG: glog, TR: tr}, nil
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
	bx.BX.Nots.Sends(fmt.Sprintf("Ценовая политика %v", bx.BX.MapCoast))

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
	bx.SKU = sku
	bx.BX.Nots.Sends(fmt.Sprintf("Получил %d артикулов из Bitrix", len(sku)))

	// FolderFiles := []string{"zara", "md", "hm", "ss"}

	// bx.zara()
	// bx.md()
	// bx.hm()
	// bx.ss()

	// FolderZara := "zara"
	// // bx.zara(FolderZara) // Парсинг
	// // ErrSub := bx.Sub(FolderZara)
	// // if ErrSub != nil {
	// // 	bx.GLOG.Err(fmt.Sprintf("%v: bx.Sub: %v", FolderZara, ErrSub))
	// // 	return
	// // }
	// // ErrTr := bx.Trans(FolderZara)
	// // if ErrTr != nil {
	// // 	bx.GLOG.Err(fmt.Sprintf("%v: bx.Trans: %v", FolderZara, ErrTr))
	// // 	return
	// // }
	// ErrPush := bx.Push(FolderZara)
	// if ErrPush != nil {
	// 	bx.GLOG.Err(fmt.Sprintf("%v: bx.ErrPush: %v", FolderZara, ErrPush))
	// 	return
	// }

	Folder := "md"
	// bx.md(Folder) // Парсинг
	// ErrSub := bx.Sub(Folder)
	// if ErrSub != nil {
	// 	bx.GLOG.Err(fmt.Sprintf("%v: bx.Sub: %v", Folder, ErrSub))
	// 	return
	// }
	// ErrTr := bx.Trans(Folder)
	// if ErrTr != nil {
	// 	bx.GLOG.Err(fmt.Sprintf("%v: bx.Trans: %v", Folder, ErrTr))
	// 	return
	// }
	ErrDR := bx.DeleteRepeated(Folder)
	if ErrDR != nil {
		bx.GLOG.Err(fmt.Sprintf("%v: bx.DeleteRepeated: %v", Folder, ErrDR))
		return
	}
	ErrPush := bx.Push(Folder)
	if ErrPush != nil {
		bx.GLOG.Err(fmt.Sprintf("%v: bx.ErrPush: %v", Folder, ErrPush))
		return
	}
}
