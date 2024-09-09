package zaratr

import "ClikShop/common/bases"

// Ссылка на все категории, [пример].
//
// [пример]: https://www.zara.com/tr/en/categories?ajax=true
const CategoriesURL string = "https://www.zara.com/tr/en/categories?ajax=true"

// Массив категорий
type Category struct {
	Categories []Subcategories `json:"categories"`
}

// Подкатегория
type Subcategories struct {
	Subcategories []Subcategories `json:"subcategories"` // Массив ссылок на каждую категорию
	Item                          // Наполнение для каждой категории
}
type Item struct {
	ID                 CustomIntToString `json:"id"`
	Name               string            `json:"name"`
	SectionName        string            `json:"sectionName"`
	Layout             string            `json:"layout"`
	ContentType        string            `json:"contentType"`
	GridLayout         string            `json:"gridLayout"`
	Seo                SeoCategory       `json:"seo"`
	Attributes         Attributes        `json:"attributes"`
	Key                string            `json:"key"`
	IsRedirected       bool              `json:"isRedirected"`
	IsCurrent          bool              `json:"isCurrent"`
	IsSelected         bool              `json:"isSelected"`
	HasSubcategories   bool              `json:"hasSubcategories"`
	Irrelevant         bool              `json:"irrelevant"`
	ViewOptions        ViewOptions       `json:"viewOptions"`
	MenuLevel          int               `json:"menuLevel"`
	RedirectCategoryID int               `json:"redirectCategoryId,omitempty"`

	Cat    []bases.Cat `json:"-"` // Слайс категорий
	Gender string      `json:"-"` // Гендер
}

type SeoCategory struct {
	SeoCategoryID  int    `json:"seoCategoryId"`
	Keyword        string `json:"keyword"`
	Irrelevant     bool   `json:"irrelevant"`
	IsHiddenInMenu bool   `json:"isHiddenInMenu"`
}
type Attributes struct {
	MustDisplayContent       bool           `json:"mustDisplayContent"`
	DisplayUnfolded          bool           `json:"displayUnfolded"`
	ShowSubcategories        bool           `json:"showSubcategories"`
	GridForcedView           GridForcedView `json:"gridForcedView"`
	ShowExtraImagesOnHover   bool           `json:"showExtraImagesOnHover"`
	IsMarketingMessageHidden bool           `json:"isMarketingMessageHidden"`
	IsDivider                bool           `json:"isDivider"`
}
type GridForcedView struct {
	Global string `json:"global"`
}
type ViewOptions struct {
	Zoom      string `json:"zoom"`
	IsDefault bool   `json:"isDefault"`
	IsForced  bool   `json:"isForced"`
}
