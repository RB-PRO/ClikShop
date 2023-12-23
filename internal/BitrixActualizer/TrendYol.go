package actualizer

import (
	"fmt"
	"time"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/trendyol"
	"github.com/cheggaaa/pb"
)

func (bx *bitrixActualizer) trendyol(folder string) error {

	MakeDir(folder)

	ShopID := 106871 // SneakSup
	ProductGroupIDs, ErrGroup := trendyol.Pages(ShopID)
	if ErrGroup != nil {
		return fmt.Errorf("trendyol.Pages: %v", ErrGroup)
	}

	BarProducts := pb.StartNew(len(ProductGroupIDs))
	BarProducts.Prefix("trendyol")
	var Products bases.Variety2
	for _, ProductGroupID := range ProductGroupIDs {
		// fmt.Println(iProductGroupID, len(ProductGroupIDs))
		Product, ErrProduct := trendyol.Product(ProductGroupID, ShopID)
		if ErrProduct != nil {
			// panic(ErrProduct)
			// fmt.Println(ErrProduct)
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
	BarProducts.Finish()
	Products.SaveJson(fmt.Sprintf("%s/trendyol_%d", folder, ShopID))

	return nil
}
