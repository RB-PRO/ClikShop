package louisvuitton

// Ответ на запрос метода получения ответа метода запроса страниц
type PageResponse struct {
	Error           error  `json:"error,omitempty"` // Ошибки, которые могут появляться в ответе от сервера
	Cfmt            bool   `json:"cfmt,omitempty"`
	Ttr             int    `json:"ttr,omitempty"`
	Anc             int    `json:"anc,omitempty"`
	Page            int    `json:"page,omitempty"`
	NbPages         int    `json:"nbPages,omitempty"` // Сколько всего страниц. Если ответ 10, то страницы: [0, 9]
	HitsPerPage     int    `json:"hitsPerPage,omitempty"`
	NbHits          int    `json:"nbHits,omitempty"`
	Chunk           string `json:"chunk,omitempty"`
	CategoryType    string `json:"categoryType,omitempty"`
	CategoryID      string `json:"categoryId,omitempty"`
	HeaderName      string `json:"headerName,omitempty"`
	HighEndTemplate bool   `json:"highEndTemplate,omitempty"`
	Facets          struct {
		FacetList []struct {
			FacetName        string  `json:"facetName,omitempty"`
			FacetDisplayName string  `json:"facetDisplayName,omitempty"`
			Order            int     `json:"order,omitempty"`
			FilterType       string  `json:"filterType,omitempty"`
			FrontEligible    bool    `json:"frontEligible,omitempty"`
			MultiSelect      bool    `json:"multiSelect,omitempty"`
			FilterResetURL   string  `json:"filterResetURL,omitempty"`
			DispatchFilter   bool    `json:"dispatchFilter,omitempty"`
			RangeMin         float64 `json:"rangeMin,omitempty"`
			RangeMax         float64 `json:"rangeMax,omitempty"`
			SelectedMin      float64 `json:"selectedMin,omitempty"`
			SelectedMax      float64 `json:"selectedMax,omitempty"`
			PriceFilterReset bool    `json:"priceFilterReset,omitempty"`
		} `json:"facetList,omitempty"`
		ResetAllURLCode         string `json:"resetAllUrlCode,omitempty"`
		OneOrMoreFilterSelected bool   `json:"oneOrMoreFilterSelected,omitempty"`
		TwoOrMoreFilterSelected bool   `json:"twoOrMoreFilterSelected,omitempty"`
	} `json:"facets,omitempty"`
	Hits []HitsPage `json:"hits,omitempty"` // Список всех товаров
}

// Сами товары
type HitsPage struct {
	ProductID string `json:"productId,omitempty"`
	Name      string `json:"name,omitempty"`
	URL       string `json:"url,omitempty"`
	Color     string `json:"color,omitempty"`
	Image     []struct {
		ContentURL string `json:"contentUrl,omitempty"`
	} `json:"image,omitempty"`
	AdditionalProperty []struct {
		Name string `json:"name,omitempty"`
		// Value string `json:"value,omitempty"`
		Value interface{} `json:"value,omitempty"`
	} `json:"additionalProperty,omitempty"`
	Category [][]struct {
		Identifier    string `json:"identifier,omitempty"`
		Name          string `json:"name,omitempty"`
		AlternateName string `json:"alternateName,omitempty"`
		URL           string `json:"url,omitempty"`
	} `json:"category,omitempty"`
	Offers struct {
		PriceSpecification []struct {
			Identifier string `json:"identifier,omitempty"`
			Price      string `json:"price,omitempty"`
		} `json:"priceSpecification,omitempty"`
	} `json:"offers,omitempty"`
	Material                  string `json:"material,omitempty"`
	Identifier                string `json:"identifier,omitempty"`
	DisambiguatingDescription string `json:"disambiguatingDescription,omitempty"`
	IsSimilarTo               []struct {
		Name  string `json:"name,omitempty"`
		URL   string `json:"url,omitempty"`
		Image []struct {
			ContentURL string `json:"contentUrl,omitempty"`
		} `json:"image,omitempty"`
		AdditionalProperty []struct {
			Name string `json:"name,omitempty"`
			// Value string `json:"value,omitempty"`
			Value interface{} `json:"value,omitempty"`
		} `json:"additionalProperty,omitempty"`
		Selected   bool   `json:"selected,omitempty"`
		Identifier string `json:"identifier,omitempty"`
	} `json:"isSimilarTo,omitempty"`
}
