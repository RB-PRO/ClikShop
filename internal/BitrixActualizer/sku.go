package actualizer

// Полуть мапу всех артикулов товаров из магазина в Bitrix
func (bx *bitrixActualizer) MapSKU() (map[string]bool, error) {

	// Получаем список ID всех товаров
	skus, ErrProducts := bx.BX.SKUs()
	if ErrProducts != nil {
		return nil, ErrProducts
	}

	skuMap := make(map[string]bool)

	for _, sku := range skus {
		skuMap[sku] = true
	}

	return skuMap, nil
}
