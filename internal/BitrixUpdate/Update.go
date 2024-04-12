package bitrixupdate

import (
	"fmt"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/apibitrix"
)

// Обновить цены и наличие по ОДНОМУ товару
func (bx *BitrixUpdator) UpdateProduct(ProductID string) error {

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
	case strings.Contains(Link, "trendyol"):
		variationReq, ErrUpdate = bx.UpdateTrandYol(ProductsDetail)
		if ErrUpdate != nil {
			// fmt.Println(fmt.Errorf("bitrix: Update: Zara: %w", ErrUpdate), Link)
			return fmt.Errorf("update: trendyol: %w", ErrUpdate)
		}
	default:
		return fmt.Errorf("update: Не знаю, какую логику применить к товару %s", ProductsDetail.Products[0].ID)
	}

	// Запрос на обновление даннных
	// fmt.Println("len(variationReq)", len(variationReq))
	if len(variationReq) != 0 {
		//for i := range variationReq {
		//	fmt.Printf("%d. %+v\n", i, variationReq[i])
		//}
		bx.BX.Log.Info(fmt.Sprintf("Для товара с ID %s https://213.226.124.16/bitrix/admin/iblock_element_edit.php?IBLOCK_ID=15&type=aspro_lite_catalog&lang=ru&ID=%s&find_section_section=0&WF=Y, с ссылкой на донора %s была подготовлена структура на обновление: %+v",
			ProductID, ProductID, Link, variationReq))
		_, ErrVariation := bx.BX.Variation(variationReq)
		if ErrVariation != nil {
			bx.BX.Log.Err(fmt.Errorf("bitrix: Variation: %w", ErrVariation))
			return fmt.Errorf("bitrix: Variation: %w", ErrVariation)
		}
	}

	return nil
}
