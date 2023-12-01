package bitrixupdate

import (
	"fmt"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/apibitrix"
)

// Обновить цены и наличие по ОДНОМУ товару
func (bx *bitrixUpdator) UpdateProduct(ProductID string) error {

	// Получить подробнее о товаре
	ProductsDetail, ErrProduct := bx.BX.Product([]string{ProductID})
	if ErrProduct != nil {
		return fmt.Errorf("bx.Product: %w", ErrProduct)
	}
	if len(ProductsDetail.Products) == 0 {
		return fmt.Errorf("bx.Product: len(ProductsDetail.Products) == 0")
	}

	// Запрос на обновление цен и наличия
	var variationReq []apibitrix.Variation_Request
	var ErrUpdate error

	// Смотрим ссылку  для определения источника того, откуда пришёл товар
	Link := ProductsDetail.Products[0].Link
	switch {
	case strings.Contains(Link, "massimodutti"):
		variationReq, ErrUpdate = bx.UpdateMassimoDutti(ProductsDetail)
		if ErrUpdate != nil {
			return fmt.Errorf("update: MD: %w", ErrUpdate)
		}
	case strings.Contains(Link, "hm.com"):
		variationReq, ErrUpdate = bx.UpdateHandM(ProductsDetail)
		if ErrUpdate != nil {
			return fmt.Errorf("update: HM: %w", ErrUpdate)
		}
	case strings.Contains(Link, "zara"):
		variationReq, ErrUpdate = bx.UpdateZara(ProductsDetail)
		if ErrUpdate != nil {
			// fmt.Println(fmt.Errorf("bitrix: Update: Zara: %w", ErrUpdate), Link)
			return fmt.Errorf("update: Zara: %w", ErrUpdate)
		}
	case strings.Contains(Link, "sneaksup.com"):
		variationReq, ErrUpdate = bx.UpdateSS(ProductsDetail)
		if ErrUpdate != nil {
			// fmt.Println(fmt.Errorf("bitrix: Update: Zara: %w", ErrUpdate), Link)
			return fmt.Errorf("update: SS: %w", ErrUpdate)
		}
	default:
		return fmt.Errorf("update: Не знаю, какую логику применить к товару %s", ProductsDetail.Products[0].ID)
	}

	// Запрос на обновление даннных
	// fmt.Println("len(variationReq)", len(variationReq))
	if len(variationReq) != 0 {
		for i := range variationReq {
			fmt.Printf("%d. %+v\n", i, variationReq[i])
		}
		_, ErrVariation := bx.BX.Variation(variationReq)
		if ErrVariation != nil {
			return fmt.Errorf("bitrix: Variation: %w", ErrVariation)
		}
	}

	return nil
}
