package settings

import (
	"fmt"

	"github.com/RB-PRO/ClikShop/pkg/apibitrix"
)

func EDITSIZES() {
	bx, _ := apibitrix.NewBitrixUser()

	// req := []apibitrix.VariationSize_Request{
	// 	{
	// 		ID:   "418084",
	// 		Size: "XL",
	// 	},
	// }

	// ErrProducts := bx.UpdateSizeVariation(req)
	// if ErrProducts != nil {
	// 	panic(ErrProducts)
	// }

	// Получаем списки товаров
	ProductsID, ErrProducts := bx.Products()
	if ErrProducts != nil {
		panic(ErrProducts)
	}

	fmt.Println(len(ProductsID))

	fmt.Println()
	fmt.Println(getdatafromsizes(bx, "418079"))
	fmt.Println()
	fmt.Println()

	for iProductID, ProductID := range ProductsID {
		break
		mm := getdatafromsizes(bx, ProductID)

		fmt.Println(mm)
		fmt.Println()

		//
		if iProductID == 1 {
			break
		}
	}
}

func getdatafromsizes(bx *apibitrix.BitrixUser, ProductID string) map[string]map[string]string {
	ProdResp, _ := bx.Product([]string{ProductID})
	fmt.Println(ProdResp)
	mm := make(map[string]map[string]string)
	for _, prods := range ProdResp.Products {
		for _, color := range prods.Colors {
			if _, ok := mm[color.ColorEng]; !ok {
				mm[color.ColorEng] = make(map[string]string)
			}
			mm[color.ColorEng][color.Size] = color.ID
		}
	}
	return mm
}
