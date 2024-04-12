package trendyol

type PageStruct struct {
	// IsSuccess  bool `json:"isSuccess"`
	// StatusCode int  `json:"statusCode"`
	// Error      any  `json:"error"`
	Result struct {
		// SlpName  string `json:"slpName"`
		Products []struct {
			ID             int `json:"id"`             // ID товара
			ProductGroupID int `json:"productGroupId"` // ID категории
			// Name           string   `json:"name"`
			// Images         []string `json:"images"`
			// ImageAlt       string   `json:"imageAlt"`
			// Brand          struct {
			// 	ID   int    `json:"id"`
			// 	Name string `json:"name"`
			// } `json:"brand"`
			// Tax int `json:"tax"`
			// // RatingScore struct {
			// // 	// AverageRating int `json:"averageRating"`
			// // 	TotalCount int `json:"totalCount"`
			// // } `json:"ratingScore,omitempty"`
			// ShowSexualContent bool   `json:"showSexualContent"`
			// HasReviewPhoto    bool   `json:"hasReviewPhoto"`
			// CardType          string `json:"cardType"`
			// Sections          []struct {
			// 	ID string `json:"id"`
			// } `json:"sections"`
			// Variants []struct {
			// 	AttributeValue string `json:"attributeValue"`
			// 	AttributeName  string `json:"attributeName"`
			// 	Price          struct {
			// 		// DiscountedPrice int `json:"discountedPrice"`
			// 		BuyingPrice   int `json:"buyingPrice"`
			// 		OriginalPrice int `json:"originalPrice"`
			// 		SellingPrice  int `json:"sellingPrice"`
			// 	} `json:"price"`
			// 	ListingID  string `json:"listingId"`
			// 	CampaignID int    `json:"campaignId"`
			// 	MerchantID int    `json:"merchantId"`
			// 	// DiscountedPriceInfo  string `json:"discountedPriceInfo"`
			// 	HasCollectableCoupon bool `json:"hasCollectableCoupon"`
			// 	SameDayShipping      bool `json:"sameDayShipping"`
			// } `json:"variants"`
			// CategoryHierarchy string `json:"categoryHierarchy"`
			// CategoryID        int    `json:"categoryId"`
			// CategoryName      string `json:"categoryName"`
			// URL               string `json:"url"`
			// MerchantID        int    `json:"merchantId"`
			// CampaignID        int    `json:"campaignId"`
			// Price             struct {
			// 	SellingPrice  int `json:"sellingPrice"`
			// 	OriginalPrice int `json:"originalPrice"`
			// 	// DiscountedPrice int `json:"discountedPrice"`
			// 	BuyingPrice int `json:"buyingPrice"`
			// } `json:"price"`
			// Promotions           []any  `json:"promotions"`
			// RushDeliveryDuration int    `json:"rushDeliveryDuration"`
			// FreeCargo            bool   `json:"freeCargo"`
			// CampaignName         string `json:"campaignName"`
			// ListingID            string `json:"listingId"`
			// WinnerVariant        string `json:"winnerVariant"`
			// ItemNumber           int    `json:"itemNumber"`
			// // DiscountedPriceInfo         string `json:"discountedPriceInfo"`
			// HasVideoContent             bool `json:"hasVideoContent"`
			// HasCrossPromotion           bool `json:"hasCrossPromotion"`
			// HasCollectableCoupon        bool `json:"hasCollectableCoupon"`
			// SameDayShipping             bool `json:"sameDayShipping"`
			// IsLegalRequirementConfirmed bool `json:"isLegalRequirementConfirmed"`
			// Badges                      []struct {
			// 	Title string `json:"title"`
			// 	Type  string `json:"type"`
			// } `json:"badges"`
			// Stamps []struct {
			// 	ImageURL    string  `json:"imageUrl"`
			// 	Type        string  `json:"type"`
			// 	Position    string  `json:"position"`
			// 	AspectRatio float64 `json:"aspectRatio"`
			// 	Priority    int     `json:"priority"`
			// } `json:"stamps,omitempty"`
			// LowestPriceDuration int `json:"lowestPriceDuration,omitempty"`
			// PriceLabel          struct {
			// 	Name  string `json:"name"`
			// 	Value string `json:"value"`
			// 	Type  int    `json:"type"`
			// } `json:"priceLabel,omitempty"`
		} `json:"products"`
		TotalCount int `json:"totalCount"`
		// RoughTotalCount string `json:"roughTotalCount"`
		// SearchStrategy  string `json:"searchStrategy"`
		// Title           string `json:"title"`
		// UxLayout        string `json:"uxLayout"`
		// QueryTerm       string `json:"queryTerm"`
		// PageIndex       int    `json:"pageIndex"`
		// Widgets         []any  `json:"widgets"`
	} `json:"result"`
	// Headers struct {
	// 	Tysidecarcachable string `json:"tysidecarcachable"`
	// } `json:"headers"`
}
