package bitrixupdate

import (
	"fmt"
	"strings"
)

// Обновить цены и наличие по ОДНОМУ товару
func (bx *BitrixUser) UpdateProduct(ProductID string) error {

	// Получить подробнее о товаре
	ProductsDetail, ErrProduct := bx.Product([]string{ProductID})
	if ErrProduct != nil {
		return fmt.Errorf("bitrix: UpdateProduct: %w", ErrProduct)
	}
	if len(ProductsDetail.Products) == 0 {
		return fmt.Errorf("bitrix: UpdateProduct: %w", ErrProduct)
	}

	// Запрос на обновление цен и наличия
	var variationReq []Variation_Request
	var ErrUpdate error

	// Смотрим ссылку  для определения источника того, откуда пришёл товар
	Link := ProductsDetail.Products[0].Link
	switch {
	case strings.Contains(Link, "massimodutti"):
		variationReq, ErrUpdate = bx.UpdateMassimoDutti(ProductsDetail)
		if ErrUpdate != nil {
			bx.Nots.Sends(fmt.Sprintf("bitrix: UpdateMassimoDutti: %v", ErrUpdate))
			return fmt.Errorf("bitrix: UpdateMassimoDutti: %w", ErrUpdate)
		}
	default:
		return fmt.Errorf("bitrix: UpdateProduct: Не знаю, какую логику применить к товару %s", ProductsDetail.Products[0].ID)
	}

	///////////////////////////////////////////////////////

	fmt.Printf("\nvariationReq\n")
	for i := range variationReq {
		fmt.Printf("%+v\n", variationReq[i])
	}
	// Запрос на обновление даннных
	_, ErrVariation := bx.Variation(variationReq)
	if ErrVariation != nil {
		return fmt.Errorf("bitrix: Variation: %w", ErrVariation)
	}

	return nil
}
