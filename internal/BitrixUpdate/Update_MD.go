package bitrixupdate

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	massimodutti "github.com/RB-PRO/SanctionedClothing/pkg/MassimoDutti"
	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

// Обновить цены и наличие по ОДНОМУ товару
func (bx *BitrixUser) UpdateMassimoDutti(ProductsDetail Product_Response) ([]Variation_Request, error) {

	Link := ProductsDetail.Products[0].Link // Основная ссылка на товар
	fmt.Println(Link)
	// Получение ID товара в системе massimodutti. Оно же Toucher
	Link = strings.ReplaceAll(Link, "https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/", "")
	Link = strings.ReplaceAll(Link, "/detail?languageId=-1&appId=1", "")
	ID, ErrAtoi := strconv.Atoi(Link)
	if ErrAtoi != nil {
		return nil, fmt.Errorf("UpdateMassimoDutti: Atoi: %w", ErrAtoi)
	}

	// Делаем запрос на получение данных
	touch, ErrToucher := massimodutti.Toucher(ID)
	if ErrToucher != nil {
		return nil, fmt.Errorf("UpdateMassimoDutti: Toucher: %w", ErrToucher)
	}
	var Product bases.Product2
	Product = massimodutti.Touch2Product2(Product, touch)

	// Алгоритм обхода по результатам bx.Product в соответствии с massimodutti.Toucher
	// с целью созданию нового запросника для обновления данных в bitrix. Сложность o(n*n) - ужасная
	variationReq := make([]Variation_Request, 0)

	BxColors := ProductsDetail.Products[0].Colors
	for _, BXproduct := range BxColors {
		for iItem, ProdItem := range Product.Item {
			BxCr := EditColorName(BXproduct.ColorEng) // Citrix Color Name
			PrCr := EditColorName(ProdItem.ColorEng)  // Product Color Name

			// Вероятность того, что слова похожи - BxCr, PrCr
			similarity := strutil.Similarity(BxCr, PrCr, metrics.NewLevenshtein())
			Product.Item[iItem].Similarity = similarity // Сохраняем результат

		}

		// Сортирвоать по убыванию
		sort.Slice(Product.Item, func(i, j int) bool {
			return Product.Item[i].Similarity > Product.Item[j].Similarity
		})

		// Если нет никаких товаров, то пропускаем
		if len(Product.Item) == 0 {
			continue
		}

		// Проверка того, что similarity у первых элементов одинаково
		if len(Product.Item) > 1 {
			if Product.Item[0].Similarity == Product.Item[1].Similarity {
				fmt.Printf("В товаре %s совпадают параметры похожести Similarity. %.3f для цвета %s и %.3f для цвета %s. А в Битрикс - %s.\n",
					ProductsDetail.Products[0].ID, Product.Item[0].Similarity, Product.Item[0].ColorEng, Product.Item[1].Similarity, Product.Item[1].ColorEng, BXproduct.ColorEng)

				bx.Nots.Sends(fmt.Sprintf("В товаре %s совпадают параметры похожести Similarity. %.3f для цвета %s и %.3f для цвета %s. А в Битрикс - %s.\n",
					ProductsDetail.Products[0].ID, Product.Item[0].Similarity, Product.Item[0].ColorEng, Product.Item[1].Similarity, Product.Item[1].ColorEng, BXproduct.ColorEng))
			}
		}

		// Если вариации есть и необходимо обновлние
		if len(Product.Item) >= 1 {
			// Смотрим все размер на соответствие
			for _, ProdColor := range Product.Item[0].Size {
				if strings.Contains(EditColorName(BXproduct.Size), EditColorName(ProdColor.Val)) {
					// fmt.Println(Product.Item[0].Price,
					// 	bx.MapCoast[Product.Manufacturer].Walrus,
					// 	float64(bx.MapCoast[Product.Manufacturer].Delivery))
					variationReq = append(variationReq, Variation_Request{
						ID: BXproduct.ID,
						Price: Product.Item[0].Price*bx.MapCoast[Product.Manufacturer].Walrus +
							float64(bx.MapCoast[Product.Manufacturer].Delivery),
						Availability: ProdColor.IsExit,
					})
				}
			}
		}

		// for iItem, ProdItem := range Product.Item {
		// 	BxCr := EditColorName(BXproduct.ColorEng) // Citrix Color Name
		// 	PrCr := EditColorName(ProdItem.ColorEng)  // Product Color Name
		// }

	}

	return variationReq, nil
}
