package apibitrix

import (
	"fmt"

	notification "github.com/RB-PRO/ClikShop/pkg/Notification"
	"github.com/RB-PRO/ClikShop/pkg/cbbank"
	"github.com/RB-PRO/ClikShop/pkg/gol"
)

const bitrixURL string = "https://clikshop.ru/parser/api/%s/index.php"

type Bitrix interface {
	NewBitrixUser() *BitrixUser
	Coasts() (map[string]CoastMap, error)
	Product(Values []string) (ProdResp Product_Response, Err error)
	Products() ([]string, error)
	Variation(VrtReq []Variation_Request) (VrtResp Variation_Response, Err error)
}

// Структура запросника к битриксу
type BitrixUser struct {
	MapCoast map[string]CoastMap // Мапа цен на товары
	Nots     *notification.Notification
	Log      *gol.Gol // Логгирование
	CB       *cbbank.CentrakBank
}

// Создать клиента для работы с данными битрикс
func NewBitrixUser() (*BitrixUser, error) {
	glog, ErrNewLogs := gol.NewGol("logs/")
	if ErrNewLogs != nil {
		return nil, fmt.Errorf("gol.NewGol: %v", ErrNewLogs)
	}
	return &BitrixUser{Log: glog}, nil
}
