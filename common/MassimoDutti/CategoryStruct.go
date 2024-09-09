package massimodutti

// Структура категорий, которую можно получить, выполни [запрос].
//
// [запрос]: https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category?languageId=-1&typeCatalog=1&appId=1
type Categories struct {
	Categories []struct {
		ID               int         `json:"id"`
		Name             string      `json:"name"`
		NameEn           string      `json:"nameEn"`
		ShortDescription interface{} `json:"shortDescription"`
		Description      string      `json:"description"`
		Keywords         interface{} `json:"keywords"`
		Key              string      `json:"key"`
		NumberOfProducts interface{} `json:"numberOfProducts"`
		Type             string      `json:"type"`
		ViewCategoryID   int         `json:"viewCategoryId"`
		Subcategories    []struct {
			ID               int         `json:"id"`
			Name             string      `json:"name"`
			NameEn           string      `json:"nameEn"`
			ShortDescription interface{} `json:"shortDescription"`
			Description      string      `json:"description"`
			Keywords         interface{} `json:"keywords"`
			Key              string      `json:"key"`
			NumberOfProducts interface{} `json:"numberOfProducts"`
			Type             string      `json:"type"`
			ViewCategoryID   int         `json:"viewCategoryId"`
			Subcategories    []struct {
				ID               int           `json:"id"`
				Name             string        `json:"name"`
				NameEn           string        `json:"nameEn"`
				ShortDescription interface{}   `json:"shortDescription"`
				Description      interface{}   `json:"description"`
				Keywords         interface{}   `json:"keywords"`
				Key              string        `json:"key"`
				NumberOfProducts interface{}   `json:"numberOfProducts"`
				Type             string        `json:"type"`
				ViewCategoryID   int           `json:"viewCategoryId"`
				Subcategories    []interface{} `json:"subcategories"`
				Attachments      []struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Path string `json:"path"`
					Type string `json:"type"`
				} `json:"attachments"`
				Sequence    int           `json:"sequence"`
				OldIds      []interface{} `json:"oldIds"`
				CategoryURL string        `json:"categoryUrl"`
				SeoCategory struct {
					Name                 string      `json:"name"`
					Title                string      `json:"title"`
					MetaDescription      string      `json:"metaDescription"`
					MainHeader           string      `json:"mainHeader"`
					LongDescription      interface{} `json:"longDescription"`
					Keywords             interface{} `json:"keywords"`
					SecondaryTitle       string      `json:"secondaryTitle"`
					SecondaryDescription string      `json:"secondaryDescription"`
				} `json:"seoCategory"`
				CategoryURLTranslations []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"categoryUrlTranslations,omitempty"`
				CategoryURLParam string `json:"categoryUrlParam,omitempty"`
			} `json:"subcategories"`
			Attachments []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				Path string `json:"path"`
				Type string `json:"type"`
			} `json:"attachments"`
			Sequence         int           `json:"sequence"`
			OldIds           []interface{} `json:"oldIds"`
			CategoryURL      string        `json:"categoryUrl,omitempty"`
			CategoryURLParam string        `json:"categoryUrlParam,omitempty"`
			SeoCategory      struct {
				Name                 string      `json:"name"`
				Title                string      `json:"title"`
				MetaDescription      string      `json:"metaDescription"`
				MainHeader           string      `json:"mainHeader"`
				LongDescription      string      `json:"longDescription"`
				Keywords             interface{} `json:"keywords"`
				SecondaryTitle       string      `json:"secondaryTitle"`
				SecondaryDescription string      `json:"secondaryDescription"`
			} `json:"seoCategory"`
			CategoryURLTranslations []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"categoryUrlTranslations,omitempty"`
		} `json:"subcategories"`
		Attachments []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Path string `json:"path"`
			Type string `json:"type"`
		} `json:"attachments"`
		Sequence                int           `json:"sequence"`
		OldIds                  []interface{} `json:"oldIds"`
		CategoryURL             string        `json:"categoryUrl"`
		CategoryURLParam        string        `json:"categoryUrlParam,omitempty"`
		CategoryURLTranslations []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"categoryUrlTranslations"`
		SeoCategory struct {
			Name                 string      `json:"name"`
			Title                string      `json:"title"`
			MetaDescription      string      `json:"metaDescription"`
			MainHeader           string      `json:"mainHeader"`
			LongDescription      string      `json:"longDescription"`
			Keywords             interface{} `json:"keywords"`
			SecondaryTitle       string      `json:"secondaryTitle"`
			SecondaryDescription string      `json:"secondaryDescription"`
		} `json:"seoCategory"`
	} `json:"categories"`
}
