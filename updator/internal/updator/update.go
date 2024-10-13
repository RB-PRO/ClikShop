package updator

import (
	"ClikShop/common/apibitrix"
	"ClikShop/common/bases"
	"fmt"
	"strings"
)

// Ключ для хэш-мапы для определения каждой вариации
type key struct {
	size  string
	color string
}

// Обновить цены и наличие по ОДНОМУ товару
func (s *Service) updateProduct(ProductID string, priceFunc func(brand string, price float64) float64) error {

	// Получить подробнее о товаре
	ProductsDetail, ErrProduct := s.BitrixService.Product([]string{ProductID})
	if ErrProduct != nil {
		return fmt.Errorf("bx.Product: %w", ErrProduct)
	}
	if len(ProductsDetail.Products) == 0 {
		return fmt.Errorf("bx.Product: len(ProductsDetail.Products) == 0")
	}

	// Запрос на обновление цен и наличия
	var variationReq []apibitrix.Variation_Request
	var err error

	// Смотрим ссылку для определения источника того, откуда пришёл товар
	link := ProductsDetail.Products[0].Link
	s.Gol.Infof("Обновляю товар по ссылке %s", link)
	switch {
	case strings.Contains(link, bases.TagMD):
		variationReq, err = s.updateMassimoDutti(ProductsDetail, priceFunc)
		if err != nil {
			return fmt.Errorf("update: MD: %w", err)
		}
	case strings.Contains(link, bases.TagHM):
		variationReq, err = s.UpdateHandM(ProductsDetail, priceFunc)
		if err != nil {
			return fmt.Errorf("update: HM: %w", err)
		}
	case strings.Contains(link, bases.TagZara):
		variationReq, err = s.UpdateZara(ProductsDetail, priceFunc)
		if err != nil {
			return fmt.Errorf("update: Zara: %w", err)
		}
	case strings.Contains(link, bases.TagSS):
		variationReq, err = s.UpdateSS(ProductsDetail, priceFunc)
		if err != nil {
			return fmt.Errorf("update: SS: %w", err)
		}
	case strings.Contains(link, bases.TagTY):
		variationReq, err = s.UpdateTrendYol(ProductsDetail, priceFunc)
		if err != nil {
			return fmt.Errorf("update: trendyol: %w", err)
		}
	default:
		return fmt.Errorf("update: Не знаю, какую логику применить к товару %s", ProductsDetail.Products[0].ID)
	}

	// Запрос на обновление даннных
	if len(variationReq) != 0 {
		s.Gol.Info(fmt.Sprintf("Для товара с ID %s https://213.226.124.16/bitrix/admin/iblock_element_edit.php?IBLOCK_ID=15&type=aspro_lite_catalog&lang=ru&ID=%s&find_section_section=0&WF=Y, с ссылкой на донора %s была подготовлена структура на обновление: %+v",
			ProductID, ProductID, link, variationReq))

		if _, err := s.BitrixService.Variation(variationReq); err != nil {
			s.Gol.Err(fmt.Errorf("bitrix: Variation: %w", err))
			return fmt.Errorf("bitrix: Variation: %w", err)
		}
	}

	return nil
}
