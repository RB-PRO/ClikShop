package zaratr

import "time"

// Ссылка на товар, [пример].
//
// [пример]: https://www.zara.com/tr/en/ribbed-strappy-vest-top-p03253306.html?ajax=true
const TouchURL string = "https://www.zara.com/tr/en/%v.html?ajax=true"

type Touch struct {
	Product struct {
		ID    int    `json:"id"`
		Type  string `json:"type"`
		Kind  string `json:"kind"`
		State string `json:"state"`
		Brand struct {
			BrandID        int    `json:"brandId"`
			BrandGroupID   int    `json:"brandGroupId"`
			BrandGroupCode string `json:"brandGroupCode"`
		} `json:"brand"`
		Name   string `json:"name"`
		Detail struct {
			Reference        string `json:"reference"`
			DisplayReference string `json:"displayReference"`
			Colors           []struct {
				ID        string `json:"id"`
				HexCode   string `json:"hexCode"`
				ProductID int    `json:"productId"`
				Name      string `json:"name"`
				Reference string `json:"reference"`
				StylingID string `json:"stylingId"`
				Xmedia    []struct {
					Datatype       string   `json:"datatype"`
					Set            int      `json:"set"`
					Type           string   `json:"type"`
					Kind           string   `json:"kind"`
					Path           string   `json:"path"`
					Name           string   `json:"name"`
					Width          int      `json:"width"`
					Height         int      `json:"height"`
					Timestamp      string   `json:"timestamp"`
					AllowedScreens []string `json:"allowedScreens"`
					Gravity        string   `json:"gravity"`
					ExtraInfo      struct {
						OriginalName string `json:"originalName"`
					} `json:"extraInfo"`
					Order int `json:"order"`
				} `json:"xmedia"`
				ShopcartMedia []struct {
					Datatype       string   `json:"datatype"`
					Set            int      `json:"set"`
					Type           string   `json:"type"`
					Kind           string   `json:"kind"`
					Path           string   `json:"path"`
					Name           string   `json:"name"`
					Width          int      `json:"width"`
					Height         int      `json:"height"`
					Timestamp      string   `json:"timestamp"`
					AllowedScreens []string `json:"allowedScreens"`
					Gravity        string   `json:"gravity"`
					ExtraInfo      struct {
						OriginalName string `json:"originalName"`
					} `json:"extraInfo"`
				} `json:"shopcartMedia"`
				Price int `json:"price"`
				Sizes []struct {
					Availability     string        `json:"availability"`
					EquivalentSizeID int           `json:"equivalentSizeId"`
					ID               int           `json:"id"`
					Name             string        `json:"name"`
					Price            int           `json:"price"`
					Reference        string        `json:"reference"`
					Sku              int           `json:"sku"`
					Attributes       []interface{} `json:"attributes"`
					Demand           string        `json:"demand"`
				} `json:"sizes"`
				Description    string `json:"description"`
				RawDescription string `json:"rawDescription"`
				ExtraInfo      struct {
					IsStockInStoresAvailable bool `json:"isStockInStoresAvailable"`
					HighlightPrice           bool `json:"highlightPrice"`
				} `json:"extraInfo"`
				TagTypes []struct {
					DisplayName string `json:"displayName"`
					Type        string `json:"type"`
				} `json:"tagTypes"`
				Attributes []interface{} `json:"attributes"`
				MainImgs   []struct {
					Datatype       string   `json:"datatype"`
					Set            int      `json:"set"`
					Type           string   `json:"type"`
					Kind           string   `json:"kind"`
					Path           string   `json:"path"`
					Name           string   `json:"name"`
					Width          int      `json:"width"`
					Height         int      `json:"height"`
					Timestamp      string   `json:"timestamp"`
					AllowedScreens []string `json:"allowedScreens"`
					Gravity        string   `json:"gravity"`
					ExtraInfo      struct {
						OriginalName string `json:"originalName"`
					} `json:"extraInfo"`
					Order int `json:"order"`
				} `json:"mainImgs"`
				PriceUnavailable bool `json:"priceUnavailable,omitempty"`
			} `json:"colors"`
			ColorSelectorLabel string        `json:"colorSelectorLabel"`
			MultipleColorLabel string        `json:"multipleColorLabel"`
			RelatedProducts    []interface{} `json:"relatedProducts"`
		} `json:"detail"`
		Section       int    `json:"section"`
		SectionName   string `json:"sectionName"`
		FamilyName    string `json:"familyName"`
		SubfamilyName string `json:"subfamilyName"`
		ExtraInfo     struct {
			IsSizeRecommender          bool   `json:"isSizeRecommender"`
			HasSpecialReturnConditions bool   `json:"hasSpecialReturnConditions"`
			HasInteractiveSizeGuide    bool   `json:"hasInteractiveSizeGuide"`
			ExtraDetailTitle           string `json:"extraDetailTitle"`
			IsBracketingRestricted     bool   `json:"isBracketingRestricted"`
			HasTipsOnExtraDetail       bool   `json:"hasTipsOnExtraDetail"`
			HighlightPrice             bool   `json:"highlightPrice"`
		} `json:"extraInfo"`
		Seo struct {
			Keyword     string `json:"keyword"`
			Description string `json:"description"`
			BreadCrumb  []struct {
				Text          string `json:"text"`
				Keyword       string `json:"keyword,omitempty"`
				ID            int    `json:"id"`
				SeoCategoryID int    `json:"seoCategoryId,omitempty"`
				Layout        string `json:"layout,omitempty"`
			} `json:"breadCrumb"`
			SeoProductID     string `json:"seoProductId"`
			DiscernProductID int    `json:"discernProductId"`
			KeyWordI18N      []struct {
				LangID  int    `json:"langId"`
				Keyword string `json:"keyword"`
			} `json:"keyWordI18n"`
		} `json:"seo"`
		FirstVisibleDate time.Time     `json:"firstVisibleDate"`
		Attributes       []interface{} `json:"attributes"`
		SizeGuide        struct {
			Enabled bool `json:"enabled"`
		} `json:"sizeGuide"`
		Xmedia                  []interface{} `json:"xmedia"`
		ProductTag              []interface{} `json:"productTag"`
		HasInteractiveSizeGuide bool          `json:"hasInteractiveSizeGuide"`
	} `json:"product"`
	ShowNativeAppBanner bool `json:"showNativeAppBanner"`
	ProductMetaData     []struct {
		Sku          string  `json:"sku"`
		Name         string  `json:"name"`
		Brand        string  `json:"brand"`
		Description  string  `json:"description"`
		Price        float64 `json:"price"`
		Availability string  `json:"availability"`
		Images       []struct {
			Datatype       string   `json:"datatype"`
			Set            int      `json:"set"`
			Type           string   `json:"type"`
			Kind           string   `json:"kind"`
			Path           string   `json:"path"`
			Name           string   `json:"name"`
			Width          int      `json:"width"`
			Height         int      `json:"height"`
			Timestamp      string   `json:"timestamp"`
			AllowedScreens []string `json:"allowedScreens"`
			Gravity        string   `json:"gravity"`
			ExtraInfo      struct {
				OriginalName string `json:"originalName"`
			} `json:"extraInfo"`
			Order int `json:"order"`
		} `json:"images"`
		URL string `json:"url"`
	} `json:"productMetaData"`
	ParentID int `json:"parentId"`
	Category struct {
	} `json:"category"`
	Categories  []interface{} `json:"categories"`
	BackURL     string        `json:"backUrl"`
	KeyWordI18N []struct {
		LangID  int    `json:"langId"`
		Keyword string `json:"keyword"`
	} `json:"keyWordI18n"`
	DocInfo struct {
		LastModified  time.Time `json:"lastModified"`
		Title         string    `json:"title"`
		Description   string    `json:"description"`
		Keywords      string    `json:"keywords"`
		PageID        string    `json:"pageId"`
		SeoAttributes string    `json:"seoAttributes"`
		RelData       struct {
			CanonicalURL   string `json:"canonicalUrl"`
			AlternatesData []struct {
				Lang string `json:"lang"`
				Href string `json:"href"`
			} `json:"alternatesData"`
		} `json:"relData"`
		HTMLAttributes struct {
		} `json:"htmlAttributes"`
	} `json:"docInfo"`
}
