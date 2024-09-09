package patrendyol

import (
	"fmt"
	"time"

	"ClikShop/common/bases"
	"ClikShop/common/trendyol"
)

func Parsing() {
	ShopID := 106871
	ProductGroupIDs, ErrGroup := trendyol.Pages(ShopID)
	if ErrGroup != nil {
		panic(ErrGroup)
	}

	var Products bases.Variety2
	for iProductGroupID, ProductGroupID := range ProductGroupIDs {
		fmt.Println(iProductGroupID, len(ProductGroupIDs))
		Product, ErrProduct := trendyol.Product(ProductGroupID, ShopID)
		if ErrProduct != nil {
			// panic(ErrProduct)
			fmt.Println(ErrProduct)
			continue
		}

		// Да-да, может быть такое, что вариаций у товара не будет.
		// Например это может возникнуть, когда продавец вариаций товаров не оригинальный
		if len(Product.Item) != 0 {
			Products.Product = append(Products.Product, Product)
		}

		time.Sleep(time.Millisecond * 100)
	}

	Products.SaveJson2("patrendyol.json")
}
