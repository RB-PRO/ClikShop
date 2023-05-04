package bases

// Структура массива товаров
type Variety2 struct {
	Product []Product2 // Массив продуктов
}

// Категория Name Slug
type Cat struct { // Категория товаров
	Name string // Название подкатегории
	Slug string // транслитом категория
	ID   int    // ID товара
}

// Структура товара
type Product2 struct {
	Cat []Cat // Подкатегория

	Name         string // Название товара
	FullName     string // Полное название товара
	Link         string // Сссылка на товар базового цвета
	Article      string // Артикул
	Manufacturer string // Производитель

	// Используется для tag
	GenderLabel string

	Size []string // Все возможные размеры

	Description struct { // Описание товара
		Eng string
		Rus string
	}

	// Описание товара по значению "цвет"
	// "Цвет" будет определять, как вариацию товара
	// "Цвет на русском"
	Item           []ColorItem
	Specifications map[string]string // Остальные характеристики

	Upload bool // Загружено или нет
}

// Цвета
type ColorItem struct {
	ColorCode string   // Цвет ключ-значение
	ColorEng  string   // Цвет на английском
	Link      string   // Ссылка на товар нужного цвета
	Price     float64  // Цена
	Size      []Size   // Размеры
	Image     []string // Картинки
}

type Size struct {
	Val    string // Размер одежды
	IsExit bool   // Есть в наличии
}
