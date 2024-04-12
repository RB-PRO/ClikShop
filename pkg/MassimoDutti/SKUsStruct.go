package massimodutti

type ID struct {
	SortedProductIdsByPricesAsc []int `json:"sortedProductIdsByPricesAsc"`
	ProductIds                  []int `json:"productIds"` // Список ID товаров
}
