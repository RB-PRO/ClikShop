package louisvuitton

// Структура ответа на запрос информации о конкретном товаре
type TouchResponse struct {
	Model []struct {
		CommercialTag   string `json:"commercialTag"`   // новинка
		SizeDisplayName string `json:"sizeDisplayName"` // Размер

		ProductVisible bool   `json:"productVisible,omitempty"`
		MacroColor     string `json:"macroColor,omitempty"`
		Links          struct {
			Look struct {
				Href string `json:"href,omitempty"`
			} `json:"look,omitempty"`
			Parent struct {
				Href string `json:"href,omitempty"`
			} `json:"parent,omitempty"`
			Self struct {
				Href string `json:"href,omitempty"`
			} `json:"self,omitempty"`
			PersoDetails struct {
				Href string `json:"href,omitempty"`
			} `json:"perso-details,omitempty"`
			APIRelsRusRuAddToCartBase001054M781431 struct {
				Name string `json:"name,omitempty"`
				Href string `json:"href,omitempty"`
			} `json:"/api/rels/rus-ru/add-to-cart-base/001054/M78143/1,omitempty"`
		} `json:"_links,omitempty"`
		Type   string `json:"@type,omitempty"`
		Height struct {
			Type     string  `json:"@type,omitempty"`
			Value    float64 `json:"value,omitempty"`
			UnitText string  `json:"unitText,omitempty"`
		} `json:"height,omitempty"`
		Width struct {
			Type     string  `json:"@type,omitempty"`
			Value    float64 `json:"value,omitempty"`
			UnitText string  `json:"unitText,omitempty"`
		} `json:"width,omitempty"`
		ProductDetailTabs  []any  `json:"productDetailTabs,omitempty"`
		Name               string `json:"name,omitempty"`
		PackagingVariation string `json:"packagingVariation,omitempty"`
		MainMaterial2      string `json:"mainMaterial2,omitempty"`
		Offers             struct {
			PriceSpecification struct {
				PriceCurrency string `json:"priceCurrency,omitempty"`
				Price         int    `json:"price,omitempty"`
			} `json:"priceSpecification,omitempty"`
			Type  string `json:"@type,omitempty"`
			Price string `json:"price,omitempty"`
		} `json:"offers,omitempty"`
		DisambiguatingDescription string `json:"disambiguatingDescription,omitempty"`
		Image                     []struct {
			PlayerType string `json:"playerType,omitempty"`
			ContentURL string `json:"contentUrl,omitempty"`
			Type       string `json:"@type,omitempty"`
			Name       string `json:"name,omitempty"`
			Context    string `json:"@context,omitempty"`
		} `json:"image,omitempty"`
		URL             string `json:"url,omitempty"`
		Material        string `json:"material,omitempty"`
		HighEndTemplate bool   `json:"highEndTemplate,omitempty"`
		Depth           struct {
			Type     string  `json:"@type,omitempty"`
			Value    float64 `json:"value,omitempty"`
			UnitText string  `json:"unitText,omitempty"`
		} `json:"depth,omitempty"`
		ProductID string `json:"productId,omitempty"`
		// AdditionalProperty []struct {
		// 	Type  string `json:"@type,omitempty"`
		// 	Name  string `json:"name,omitempty"`
		// 	Value string `json:"value,omitempty"`
		// } `json:"additionalProperty,omitempty"`
		Category []struct {
			Type          string `json:"@type,omitempty"`
			URL           string `json:"url,omitempty"`
			AlternateName string `json:"alternateName,omitempty"`
			Name          string `json:"name,omitempty"`
			Identifier    string `json:"identifier,omitempty"`
		} `json:"category,omitempty"`
		CareLink struct {
			Sections []any `json:"sections,omitempty"`
		} `json:"careLink,omitempty"`
		Disclaimers struct {
			CustomDisclaimer bool   `json:"customDisclaimer,omitempty"`
			Value            string `json:"value,omitempty"`
			PopInContent     string `json:"popInContent,omitempty"`
			PopInTitle       string `json:"popInTitle,omitempty"`
		} `json:"disclaimers,omitempty"`
		Collectibles                bool   `json:"collectibles,omitempty"`
		Color                       string `json:"color,omitempty"`
		Context                     string `json:"@context,omitempty"`
		Identifier                  string `json:"identifier,omitempty"`
		DefaultFamilySapDisplayName string `json:"defaultFamilySapDisplayName,omitempty"`
	} `json:"model,omitempty"`
	DisambiguatingDescription string `json:"disambiguatingDescription,omitempty"`
	Links                     struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		Push struct {
			Href string `json:"href,omitempty"`
		} `json:"push,omitempty"`
		CareLink struct {
			Href string `json:"href,omitempty"`
		} `json:"care-link,omitempty"`
		Availability struct {
			Href string `json:"href,omitempty"`
		} `json:"availability,omitempty"`
	} `json:"_links,omitempty"`
	Type        string `json:"@type,omitempty"`
	Sku         string `json:"sku,omitempty"`
	IsSimilarTo []any  `json:"isSimilarTo,omitempty"`
	ProductTips []any  `json:"productTips,omitempty"`
	URL         string `json:"url,omitempty"`
	Material    string `json:"material,omitempty"`
	Category    []struct {
		Type          string `json:"@type,omitempty"`
		URL           string `json:"url,omitempty"`
		AlternateName string `json:"alternateName,omitempty"`
		Name          string `json:"name,omitempty"`
		Identifier    string `json:"identifier,omitempty"`
	} `json:"category,omitempty"`
	// AdditionalProperty []struct {
	// 	Type  string `json:"@type,omitempty"`
	// 	Name  string `json:"name,omitempty"`
	// 	Value bool   `json:"value,omitempty"`
	// } `json:"additionalProperty,omitempty"`
	Name        string `json:"name,omitempty"`
	IsRelatedTo []any  `json:"isRelatedTo,omitempty"`
	Context     string `json:"@context,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
}
