package entity

type Meta struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
	// Value string `json:"value,omitempty"`
	Val interface{} `json:"value"`
}

// type Date struct {
// 	timeDateValue string
// }

// // UnmarshalJSON is used to convert the timestamp from JSON
// func (t *Date) UnmarshalJSON(s []byte) error {
// 	var errorTime error
// 	stringTime := strings.ReplaceAll(string(s), "\"", "")
// 	if len(stringTime) > 10 { // 04.01.2023 16:55
// 		t.timeDateValue, errorTime = time.Parse("02.01.2006 15:04", stringTime)
// 	} else if len(stringTime) == 6 { // 05\/01
// 		t.timeDateValue, errorTime = time.Parse("02\\/01", stringTime)
// 	} else {
// 		t.timeDateValue, errorTime = time.Parse("02.01.2006", stringTime)
// 	}
// 	return errorTime
// }

// // MarshalJSON is used to convert the timestamp to JSON
// func (t Date) MarshalJSON() ([]byte, error) {
// 	return []byte(t.timeDateValue.Format("02.01.2006")), nil
// }
