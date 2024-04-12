package zaratr

import (
	"time"

	"github.com/RB-PRO/ClikShop/pkg/bases"
)

// Ссылка на товар, [пример].
//
// [пример]: https://www.zara.com/tr/en/ribbed-strappy-vest-top-p03253306.html?ajax=true
const TouchURL string = "https://www.zara.com/tr/en/%v.html?ajax=true"

type Touch struct {
	Location string `json:"location,omitempty"` // Новая ссылка на товар

	// массив категорий
	Cat []bases.Cat `json:"-"`

	NoIndex bool `json:"noIndex"`
	MkSpots struct {
		ESpotCommonStyles struct {
			Key     string `json:"key"`
			Content struct {
				Content string `json:"content"`
			} `json:"content"`
		} `json:"ESpot_CommonStyles"`
		ESpotCopyright struct {
			Key     string `json:"key"`
			Content struct {
				Content string `json:"content"`
			} `json:"content"`
		} `json:"ESpot_Copyright"`
		ESpotFooterLinks struct {
			Type    string `json:"type"`
			Content struct {
				Content string `json:"content"`
			} `json:"content"`
			Key string `json:"key"`
		} `json:"ESpot_Footer_Links"`
		ESpotVirtualGiftCardPreview             any `json:"ESpot_VirtualGiftCard_Preview"`
		ESpotProductPageSpecialReturnConditions struct {
			Type    string `json:"type"`
			Content struct {
				Content string `json:"content"`
			} `json:"content"`
			Key string `json:"key"`
		} `json:"ESpot_ProductPage_SpecialReturnConditions"`
		ESpotSocialNetworkFooter struct {
			Type    string `json:"type"`
			Content struct {
				Content string `json:"content"`
			} `json:"content"`
			Key string `json:"key"`
		} `json:"ESpot_SocialNetwork_Footer"`
	} `json:"mkSpots"`
	Product struct {
		ID    CustomIntToString `json:"id"`
		Type  string            `json:"type"`
		Kind  string            `json:"kind"`
		State string            `json:"state"`
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
					Availability     string `json:"availability"`
					EquivalentSizeID int    `json:"equivalentSizeId"`
					ID               int    `json:"id"`
					Name             string `json:"name"`
					Price            int    `json:"price"`
					Reference        string `json:"reference"`
					Sku              int    `json:"sku"`
					Attributes       []struct {
						Type       string `json:"type"`
						Identifier string `json:"identifier"`
						Values     []any  `json:"values"`
					} `json:"attributes"`
					Demand string `json:"demand"`
				} `json:"sizes"`
				Description    string `json:"description"`
				RawDescription string `json:"rawDescription"`
				ExtraInfo      struct {
					IsStockInStoresAvailable bool `json:"isStockInStoresAvailable"`
					HighlightPrice           bool `json:"highlightPrice"`
				} `json:"extraInfo,omitempty"`
				TagTypes []struct {
					DisplayName string `json:"displayName"`
					Type        string `json:"type"`
				} `json:"tagTypes"`
				Attributes []any `json:"attributes"`
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
				// ExtraInfo0       struct {
				// 	IsStockInStoresAvailable bool   `json:"isStockInStoresAvailable"`
				// 	FitSizeMessage           string `json:"fitSizeMessage"`
				// 	HighlightPrice           bool   `json:"highlightPrice"`
				// } `json:"extraInfo,omitempty"`
				// ExtraInfo1 struct {
				// 	IsStockInStoresAvailable bool   `json:"isStockInStoresAvailable"`
				// 	FitSizeMessage           string `json:"fitSizeMessage"`
				// 	HighlightPrice           bool   `json:"highlightPrice"`
				// } `json:"extraInfo,omitempty"`
				// ExtraInfo2 struct {
				// 	IsStockInStoresAvailable bool   `json:"isStockInStoresAvailable"`
				// 	FitSizeMessage           string `json:"fitSizeMessage"`
				// 	HighlightPrice           bool   `json:"highlightPrice"`
				// } `json:"extraInfo,omitempty"`
			} `json:"colors"`
			ColorSelectorLabel string `json:"colorSelectorLabel"`
			MultipleColorLabel string `json:"multipleColorLabel"`
			RelatedProducts    []any  `json:"relatedProducts"`
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
		FirstVisibleDate time.Time `json:"firstVisibleDate"`
		Attributes       []any     `json:"attributes"`
		SizeGuide        struct {
			Enabled bool `json:"enabled"`
		} `json:"sizeGuide"`
		Xmedia                  []any `json:"xmedia"`
		ProductTag              []any `json:"productTag"`
		HasInteractiveSizeGuide bool  `json:"hasInteractiveSizeGuide"`
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
	Categories  []any  `json:"categories"`
	BackURL     string `json:"backUrl"`
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
	BreadCrumbs []struct {
		Text          string `json:"text"`
		Keyword       string `json:"keyword,omitempty"`
		ID            int    `json:"id"`
		SeoCategoryID int    `json:"seoCategoryId,omitempty"`
		Layout        string `json:"layout,omitempty"`
	} `json:"breadCrumbs"`
	AnalyticsData struct {
		AppVersion string `json:"appVersion"`
		PageType   string `json:"pageType"`
		Page       struct {
			Language string `json:"language"`
			Shop     string `json:"shop"`
			Currency string `json:"currency"`
		} `json:"page"`
		TrackerUA         string  `json:"trackerUA"`
		AnonymizeIP       string  `json:"anonymizeIp"`
		Hostname          string  `json:"hostname"`
		CategoryName      string  `json:"categoryName"`
		ColorCode         string  `json:"colorCode"`
		MainPrice         float64 `json:"mainPrice"`
		ColorRef          string  `json:"colorRef"`
		ProductID         int     `json:"productId"`
		ProductRef        string  `json:"productRef"`
		ProductName       string  `json:"productName"`
		Section           string  `json:"section"`
		StylingID         string  `json:"stylingId"`
		Family            string  `json:"family"`
		Subfamily         string  `json:"subfamily"`
		CatentryID        int     `json:"catentryId"`
		LowOnStockProduct bool    `json:"lowOnStockProduct"`
		Brand             int     `json:"brand"`
	} `json:"analyticsData"`
	MobileApp struct {
		Msg        string `json:"msg"`
		IOSURI     string `json:"iOSUri"`
		AndroidURI string `json:"androidUri"`
	} `json:"mobileApp"`
	IsSharedProduct          bool `json:"isSharedProduct"`
	GiftCardExpirationMonths int  `json:"giftCardExpirationMonths"`
	Sections                 []struct {
		ID             int      `json:"id"`
		Name           string   `json:"name"`
		Description    string   `json:"description"`
		AvailableFor   []string `json:"availableFor"`
		EngDescription string   `json:"engDescription"`
	} `json:"sections"`
	IsRgpdEnabled         bool   `json:"isRgpdEnabled"`
	ShowSizeGuideInfoLink bool   `json:"showSizeGuideInfoLink"`
	Workgroups            []any  `json:"workgroups"`
	ChatView              any    `json:"chatView"`
	ViewName              string `json:"viewName"`
	UserKind              string `json:"userKind"`
	ClientAppConfig       struct {
		Version           string `json:"version"`
		LangID            int    `json:"langId"`
		StoreID           int    `json:"storeId"`
		LangCode          string `json:"langCode"`
		StoreCode         string `json:"storeCode"`
		StoreCountryCode  string `json:"storeCountryCode"`
		Locale            string `json:"locale"`
		AppAssetsBasePath string `json:"appAssetsBasePath"`
		AppLinks          struct {
			Ios     string `json:"ios"`
			Android string `json:"android"`
		} `json:"appLinks"`
		UniversalLinks struct {
			Ios     string `json:"ios"`
			Android string `json:"android"`
		} `json:"universalLinks"`
		ClientSideNavigationTimeout int `json:"clientSideNavigationTimeout"`
		I18NConfig                  struct {
			CacheEnabled   bool   `json:"cacheEnabled"`
			Version        int    `json:"version"`
			URL            string `json:"url"`
			DefaultMessage string `json:"defaultMessage"`
		} `json:"i18nConfig"`
		OriginalURL  string `json:"originalUrl"`
		ImageBaseURL string `json:"imageBaseUrl"`
		VideoBaseURL string `json:"videoBaseUrl"`
		Domains      struct {
			Desktop struct {
				Dynamic struct {
					Base string `json:"base"`
					Cn   string `json:"cn"`
					Xn   string `json:"xn"`
				} `json:"dynamic"`
				Static struct {
					Base string `json:"base"`
					Cn   string `json:"cn"`
					Xn   string `json:"xn"`
				} `json:"static"`
				Ports struct {
					Plain int `json:"plain"`
					Ssl   int `json:"ssl"`
				} `json:"ports"`
			} `json:"desktop"`
		} `json:"domains"`
		IsSsl       bool `json:"isSsl"`
		ServerPorts struct {
			Plain int `json:"plain"`
			Ssl   int `json:"ssl"`
		} `json:"serverPorts"`
		FormatterConfig struct {
			Currency           string  `json:"currency"`
			Symbol             string  `json:"symbol"`
			CurrencyFormat     string  `json:"currencyFormat"`
			CurrencyDecimals   int     `json:"currencyDecimals"`
			CurrencyCode       string  `json:"currencyCode"`
			CurrencySymbol     string  `json:"currencySymbol"`
			CurrencyRateToEuro float64 `json:"currencyRateToEuro"`
			Formats            struct {
				Number struct {
					DecimalSeparator   string `json:"decimalSeparator"`
					ThousandsSeparator string `json:"thousandsSeparator"`
				} `json:"number"`
				Date struct {
					ShortDate string `json:"shortDate"`
					LongDate  string `json:"longDate"`
				} `json:"date"`
			} `json:"formats"`
		} `json:"formatterConfig"`
		FormatterConfigByCur struct {
			Try struct {
				Currency           string  `json:"currency"`
				Symbol             string  `json:"symbol"`
				CurrencyFormat     string  `json:"currencyFormat"`
				CurrencyDecimals   int     `json:"currencyDecimals"`
				CurrencyCode       string  `json:"currencyCode"`
				CurrencySymbol     string  `json:"currencySymbol"`
				CurrencyRateToEuro float64 `json:"currencyRateToEuro"`
				Formats            struct {
					Number struct {
						DecimalSeparator   string `json:"decimalSeparator"`
						ThousandsSeparator string `json:"thousandsSeparator"`
					} `json:"number"`
					Date struct {
						ShortDate string `json:"shortDate"`
						LongDate  string `json:"longDate"`
					} `json:"date"`
				} `json:"formats"`
			} `json:"TRY"`
		} `json:"formatterConfigByCur"`
		XmediaNetworkPerformance struct {
			Dpr       int  `json:"dpr"`
			Interval  int  `json:"interval"`
			Enabled   bool `json:"enabled"`
			Threshold struct {
				Min float64 `json:"min"`
			} `json:"threshold"`
			FileURL string `json:"fileUrl"`
		} `json:"xmediaNetworkPerformance"`
		XmediaFormats []struct {
			Datatype      string `json:"datatype"`
			ID            int    `json:"id"`
			Set           int    `json:"set"`
			Type          string `json:"type"`
			Name          string `json:"name"`
			Description   string `json:"description"`
			Codecs        string `json:"codecs"`
			Extension     string `json:"extension"`
			Width         int    `json:"width"`
			IsSeoRelevant bool   `json:"isSeoRelevant"`
		} `json:"xmediaFormats"`
		XmediaTransformations []struct {
			Meta struct {
				Xmedia []struct {
					AllowedScreens []string `json:"allowedScreens"`
					Datatype       string   `json:"datatype"`
					Height         int      `json:"height"`
					Kind           string   `json:"kind"`
					Name           string   `json:"name"`
					Path           string   `json:"path"`
					Set            int      `json:"set"`
					Timestamp      string   `json:"timestamp"`
					Type           string   `json:"type"`
					Width          int      `json:"width"`
				} `json:"xmedia"`
			} `json:"meta"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"xmediaTransformations"`
		IsOpenProductPageInNewTab bool `json:"isOpenProductPageInNewTab"`
		Store                     struct {
			ID                         int    `json:"id"`
			CatalogID                  int    `json:"catalogId"`
			SearchLang                 string `json:"searchLang"`
			ColbensonURL               string `json:"colbensonUrl"`
			IsOpenForSale              bool   `json:"isOpenForSale"`
			ShowPrivacityPolicyCookie  bool   `json:"showPrivacityPolicyCookie"`
			ShowAdCookiesPreferences   bool   `json:"showAdCookiesPreferences"`
			IsDataConsentPolicyEnabled bool   `json:"isDataConsentPolicyEnabled"`
			PrivacyURL                 string `json:"privacyUrl"`
			Privacy                    struct {
				URL     string `json:"url"`
				Version string `json:"version"`
			} `json:"privacy"`
			EguiData struct {
				AssignTitleBaseURL          string `json:"assignTitleBaseUrl"`
				DonationEntitiesListBaseURL string `json:"donationEntitiesListBaseUrl"`
				IsEnabled                   bool   `json:"isEnabled"`
			} `json:"eguiData"`
			WarehouseID                 int    `json:"warehouseId"`
			PhoneCountryCode            string `json:"phoneCountryCode"`
			IsStockInStoresAvailable    bool   `json:"isStockInStoresAvailable"`
			IsElectronicInvoiceActive   bool   `json:"isElectronicInvoiceActive"`
			AvailableStoreReturnProcess bool   `json:"availableStoreReturnProcess"`
			Eula                        struct {
				URL string `json:"url"`
			} `json:"eula"`
			CountryCode         string `json:"countryCode"`
			ShouldShowState     bool   `json:"shouldShowState"`
			CountryName         string `json:"countryName"`
			RelatedStores       []any  `json:"relatedStores"`
			LinkToRelatedStores bool   `json:"linkToRelatedStores"`
			SupportedLanguages  []struct {
				ID                     int    `json:"id"`
				Code                   string `json:"code"`
				Locale                 string `json:"locale"`
				Name                   string `json:"name"`
				CountryName            string `json:"countryName"`
				IsSeoIrrelevant        bool   `json:"isSeoIrrelevant"`
				IsSeoProductIrrelevant bool   `json:"isSeoProductIrrelevant"`
				Direction              string `json:"direction"`
				Formats                struct {
					Number struct {
						DecimalSeparator   string `json:"decimalSeparator"`
						ThousandsSeparator string `json:"thousandsSeparator"`
					} `json:"number"`
					Date struct {
						ShortDate string `json:"shortDate"`
						LongDate  string `json:"longDate"`
					} `json:"date"`
				} `json:"formats"`
				IsRtl bool `json:"isRtl,omitempty"`
			} `json:"supportedLanguages"`
			GiftCardStepSliderValue       int      `json:"giftCardStepSliderValue"`
			VirtualGiftCardMinTimeToSend  int      `json:"virtualGiftCardMinTimeToSend"`
			VirtualGiftCardMaxDaysToSend  int      `json:"virtualGiftCardMaxDaysToSend"`
			IsGiftcardDetailRequired      bool     `json:"isGiftcardDetailRequired"`
			IsVirtualGiftCardShareAllowed bool     `json:"isVirtualGiftCardShareAllowed"`
			SharingMethods                []string `json:"sharingMethods"`
			ShowContactLinkInHelpMenu     bool     `json:"showContactLinkInHelpMenu"`
			IsShowTaxesRequired           bool     `json:"isShowTaxesRequired"`
			IsShowPriceTaxMessageRequired bool     `json:"isShowPriceTaxMessageRequired"`
			IsTaxIncluded                 bool     `json:"isTaxIncluded"`
			IsFacebookConversionEnabled   bool     `json:"isFacebookConversionEnabled"`
			IsRefundBankSearchAvailable   bool     `json:"isRefundBankSearchAvailable"`
			IsRepentanceEnabled           bool     `json:"isRepentanceEnabled"`
			IsPaperlessShipmentEnabled    bool     `json:"isPaperlessShipmentEnabled"`
			IsCompanyAllowed              bool     `json:"isCompanyAllowed"`
			IsDonationEnabled             bool     `json:"isDonationEnabled"`
			IsDonationFilterEnabled       bool     `json:"isDonationFilterEnabled"`
			IsWalletAvailable             bool     `json:"isWalletAvailable"`
			GeneratePermanentSeoURL       bool     `json:"generatePermanentSeoUrl"`
			UseXmediaRealWidth            bool     `json:"useXmediaRealWidth"`
			ShowContactOnUserMenu         bool     `json:"showContactOnUserMenu"`
			IsDarkModeEnabled             bool     `json:"isDarkModeEnabled"`
			ZaraIDQR                      struct {
				IsEnabled bool `json:"isEnabled"`
				Colors    []struct {
					BackgroundColor string `json:"backgroundColor"`
					TextColor       string `json:"textColor"`
				} `json:"colors"`
			} `json:"zaraIdQR"`
			IsNewsletterSingleOptInEnabled    bool `json:"isNewsletterSingleOptInEnabled"`
			IsNewsletterSingleOptOutEnabled   bool `json:"isNewsletterSingleOptOutEnabled"`
			IsMultiOrderReturnsExchangeActive bool `json:"isMultiOrderReturnsExchangeActive"`
			IsMultiOrderReturnsReturnActive   bool `json:"isMultiOrderReturnsReturnActive"`
			IsMultiOrderReturnsDetailActive   bool `json:"isMultiOrderReturnsDetailActive"`
			IsReturnRequestFormProcessEnabled bool `json:"isReturnRequestFormProcessEnabled"`
			IsReturnInStoreMessageVisible     bool `json:"isReturnInStoreMessageVisible"`
			IsOnlineExchangeAllowed           bool `json:"isOnlineExchangeAllowed"`
			IsReturnRequestListAvailable      bool `json:"isReturnRequestListAvailable"`
			IsPromotioniOSActive              bool `json:"isPromotioniOSActive"`
			IsPromotionAndroidActive          bool `json:"isPromotionAndroidActive"`
			IsPrefixPhoneDisabled             bool `json:"isPrefixPhoneDisabled"`
			IsABTestingAllowed                bool `json:"isABTestingAllowed"`
			HidePriceNotAvailableMessage      bool `json:"hidePriceNotAvailableMessage"`
			VirtualGiftCard                   struct {
				Share struct {
					EnabledChannels []any `json:"enabledChannels"`
				} `json:"share"`
				InstantShipping struct {
					EnabledChannels []any `json:"enabledChannels"`
					IsEnabled       bool  `json:"isEnabled"`
				} `json:"instantShipping"`
			} `json:"virtualGiftCard"`
			GitftCardMonthsToExpire int    `json:"gitftCardMonthsToExpire"`
			StockBaseURL            string `json:"stockBaseUrl"`
			Sections                []struct {
				ID             int      `json:"id"`
				Name           string   `json:"name"`
				Description    string   `json:"description"`
				AvailableFor   []string `json:"availableFor"`
				EngDescription string   `json:"engDescription"`
			} `json:"sections"`
			DisplayFuturePrice bool `json:"displayFuturePrice"`
			PhysicalStores     struct {
				HasPhysicalStores    bool     `json:"hasPhysicalStores"`
				StoreServices        []string `json:"storeServices"`
				StoreServicesBaseURL string   `json:"storeServicesBaseUrl"`
				Services             struct {
					StockInStore struct {
						BaseURL         string   `json:"baseUrl"`
						EnabledChannels []string `json:"enabledChannels"`
					} `json:"stockInStore"`
				} `json:"services"`
			} `json:"physicalStores"`
			IsRegisterEnabled   bool `json:"isRegisterEnabled"`
			AccountVerification struct {
				IsEnabledForGuestUser                bool   `json:"isEnabledForGuestUser"`
				IsEnabledForRegisteredUser           bool   `json:"isEnabledForRegisteredUser"`
				VerificationMethodAvailableForSignUp string `json:"verificationMethodAvailableForSignUp"`
			} `json:"accountVerification"`
			IsLiteRegisterEnabled      bool `json:"isLiteRegisterEnabled"`
			IsUsePhoneAsLogonID        bool `json:"isUsePhoneAsLogonId"`
			IsAddressEvaluationEnabled bool `json:"isAddressEvaluationEnabled"`
			IsDeleteAccountActive      bool `json:"isDeleteAccountActive"`
			IsRefundHelpEnabled        bool `json:"isRefundHelpEnabled"`
			IsLiveTrackingEnabled      bool `json:"isLiveTrackingEnabled"`
			Wishlist                   struct {
				IsEnabled bool `json:"isEnabled"`
			} `json:"wishlist"`
			IsRegionGroupEnabled bool `json:"isRegionGroupEnabled"`
			Styles               struct {
				Colors struct {
					PriceColors struct {
						SalesPrice struct {
							BackgroundColorHexCode         string `json:"backgroundColorHexCode"`
							TextColorHexCode               string `json:"textColorHexCode"`
							DarkModeBackgroundColorHexCode string `json:"darkModeBackgroundColorHexCode"`
							DarkModeTextColorHexCode       string `json:"darkModeTextColorHexCode"`
						} `json:"salesPrice"`
						FuturePrice struct {
							BackgroundColorHexCode         string `json:"backgroundColorHexCode"`
							TextColorHexCode               string `json:"textColorHexCode"`
							DarkModeBackgroundColorHexCode string `json:"darkModeBackgroundColorHexCode"`
							DarkModeTextColorHexCode       string `json:"darkModeTextColorHexCode"`
						} `json:"futurePrice"`
						HighlightPrice struct {
							BackgroundColorHexCode         string `json:"backgroundColorHexCode"`
							TextColorHexCode               string `json:"textColorHexCode"`
							DarkModeBackgroundColorHexCode string `json:"darkModeBackgroundColorHexCode"`
							DarkModeTextColorHexCode       string `json:"darkModeTextColorHexCode"`
						} `json:"highlightPrice"`
					} `json:"priceColors"`
					FreeShippingMethod struct {
						TextColorHexCodeLight string `json:"textColorHexCodeLight"`
						TextColorHexCodeDark  string `json:"textColorHexCodeDark"`
					} `json:"freeShippingMethod"`
				} `json:"colors"`
				Checkout struct {
					Summary string `json:"summary"`
				} `json:"checkout"`
			} `json:"styles"`
			HasUserPreferredLanguage bool `json:"hasUserPreferredLanguage"`
			CartRelatedProducts      struct {
				IsEnabled         bool     `json:"isEnabled"`
				SupportedChannels []string `json:"supportedChannels"`
				MaxItems          int      `json:"maxItems"`
			} `json:"cartRelatedProducts"`
			SizeSelectorAfterAdd      []string `json:"sizeSelectorAfterAdd"`
			RecommendProviderLocation struct {
				Global struct {
					EnabledSections []any  `json:"enabledSections"`
					Provider        string `json:"provider"`
				} `json:"global"`
				Grid struct {
					EnabledSections []any  `json:"enabledSections"`
					Provider        string `json:"provider"`
				} `json:"grid"`
				PdpGrid struct {
					EnabledSections []string `json:"enabledSections"`
					Provider        string   `json:"provider"`
				} `json:"pdpGrid"`
				PdpToast struct {
					EnabledSections []string `json:"enabledSections"`
					Provider        string   `json:"provider"`
				} `json:"pdpToast"`
				Cart struct {
					EnabledSections []any  `json:"enabledSections"`
					Provider        string `json:"provider"`
				} `json:"cart"`
				SearchHome struct {
					EnabledSections []string `json:"enabledSections"`
					Provider        string   `json:"provider"`
				} `json:"searchHome"`
			} `json:"recommendProviderLocation"`
			DiscountDisclaimer        string `json:"discountDisclaimer"`
			DisplayOriginalPrice      bool   `json:"displayOriginalPrice"`
			DisplayDiscountPercentage bool   `json:"displayDiscountPercentage"`
			AddressSearchEngine       struct {
				Daum struct {
					ClientServiceURL string `json:"clientServiceUrl"`
				} `json:"daum"`
			} `json:"addressSearchEngine"`
			DeviceFingerprint struct {
				AlipayJavascriptRiskURL         string `json:"alipayJavascriptRiskUrl"`
				DeviceFingerPrintFlashActive    bool   `json:"deviceFingerPrintFlashActive"`
				FraudCybersourceBasicMerchantID string `json:"fraudCybersourceBasicMerchantId"`
				GiftcardFraudCheckActive        bool   `json:"giftcardFraudCheckActive"`
				Hostname                        string `json:"hostname"`
				MerchantID                      string `json:"merchantId"`
				OrganizationID                  string `json:"organizationId"`
			} `json:"deviceFingerprint"`
			Support struct {
				AbTesting struct {
					AppsClientKey        string   `json:"appsClientKey"`
					Enabled              bool     `json:"enabled"`
					EnabledChannels      []string `json:"enabledChannels"`
					WebMobileClientKey   string   `json:"webMobileClientKey"`
					WebStandardClientKey string   `json:"webStandardClientKey"`
				} `json:"abTesting"`
				AccessibilityAid struct {
					EnabledChannels struct {
					} `json:"enabledChannels"`
				} `json:"accessibilityAid"`
				AdoptLegalChangesInOrderInfo bool `json:"adoptLegalChangesInOrderInfo"`
				Chat                         struct {
					IntegratedChatLangIds  []int  `json:"integratedChatLangIds"`
					IntegratedChatURL      string `json:"integratedChatUrl"`
					IsChatEnabled          bool   `json:"isChatEnabled"`
					IsMochatEnabled        bool   `json:"isMochatEnabled"`
					ItxWebChatMainURL      string `json:"itxWebChatMainUrl"`
					RegisteredChatBasePath string `json:"registeredChatBasePath"`
				} `json:"chat"`
				ClickToCall struct {
					ClickToCallBaseURL string `json:"clickToCallBaseUrl"`
					ClickToCallLangs   []any  `json:"clickToCallLangs"`
					ClickToCallLangsID []any  `json:"clickToCallLangsId"`
				} `json:"clickToCall"`
				Contact struct {
					EnabledChannels []any `json:"enabledChannels"`
				} `json:"contact"`
				ContactLinkInHelpMenu struct {
					EnabledChannels []string `json:"enabledChannels"`
				} `json:"contactLinkInHelpMenu"`
				IsContactLegalMessageRequired bool `json:"isContactLegalMessageRequired"`
				IsContactPopupEnable          bool `json:"isContactPopupEnable"`
				ConversionIntegration         struct {
					AdWords struct {
						AccountID      string `json:"accountId"`
						BaseImageURL   string `json:"baseImageUrl"`
						Color          string `json:"color"`
						Enabled        bool   `json:"enabled"`
						Format         string `json:"format"`
						Label          string `json:"label"`
						ScriptURL      string `json:"scriptUrl"`
						ScriptURLAsync string `json:"scriptUrlAsync"`
					} `json:"adWords"`
					DoubleClick struct {
						Enabled bool `json:"enabled"`
					} `json:"doubleClick"`
					Facebook struct {
						AccountID string `json:"accountId"`
						Enabled   bool   `json:"enabled"`
						ScriptURL string `json:"scriptUrl"`
					} `json:"facebook"`
					Yahoo struct {
						AccountID    string `json:"accountId"`
						BaseImageURL string `json:"baseImageUrl"`
						ConversionID string `json:"conversionId"`
						Enabled      bool   `json:"enabled"`
						Label        string `json:"label"`
						ScriptURL    string `json:"scriptUrl"`
					} `json:"yahoo"`
				} `json:"conversionIntegration"`
				Donation struct {
					IsTermsLinkEnabled bool `json:"isTermsLinkEnabled"`
				} `json:"donation"`
				DropPoints struct {
					ShowCode bool `json:"showCode"`
				} `json:"dropPoints"`
				ForceHTTPS  []string `json:"forceHttps"`
				FraudConfig struct {
					IsRiskifiedActive bool `json:"isRiskifiedActive"`
				} `json:"fraudConfig"`
				IsGiftCardExpirationDisclaimerRequired bool `json:"isGiftCardExpirationDisclaimerRequired"`
				GoogleServices                         struct {
					Key string `json:"key"`
				} `json:"googleServices"`
				Tracking struct {
					FinalMilestones []string `json:"finalMilestones"`
					LiveTracking    struct {
						EnabledChannels []any `json:"enabledChannels"`
					} `json:"liveTracking"`
					MilestonesOrder []string `json:"milestonesOrder"`
				} `json:"tracking"`
				Gtm struct {
					AccountID string `json:"accountId"`
					Enabled   bool   `json:"enabled"`
				} `json:"gtm"`
				OrderList struct {
					APIVersion string `json:"apiVersion"`
				} `json:"orderList"`
				HelpCenter struct {
					EnabledChannels []string `json:"enabledChannels"`
				} `json:"helpCenter"`
				LegalDocuments []struct {
					Kind                  string `json:"kind"`
					Label                 string `json:"label"`
					URL                   string `json:"url"`
					Version               string `json:"version"`
					ShowWarningDuringDays int    `json:"showWarningDuringDays"`
					VisibleAt             []any  `json:"visibleAt"`
				} `json:"legalDocuments"`
				Documents                   []any  `json:"documents"`
				MiniContactAvailableContext string `json:"miniContactAvailableContext"`
				Naizfit                     struct {
					WebScriptURL    string `json:"webScriptUrl"`
					AppsScriptURL   string `json:"appsScriptUrl"`
					EnabledChannels []any  `json:"enabledChannels"`
				} `json:"naizfit"`
				OnlineExchange struct {
					IsEnabled                           bool     `json:"isEnabled"`
					EnabledChannels                     []string `json:"enabledChannels"`
					MaxExchangeUnitsCount               int      `json:"maxExchangeUnitsCount"`
					IsMobileNewProcessForGuestAvailable bool     `json:"isMobileNewProcessForGuestAvailable"`
					IsNewWindowForGuestAvailable        bool     `json:"isNewWindowForGuestAvailable"`
					IsShippingEditable                  bool     `json:"isShippingEditable"`
				} `json:"onlineExchange"`
				OrderProcess struct {
					EdgeImplementationStatus struct {
						WebMobile   string `json:"webMobile"`
						WebStandard string `json:"webStandard"`
					} `json:"edgeImplementationStatus"`
					IsFullBillingAddresNeeded  bool   `json:"isFullBillingAddresNeeded"`
					RestylingCheckoutStatus    string `json:"restylingCheckoutStatus"`
					RestylingCheckoutURL       string `json:"restylingCheckoutUrl"`
					RestylingLegacyCheckoutURL string `json:"restylingLegacyCheckoutUrl"`
				} `json:"orderProcess"`
				Payment struct {
					CardinalDeviceDataCollectionURL    string `json:"cardinalDeviceDataCollectionUrl"`
					CreditCardExpirationMonthsThresold int    `json:"creditCardExpirationMonthsThresold"`
					IsShowPaymentExchangeWarningEnable bool   `json:"isShowPaymentExchangeWarningEnable"`
					KcpBinaryInstallerURL              string `json:"kcpBinaryInstallerUrl"`
					KcpJsURL                           string `json:"kcpJsUrl"`
					KcpWebPluginURL                    string `json:"kcpWebPluginUrl"`
					KlarnaWidgetURL                    string `json:"klarnaWidgetUrl"`
					OfflineExpirationDelayTime         int    `json:"offlineExpirationDelayTime"`
				} `json:"payment"`
				ProductsCategoryNamePosition int `json:"productsCategoryNamePosition"`
				ProductsSearch               struct {
					Provider                 string `json:"provider"`
					MaxPrefetchedNextQueries int    `json:"maxPrefetchedNextQueries"`
					SearchByImage            struct {
						Enabled bool   `json:"enabled"`
						Host    string `json:"host"`
						APIKey  string `json:"apiKey"`
					} `json:"searchByImage"`
					Urls struct {
						Autocomplete string `json:"autocomplete"`
						Empathize    string `json:"empathize"`
						Ping         string `json:"ping"`
						Search       string `json:"search"`
						NextQueries  string `json:"nextQueries"`
					} `json:"urls"`
					Filtering struct {
						EnabledStatus struct {
							WebMobile   string `json:"webMobile"`
							WebStandard string `json:"webStandard"`
							IOS         string `json:"iOS"`
							Android     string `json:"android"`
						} `json:"enabledStatus"`
						AllowedFacets []string `json:"allowedFacets"`
					} `json:"filtering"`
					Query []struct {
						Name  string   `json:"name"`
						Value []string `json:"value"`
					} `json:"query"`
					EngineName string `json:"engineName"`
					Engines    []struct {
						Name string `json:"name"`
						Urls struct {
							Autocomplete string `json:"autocomplete"`
							Empathize    string `json:"empathize"`
							Ping         string `json:"ping"`
							Search       string `json:"search"`
							NextQueries  string `json:"nextQueries"`
						} `json:"urls"`
						Query []struct {
							Name  string   `json:"name"`
							Value []string `json:"value"`
						} `json:"query"`
					} `json:"engines"`
				} `json:"productsSearch"`
				HelpSearch struct {
					AppID                       string `json:"appId"`
					APIKey                      string `json:"apiKey"`
					IndexName                   string `json:"indexName"`
					MinCharsToQuery             int    `json:"minCharsToQuery"`
					MinMillisToQuery            int    `json:"minMillisToQuery"`
					ShowSuggestionWhenNoResults bool   `json:"showSuggestionWhenNoResults"`
				} `json:"helpSearch"`
				Qubit struct {
					IsQubitEnabled bool   `json:"isQubitEnabled"`
					QubitScriptURL string `json:"qubitScriptUrl"`
				} `json:"qubit"`
				Rgpd struct {
					IsEnabled bool `json:"isEnabled"`
					ShowPopup bool `json:"showPopup"`
				} `json:"rgpd"`
				ShowAndroidLegacyCartPercent int  `json:"showAndroidLegacyCartPercent"`
				ShowPrivacyLinks             bool `json:"showPrivacyLinks"`
				StockOutSubscription         struct {
					ShouldConfirmEmail bool `json:"shouldConfirmEmail"`
				} `json:"stockOutSubscription"`
				TicketToBill struct {
					CaptchaURL       string `json:"captchaUrl"`
					CreateInvoiceURL string `json:"createInvoiceUrl"`
					IsEnabled        bool   `json:"isEnabled"`
					TicketImageURL   string `json:"ticketImageUrl"`
				} `json:"ticketToBill"`
				WebClientPerformanceMonitoring struct {
					Enabled        bool   `json:"enabled"`
					WebMobileKey   string `json:"webMobileKey"`
					WebStandardKey string `json:"webStandardKey"`
				} `json:"webClientPerformanceMonitoring"`
				WideEyes struct {
					APIKey      string `json:"apiKey"`
					Host        string `json:"host"`
					ShowSimilar struct {
						EnabledChannels []string `json:"enabledChannels"`
					} `json:"showSimilar"`
					WearItWith struct {
						EnabledChannels []string `json:"enabledChannels"`
						EnabledSections []int    `json:"enabledSections"`
					} `json:"wearItWith"`
				} `json:"wideEyes"`
				ItxRestRelatedProducts struct {
					EnabledChannels []string `json:"enabledChannels"`
					EnabledSections []string `json:"enabledSections"`
				} `json:"itxRestRelatedProducts"`
				Checkout struct {
					DisabledCartContinue         bool `json:"disabledCartContinue"`
					ForceShippingMethodSelection bool `json:"forceShippingMethodSelection"`
					QuickPurchaseEnabled         bool `json:"quickPurchaseEnabled"`
					PostPayment                  struct {
						SupportedChannels []any `json:"supportedChannels"`
					} `json:"postPayment"`
					GenericPunchout struct {
						SupportedChannels []any `json:"supportedChannels"`
					} `json:"genericPunchout"`
					DeliveryGroupsEnabled struct {
						SupportedChannels []string `json:"supportedChannels"`
					} `json:"deliveryGroupsEnabled"`
					GetPurchaseAttempt struct {
						SupportedChannels []any `json:"supportedChannels"`
					} `json:"getPurchaseAttempt"`
					DeliveryGroup struct {
						DefaultVariant string `json:"defaultVariant"`
						PostShipping   struct {
							SupportedChannels []string `json:"supportedChannels"`
						} `json:"postShipping"`
						ShippingByDelivery struct {
							SupportedChannels []any `json:"supportedChannels"`
						} `json:"shippingByDelivery"`
					} `json:"deliveryGroup"`
				} `json:"checkout"`
				MultiWishlist struct {
					EnabledChannels []string `json:"enabledChannels"`
				} `json:"multiWishlist"`
				WishlistActiveChannels           []string `json:"wishlistActiveChannels"`
				WishlistOnUserMenuActiveChannels []any    `json:"wishlistOnUserMenuActiveChannels"`
				WishlistSharingActiveChannels    []string `json:"wishlistSharingActiveChannels"`
				BuyLaterActiveChannels           []string `json:"buyLaterActiveChannels"`
				CategoryGrid                     struct {
					WebMobile struct {
						ClientRows      int `json:"clientRows"`
						NumPreloadMedia int `json:"numPreloadMedia"`
					} `json:"webMobile"`
				} `json:"categoryGrid"`
				IsCookieMigrationEnabled bool `json:"isCookieMigrationEnabled"`
				IsNewAddressFormsEnabled bool `json:"isNewAddressFormsEnabled"`
				ClientTelemetry          struct {
					PageViewsEnabledChannels         []string `json:"pageViewsEnabledChannels"`
					AddToCartEnabledChannels         []string `json:"addToCartEnabledChannels"`
					PurchaseConfirmedEnabledChannels []string `json:"purchaseConfirmedEnabledChannels"`
				} `json:"clientTelemetry"`
				IsSRPLSSubscriptionEnabled bool `json:"isSRPLSSubscriptionEnabled"`
				Cookies                    struct {
					CookiesConsent struct {
						EnabledChannels []string `json:"enabledChannels"`
						OneTrust        struct {
							Ids struct {
								Web       string `json:"web"`
								WebMobile string `json:"web-mobile"`
								Ios       string `json:"ios"`
								Android   string `json:"android"`
								MiniP     string `json:"mini-p"`
							} `json:"ids"`
						} `json:"oneTrust"`
					} `json:"cookiesConsent"`
				} `json:"cookies"`
				CookiesConfig struct {
					CookiesConsent struct {
						EnabledChannels []string `json:"enabledChannels"`
						OneTrust        struct {
							Ids struct {
								Web       string `json:"web"`
								WebMobile string `json:"web-mobile"`
								Ios       string `json:"ios"`
								Android   string `json:"android"`
								MiniP     string `json:"mini-p"`
							} `json:"ids"`
						} `json:"oneTrust"`
					} `json:"cookiesConsent"`
				} `json:"cookiesConfig"`
				Returns struct {
					PaymentMethodsWarning struct {
						EnabledChannels []any `json:"enabledChannels"`
						MethodsAllowed  []any `json:"methodsAllowed"`
					} `json:"paymentMethodsWarning"`
					ReturnRequestForm struct {
						EnabledChannels []any `json:"enabledChannels"`
					} `json:"returnRequestForm"`
				} `json:"returns"`
				Tempe3DViewer struct {
					Urls struct {
						React      string `json:"react"`
						Standalone string `json:"standalone"`
					} `json:"urls"`
				} `json:"tempe3DViewer"`
				IsIOSNewGridEnabled   bool `json:"isIOSNewGridEnabled"`
				ProductRecommendation struct {
					Enabled bool `json:"enabled"`
				} `json:"productRecommendation"`
				AccountVerification struct {
					RegistrationProcessActiveChannels       []any  `json:"registrationProcessActiveChannels"`
					RegistrationProcessV2ActiveChannels     []any  `json:"registrationProcessV2ActiveChannels"`
					RegistrationProcessV2VerificationMethod string `json:"registrationProcessV2VerificationMethod"`
					RegisteredActiveChannels                []any  `json:"registeredActiveChannels"`
				} `json:"accountVerification"`
				LiveStreaming struct {
					ScriptUrls struct {
						Web  string `json:"web"`
						Apps string `json:"apps"`
					} `json:"scriptUrls"`
					EnabledChannels []any `json:"enabledChannels"`
				} `json:"liveStreaming"`
				Legal struct {
					TermsAndConditions struct {
						Kind                  string `json:"kind"`
						Label                 string `json:"label"`
						URL                   string `json:"url"`
						Version               string `json:"version"`
						ShowWarningDuringDays int    `json:"showWarningDuringDays"`
						VisibleAt             []any  `json:"visibleAt"`
					} `json:"TERMS_AND_CONDITIONS"`
					PrivacyPolicy struct {
						Kind                  string   `json:"kind"`
						Label                 string   `json:"label"`
						URL                   string   `json:"url"`
						Version               string   `json:"version"`
						ShowWarningDuringDays int      `json:"showWarningDuringDays"`
						VisibleAt             []string `json:"visibleAt"`
					} `json:"PRIVACY_POLICY"`
					GiftcardTerms struct {
						Kind                  string `json:"kind"`
						Label                 string `json:"label"`
						URL                   string `json:"url"`
						Version               string `json:"version"`
						ShowWarningDuringDays int    `json:"showWarningDuringDays"`
						VisibleAt             []any  `json:"visibleAt"`
					} `json:"GIFTCARD_TERMS"`
					CancellationConditions struct {
						Kind                  string `json:"kind"`
						Label                 string `json:"label"`
						URL                   string `json:"url"`
						Version               string `json:"version"`
						ShowWarningDuringDays int    `json:"showWarningDuringDays"`
						VisibleAt             []any  `json:"visibleAt"`
					} `json:"CANCELLATION_CONDITIONS"`
				} `json:"legal"`
			} `json:"support"`
			GiftCardTerms struct {
				URL     string `json:"url"`
				Version string `json:"version"`
			} `json:"giftCardTerms"`
			GeoInfo struct {
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Bounds struct {
					Northeast struct {
						Lat float64 `json:"lat"`
						Lng float64 `json:"lng"`
					} `json:"northeast"`
					Southwest struct {
						Lat float64 `json:"lat"`
						Lng float64 `json:"lng"`
					} `json:"southwest"`
				} `json:"bounds"`
			} `json:"geoInfo"`
			SizeRecommender struct {
				SizeRecommenderDesktopScript  string `json:"sizeRecommenderDesktopScript"`
				IsSizeRecommenderEnabled      bool   `json:"isSizeRecommenderEnabled"`
				SizeRecommenderMobileScript   string `json:"sizeRecommenderMobileScript"`
				SizeRecommenderPurchaseScript string `json:"sizeRecommenderPurchaseScript"`
			} `json:"sizeRecommender"`
			InteractiveSizeGuide struct {
				EnabledChannels []string `json:"enabledChannels"`
				EnabledSections []string `json:"enabledSections"`
			} `json:"interactiveSizeGuide"`
			Locale struct {
				CurrencyCode     string  `json:"currencyCode"`
				CurrencySymbol   string  `json:"currencySymbol"`
				CurrencyFormat   string  `json:"currencyFormat"`
				CurrencyDecimals int     `json:"currencyDecimals"`
				CurrencyRate     float64 `json:"currencyRate"`
				Currencies       []struct {
					CurrencyCode     string `json:"currencyCode"`
					CurrencySymbol   string `json:"currencySymbol"`
					CurrencyFormat   string `json:"currencyFormat"`
					CurrencyDecimals int    `json:"currencyDecimals"`
					ConversionRates  []struct {
						From string  `json:"from"`
						To   string  `json:"to"`
						Rate float64 `json:"rate"`
					} `json:"conversionRates"`
				} `json:"currencies"`
				IsBankBicMandatory bool `json:"isBankBicMandatory"`
				IsBankInnMandatory bool `json:"isBankInnMandatory"`
				IsBankSwift        bool `json:"isBankSwift"`
				IsCompoundName     bool `json:"isCompoundName"`
				IsLastNameFirst    bool `json:"isLastNameFirst"`
			} `json:"locale"`
		} `json:"store"`
		WechatAppID string `json:"wechatAppId"`
		Geo         struct {
			MapsService string `json:"mapsService"`
			Gmaps       struct {
				User                        string `json:"user"`
				Key                         string `json:"key"`
				Channel                     string `json:"channel"`
				IsAddressAutocompleteActive bool   `json:"isAddressAutocompleteActive"`
				AutocompleteKey             string `json:"autocompleteKey"`
			} `json:"gmaps"`
		} `json:"geo"`
		Sem struct {
			Exelution struct {
				Enabled bool `json:"enabled"`
			} `json:"exelution"`
		} `json:"sem"`
		Cis struct {
			Messaging struct {
				SubscribeURL struct {
					Base string `json:"base"`
					Cn   string `json:"cn"`
				} `json:"subscribeUrl"`
				RenewSubscriptionURL struct {
					Base string `json:"base"`
					Cn   string `json:"cn"`
				} `json:"renewSubscriptionUrl"`
			} `json:"messaging"`
		} `json:"cis"`
		IsDevEnv        bool     `json:"isDevEnv"`
		Channel         string   `json:"channel"`
		EnabledFeatures []string `json:"enabledFeatures"`
		CookiesConsent  struct {
			IsEnabled  bool   `json:"isEnabled"`
			OneTrustID string `json:"oneTrustId"`
		} `json:"cookiesConsent"`
		Trkpln struct {
			ClientID string `json:"clientId"`
		} `json:"trkpln"`
		Env string `json:"env"`
	} `json:"clientAppConfig"`
	RenderingEngine string `json:"renderingEngine"`
	AppVersion      string `json:"appVersion"`
	I18NVersion     int    `json:"i18nVersion"`
}
