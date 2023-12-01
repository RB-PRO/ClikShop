package massimodutti_test

import (
	"fmt"
	"testing"

	massimodutti "github.com/RB-PRO/ClikShop/pkg/MassimoDutti"
	"github.com/RB-PRO/ClikShop/pkg/bases"
)

func TestParsing(t *testing.T) {

	// Получить все категории товаров
	categs, ErrCategory := massimodutti.Category()
	if ErrCategory != nil {
		t.Error(ErrCategory)
	}
	fmt.Println("> Category: Получили все категории")
	fmt.Println("Категория:", categs.Categories[0].Subcategories[8].ID, categs.Categories[0].Subcategories[8].Name)

	// Получить все ID товаров категории
	SKUs, ErrSKUs := massimodutti.SKUs(categs.Categories[0].Subcategories[8].ID)
	if ErrSKUs != nil {
		t.Error(ErrSKUs)
	}
	fmt.Println("> SKUs: Получили все ID товаров")
	fmt.Println("Всего товаров:", len(SKUs.ProductIds))
	fmt.Println("Будем рассматривать товар с ID", SKUs.ProductIds[1])

	// Получить обширную информацию по товарам. На входе - ID товара
	line, ErrLine := massimodutti.Lines([]int{SKUs.ProductIds[1]})
	if ErrLine != nil {
		t.Error(ErrLine)
	}
	fmt.Println("> Lines: Получили line для данной категории")
	fmt.Println("Рассмотрим товар с ID", line.Products[0].ID, massimodutti.URL+line.Products[0].ProductURL)

	// Преобразуем в структур продекта, удобную для нас
	prod2 := massimodutti.Line2Product2(line, []bases.Cat{})

	// Получить инфомрацию о конкретном товаре
	productTouch, ErrTouch := massimodutti.Toucher(line.Products[0].ID)
	if ErrTouch != nil {
		t.Error(ErrTouch)
	}
	fmt.Println("> Toucher: Получили информацию о товаре")
	fmt.Println("Название товара", productTouch.Name)
	fmt.Println()

	// Преобразуем структу точечного парсинга товара
	prod2[0] = massimodutti.Touch2Product2(prod2[0], productTouch)
	prod2[0].Size = bases.EditProdSize(prod2[0])

	fmt.Println(bases.ProdStr(prod2[0]))
}
