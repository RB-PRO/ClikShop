package entity

// ProductAttribute product attribute properties
type ProductAttribute struct {
	/*
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Type        string `json:"type"`
		OrderBy     string `json:"order_by"`
		HasArchives bool   `json:"has_archives"`
	*/

	ID        int      `json:"id"`
	Variation bool     `json:"variation"`
	Visible   bool     `json:"visible"`
	Options   []string `json:"options"`
}
