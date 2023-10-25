package bitrixupdate

const bitrixURL string = "https://clikshop.ru/parser/api/%s/index.php"

// Структура запросника к битриксу
type BitrixUser struct{}

// Создать клиента для работы с данными битрикс
func NewBitrixUser() *BitrixUser {
	return &BitrixUser{}
}
