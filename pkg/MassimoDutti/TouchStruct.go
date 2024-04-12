package massimodutti

type Touch struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	NameEn string `json:"nameEn"`
	// Image       interface{} `json:"image"`
	IsBuyable   bool   `json:"isBuyable"`
	OnSpecial   bool   `json:"onSpecial"`
	BackSoon    string `json:"backSoon"`
	UnitsLot    int    `json:"unitsLot"`
	IsTop       int    `json:"isTop"`
	SizeSystem  string `json:"sizeSystem"`
	SubFamily   string `json:"subFamily"`
	ProductType string `json:"productType"`
	// BundleColors []interface{} `json:"bundleColors"`
	// Tags         []interface{} `json:"tags"`
	Attributes []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
		Type  string `json:"type"`
		// Identifier interface{} `json:"identifier"`
	} `json:"attributes"`
	// RelatedCategories      []interface{} `json:"relatedCategories"`
	// Attachments            []interface{} `json:"attachments"`
	// BundleProductSummaries []interface{} `json:"bundleProductSummaries"`
	BundleProductSummaries []struct {
		ID           int           `json:"id"`
		Type         string        `json:"type"`
		Name         string        `json:"name"`
		NameEn       string        `json:"nameEn"`
		Image        interface{}   `json:"image"`
		IsBuyable    bool          `json:"isBuyable"`
		OnSpecial    bool          `json:"onSpecial"`
		BackSoon     string        `json:"backSoon"`
		UnitsLot     int           `json:"unitsLot"`
		IsTop        int           `json:"isTop"`
		SizeSystem   string        `json:"sizeSystem"`
		SubFamily    string        `json:"subFamily"`
		ProductType  string        `json:"productType"`
		BundleColors []interface{} `json:"bundleColors"`
		Tags         []interface{} `json:"tags"`
		Attributes   []struct {
			ID         string      `json:"id"`
			Name       string      `json:"name"`
			Value      string      `json:"value"`
			Type       string      `json:"type"`
			Identifier interface{} `json:"identifier"`
		} `json:"attributes"`
		RelatedCategories      []interface{} `json:"relatedCategories"`
		Attachments            []interface{} `json:"attachments"`
		BundleProductSummaries []interface{} `json:"bundleProductSummaries"`
		Detail                 struct {
			Description      string      `json:"description"`
			LongDescription  string      `json:"longDescription"`
			Reference        string      `json:"reference"`
			DisplayReference string      `json:"displayReference"`
			DefaultImageType interface{} `json:"defaultImageType"`
			Composition      []struct {
				Part        string `json:"part"`
				Composition []struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Percentage  string `json:"percentage"`
				} `json:"composition"`
			} `json:"composition"`
			CompositionByZone []struct {
				Part  string `json:"part"`
				Zones []struct {
					Zone        string `json:"zone"`
					ZoneName    string `json:"zoneName"`
					Composition []struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						Description string `json:"description"`
						Percentage  string `json:"percentage"`
					} `json:"composition"`
				} `json:"zones"`
			} `json:"compositionByZone"`
			Care []struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"care"`
			Colors []struct {
				ID         string      `json:"id"`
				Name       string      `json:"name"`
				ModelHeigh interface{} `json:"modelHeigh"`
				ModelName  interface{} `json:"modelName"`
				ModelSize  interface{} `json:"modelSize"`
				Image      struct {
					Timestamp          string   `json:"timestamp"`
					URL                string   `json:"url"`
					Aux                []string `json:"aux"`
					Type               []string `json:"type"`
					Style              []string `json:"style"`
					AvailableEstilismo bool     `json:"availableEstilismo"`
				} `json:"image"`
				Sizes []struct {
					Sku             int         `json:"sku"`
					Name            string      `json:"name"`
					Description     interface{} `json:"description"`
					Partnumber      string      `json:"partnumber"`
					IsBuyable       bool        `json:"isBuyable"`
					BackSoon        string      `json:"backSoon"`
					MastersSizeID   string      `json:"mastersSizeId"`
					Position        int         `json:"position"`
					Price           string      `json:"price"`
					OldPrice        interface{} `json:"oldPrice"`
					VisibilityValue string      `json:"visibilityValue"`
				} `json:"sizes"`
				IsContinuity bool `json:"isContinuity"`
				JoinLifeInfo struct {
					DescJoinLife      interface{} `json:"descJoinLife"`
					IsJoinLife        bool        `json:"isJoinLife"`
					JoinLifeID        string      `json:"joinLifeId"`
					ShortDescJoinLife string      `json:"shortDescJoinLife"`
				} `json:"joinLifeInfo"`
				CompositionDetail interface{} `json:"compositionDetail"`
				Sustainability    struct {
					Show                     bool `json:"show"`
					SyntheticFiberPercentage struct {
						Name string `json:"name"`
					} `json:"syntheticFiberPercentage"`
				} `json:"sustainability"`
				Traceability struct {
					Show    bool `json:"show"`
					Weaving struct {
						Name    string   `json:"name"`
						Country []string `json:"country"`
					} `json:"weaving"`
					DyeingPrinting struct {
						Name    string   `json:"name"`
						Country []string `json:"country"`
					} `json:"dyeingPrinting"`
					Confection struct {
						Name    string   `json:"name"`
						Country []string `json:"country"`
					} `json:"confection"`
					Assembly struct {
						Name    string        `json:"name"`
						Country []interface{} `json:"country"`
					} `json:"assembly"`
					Pricking struct {
						Name    string        `json:"name"`
						Country []interface{} `json:"country"`
					} `json:"pricking"`
					Finish struct {
						Name    string        `json:"name"`
						Country []interface{} `json:"country"`
					} `json:"finish"`
				} `json:"traceability"`
				ExtraInfo struct {
					FitSizeMessage interface{} `json:"fitSizeMessage"`
				} `json:"extraInfo"`
				JoinLifeLabelInfo struct {
					Show  bool          `json:"show"`
					Areas []interface{} `json:"areas"`
				} `json:"joinLifeLabelInfo"`
				ColFilter []interface{} `json:"colFilter"`
			} `json:"colors"`
			RelatedProducts  []interface{} `json:"relatedProducts"`
			XmediaDefaultSet interface{}   `json:"xmediaDefaultSet"`
			Xmedia           []struct {
				Path        string `json:"path"`
				XmediaItems []struct {
					Medias []struct {
						Format    int         `json:"format"`
						Clazz     int         `json:"clazz"`
						IDMedia   string      `json:"idMedia"`
						ExtraInfo interface{} `json:"extraInfo"`
						Timestamp int64       `json:"timestamp"`
					} `json:"medias"`
					Set int `json:"set"`
				} `json:"xmediaItems"`
				ColorCode       string `json:"colorCode"`
				XmediaLocations []struct {
					Locations []struct {
						MediaLocations []string `json:"mediaLocations"`
						// Location       int      `json:"location"`
						Location interface{} `json:"location"`
					} `json:"locations"`
					Set int `json:"set"`
				} `json:"xmediaLocations"`
			} `json:"xmedia"`
			SkuDimensions struct {
			} `json:"skuDimensions"`
			Dimensions struct {
			} `json:"dimensions"`
			FamilyInfo struct {
				FamilyID   int    `json:"familyId"`
				FamilyCode int    `json:"familyCode"`
				FamilyName string `json:"familyName"`
			} `json:"familyInfo"`
			SubfamilyInfo struct {
				SubFamilyID   int    `json:"subFamilyId"`
				SubFamilyCode int    `json:"subFamilyCode"`
				SubFamilyName string `json:"subFamilyName"`
			} `json:"subfamilyInfo"`
			IsJoinLife        bool          `json:"isJoinLife"`
			CompositionDetail interface{}   `json:"compositionDetail"`
			Promotions        []interface{} `json:"promotions"`
		} `json:"detail"`
		Field5                 string        `json:"field5"`
		Sequence               float64       `json:"sequence"`
		Section                string        `json:"section"`
		Family                 string        `json:"family"`
		SectionName            string        `json:"sectionName"`
		SectionNameEN          string        `json:"sectionNameEN"`
		FamilyName             string        `json:"familyName"`
		FamilyNameEN           string        `json:"familyNameEN"`
		SubFamilyName          string        `json:"subFamilyName"`
		SubFamilyNameEN        string        `json:"subFamilyNameEN"`
		StartDate              string        `json:"startDate"`
		ProductURL             string        `json:"productUrl"`
		GridElemType           string        `json:"gridElemType"`
		AvailabilityDate       string        `json:"availabilityDate"`
		ProductURLTranslations []interface{} `json:"productUrlTranslations"`
	} `json:"bundleProductSummaries,omitempty"`
	Detail struct {
		Description      string      `json:"description"`
		LongDescription  string      `json:"longDescription"`
		Reference        string      `json:"reference"`
		DisplayReference string      `json:"displayReference"`
		DefaultImageType interface{} `json:"defaultImageType"`
		Composition      []struct {
			Part        string `json:"part"`
			Composition []struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Percentage  string `json:"percentage"`
			} `json:"composition"`
		} `json:"composition"`
		CompositionByZone []interface{} `json:"compositionByZone"`
		Care              []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"care"`
		Colors []struct {
			ID         string      `json:"id"`
			Name       string      `json:"name"`
			ModelHeigh interface{} `json:"modelHeigh"`
			ModelName  interface{} `json:"modelName"`
			ModelSize  interface{} `json:"modelSize"`
			Image      struct {
				Timestamp          string        `json:"timestamp"`
				URL                string        `json:"url"`
				Aux                []string      `json:"aux"`
				Type               []string      `json:"type"`
				Style              []interface{} `json:"style"`
				AvailableEstilismo bool          `json:"availableEstilismo"`
			} `json:"image"`
			Sizes []struct {
				Sku             int         `json:"sku"`
				Name            string      `json:"name"`
				Description     interface{} `json:"description"`
				Partnumber      string      `json:"partnumber"`
				IsBuyable       bool        `json:"isBuyable"`
				BackSoon        string      `json:"backSoon"`
				MastersSizeID   string      `json:"mastersSizeId"`
				Position        int         `json:"position"`
				Price           string      `json:"price"`
				OldPrice        interface{} `json:"oldPrice"`
				VisibilityValue string      `json:"visibilityValue"`
			} `json:"sizes"`
			IsContinuity bool `json:"isContinuity"`
			JoinLifeInfo struct {
				DescJoinLife      interface{} `json:"descJoinLife"`
				IsJoinLife        bool        `json:"isJoinLife"`
				JoinLifeID        string      `json:"joinLifeId"`
				ShortDescJoinLife string      `json:"shortDescJoinLife"`
			} `json:"joinLifeInfo"`
			CompositionDetail interface{} `json:"compositionDetail"`
			Sustainability    struct {
				Show                     bool `json:"show"`
				SyntheticFiberPercentage struct {
					Name string `json:"name"`
				} `json:"syntheticFiberPercentage"`
			} `json:"sustainability"`
			Traceability struct {
				Show    bool `json:"show"`
				Weaving struct {
					Name    string   `json:"name"`
					Country []string `json:"country"`
				} `json:"weaving"`
				DyeingPrinting struct {
					Name    string   `json:"name"`
					Country []string `json:"country"`
				} `json:"dyeingPrinting"`
				Confection struct {
					Name    string   `json:"name"`
					Country []string `json:"country"`
				} `json:"confection"`
				Assembly struct {
					Name    string        `json:"name"`
					Country []interface{} `json:"country"`
				} `json:"assembly"`
				Pricking struct {
					Name    string        `json:"name"`
					Country []interface{} `json:"country"`
				} `json:"pricking"`
				Finish struct {
					Name    string        `json:"name"`
					Country []interface{} `json:"country"`
				} `json:"finish"`
			} `json:"traceability"`
			ExtraInfo struct {
				FitSizeMessage interface{} `json:"fitSizeMessage"`
			} `json:"extraInfo"`
			JoinLifeLabelInfo struct {
				Show  bool          `json:"show"`
				Areas []interface{} `json:"areas"`
			} `json:"joinLifeLabelInfo"`
			ColFilter []interface{} `json:"colFilter"`
		} `json:"colors"`
		RelatedProducts  []interface{} `json:"relatedProducts"`
		XmediaDefaultSet interface{}   `json:"xmediaDefaultSet"`
		Xmedia           []struct {
			Path        string `json:"path"`
			XmediaItems []struct {
				Medias []struct {
					Format    int    `json:"format"`
					Clazz     int    `json:"clazz"`
					IDMedia   string `json:"idMedia"`
					ExtraInfo struct {
						OriginalName string `json:"originalName"`
					} `json:"extraInfo"`
					Timestamp int64 `json:"timestamp"`
				} `json:"medias"`
				Set int `json:"set"`
			} `json:"xmediaItems"`
			ColorCode       string `json:"colorCode"`
			XmediaLocations []struct {
				Locations []struct {
					MediaLocations []string `json:"mediaLocations"`
					// Location       int      `json:"location"`
					Location interface{} `json:"location"`
				} `json:"locations"`
				Set int `json:"set"`
			} `json:"xmediaLocations"`
		} `json:"xmedia"`
		SkuDimensions struct {
		} `json:"skuDimensions"`
		Dimensions struct {
		} `json:"dimensions"`
		FamilyInfo struct {
			FamilyID   int    `json:"familyId"`
			FamilyCode int    `json:"familyCode"`
			FamilyName string `json:"familyName"`
		} `json:"familyInfo"`
		SubfamilyInfo struct {
			SubFamilyID   int    `json:"subFamilyId"`
			SubFamilyCode int    `json:"subFamilyCode"`
			SubFamilyName string `json:"subFamilyName"`
		} `json:"subfamilyInfo"`
		IsJoinLife        bool          `json:"isJoinLife"`
		CompositionDetail interface{}   `json:"compositionDetail"`
		Promotions        []interface{} `json:"promotions"`
	} `json:"detail,omitempty"`
	Field5          string `json:"field5"`
	Section         string `json:"section"`
	Family          string `json:"family"`
	SectionName     string `json:"sectionName"`
	SectionNameEN   string `json:"sectionNameEN"`
	FamilyName      string `json:"familyName"`
	FamilyNameEN    string `json:"familyNameEN"`
	SubFamilyName   string `json:"subFamilyName"`
	SubFamilyNameEN string `json:"subFamilyNameEN"`
	StartDate       string `json:"startDate"`
	// IsSales          interface{} `json:"isSales"`
	Keywords    string `json:"keywords"`
	MainColorid string `json:"mainColorid"`
	// FamilyCode       interface{} `json:"familyCode"`
	// SubFamilyCode    interface{} `json:"subFamilyCode"`
	ProductURL       string `json:"productUrl"`
	GridElemType     string `json:"gridElemType"`
	AvailabilityDate string `json:"availabilityDate"`
	VisibilityValue  string `json:"visibilityValue"`
	RueiData         struct {
		StoreLangRUEI     string `json:"StoreLangRUEI"`
		StoreTypeRUEI     string `json:"StoreTypeRUEI"`
		OperationTypeRUEI string `json:"OperationTypeRUEI"`
		OperationRUEI     string `json:"OperationRUEI"`
		StoreIDRUEI       string `json:"StoreIdRUEI"`
	} `json:"rueiData"`
}

type DetailsOfImageAndSize struct {
	DisplayReference string `json:"displayReference"`
	Colors           []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Sizes []struct {
			Sku       int    `json:"sku"`
			Name      string `json:"name"`
			IsBuyable bool   `json:"isBuyable"`
			Price     string `json:"price"`
		} `json:"sizes"`
		IsContinuity bool `json:"isContinuity"`
	} `json:"colors,omitempty"`
	// RelatedProducts  []interface{} `json:"relatedProducts"`
	// XmediaDefaultSet interface{}   `json:"xmediaDefaultSet"`
	Xmedia []struct {
		Path        string `json:"path"`
		XmediaItems []struct {
			Medias []struct {
				IDMedia string `json:"idMedia"`
			} `json:"medias"`
		} `json:"xmediaItems"`
		ColorCode string `json:"colorCode"`
	} `json:"xmedia"`
	IsJoinLife bool `json:"isJoinLife"`
}
