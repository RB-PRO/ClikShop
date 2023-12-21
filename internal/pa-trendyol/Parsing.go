package patrendyol

import (
	"fmt"
	"time"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/trendyol"
)

func Parsing() {
	ShopID := "106871"
	ProductGroupIDs, ErrGroup := trendyol.Pages(ShopID)
	if ErrGroup != nil {
		panic(ErrGroup)
	}

	var Products bases.Variety2
	for iProductGroupID, ProductGroupID := range ProductGroupIDs {
		fmt.Println(iProductGroupID, len(ProductGroupIDs))
		Product, ErrProduct := trendyol.Product(ProductGroupID)
		if ErrProduct != nil {
			// panic(ErrProduct)
			fmt.Println(ErrProduct)
			continue
		}

		Products.Product = append(Products.Product, Product)
		time.Sleep(time.Millisecond * 100)
	}

	Products.SaveJson2("patrendyol.json")
}
