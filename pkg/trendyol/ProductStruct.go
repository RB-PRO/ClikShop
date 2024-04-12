package trendyol

type ProductStruct struct {
	IsSuccess  bool `json:"isSuccess"`
	StatusCode int  `json:"statusCode"`
	Error      any  `json:"error"`
	Result     struct {
		Attributes []struct {
			Key struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"key"`
			Value struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"value"`
			Starred     bool   `json:"starred"`
			Description string `json:"description"`
			MediaUrls   []any  `json:"mediaUrls"`
		} `json:"attributes"`
		// SelectedVasAttributes []any `json:"selectedVasAttributes"`
		// AlternativeVariants   []any `json:"alternativeVariants"`
		Variants []struct {
			AttributeID    int    `json:"attributeId"`
			AttributeName  string `json:"attributeName"`
			AttributeType  string `json:"attributeType"`
			AttributeValue string `json:"attributeValue"`
			Stamps         []struct {
				Type int    `json:"type"`
				Text string `json:"text"`
			} `json:"stamps"`
			Price struct {
				ProfitMargin    int `json:"profitMargin"`
				DiscountedPrice struct {
					Text  string  `json:"text"`
					Value float64 `json:"value"`
				} `json:"discountedPrice"`
				SellingPrice struct {
					Text  string  `json:"text"`
					Value float64 `json:"value"`
				} `json:"sellingPrice"`
				OriginalPrice struct {
					Text  string  `json:"text"`
					Value float64 `json:"value"`
				} `json:"originalPrice"`
				Currency string `json:"currency"`
			} `json:"price"`
			FulfilmentType           string `json:"fulfilmentType"`
			AttributeBeautifiedValue string `json:"attributeBeautifiedValue"`
			IsWinner                 bool   `json:"isWinner"`
			ListingID                string `json:"listingId"`
			Stock                    int    `json:"stock"`
			Sellable                 bool   `json:"sellable"`
			AvailableForClaim        bool   `json:"availableForClaim"`
			Barcode                  string `json:"barcode"`
			ItemNumber               int    `json:"itemNumber"`
			DiscountedPriceInfo      string `json:"discountedPriceInfo"`
			HasCollectable           bool   `json:"hasCollectable"`
			UnitInfo                 struct {
				UnitPrice     bool `json:"unitPrice"`
				UnitPriceText bool `json:"unitPriceText"`
			} `json:"unitInfo"`
			RushDeliveryMerchantListingExist bool `json:"rushDeliveryMerchantListingExist"`
			// LowerPriceMerchantListingExist   bool     `json:"lowerPriceMerchantListingExist"`
			// FastDeliveryOptions              []any    `json:"fastDeliveryOptions"`
			GroupTagIds []string `json:"groupTagIds"`
		} `json:"variants"`
		OtherMerchants []struct {
			URL          string `json:"url"`
			ReviewsURL   string `json:"reviewsUrl"`
			QuestionsURL string `json:"questionsUrl"`
			Promotions   []struct {
				PromotionRemainingTime string `json:"promotionRemainingTime"`
				IsOnlyAz               bool   `json:"isOnlyAz"`
				IsLimitSatisfied       bool   `json:"isLimitSatisfied"`
				Type                   int    `json:"type"`
				Text                   string `json:"text"`
				ID                     int    `json:"id"`
				PromotionDiscountType  string `json:"promotionDiscountType"`
				Link                   string `json:"link"`
			} `json:"promotions"`
			DiscountedPriceInfo string `json:"discountedPriceInfo"`
			IsSellable          bool   `json:"isSellable"`
			IsBasketDiscount    bool   `json:"isBasketDiscount"`
			HasStock            bool   `json:"hasStock"`
			Price               struct {
				ProfitMargin int `json:"profitMargin"`
				// DiscountedPrice struct {
				// 	Text  string  `json:"text"`
				// 	Value float64 `json:"value"`
				// } `json:"discountedPrice"`
				// SellingPrice struct {
				// 	Text  string  `json:"text"`
				// 	Value float64 `json:"value"`
				// } `json:"sellingPrice"`
				// OriginalPrice struct {
				// 	Text  string  `json:"text"`
				// 	Value float64 `json:"value"`
				// } `json:"originalPrice"`
				Currency string `json:"currency"`
			} `json:"price"`
			IsFreeCargo        bool `json:"isFreeCargo"`
			IsLongTermDelivery bool `json:"isLongTermDelivery"`
			Merchant           struct {
				IsSearchableMerchant            bool   `json:"isSearchableMerchant"`
				Stickers                        []any  `json:"stickers"`
				MerchantBadges                  []any  `json:"merchantBadges"`
				MerchantMarkers                 []any  `json:"merchantMarkers"`
				ID                              int    `json:"id"`
				Name                            string `json:"name"`
				OfficialName                    string `json:"officialName"`
				CityName                        string `json:"cityName"`
				CentralRegistrationSystemNumber string `json:"centralRegistrationSystemNumber"`
				RegisteredEmailAddress          string `json:"registeredEmailAddress"`
				TaxNumber                       string `json:"taxNumber"`
				// SellerScore                     float64 `json:"sellerScore"`
				// SellerScoreColor                string  `json:"sellerScoreColor"`
				DeliveryProviderName       string `json:"deliveryProviderName"`
				CorporateInvoiceApplicable bool   `json:"corporateInvoiceApplicable"`
				LocationBasedSales         bool   `json:"locationBasedSales"`
				BulkSalesLimit             int    `json:"bulkSalesLimit"`
			} `json:"merchant,omitempty"`
			DeliveryInformation struct {
				IsRushDelivery      bool   `json:"isRushDelivery"`
				DeliveryDate        string `json:"deliveryDate"`
				FastDeliveryOptions []any  `json:"fastDeliveryOptions"`
			} `json:"deliveryInformation"`
			CargoRemainingDays int  `json:"cargoRemainingDays"`
			IsBlacklist        bool `json:"isBlacklist"`
			// Merchant0          struct {
			// 	IsSearchableMerchant            bool   `json:"isSearchableMerchant"`
			// 	Stickers                        []any  `json:"stickers"`
			// 	MerchantBadges                  []any  `json:"merchantBadges"`
			// 	MerchantMarkers                 []any  `json:"merchantMarkers"`
			// 	ID                              int    `json:"id"`
			// 	Name                            string `json:"name"`
			// 	OfficialName                    string `json:"officialName"`
			// 	CityName                        string `json:"cityName"`
			// 	CentralRegistrationSystemNumber string `json:"centralRegistrationSystemNumber"`
			// 	RegisteredEmailAddress          string `json:"registeredEmailAddress"`
			// 	TaxNumber                       string `json:"taxNumber"`
			// 	SellerScore                     int    `json:"sellerScore"`
			// 	SellerScoreColor                string `json:"sellerScoreColor"`
			// 	DeliveryProviderName            string `json:"deliveryProviderName"`
			// 	CorporateInvoiceApplicable      bool   `json:"corporateInvoiceApplicable"`
			// 	LocationBasedSales              bool   `json:"locationBasedSales"`
			// } `json:"merchant,omitempty"`
			// Merchant1 struct {
			// 	IsSearchableMerchant            bool    `json:"isSearchableMerchant"`
			// 	Stickers                        []any   `json:"stickers"`
			// 	MerchantBadges                  []any   `json:"merchantBadges"`
			// 	MerchantMarkers                 []any   `json:"merchantMarkers"`
			// 	ID                              int     `json:"id"`
			// 	Name                            string  `json:"name"`
			// 	OfficialName                    string  `json:"officialName"`
			// 	CityName                        string  `json:"cityName"`
			// 	CentralRegistrationSystemNumber string  `json:"centralRegistrationSystemNumber"`
			// 	RegisteredEmailAddress          string  `json:"registeredEmailAddress"`
			// 	TaxNumber                       string  `json:"taxNumber"`
			// 	SellerScore                     float64 `json:"sellerScore"`
			// 	SellerScoreColor                string  `json:"sellerScoreColor"`
			// 	DeliveryProviderName            string  `json:"deliveryProviderName"`
			// 	CorporateInvoiceApplicable      bool    `json:"corporateInvoiceApplicable"`
			// 	LocationBasedSales              bool    `json:"locationBasedSales"`
			// } `json:"merchant,omitempty"`
		} `json:"otherMerchants"`
		Campaign struct {
			ID                 int    `json:"id"`
			Name               string `json:"name"`
			StartDate          string `json:"startDate"`
			EndDate            string `json:"endDate"`
			IsMultipleSupplied bool   `json:"isMultipleSupplied"`
			StockTypeID        int    `json:"stockTypeId"`
			URL                string `json:"url"`
			ShowTimer          bool   `json:"showTimer"`
		} `json:"campaign"`
		Category struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			Hierarchy      string `json:"hierarchy"`
			Refundable     bool   `json:"refundable"`
			BeautifiedName string `json:"beautifiedName"`
			IsVASEnabled   bool   `json:"isVASEnabled"`
		} `json:"category"`
		Brand struct {
			IsVirtual      bool   `json:"isVirtual"`
			BeautifiedName string `json:"beautifiedName"`
			ID             int    `json:"id"`
			Name           string `json:"name"`
			Path           string `json:"path"`
		} `json:"brand"`
		Color     string `json:"color"`
		MetaBrand struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			BeautifiedName string `json:"beautifiedName"`
			IsVirtual      bool   `json:"isVirtual"`
			Path           string `json:"path"`
		} `json:"metaBrand"`
		ShowVariants         bool  `json:"showVariants"`
		ShowSexualContent    bool  `json:"showSexualContent"`
		BrandCategoryBanners []any `json:"brandCategoryBanners"`
		AllVariants          []struct {
			ItemNumber int     `json:"itemNumber"`
			Value      string  `json:"value"`
			InStock    bool    `json:"inStock"`
			Currency   string  `json:"currency"`
			Barcode    string  `json:"barcode"`
			Price      float64 `json:"price"`
		} `json:"allVariants"`
		// OtherMerchantVariants                      []any  `json:"otherMerchantVariants"`
		InstallmentBanner                          string `json:"installmentBanner"`
		InstallmentText                            string `json:"installmentText"`
		Token                                      string `json:"token"`
		IsThereAnyCorporateInvoiceInOtherMerchants bool   `json:"isThereAnyCorporateInvoiceInOtherMerchants"`
		// AdvertProduct                              any    `json:"advertProduct"`
		// CategoryTopRankings                        any    `json:"categoryTopRankings"`
		IsVasEnabled     bool `json:"isVasEnabled"`
		OriginalCategory struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			Hierarchy      string `json:"hierarchy"`
			Refundable     bool   `json:"refundable"`
			BeautifiedName string `json:"beautifiedName"`
			IsVASEnabled   bool   `json:"isVASEnabled"`
		} `json:"originalCategory"`
		Landings            []any  `json:"landings"`
		ID                  int    `json:"id"`
		ProductCode         string `json:"productCode"`
		Name                string `json:"name"`
		NameWithProductCode string `json:"nameWithProductCode"`
		ContentDescriptions []struct {
			Description string `json:"description"`
			Bold        bool   `json:"bold"`
		} `json:"contentDescriptions"`
		Faq         []any `json:"faq"`
		Description []struct {
			Text     string `json:"text"`
			Priority int    `json:"priority"`
		} `json:"description"`
		ProductGroupID int    `json:"productGroupId"`
		Tax            int    `json:"tax"`
		BusinessUnit   string `json:"businessUnit"`
		BusinessUnitID int    `json:"businessUnitId"`
		MaxInstallment int    `json:"maxInstallment"`
		Gender         struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"gender"`
		WebGender struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"webGender"`
		URL              string   `json:"url"`
		Images           []string `json:"images"`
		IsSellable       bool     `json:"isSellable"`
		IsBasketDiscount bool     `json:"isBasketDiscount"`
		HasStock         bool     `json:"hasStock"`
		Price            struct {
			ProfitMargin    int `json:"profitMargin"`
			DiscountedPrice struct {
				Text  string  `json:"text"`
				Value float64 `json:"value"`
			} `json:"discountedPrice"`
			SellingPrice struct {
				Text  string  `json:"text"`
				Value float64 `json:"value"`
			} `json:"sellingPrice"`
			OriginalPrice struct {
				Text  string  `json:"text"`
				Value float64 `json:"value"`
			} `json:"originalPrice"`
			Currency string `json:"currency"`
		} `json:"price"`
		IsFreeCargo        bool `json:"isFreeCargo"`
		IsLongTermDelivery bool `json:"isLongTermDelivery"`
		Promotions         []struct {
			PromotionRemainingTime string `json:"promotionRemainingTime"`
			IsOnlyAz               bool   `json:"isOnlyAz"`
			IsLimitSatisfied       bool   `json:"isLimitSatisfied"`
			Type                   int    `json:"type"`
			Text                   string `json:"text"`
			ID                     int    `json:"id"`
			PromotionDiscountType  string `json:"promotionDiscountType"`
			Link                   string `json:"link"`
		} `json:"promotions"`
		Merchant struct {
			IsSearchableMerchant bool `json:"isSearchableMerchant"`
			// Stickers                        []any  `json:"stickers"`
			// MerchantBadges                  []any  `json:"merchantBadges"`
			// MerchantMarkers                 []any  `json:"merchantMarkers"`
			ID                              int    `json:"id"`
			Name                            string `json:"name"`
			OfficialName                    string `json:"officialName"`
			CityName                        string `json:"cityName"`
			CentralRegistrationSystemNumber string `json:"centralRegistrationSystemNumber"`
			RegisteredEmailAddress          string `json:"registeredEmailAddress"`
			TaxNumber                       string `json:"taxNumber"`
			// SellerScore                     float64 `json:"sellerScore"`
			// SellerScoreColor                string  `json:"sellerScoreColor"`
			DeliveryProviderName       string `json:"deliveryProviderName"`
			CorporateInvoiceApplicable bool   `json:"corporateInvoiceApplicable"`
			LocationBasedSales         bool   `json:"locationBasedSales"`
			SellerLink                 string `json:"sellerLink"`
		} `json:"merchant"`
		MerchantListings []struct {
			Beta              bool   `json:"beta"`
			DeliveryStartDate string `json:"deliveryStartDate"`
			DeliveryEndDate   string `json:"deliveryEndDate"`
			Merchant          struct {
				ID                              int    `json:"id"`
				Name                            string `json:"name"`
				OfficialName                    string `json:"officialName"`
				CityName                        string `json:"cityName"`
				CentralRegistrationSystemNumber string `json:"centralRegistrationSystemNumber"`
				RegisteredEmailAddress          string `json:"registeredEmailAddress"`
				TaxNumber                       string `json:"taxNumber"`
				// SellerScore                     float64 `json:"sellerScore"`
				// SellerScoreColor                string  `json:"sellerScoreColor"`
				DistrictName string `json:"districtName"`
				CountryName  string `json:"countryName"`
				Address      string `json:"address"`
				LogoURL      string `json:"logoUrl"`
				// Stickers                   []any  `json:"stickers"`
				// MerchantBadges             []any  `json:"merchantBadges"`
				// CorporateInvoiceApplicable bool   `json:"corporateInvoiceApplicable"`
				// Affiliate                  string `json:"affiliate"`
				// MerchantMarkers            []any  `json:"merchantMarkers"`
				// ComponentCodes             []any  `json:"componentCodes"`
				LocationBasedSales       bool `json:"locationBasedSales"`
				MpAgreedDTD              bool `json:"mpAgreedDTD"`
				MpEstimatedPackageDTD    bool `json:"mpEstimatedPackageDTD"`
				AllMpEstimatedPackageDTD bool `json:"allMpEstimatedPackageDTD"`
				TexApplicable            bool `json:"texApplicable"`
				DropShipmentSupplier     bool `json:"dropShipmentSupplier"`
			} `json:"merchant,omitempty"`
			Campaign struct {
				ID                 int    `json:"id"`
				Name               string `json:"name"`
				StockTypeID        int    `json:"stockTypeId"`
				StartDate          string `json:"startDate"`
				EndDate            string `json:"endDate"`
				Tags               []any  `json:"tags"`
				IsMultipleSupplied bool   `json:"isMultipleSupplied"`
			} `json:"campaign"`
			Promotions []struct {
				ID                        int    `json:"id"`
				Name                      string `json:"name"`
				DiscountType              int    `json:"discountType"`
				PromotionDiscountType     string `json:"promotionDiscountType"`
				PromotionEndDate          string `json:"promotionEndDate"`
				PromotionRemainingDays    int    `json:"promotionRemainingDays"`
				PromotionRemainingHours   int    `json:"promotionRemainingHours"`
				PromotionRemainingMinutes int    `json:"promotionRemainingMinutes"`
				IsLimitSatisfied          bool   `json:"isLimitSatisfied"`
				ShortName                 string `json:"shortName"`
				MocPerUser                int    `json:"mocPerUser"`
				IsOnlyAz                  bool   `json:"isOnlyAz"`
			} `json:"promotions"`
			CrossPromotions []any `json:"crossPromotions"`
			Variants        []struct {
				ListingID      string `json:"listingId"`
				FulfilmentType string `json:"fulfilmentType"`
				ItemNumber     int    `json:"itemNumber"`
				Barcode        string `json:"barcode"`
				Quantity       int    `json:"quantity"`
				Price          struct {
					// BuyingPrice              int     `json:"buyingPrice"`
					// SellingPrice             float64 `json:"sellingPrice"`
					OriginalPrice float64 `json:"originalPrice"`
					// ManipulatedOriginalPrice float64 `json:"manipulatedOriginalPrice"`
					// DiscountedPrice float64 `json:"discountedPrice"`
					// Currency        string  `json:"currency"`
					// ProfitMargin int `json:"profitMargin"`
				} `json:"price"`
				FreeCargo         bool `json:"freeCargo"`
				AvailableForClaim bool `json:"availableForClaim"`
				// DiscountedPriceInfo  string `json:"discountedPriceInfo"`
				HasCollectable       bool `json:"hasCollectable"`
				Deci                 int  `json:"deci"`
				RushDeliveryDuration int  `json:"rushDeliveryDuration"`
				VariantAttributes    []struct {
					AttributeID    int    `json:"attributeId"`
					AttributeName  string `json:"attributeName"`
					AttributeValue string `json:"attributeValue"`
					AttributeType  string `json:"attributeType"`
				} `json:"variantAttributes"`
				Scores []struct {
					Key        string  `json:"key"`
					Value      float64 `json:"value"`
					CampaignID int     `json:"campaignId"`
				} `json:"scores"`
				MaxProductSaleQuantity int `json:"maxProductSaleQuantity"`
				UnitInfo               struct {
				} `json:"unitInfo"`
				// SupplementaryServices            []any    `json:"supplementaryServices"`
				GroupTagIds                      []string `json:"groupTagIds"`
				RushDeliveryMerchantListingExist bool     `json:"rushDeliveryMerchantListingExist"`
				RushDeliveryMerchantListing      struct {
					Exist bool `json:"exist"`
				} `json:"rushDeliveryMerchantListing"`
				LowerPriceMerchantListingExist bool `json:"lowerPriceMerchantListingExist"`
				HasGroupDeal                   bool `json:"hasGroupDeal"`
				// FastDeliveryOptions            []any `json:"fastDeliveryOptions"`
				OverPriced          bool `json:"overPriced"`
				Sellable            bool `json:"sellable"`
				ScheduledDelivery   bool `json:"scheduledDelivery"`
				IsFlash             bool `json:"isFlash"`
				IsWinner            bool `json:"isWinner"`
				IsDiscountedListing bool `json:"isDiscountedListing"`
			} `json:"variants"`
			Stamps []struct {
				Type          string  `json:"type"`
				ImageURL      string  `json:"imageUrl"`
				Position      string  `json:"position"`
				AspectRatio   float64 `json:"aspectRatio"`
				Priority      int     `json:"priority"`
				PriceTagStamp bool    `json:"priceTagStamp"`
			} `json:"stamps"`
			// AlternativeVariants    []any  `json:"alternativeVariants"`
			MaxProductSaleQuantity int    `json:"maxProductSaleQuantity"`
			CargoStartDate         string `json:"cargoStartDate"`
			CargoRemainingDays     int    `json:"cargoRemainingDays"`
			AllVariants            []struct {
				ItemNumber int    `json:"itemNumber"`
				Value      string `json:"value"`
				InStock    bool   `json:"inStock"`
			} `json:"allVariants,omitempty"`
			OtherMerchantVariants []struct {
				ItemNumber int    `json:"itemNumber"`
				Value      string `json:"value"`
				InStock    bool   `json:"inStock"`
			} `json:"otherMerchantVariants,omitempty"`
			HasCheapestVariant bool `json:"hasCheapestVariant"`
			AgreedDeliveryDays int  `json:"agreedDeliveryDays"`
			// BuyMorePayLessPromotions []any `json:"buyMorePayLessPromotions"`
			CustomValues []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"customValues"`
			// SearchableTags []any `json:"searchableTags"`
			FreeCargo bool `json:"freeCargo"`
			IsWinner  bool `json:"isWinner"`
			// Merchant0      struct {
			// 	ID                              int     `json:"id"`
			// 	Name                            string  `json:"name"`
			// 	OfficialName                    string  `json:"officialName"`
			// 	CityName                        string  `json:"cityName"`
			// 	CentralRegistrationSystemNumber string  `json:"centralRegistrationSystemNumber"`
			// 	RegisteredEmailAddress          string  `json:"registeredEmailAddress"`
			// 	TaxNumber                       string  `json:"taxNumber"`
			// 	SellerScore                     float64 `json:"sellerScore"`
			// 	SellerScoreColor                string  `json:"sellerScoreColor"`
			// 	DistrictName                    string  `json:"districtName"`
			// 	CountryName                     string  `json:"countryName"`
			// 	Address                         string  `json:"address"`
			// 	LogoURL                         string  `json:"logoUrl"`
			// 	Stickers                        []any   `json:"stickers"`
			// 	MerchantBadges                  []any   `json:"merchantBadges"`
			// 	CorporateInvoiceApplicable      bool    `json:"corporateInvoiceApplicable"`
			// 	Affiliate                       string  `json:"affiliate"`
			// 	MerchantMarkers                 []any   `json:"merchantMarkers"`
			// 	ComponentCodes                  []any   `json:"componentCodes"`
			// 	LocationBasedSales              bool    `json:"locationBasedSales"`
			// 	MpAgreedDTD                     bool    `json:"mpAgreedDTD"`
			// 	MpEstimatedPackageDTD           bool    `json:"mpEstimatedPackageDTD"`
			// 	AllMpEstimatedPackageDTD        bool    `json:"allMpEstimatedPackageDTD"`
			// 	BulkSalesLimit                  int     `json:"bulkSalesLimit"`
			// 	TexApplicable                   bool    `json:"texApplicable"`
			// 	DropShipmentSupplier            bool    `json:"dropShipmentSupplier"`
			// } `json:"merchant,omitempty"`
			// Merchant1 struct {
			// 	ID                              int    `json:"id"`
			// 	Name                            string `json:"name"`
			// 	OfficialName                    string `json:"officialName"`
			// 	CityName                        string `json:"cityName"`
			// 	CentralRegistrationSystemNumber string `json:"centralRegistrationSystemNumber"`
			// 	RegisteredEmailAddress          string `json:"registeredEmailAddress"`
			// 	TaxNumber                       string `json:"taxNumber"`
			// 	SellerScore                     int    `json:"sellerScore"`
			// 	SellerScoreColor                string `json:"sellerScoreColor"`
			// 	DistrictName                    string `json:"districtName"`
			// 	CountryName                     string `json:"countryName"`
			// 	Address                         string `json:"address"`
			// 	LogoURL                         string `json:"logoUrl"`
			// 	Stickers                        []any  `json:"stickers"`
			// 	MerchantBadges                  []any  `json:"merchantBadges"`
			// 	CorporateInvoiceApplicable      bool   `json:"corporateInvoiceApplicable"`
			// 	Affiliate                       string `json:"affiliate"`
			// 	MerchantMarkers                 []any  `json:"merchantMarkers"`
			// 	ComponentCodes                  []any  `json:"componentCodes"`
			// 	LocationBasedSales              bool   `json:"locationBasedSales"`
			// 	MpAgreedDTD                     bool   `json:"mpAgreedDTD"`
			// 	MpEstimatedPackageDTD           bool   `json:"mpEstimatedPackageDTD"`
			// 	AllMpEstimatedPackageDTD        bool   `json:"allMpEstimatedPackageDTD"`
			// 	BulkSalesLimit                  int    `json:"bulkSalesLimit"`
			// 	TexApplicable                   bool   `json:"texApplicable"`
			// 	DropShipmentSupplier            bool   `json:"dropShipmentSupplier"`
			// } `json:"merchant,omitempty"`
		} `json:"merchantListings"`
		DeliveryInformation struct {
			IsRushDelivery bool   `json:"isRushDelivery"`
			DeliveryDate   string `json:"deliveryDate"`
			// FastDeliveryOptions []any  `json:"fastDeliveryOptions"`
		} `json:"deliveryInformation"`
		CargoRemainingDays int  `json:"cargoRemainingDays"`
		IsMarketplace      bool `json:"isMarketplace"`
		ProductStamps      []struct {
			Type          string  `json:"type"`
			ImageURL      string  `json:"imageUrl"`
			Position      string  `json:"position"`
			AspectRatio   float64 `json:"aspectRatio"`
			Priority      int     `json:"priority"`
			PriceTagStamp bool    `json:"priceTagStamp"`
		} `json:"productStamps"`
		HasHTMLContent    bool   `json:"hasHtmlContent"`
		FavoriteCount     int    `json:"favoriteCount"`
		UxLayout          string `json:"uxLayout"`
		IsDigitalGood     bool   `json:"isDigitalGood"`
		IsRunningOut      bool   `json:"isRunningOut"`
		ScheduledDelivery bool   `json:"scheduledDelivery"`
		// RatingScore       struct {
		// 	AverageRating     int `json:"averageRating"`
		// 	TotalRatingCount  int `json:"totalRatingCount"`
		// 	TotalCommentCount int `json:"totalCommentCount"`
		// } `json:"ratingScore"`
		ShowStarredAttributes    bool   `json:"showStarredAttributes"`
		ReviewsURL               string `json:"reviewsUrl"`
		QuestionsURL             string `json:"questionsUrl"`
		SellerQuestionEnabled    bool   `json:"sellerQuestionEnabled"`
		SizeChartURL             string `json:"sizeChartUrl"`
		SizeExpectationAvailable bool   `json:"sizeExpectationAvailable"`
		CrossPromotionAward      struct {
			AwardType  any `json:"awardType"`
			AwardValue any `json:"awardValue"`
			ContentID  int `json:"contentId"`
			MerchantID int `json:"merchantId"`
		} `json:"crossPromotionAward"`
		RushDeliveryMerchantListingExist bool `json:"rushDeliveryMerchantListingExist"`
		IsRushDelivery                   bool `json:"isRushDelivery"`
		LowerPriceMerchantListingExist   bool `json:"lowerPriceMerchantListingExist"`
		ShowValidFlashSales              bool `json:"showValidFlashSales"`
		ShowExpiredFlashSales            bool `json:"showExpiredFlashSales"`
		WalletRebate                     struct {
			MinPrice    int     `json:"minPrice"`
			MaxPrice    int     `json:"maxPrice"`
			RebateRatio float64 `json:"rebateRatio"`
		} `json:"walletRebate"`
		IsArtWork bool `json:"isArtWork"`
		// BuyMorePayLessPromotions []any `json:"buyMorePayLessPromotions"`
		FilterableLabels []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			VisibleAgg  bool   `json:"visibleAgg"`
		} `json:"filterableLabels"`
		CustomValues []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"customValues"`
	} `json:"result"`
	Headers struct {
		Tysidecarcachable string `json:"tysidecarcachable"`
	} `json:"headers"`
}
