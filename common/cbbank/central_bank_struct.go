package cbbank

import "time"

type Bank struct {
	Date         time.Time `json:"Date"`
	PreviousDate time.Time `json:"PreviousDate"`
	PreviousURL  string    `json:"PreviousURL"`
	Timestamp    time.Time `json:"Timestamp"`
	Valute       struct {
		Aud Currency `json:"AUD"`
		Azn Currency `json:"AZN"`
		Gbp Currency `json:"GBP"`
		Amd Currency `json:"AMD"`
		Byn Currency `json:"BYN"`
		Bgn Currency `json:"BGN"`
		Brl Currency `json:"BRL"`
		Huf Currency `json:"HUF"`
		Vnd Currency `json:"VND"`
		Hkd Currency `json:"HKD"`
		Gel Currency `json:"GEL"`
		Dkk Currency `json:"DKK"`
		Aed Currency `json:"AED"`
		Usd Currency `json:"USD"`
		Eur Currency `json:"EUR"`
		Egp Currency `json:"EGP"`
		Inr Currency `json:"INR"`
		Idr Currency `json:"IDR"`
		Kzt Currency `json:"KZT"`
		Cad Currency `json:"CAD"`
		Qar Currency `json:"QAR"`
		Kgs Currency `json:"KGS"`
		Cny Currency `json:"CNY"`
		Mdl Currency `json:"MDL"`
		Nzd Currency `json:"NZD"`
		Nok Currency `json:"NOK"`
		Pln Currency `json:"PLN"`
		Ron Currency `json:"RON"`
		Xdr Currency `json:"XDR"`
		Sgd Currency `json:"SGD"`
		Tjs Currency `json:"TJS"`
		Thb Currency `json:"THB"`
		Try Currency `json:"TRY"`
		Tmt Currency `json:"TMT"`
		Uzs Currency `json:"UZS"`
		Uah Currency `json:"UAH"`
		Czk Currency `json:"CZK"`
		Sek Currency `json:"SEK"`
		Chf Currency `json:"CHF"`
		Rsd Currency `json:"RSD"`
		Zar Currency `json:"ZAR"`
		Krw Currency `json:"KRW"`
		Jpy Currency `json:"JPY"`
	} `json:"Valute"`
}
