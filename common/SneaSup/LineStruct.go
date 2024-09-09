package sneaksup

type LineStruct struct {
	Pager struct {
		TotalPages int `json:"TotalPages"`
		TotalItems int `json:"TotalItems"`
		PageIndex  int `json:"PageIndex"`
	} `json:"Pager"`
	Products         []Product `json:"Products"`
	ID               int       `json:"Id"`
	CustomProperties struct {
	} `json:"CustomProperties"`
}

// Структура товара по версии сервера SneakSup
type Product struct {
	MYLINK string

	ID           int    `json:"Id"`
	ProductIndex int    `json:"ProductIndex"`
	Name         string `json:"Name"`
	Pictures     []struct {
		ImageURL            string      `json:"ImageUrl"`
		SecondImageURL      interface{} `json:"SecondImageUrl"`
		FullSizeImageURL    string      `json:"FullSizeImageUrl"`
		Title               string      `json:"Title"`
		AlternateText       string      `json:"AlternateText"`
		OriginalImageURL    string      `json:"OriginalImageUrl"`
		ThumbImageURL       interface{} `json:"ThumbImageUrl"`
		PictureID           int         `json:"PictureId"`
		IntegrationFileName string      `json:"IntegrationFileName"`
		CustomProperties    struct {
		} `json:"CustomProperties"`
	} `json:"Pictures"`
	ShortDescription string `json:"ShortDescription"`
	ProductPrice     struct {
		OldPrice                             interface{} `json:"OldPrice"`
		Price                                string      `json:"Price"`
		PriceValue                           float64     `json:"PriceValue"`
		OldPriceValue                        float64     `json:"OldPriceValue"`
		PriceWithoutDiscount                 string      `json:"PriceWithoutDiscount"`
		PriceWithoutDiscountValue            float64     `json:"PriceWithoutDiscountValue"`
		DisableBuyButton                     bool        `json:"DisableBuyButton"`
		DisableWishlistButton                bool        `json:"DisableWishlistButton"`
		DisableAddToCompareListButton        bool        `json:"DisableAddToCompareListButton"`
		AvailableForPreOrder                 bool        `json:"AvailableForPreOrder"`
		PreOrderAvailabilityStartDateTimeUtc interface{} `json:"PreOrderAvailabilityStartDateTimeUtc"`
		IsRental                             bool        `json:"IsRental"`
		ForceRedirectionAfterAddingToCart    bool        `json:"ForceRedirectionAfterAddingToCart"`
		DisplayTaxShippingInfo               bool        `json:"DisplayTaxShippingInfo"`
		ProductDiscountPercentage            string      `json:"ProductDiscountPercentage"`
		TaxDisplayType                       int         `json:"TaxDisplayType"`
		CustomProperties                     struct {
		} `json:"CustomProperties"`
	} `json:"ProductPrice"`
	DefaultPictureModel struct {
		ImageURL            string      `json:"ImageUrl"`
		SecondImageURL      interface{} `json:"SecondImageUrl"`
		FullSizeImageURL    string      `json:"FullSizeImageUrl"`
		Title               string      `json:"Title"`
		AlternateText       string      `json:"AlternateText"`
		OriginalImageURL    string      `json:"OriginalImageUrl"`
		ThumbImageURL       interface{} `json:"ThumbImageUrl"`
		PictureID           int         `json:"PictureId"`
		IntegrationFileName interface{} `json:"IntegrationFileName"`
		CustomProperties    struct {
		} `json:"CustomProperties"`
	} `json:"DefaultPictureModel"`
	URL                           string   `json:"Url"`
	AddToCartLink                 string   `json:"AddToCartLink"`
	CategoryName                  string   `json:"CategoryName"`
	ProductCategoryBreadcrumbList []string `json:"ProductCategoryBreadcrumbList"`
	ManufacturerName              string   `json:"ManufacturerName"`
	BrandName                     string   `json:"BrandName"`
	ProductAttributes             []struct {
		ID                     int         `json:"Id"`
		ProductID              int         `json:"ProductId"`
		ProductAttributeID     int         `json:"ProductAttributeId"`
		DisplayOrder           int         `json:"DisplayOrder"`
		DefaultValue           interface{} `json:"DefaultValue"`
		ConditionAttributeXML  interface{} `json:"ConditionAttributeXml"`
		TextPrompt             string      `json:"TextPrompt"`
		ProductAttributeValues []struct {
			ID         int    `json:"Id"`
			Name       string `json:"Name"`
			Barcode    string `json:"Barcode"`
			VariantSku string `json:"VariantSku"`
			Quantity   int    `json:"Quantity"`
			InStock    bool   `json:"InStock"`
		} `json:"ProductAttributeValues"`
	} `json:"ProductAttributes"`
	Manufacturers []struct {
		Name              string `json:"Name"`
		Description       string `json:"Description"`
		MetaKeywords      string `json:"MetaKeywords"`
		MetaDescription   string `json:"MetaDescription"`
		MetaTitle         string `json:"MetaTitle"`
		SeName            string `json:"SeName"`
		FooterDescription string `json:"FooterDescription"`
		PictureModel      struct {
			ImageURL            string      `json:"ImageUrl"`
			SecondImageURL      interface{} `json:"SecondImageUrl"`
			FullSizeImageURL    string      `json:"FullSizeImageUrl"`
			Title               string      `json:"Title"`
			AlternateText       string      `json:"AlternateText"`
			OriginalImageURL    string      `json:"OriginalImageUrl"`
			ThumbImageURL       interface{} `json:"ThumbImageUrl"`
			PictureID           int         `json:"PictureId"`
			IntegrationFileName interface{} `json:"IntegrationFileName"`
			CustomProperties    struct {
			} `json:"CustomProperties"`
		} `json:"PictureModel"`
		PagingFilteringContext struct {
			PriceRangeFilter struct {
				Enabled          bool          `json:"Enabled"`
				Items            []interface{} `json:"Items"`
				RemoveFilterURL  interface{}   `json:"RemoveFilterUrl"`
				MinPrice         interface{}   `json:"MinPrice"`
				MaxPrice         interface{}   `json:"MaxPrice"`
				CustomProperties struct {
				} `json:"CustomProperties"`
				WebHelper struct {
					IsRequestBeingRedirected bool `json:"IsRequestBeingRedirected"`
					IsPostBeingDone          bool `json:"IsPostBeingDone"`
				} `json:"WebHelper"`
			} `json:"PriceRangeFilter"`
			SpecificationFilter struct {
				Enabled                 bool          `json:"Enabled"`
				AlreadyFilteredItems    []interface{} `json:"AlreadyFilteredItems"`
				NotFilteredItems        []interface{} `json:"NotFilteredItems"`
				RemoveFilterURL         interface{}   `json:"RemoveFilterUrl"`
				ShowProductCount        bool          `json:"ShowProductCount"`
				SpecOptionDisplayTypeID int           `json:"SpecOptionDisplayTypeId"`
				CustomProperties        struct {
				} `json:"CustomProperties"`
				WebHelper struct {
					IsRequestBeingRedirected bool `json:"IsRequestBeingRedirected"`
					IsPostBeingDone          bool `json:"IsPostBeingDone"`
				} `json:"WebHelper"`
			} `json:"SpecificationFilter"`
			AllowProductSorting            bool          `json:"AllowProductSorting"`
			AvailableSortOptions           []interface{} `json:"AvailableSortOptions"`
			AllowProductViewModeChanging   bool          `json:"AllowProductViewModeChanging"`
			AvailableViewModes             []interface{} `json:"AvailableViewModes"`
			AllowCustomersToSelectPageSize bool          `json:"AllowCustomersToSelectPageSize"`
			PageSizeOptions                []interface{} `json:"PageSizeOptions"`
			OrderBy                        int           `json:"OrderBy"`
			ViewMode                       interface{}   `json:"ViewMode"`
			InStock                        bool          `json:"InStock"`
			FirstItem                      int           `json:"FirstItem"`
			HasNextPage                    bool          `json:"HasNextPage"`
			HasPreviousPage                bool          `json:"HasPreviousPage"`
			LastItem                       int           `json:"LastItem"`
			PageIndex                      int           `json:"PageIndex"`
			PageNumber                     int           `json:"PageNumber"`
			PageSize                       int           `json:"PageSize"`
			TotalItems                     int           `json:"TotalItems"`
			TotalPages                     int           `json:"TotalPages"`
			PaginationType                 int           `json:"PaginationType"`
		} `json:"PagingFilteringContext"`
		FeaturedProducts []interface{} `json:"FeaturedProducts"`
		Products         []interface{} `json:"Products"`
		ID               int           `json:"Id"`
		CustomProperties struct {
		} `json:"CustomProperties"`
	} `json:"Manufacturers"`
	SpecificationAttributeModels []struct {
		SpecificationAttributeID        int         `json:"SpecificationAttributeId"`
		SpecificationAttributeGroupName interface{} `json:"SpecificationAttributeGroupName"`
		SpecificationAttributeName      string      `json:"SpecificationAttributeName"`
		ValueRaw                        string      `json:"ValueRaw"`
		ShowOnFilterSpecs               bool        `json:"ShowOnFilterSpecs"`
		ShowOnProductPage               bool        `json:"ShowOnProductPage"`
		ShowAsBadge                     bool        `json:"ShowAsBadge"`
		BadgePosition                   interface{} `json:"BadgePosition"`
		BadgeCSSClass                   interface{} `json:"BadgeCssClass"`
		BadgeText                       string      `json:"BadgeText"`
		BadgeIcon                       interface{} `json:"BadgeIcon"`
		ErpCode                         string      `json:"ErpCode"`
		ErpName                         string      `json:"ErpName"`
		OptionErpCode                   string      `json:"OptionErpCode"`
		OptionCustomValue               interface{} `json:"OptionCustomValue"`
		Name                            string      `json:"Name"`
		ErpTypeCode                     interface{} `json:"ErpTypeCode"`
		IsPeriodicDisplay               bool        `json:"IsPeriodicDisplay"`
		StartDate                       string      `json:"StartDate"`
		EndDate                         string      `json:"EndDate"`
		OptionID                        int         `json:"OptionId"`
		CustomProperties                struct {
			ColorNames string `json:"ColorNames"`
		} `json:"CustomProperties,omitempty"`
		// CustomProperties0 struct {
		// 	ColorNames   string `json:"ColorNames"`
		// 	ColorHexCode string `json:"ColorHexCode"`
		// } `json:"CustomProperties,omitempty"`
		// CustomProperties1 struct {
		// 	NewProductBadge string `json:"NewProductBadge"`
		// } `json:"CustomProperties,omitempty"`
	} `json:"SpecificationAttributeModels"`
	HasStock                                           bool   `json:"HasStock"`
	Sku                                                string `json:"Sku"`
	ProductType                                        int    `json:"ProductType"`
	StockQuantity                                      int    `json:"StockQuantity"`
	AvailableStartDateTimeUtcString                    string `json:"AvailableStartDateTimeUtcString"`
	AvailableStartDateTimeUtcHasValue                  bool   `json:"AvailableStartDateTimeUtcHasValue"`
	AvailableStartDateTimeUtcLessThenDateTimeUtcNow    bool   `json:"AvailableStartDateTimeUtcLessThenDateTimeUtcNow"`
	AvailableStartDateTimeUtcGreaterThenDateTimeUtcNow bool   `json:"AvailableStartDateTimeUtcGreaterThenDateTimeUtcNow"`
	Siblings                                           []struct {
		ID                   int           `json:"Id"`
		ColorCustomNameAttrs []interface{} `json:"ColorCustomNameAttrs"`
		ColorName            string        `json:"ColorName"`
		DefaultPictureModel  struct {
			ImageURL            string      `json:"ImageUrl"`
			SecondImageURL      interface{} `json:"SecondImageUrl"`
			FullSizeImageURL    string      `json:"FullSizeImageUrl"`
			Title               string      `json:"Title"`
			AlternateText       string      `json:"AlternateText"`
			OriginalImageURL    string      `json:"OriginalImageUrl"`
			ThumbImageURL       interface{} `json:"ThumbImageUrl"`
			PictureID           int         `json:"PictureId"`
			IntegrationFileName interface{} `json:"IntegrationFileName"`
			CustomProperties    struct {
			} `json:"CustomProperties"`
		} `json:"DefaultPictureModel"`
		SeName       string `json:"SeName"`
		URL          string `json:"Url"`
		ProductPrice struct {
			OldPrice                             interface{} `json:"OldPrice"`
			Price                                string      `json:"Price"`
			PriceValue                           float64     `json:"PriceValue"`
			OldPriceValue                        float64     `json:"OldPriceValue"`
			PriceWithoutDiscount                 string      `json:"PriceWithoutDiscount"`
			PriceWithoutDiscountValue            float64     `json:"PriceWithoutDiscountValue"`
			DisableBuyButton                     bool        `json:"DisableBuyButton"`
			DisableWishlistButton                bool        `json:"DisableWishlistButton"`
			DisableAddToCompareListButton        bool        `json:"DisableAddToCompareListButton"`
			AvailableForPreOrder                 bool        `json:"AvailableForPreOrder"`
			PreOrderAvailabilityStartDateTimeUtc interface{} `json:"PreOrderAvailabilityStartDateTimeUtc"`
			IsRental                             bool        `json:"IsRental"`
			ForceRedirectionAfterAddingToCart    bool        `json:"ForceRedirectionAfterAddingToCart"`
			DisplayTaxShippingInfo               bool        `json:"DisplayTaxShippingInfo"`
			ProductDiscountPercentage            interface{} `json:"ProductDiscountPercentage"`
			TaxDisplayType                       int         `json:"TaxDisplayType"`
			CustomProperties                     struct {
			} `json:"CustomProperties"`
		} `json:"ProductPrice"`
		Pictures []struct {
			ImageURL            string      `json:"ImageUrl"`
			SecondImageURL      interface{} `json:"SecondImageUrl"`
			FullSizeImageURL    string      `json:"FullSizeImageUrl"`
			Title               string      `json:"Title"`
			AlternateText       string      `json:"AlternateText"`
			OriginalImageURL    string      `json:"OriginalImageUrl"`
			ThumbImageURL       interface{} `json:"ThumbImageUrl"`
			PictureID           int         `json:"PictureId"`
			IntegrationFileName string      `json:"IntegrationFileName"`
			CustomProperties    struct {
			} `json:"CustomProperties"`
		} `json:"Pictures"`
	} `json:"Siblings"`
}
