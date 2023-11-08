package bitrixupdate

import (
	"fmt"
	"sort"
	"strings"

	zaratr "github.com/RB-PRO/SanctionedClothing/pkg/ZaraTR"
	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

// Обновить цены и наличие по ОДНОМУ товару
func (bx *BitrixUser) UpdateZara(ProductsDetail Product_Response) ([]Variation_Request, error) {

	// Ссылки на все вариации в подтоваре
	Link := ProductsDetail.Products[0].Link
	Code := strings.ReplaceAll(Link, "https://www.zara.com/tr/en/", "")
	Code = strings.ReplaceAll(Code, ".html?ajax=true", "")

	touch, ErrTouch := zaratr.LoadTouch(Code) // Выполняем запрос
	if ErrTouch != nil {
		return nil, fmt.Errorf("touch: %s", ErrTouch)
	}
	Prod2 := zaratr.Touch2Product2(touch) // АПереводим в структуру Product2

	// Решение задачи сличения данных из битрикса и из донора

	// Алгоритм обхода по результатам bx.Product в соответствии с massimodutti.Toucher
	// с целью созданию нового запросника для обновления данных в bitrix. Сложность o(n*n) - ужасная
	variationReq := make([]Variation_Request, 0)

	for _, BXproduct := range ProductsDetail.Products[0].Colors {
		for iColor, Color := range Prod2.Item {
			BxCr := EditColorName(BXproduct.ColorEng) // Citrix Color Name
			PrCr := EditColorName(Color.ColorEng)     // Product Color Name

			// Вероятность того, что слова похожи - BxCr, PrCr
			similarity := strutil.Similarity(BxCr, PrCr, metrics.NewLevenshtein())
			Prod2.Item[iColor].Similarity = similarity // Сохраняем результат
		}

		// Сортирвоать по убыванию
		sort.Slice(Prod2.Item, func(i, j int) bool {
			return Prod2.Item[i].Similarity > Prod2.Item[j].Similarity
		})

		// Если нет никаких товаров, то пропускаем
		if len(Prod2.Item) == 0 {
			continue
		}

		// Проверка того, что similarity у первых элементов одинаково
		if len(Prod2.Item) > 1 {
			if Prod2.Item[0].Similarity == Prod2.Item[1].Similarity {
				bx.log.Warn(fmt.Sprintf("Zara: В товаре %s совпадают параметры похожести Similarity. %.3f для цвета %s и %.3f для цвета %s. А в Битрикс - %s.\n",
					ProductsDetail.Products[0].ID, Prod2.Item[0].Similarity, Prod2.Item[0].ColorEng, Prod2.Item[1].Similarity, Prod2.Item[1].ColorEng, BXproduct.ColorEng))
				bx.Nots.Sends(fmt.Sprintf("Zara: В товаре %s совпадают параметры похожести Similarity. %.3f для цвета %s и %.3f для цвета %s. А в Битрикс - %s.\n",
					ProductsDetail.Products[0].ID, Prod2.Item[0].Similarity, Prod2.Item[0].ColorEng, Prod2.Item[1].Similarity, Prod2.Item[1].ColorEng, BXproduct.ColorEng))
				continue
			}
		}

		// Если вариации есть и необходимо обновлние
		if len(Prod2.Item) >= 1 {
			// Смотрим все размер на соответствие
			for _, ProdColor := range Prod2.Item[0].Size {
				if strings.Contains(EditColorName(BXproduct.Size), EditColorName(ProdColor.Val)) {
					// fmt.Println(Product.Item[0].Price,
					// 	bx.MapCoast[Product.Manufacturer].Walrus,
					// 	float64(bx.MapCoast[Product.Manufacturer].Delivery))
					variationReq = append(variationReq, Variation_Request{
						ID: BXproduct.ID,
						Price: bases.EditDecadense((bx.cb.Data.Valute.Try.Value/10)*Prod2.Item[0].Price*bx.MapCoast["zara"].Walrus +
							float64(bx.MapCoast["zara"].Delivery)),
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

	bx.log.Info(fmt.Sprintf("Zara: В товаре %s(%s) на обвновление идут %d товара",
		ProductsDetail.Products[0].ID, Link, len(variationReq)))
	return variationReq, nil
}
