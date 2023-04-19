package zaraapp

import zaratr "github.com/RB-PRO/SanctionedClothing/pkg/ZaraTR"

// Начать парсить и одновременно загружать товары
func Start() {
	Category, ErrorCat := zaratr.LoadCategory()
	if ErrorCat != nil {
		panic(ErrorCat)
	}

}
