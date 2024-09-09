package actualizer

import (
	"fmt"

	"ClikShop/common/apibitrix"
	"ClikShop/common/cbbank"
	"ClikShop/common/gol"
	notification "ClikShop/common/notify"
	"ClikShop/common/transrb"
	"ClikShop/common/wcprod"
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

	// Цикл по всем магазинам с последующим парсингом
	// shops := []Shop{NewHM(bx), NewMD(bx), NewZARA(bx), NewTY(bx)}
	shops := []Shop{NewMD(bx), NewZARA(bx), NewTY(bx), NewSS(bx)}
	for _, shop := range shops {

		// Парсинг товаров
		folder, ErrScrap := shop.screper()
		if ErrScrap != nil {
			ErrStr := fmt.Sprintf("shop.screper(): %s: %v", folder, ErrScrap)
			fmt.Println(ErrStr)
			bx.GLOG.Err(ErrStr)
			continue
		}

		// Вычитание товаров
		ErrSub := bx.Sub(folder)
		if ErrSub != nil {
			bx.GLOG.Err(fmt.Sprintf("%v: bx.Sub: %v", folder, ErrSub))
			return
		}

		// Удаление дубликатов
		ErrDR := bx.DeleteRepeated(folder)
		if ErrDR != nil {
			bx.GLOG.Err(fmt.Sprintf("%v: bx.DeleteRepeated: %v", folder, ErrDR))
			return
		}

		// Перевод
		ErrTr := bx.Trans(folder)
		if ErrTr != nil {
			bx.GLOG.Err(fmt.Sprintf("%v: bx.Trans: %v", folder, ErrTr))
			return
		}

		// Публикация товара
		ErrPush := bx.Push(folder)
		if ErrPush != nil {
			bx.GLOG.Err(fmt.Sprintf("%v: bx.ErrPush: %v", folder, ErrPush))
			return
		}

	}
}
