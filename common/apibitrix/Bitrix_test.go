package apibitrix

import (
	"fmt"
	"testing"

	"ClikShop/common/bases"
)

func TestRequest(t *testing.T) { // Комплексное тестирование всех методов Bitrix
	bx, _ := NewBitrixUser()

	// Products
	fmt.Println("Products:")
	ProductsID, ErrProducts := bx.Products()
	if ErrProducts != nil {
		t.Error(ErrProducts)
	}
	fmt.Println("В Bitrix всего", len(ProductsID), "товаров.")
	fmt.Println("Анализируем на примере товара с ID", ProductsID[0])
	fmt.Println()

	// Product
	var price float64 // Цена для анализа
	ProductsDetail, ErrProduct := bx.Product([]string{ProductsID[0]})
	if ErrProduct != nil {
		t.Error(ErrProduct)
	}
	if len(ProductsDetail.Products) != 1 {
		t.Error("При выполнении запроса Product с одним входным параметром", ProductsID[0], "получен запрос, с к-вом параметров равным:", len(ProductsDetail.Products))
	}
	fmt.Println("Product:")
	fmt.Println("ID:", ProductsDetail.Products[0].ID)
	fmt.Println("link:", ProductsDetail.Products[0].Link)
	for _, ProductDetail := range ProductsDetail.Products[0].Colors {
		fmt.Println("->ID:", ProductDetail.ID)
		fmt.Println("---> ColorEng:", ProductDetail.ColorEng)
		fmt.Println("---> Price:", ProductDetail.Price)
		fmt.Println("---> Size:", ProductDetail.Size)
		fmt.Println("---> Link:", ProductDetail.Link)
	}
	price = ProductsDetail.Products[0].Colors[0].Price
	fmt.Println("Анализируем на примере вариации с ID", ProductsDetail.Products[0].Colors[0].ID,
		"и ценой", ProductsDetail.Products[0].Colors[0].Price)
	fmt.Println()

	// variation, чтобы изменить значения исследуемого товара
	variationReq := make([]Variation_Request, 1)
	variationReq[0].ID = ProductsDetail.Products[0].Colors[0].ID
	variationReq[0].Price = 123.0
	variationReq[0].Availability = false
	VariationResp, ErrVariation := bx.Variation(variationReq)
	if ErrVariation != nil {
		t.Error(ErrVariation)
	}
	if len(VariationResp.Error) != 0 {
		t.Error(VariationResp.Error)
	}
	fmt.Println()

	// Product 2, для того, чтобы точно зафиксировать, что значения цены изменились!
	ProductsDetail, ErrProduct = bx.Product([]string{ProductsID[0]})
	if ErrProduct != nil {
		t.Error(ErrProduct)
	}
	if len(ProductsDetail.Products) != 1 {
		t.Error("При выполнении запроса Product с одним входным параметром", ProductsID[0], "получен запрос, с к-вом параметров равным:", len(ProductsDetail.Products))
	}
	if ProductsDetail.Products[0].Colors[0].Price == price {
		t.Error("Цена не обновилась при запросе метода variation. Была:", price, "Стала:", ProductsDetail.Products[0].Colors[0].Price)
	}

	// variation 2, чтобы вернуть всё на место
	variationReq[0].ID = ProductsDetail.Products[0].Colors[0].ID
	variationReq[0].Price = price
	variationReq[0].Availability = true
	VariationResp, ErrVariation = bx.Variation(variationReq)
	if ErrVariation != nil {
		t.Error(ErrVariation)
	}
	if len(VariationResp.Error) != 0 {
		t.Error(VariationResp.Error)
	}
	fmt.Println()

	// Product 3, для того, чтобы точно зафиксировать, что значения цены вернулись на место
	ProductsDetail, ErrProduct = bx.Product([]string{ProductsID[0]})
	if ErrProduct != nil {
		t.Error(ErrProduct)
	}
	if len(ProductsDetail.Products) != 1 {
		t.Error("При выполнении запроса Product с одним входным параметром", ProductsID[0], "получен запрос, с к-вом параметров равным:", len(ProductsDetail.Products))
	}
	if ProductsDetail.Products[0].Colors[0].Price != price {
		t.Error("Цена не обновилась при запросе метода variation 2. Должна была стать:", price, "Стала:", ProductsDetail.Products[0].Colors[0].Price)
	}
}

func TestCoasts(t *testing.T) {
	bx, _ := NewBitrixUser()
	Coasts, ErrCoasts := bx.Coasts()
	if ErrCoasts != nil {
		t.Error(ErrCoasts)
	}
	fmt.Println(Coasts)
}

func TestAdd(t *testing.T) {
	bx, _ := NewBitrixUser()
	Coasts, ErrCoasts := bx.AddProduct(bases.Product2{
		Name:         "test",
		Article:      "1234567890",
		GenderLabel:  "man",
		Manufacturer: "crocs",
		Img: []string{"https://cdn.dsmcdn.com/mnresize/1200/1800/ty1086/product/media/images/prod/SPM/PIM/20231213/10/834b9534-e5cf-3667-82ce-b52af1ca4ed5/1_org_zoom.jpg",
			"https://cdn.dsmcdn.com/mnresize/1200/1800/ty1088/product/media/images/prod/SPM/PIM/20231213/10/0ee8874c-7274-34c9-821e-b3036e95b6ba/1_org_zoom.jpg"},
		Item: []bases.ColorItem{{ColorCode: "beliy", ColorEng: "Beliy", ColorRus: "Белый", Price: 228.0,
			Size: []bases.Size{{Val: "L", IsExit: true}},
			Image: []string{"https://cdn.dsmcdn.com/mnresize/1200/1800/ty1086/product/media/images/prod/SPM/PIM/20231213/10/834b9534-e5cf-3667-82ce-b52af1ca4ed5/1_org_zoom.jpg",
				"https://cdn.dsmcdn.com/mnresize/1200/1800/ty1088/product/media/images/prod/SPM/PIM/20231213/10/0ee8874c-7274-34c9-821e-b3036e95b6ba/1_org_zoom.jpg"},
		}},
	})

	if ErrCoasts != nil {
		t.Error(ErrCoasts)
	}
	fmt.Println(Coasts)
}

func TestUpdateSize(t *testing.T) {
	bx, _ := NewBitrixUser()

	req := []VariationSize_Request{
		{
			ID:   "418084",
			Size: "XL",
		},
	}

	ErrProducts := bx.UpdateSizeVariation(req)
	if ErrProducts != nil {
		t.Error(ErrProducts)
	}
}
