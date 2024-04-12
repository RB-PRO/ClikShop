package louisvuitton

import (
	"fmt"
	"strings"
)

// Структура по ТЗ
type Product struct {
	Brand    string   // Два варианта: Для нее или Для него. В зависимости от того мужская вещь или женская
	SKU      string   // Артикул. можно взять с сайта ЛВ. Там вроде бы есть айдишники у товаров в названии картинок, но можно поискать еще где-нибудь. Выглядят они как: M12312, но есть и длиннее по количеству цифр, надо смотреть.
	Title    string   // Вытащить название товара
	Category []string // Раздел товара. Собрать

	Description string // Описание товара из раздела Об изделии и также характериситики (взять без Булит поинтов, просто в одно предложение через запятую.)

	Height float64 // Высота из раздела Об изделии
	Width  float64 // Высота из раздела Об изделии
	Length float64 // Длина из раздела Об изделии

	External_ID string // Вставить сюда значение из поля SKU товара

	Variations []Variation // Вариации товаров
}

type Variation struct {
	Photo []string // У каждого товара у ЛВ есть фотографии. Нужно их спарсить и создать папку с понятным разделением на категории товара
	Mark  bool     // Достать поле “Новинка”

	SKU         string // Артикул вариации
	Parent_UID  string //Если есть вариант у товара, то указывается родительский ID, То есть его External ID и SKU
	External_ID string // Вставить сюда значение из поля SKU товара
	Editions    string // В этом поле должны быть описаны свойства варианта товара в таком формате: Цвет:Tan;Материал:Monogram Eclipse;Размер:45

	// Цены
	// Нужно взять цену с сайта Дубайской версии в AED, перевести в евро
	PriceRus float64 // Цена в Рублях
	PriceDub float64 // Цена в Дубаи
	PriceFr  float64 // Цена во Франции
}

// Привести структу ЛВ в структуру заказчика
func TouchResponse2Product(touch TouchResponse) (prod Product) {

	// Название товара
	prod.Title = touch.Name

	// Артикул
	prod.SKU = touch.Sku

	// Категория + Brand
	category := make([]string, len(touch.Category))
	for i := range touch.Category {
		if touch.Category[i].Name == "men" {
			prod.Brand = "Для него"
		}
		if touch.Category[i].Name == "women" {
			prod.Brand = "Для нее"
		}
		category[i] = touch.Category[i].Name
	}
	prod.Category = category

	// Пол в Brand
	if prod.Brand == "" {
		prod.Brand = "Унисекс"
	}

	// External_ID
	prod.External_ID = prod.SKU

	// Описание товара
	if len(touch.Model) > 0 {
		prod.Description = touch.Model[0].DisambiguatingDescription
	}

	// Длина, высота, ширина
	if len(touch.Model) > 0 {
		prod.Height = float64(touch.Model[0].Height.Value)
		prod.Width = float64(touch.Model[0].Width.Value)
		prod.Length = float64(touch.Model[0].Depth.Value)
	}

	// Обработка вариаций
	for _, val := range touch.Model {
		var variat Variation

		// Новинка
		if val.CommercialTag == "Новинки" {
			variat.Mark = true
		}

		// Артикулы
		variat.SKU = val.Identifier
		variat.Parent_UID = prod.SKU
		variat.External_ID = prod.SKU

		// Цена
		var price float64 = float64(val.Offers.PriceSpecification.Price)
		if val.Offers.PriceSpecification.PriceCurrency == "RUB" {
			variat.PriceRus = price
		}
		if val.Offers.PriceSpecification.PriceCurrency == "Дубай" {
			variat.PriceDub = price
		}
		if val.Offers.PriceSpecification.PriceCurrency == "Франция" {
			variat.PriceFr = price
		}

		// Описание товара - Editions - Цвет:Tan;Материал:Monogram Eclipse;Размер:45
		variat.Editions = fmt.Sprintf("Цвет:%s;Материал:%s;Размер:%v", val.Color, val.Material, val.SizeDisplayName)

		// Фотографии
		for _, image := range val.Image {
			if image.Type == "ImageObject" && image.PlayerType != "freecaster" {
				var imageUrl string = image.ContentURL
				ImageUrls := strings.Split(imageUrl, " ")
				if len(ImageUrls) > 0 {
					imageUrl = ImageUrls[0]
				}
				variat.Photo = append(variat.Photo, imageUrl+" view.png")
			}
		}

		prod.Variations = append(prod.Variations, variat)
	}

	return prod
}
