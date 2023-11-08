package bitrixupdate

import (
	notification "github.com/RB-PRO/SanctionedClothing/pkg/Notification"
	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
	"github.com/RB-PRO/SanctionedClothing/pkg/gol"
)

const bitrixURL string = "https://clikshop.ru/parser/api/%s/index.php"

// Структура запросника к битриксу
type BitrixUser struct {
	MapCoast map[string]CoastMap // Мапа цен на товары
	Nots     *notification.Notification
	log      *gol.Gol // Логгирование
	cb       *cbbank.CentrakBank
}

// Создать клиента для работы с данными битрикс
func NewBitrixUser() *BitrixUser {
	return &BitrixUser{log: gol.NewGol()}
}
