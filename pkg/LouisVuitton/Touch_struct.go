package louisvuitton

// Структура ответа на запрос информации о конкретном товаре
type TouchResponse struct {
	Model []struct {
		ProductVisible bool   `json:"productVisible"`
		MacroColor     string `json:"macroColor"`
		Links          struct {
			Look struct {
				Href string `json:"href"`
			} `json:"look"`
			Parent struct {
				Href string `json:"href"`
			} `json:"parent"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			PersoDetails struct {
				Href string `json:"href"`
			} `json:"perso-details"`
			APIRelsRusRuAddToCartBaseNvprod3900007VM462711 struct {
				Name string `json:"name"`
				Href string `json:"href"`
			} `json:"/api/rels/rus-ru/add-to-cart-base/nvprod3900007v/M46271/1"`
		} `json:"_links"`
		Type   string `json:"@type"`
		Height struct {
			Type     string `json:"@type"`
			Value    int    `json:"value"`
			UnitText string `json:"unitText"`
		} `json:"height"`
		ProductDetailTabs  []any  `json:"productDetailTabs"`
		Name               string `json:"name"`
		PackagingVariation string `json:"packagingVariation"`
		MainMaterial2      string `json:"mainMaterial2"`
		Width              struct {
			Type     string `json:"@type"`
			Value    int    `json:"value"`
			UnitText string `json:"unitText"`
		} `json:"width"`
		Offers struct {
			PriceSpecification struct {
				PriceCurrency string `json:"priceCurrency"`
				Price         int    `json:"price"`
			} `json:"priceSpecification"`
			Type  string `json:"@type"`
			Price string `json:"price"`
		} `json:"offers"`
		DisambiguatingDescription string `json:"disambiguatingDescription"`
		Image                     []struct {
			ContentURL string `json:"contentUrl"`
			Type       string `json:"@type"`
			Name       string `json:"name"`
			Context    string `json:"@context"`
			PlayerType string `json:"playerType,omitempty"`
		} `json:"image"`
		URL             string `json:"url"`
		Material        string `json:"material"`
		HighEndTemplate bool   `json:"highEndTemplate"`
		Depth           struct {
			Type     string `json:"@type"`
			Value    int    `json:"value"`
			UnitText string `json:"unitText"`
		} `json:"depth"`
		ProductID          string `json:"productId"`
		AdditionalProperty []struct {
			Type string `json:"@type"`
			Name string `json:"name"`
			// Value string `json:"value"`
		} `json:"additionalProperty"`
		Category []struct {
			Type          string `json:"@type"`
			URL           string `json:"url"`
			AlternateName string `json:"alternateName"`
			Name          string `json:"name"`
			Identifier    string `json:"identifier"`
		} `json:"category"`
		CareLink struct {
			Sections []any `json:"sections"`
		} `json:"careLink"`
		Disclaimers struct {
			CustomDisclaimer bool   `json:"customDisclaimer"`
			Value            string `json:"value"`
			PopInContent     string `json:"popInContent"`
			PopInTitle       string `json:"popInTitle"`
		} `json:"disclaimers"`
		Collectibles                bool   `json:"collectibles"`
		Color                       string `json:"color"`
		Context                     string `json:"@context"`
		Identifier                  string `json:"identifier"`
		DefaultFamilySapDisplayName string `json:"defaultFamilySapDisplayName"`
	} `json:"model"`
	PanelInformations []struct {
		Type         string `json:"@type"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		Introduction string `json:"introduction"`
	} `json:"panelInformations"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Push struct {
			Href string `json:"href"`
		} `json:"push"`
		CareLink struct {
			Href string `json:"href"`
		} `json:"care-link"`
		Availability struct {
			Href string `json:"href"`
		} `json:"availability"`
	} `json:"_links"`
	Type        string `json:"@type"`
	Sku         string `json:"sku"`
	IsSimilarTo []any  `json:"isSimilarTo"`
	ProductTips []any  `json:"productTips"`
	URL         string `json:"url"`
	Material    string `json:"material"`
	Category    []struct {
		Type          string `json:"@type"`
		URL           string `json:"url"`
		AlternateName string `json:"alternateName"`
		Name          string `json:"name"`
		Identifier    string `json:"identifier"`
	} `json:"category"`
	AdditionalProperty []struct {
		Type string `json:"@type"`
		Name string `json:"name"`
		// Value bool   `json:"value"`
	} `json:"additionalProperty"`
	Name        string `json:"name"`
	IsRelatedTo []any  `json:"isRelatedTo"`
	Context     string `json:"@context"`
	Identifier  string `json:"identifier"`
}
