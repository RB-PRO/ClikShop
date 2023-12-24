package actualizer

import (
	"fmt"
	"strconv"
	"time"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/trendyol"
	"github.com/cheggaaa/pb"
)

func (bx *bitrixActualizer) trendyol(folder string) error {

	ShopIDs := []int{
		332585, // Levi's
		107483, // Aktaş Sport AS
		106871, // SneakSup
		815951, // HUGO
		804476, // BOSS
		742918, // Victoria's secret
	}
	fmt.Println(ShopIDs)

	for _, ShopID := range ShopIDs {
		bx.trendyolOne(folder, ShopID)
	}

	return nil
}

func (bx *bitrixActualizer) trendyolOne(folder string, ShopID int) error {

	MakeDir(folder)

	fmt.Println("ShopID trendyolOne", ShopID)
	ProductGroupIDs, ErrGroup := trendyol.Pages(ShopID)
	if ErrGroup != nil {
		return fmt.Errorf("trendyol.Pages: %v", ErrGroup)
	}

	fmt.Println(ProductGroupIDs)
	fmt.Println("len(ProductGroupIDs)", len(ProductGroupIDs))

	BarProducts := pb.StartNew(len(ProductGroupIDs))
	defer BarProducts.Finish()
	BarProducts.Prefix(strconv.Itoa(ShopID))
	var Products bases.Variety2
	for _, ProductGroupID := range ProductGroupIDs {
		Product, ErrProduct := trendyol.Product(ProductGroupID, ShopID)
		if ErrProduct != nil {
			// panic(ErrProduct)
			fmt.Println(ErrProduct)
			continue
		}

		// Да-да, может быть такое, что вариаций у товара не будет.
		// Например это может возникнуть, когда продавец вариаций товаров не оригинальный
		if len(Product.Item) != 0 {
			// Product
			Product = bases.EditDoubleColors(Product)
			Product.Size = bases.EditProdSize(Product)
			Product.Img = bases.EditIMG(Product)

			Products.Product = append(Products.Product, Product)
		}
		BarProducts.Increment()
		time.Sleep(time.Millisecond * 200)
	}

	Products.SaveJson(fmt.Sprintf("%s/trendyol_%d", folder, ShopID))

	return nil
}
