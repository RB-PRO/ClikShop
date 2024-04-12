package zaraapp_test

import (
	"fmt"
	"log"
	"testing"

	zaratr "github.com/RB-PRO/ClikShop/pkg/ZaraTR"
	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/cbbank"
	"github.com/RB-PRO/ClikShop/pkg/wcprod"
)

// Тестовая загрузка единичного товара
// go test -v -run ^TestSingleAddProduct$ github.com/RB-PRO/ClikShop/internal/ZaraApp
func TestSingleAddProduct(t *testing.T) {

	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		t.Error(ErrorCB)
	}
	fmt.Println("Курс лиры", cb.Data.Valute.Try.Value)

	// Загружаем товары на WC //
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		t.Error(errorInitWcAdd)
	}

	////////// Парсим один товар// Категории
	CatArr := zaratr.CatCycle() // Получить все категории
	fmt.Println("Всего", len(CatArr.Items), "категорий")

	var cat zaratr.Item
	for _, val := range CatArr.Items {
		if val.ID.String() == "2184366" {
			cat = val
			// fmt.Printf("%v - cat: %+v\n\n", ind, cat)
		}
	}
	if cat.ID.String() == "" {
		t.Error("Не нашёл товар с категорией 2184366")
	}

	// Список всех товаров
	// cat := CatArr.Items[1]
	fmt.Println("ID категории", cat.ID.String())
	fmt.Println("Категория товара:", cat.Cat) // WOMAN > SHIRTS > Satin
	// fmt.Printf("Весь товар: %v\n\n", cat.Cat) // WOMAN > SHIRTS > Satin
	line, ErrorLine := zaratr.LoadLine(fmt.Sprintf("%v", cat.ID.String()))
	if ErrorLine != nil {
		fmt.Println(ErrorLine)
	}

	/////////////

	ProductsLine := make([]zaratr.CommercialComponents, 0)
	if len(line.ProductGroups) != 0 {
		if len(line.ProductGroups) != 0 {
			if len(line.ProductGroups[0].Elements) != 0 {
				for ind := range line.ProductGroups[0].Elements[0].CommercialComponents { // Циклом обновляем категории
					if line.ProductGroups[0].Elements[0].CommercialComponents[ind].Type == "Product" { // Если это сам товар
						line.ProductGroups[0].Elements[0].CommercialComponents[ind].Cat = cat.Cat
						ProductsLine = append(ProductsLine, line.ProductGroups[0].Elements[0].CommercialComponents[ind])
					}
				}
			}
		}
	}
	fmt.Println("Всего", len(ProductsLine), "товара(ов)")

	// Сам товар
	prod := ProductsLine[0]
	var Variety bases.Variety2
	fmt.Println("Ссылка на товар", (prod.Seo.Keyword + "-p" + prod.Seo.SeoProductID))
	touch, _ := zaratr.LoadTouch(prod.Seo.Keyword + "-p" + prod.Seo.SeoProductID)
	Prod2 := zaratr.Touch2Product2(touch)
	Prod2.Cat = prod.Cat // Обновляем категнории
	Prod2.Article += "_test"
	fmt.Println("Артикул товара:", Prod2.Article)

	//fmt.Printf("%+#v", Prod2)
	fmt.Printf("В этом товаре всего %d цветов.\n", len(Prod2.Item))
	fmt.Println("Спарсили товар с параметрами:\n", bases.ProdStr(Prod2))

	Variety.Product = append(Variety.Product, Prod2)

	fmt.Println("Всего товаров для тестовой загрузки:", len(Variety.Product))

	///////////////////////////////////////////////////////////
	// Загружаем товары
	delivery := 500 // Доставка
	walrus := 1.3   // Моржа
	for i := 0; i < len(Variety.Product); i++ {
		if !Variety.Product[i].Upload {
			// Формирование адекватной цены доставки из файла
			ActualDelivery := Adding.EditDelivery(Variety.Product[i].Cat, delivery)
			Variety.Product[i] = bases.EditCoast(Variety.Product[i], cb.Data.Valute.Try.Value/10, walrus, ActualDelivery)
			//errorAddProductWC := Adding.AddProduct(wcprod.ProductTranslate(variety.Product[i])) //.AddAttr()
			Variety.Product[i], _ = Adding.YandexTranslate(Variety.Product[i])
			errorAddProductWC := Adding.AddProduct(Variety.Product[i]) //.AddAttr()
			if errorAddProductWC != nil {
				Variety.Product[i].Upload = true
			}
		}
	}
}

// go clean -testcache
// go test -v -run ^TestSingleAddProductLink$ github.com/RB-PRO/ClikShop/internal/ZaraApp
func TestSingleAddProductLink(t *testing.T) {

	// Нало работы с центральным банком
	cb, ErrorCB := cbbank.New() // Получить курс валюты
	if ErrorCB != nil {
		t.Error(ErrorCB)
	}
	log.Println("Курс лиры", cb.Data.Valute.Try.Value/10)

	// Загружаем товары на WC //
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		t.Error(errorInitWcAdd)
	}
	// "asymmetric-double-breasted-waistcoat---limited-edition-p07655728"
	// "metallic-block-heel-sandals-p13344110"
	// "lace-up-denim-wedge-heel-shoes-p12280210"
	touch, _ := zaratr.LoadTouch("100-linen-blazer-with-printed-cuffs-p07726707")
	Prod2 := zaratr.Touch2Product2(touch)

	// Prod2.Item[0].Image = Prod2.Item[0].Image[2:]
	for i, p := range Prod2.Item {
		for ii, pp := range p.Image {
			fmt.Printf("%d:%d. '%s'\n", i, ii, pp)
		}
	}
	delivery := 500 // Доставка
	walrus := 1.3   // Моржа
	Prod2 = bases.EditCoast(Prod2, cb.Data.Valute.Try.Value/10, walrus, delivery)
	//errorAddProductWC := Adding.AddProduct(wcprod.ProductTranslate(variety.Product[i])) //.AddAttr()
	// Prod2, _ = Adding.YandexTranslate(Prod2)
	Prod2.Article += "_test"

	fmt.Println("Спарсили товар с параметрами:\n", bases.ProdStr(Prod2))

	errorAddProductWC := Adding.AddProduct(Prod2) //.AddAttr()
	if errorAddProductWC != nil {
		t.Error(errorAddProductWC)
	}
	Prod2.Upload = true
}
