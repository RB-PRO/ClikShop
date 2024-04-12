package trendyol

type GroupStruct struct {
	// IsSuccess  bool `json:"isSuccess"`
	// StatusCode int  `json:"statusCode"`
	// Error      any  `json:"error"`
	Result struct {
		SlicingAttributes []struct {
			Brand struct {
				BeautifiedName string `json:"beautifiedName"`
				ID             int    `json:"id"`
				// Name           string `json:"name"`
				// IsVirtual      bool   `json:"isVirtual"`
				// Path           string `json:"path"`
			} `json:"brand"`
			Attributes []struct {
				Contents []struct {
					URL      string `json:"url"`
					ID       int    `json:"id"`
					ImageURL string `json:"imageUrl"`
					Name     string `json:"name"`
					// Price    struct {
					// 	DiscountedPrice struct {
					// 		Text  string  `json:"text"`
					// 		Value float64 `json:"value"`
					// 	} `json:"discountedPrice"`
					// 	OriginalPrice struct {
					// 		Text  string  `json:"text"`
					// 		Value float64 `json:"value"`
					// 	} `json:"originalPrice"`
					// 	SellingPrice struct {
					// 		Text  string  `json:"text"`
					// 		Value float64 `json:"value"`
					// 	} `json:"sellingPrice"`
					// } `json:"price"`
				} `json:"contents"`
				Name           string `json:"name"`
				BeautifiedName string `json:"beautifiedName"`
			} `json:"attributes"`
			Type        string `json:"type"`
			DisplayName string `json:"displayName"`
			Order       int    `json:"order"`
			DisplayType int    `json:"displayType"`
		} `json:"slicingAttributes"`
	} `json:"result"`
	Headers struct {
	} `json:"headers"`
}
