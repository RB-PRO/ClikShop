package zaratr

import (
	"encoding/json"
	"errors"
	"strconv"

	"ClikShop/common/bases"
)

// Ссылка на категорию, [пример].
//
// [пример]: https://www.zara.com/tr/en/category/2184443/products
const LineURL string = "https://www.zara.com/tr/en/category/%v/products"

// Структура всех записей по ссылке всех товаров из категории
type Line struct {
	ProductGroups []ProductGroups `json:"productGroups"`
}

// Группы товаров
type ProductGroups struct {
	Type            string     `json:"type"`
	Elements        []Elements `json:"elements"`
	HasStickyBanner bool       `json:"hasStickyBanner"`
}

// Товар
type Elements struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name,omitempty"`
	Type                 string                 `json:"type,omitempty"`
	Layout               string                 `json:"layout,omitempty"`
	CommercialComponents []CommercialComponents `json:"commercialComponents,omitempty"`
	Animations           []interface{}          `json:"animations,omitempty"`
	HasStickyBanner      bool                   `json:"hasStickyBanner,omitempty"`
	NeedsSeparator       bool                   `json:"needsSeparator,omitempty"`
	Header               string                 `json:"header,omitempty"`
	Description          string                 `json:"description,omitempty"`
	PreserveInZoom2      bool                   `json:"preserveInZoom2,omitempty"`

	Cat []bases.Cat `json:"-"`
}

type Seo struct {
	Keyword          string `json:"keyword"`
	SeoProductID     string `json:"seoProductId"`
	DiscernProductID int    `json:"discernProductId"`
}
type ColorInfo struct {
	MainColorHexCode              string `json:"mainColorHexCode"`
	ShouldUseColorcutInColorLabel bool   `json:"shouldUseColorcutInColorLabel"`
}

type ID_int_or_string string // Тип для ID, который возможно int или string
type CustomIntToString struct {
	value string
}

func (cis *CustomIntToString) String() string {
	return cis.value
}
func (cis *CustomIntToString) Int() (int, error) {
	val, err := strconv.Atoi(cis.value)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// Кастомное декодирование JSON для ключа SearchID
func (cis *CustomIntToString) UnmarshalJSON(data []byte) error {
	if data[0] == 34 { // Если первый символ - Кавычка
		err := json.Unmarshal(data, &cis.value)
		if err != nil {
			return errors.New("CustomIntToString: UnmarshalJSON: Find 34: " + err.Error())
		}
	} else {
		// Добавление Кавычек в начале и в конце массива byte
		newData := make([]byte, 1)
		newData[0] = 34
		newData = append(newData, data...)
		newData = append(newData, 34)
		//newData[len(data)] = 34

		err := json.Unmarshal(newData, &cis.value)
		if err != nil {
			return errors.New("CustomIntToString: UnmarshalJSON: Find't 34: " + err.Error())
		}
	}
	return nil
}

type CommercialComponents struct {
	ID                     CustomIntToString             `json:"id"`
	Reference              string                        `json:"reference"`
	Type                   string                        `json:"type"`
	Kind                   string                        `json:"kind"`
	Brand                  Brand                         `json:"brand"`
	Xmedia                 []Xmedia                      `json:"xmedia"`
	Name                   string                        `json:"name"`
	Description            string                        `json:"description"`
	Price                  int                           `json:"price"`
	Section                int                           `json:"section"`
	SectionName            string                        `json:"sectionName"`
	FamilyName             string                        `json:"familyName"`
	SubfamilyName          string                        `json:"subfamilyName"`
	Detail                 Detail                        `json:"detail"`
	Seo                    Seo                           `json:"seo"`
	Availability           string                        `json:"availability"`
	TagTypes               []interface{}                 `json:"tagTypes"`
	ExtraInfo              ExtraInfoCommercialComponents `json:"extraInfo"`
	ColorInfo              ColorInfo                     `json:"colorInfo"`
	GridPosition           int                           `json:"gridPosition"`
	ZoomedGridPosition     int                           `json:"zoomedGridPosition"`
	PreservedBlockPosition int                           `json:"preservedBlockPosition"`
	AthleticzPosition      int                           `json:"athleticzPosition"`
	ProductTag             []interface{}                 `json:"productTag"`
	ColorList              string                        `json:"colorList"`
	IsDivider              bool                          `json:"isDivider"`
	HasXmediaDouble        bool                          `json:"hasXmediaDouble"`
	ShowExtraImageOnHover  bool                          `json:"showExtraImageOnHover"`
	ShowAvailability       bool                          `json:"showAvailability"`
	PriceUnavailable       bool                          `json:"priceUnavailable"`

	// массив категорий
	Cat    []bases.Cat `json:"-"`
	Gender string      `json:"-"`
}
type ExtraInfoCommercialComponents struct {
	IsDivider       bool `json:"isDivider"`
	HighlightPrice  bool `json:"highlightPrice"`
	HideProductInfo bool `json:"hideProductInfo"`
}
type Xmedia struct {
	Datatype       string          `json:"datatype"`
	Set            int             `json:"set"`
	Type           string          `json:"type"`
	Kind           string          `json:"kind"`
	Path           string          `json:"path"`
	Name           string          `json:"name"`
	Width          int             `json:"width"`
	Height         int             `json:"height"`
	Timestamp      string          `json:"timestamp"`
	AllowedScreens []string        `json:"allowedScreens"`
	ExtraInfo      ExtraInfoXmedia `json:"extraInfo"`
}
type ExtraInfoXmedia struct {
	OriginalName string `json:"originalName"`
}

type Colors struct {
	ID                 string          `json:"id"`
	ProductID          int             `json:"productId"`
	Name               string          `json:"name"`
	StylingID          string          `json:"stylingId"`
	OutfitID           string          `json:"outfitId"`
	Xmedia             []Xmedia        `json:"xmedia"`
	Price              int             `json:"price"`
	Availability       string          `json:"availability"`
	Reference          string          `json:"reference"`
	ExtraInfo          ExtraInfoColors `json:"extraInfo"`
	CanonicalReference string          `json:"canonicalReference"`
}
type ExtraInfoColors struct {
	HighlightPrice bool `json:"highlightPrice"`
}

type Brand struct {
	BrandID        int    `json:"brandId"`
	BrandGroupID   int    `json:"brandGroupId"`
	BrandGroupCode string `json:"brandGroupCode"`
}

type Detail struct {
	Reference        string   `json:"reference"`
	DisplayReference string   `json:"displayReference"`
	Colors           []Colors `json:"colors"`
}
