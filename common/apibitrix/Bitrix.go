package apibitrix

const bitrixURL string = "https://clikshop.ru/parser/api/%s/index.php"

type Service struct{}

func New() *Service {
	return &Service{}
}

type Bitrix interface {
	Coasts() (map[string]CoastMap, error)
	Product(Values []string) (ProdResp Product_Response, Err error)
	Products() ([]string, error)
	Variation(VrtReq []Variation_Request) (VrtResp Variation_Response, Err error)
}
