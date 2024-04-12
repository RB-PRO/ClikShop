package clicker

import (
	"github.com/RB-PRO/ClikShop/pkg/wcprod"
	wc "github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity"
)

type Hand struct {
	URL  string // Ссылка на товар в ClikShop
	Link string // Ссылка на источник
}

func Hands(Adding *wcprod.WcAdd) (hands []Hand, ErrorHand error) {
	var IsLastPages bool
	var TecalPage int
	for !IsLastPages {
		var ErrorAllProd error
		var items []entity.Product
		items, _, _, IsLastPages, ErrorAllProd = Adding.WooClient.Services.Product.All(wc.ProductsQueryParams{}, TecalPage, 100)
		if ErrorAllProd != nil {
			return nil, ErrorAllProd
		}
		for _, prod := range items {
			var Link string
			for _, val := range prod.MetaData {
				if val.Key == "linkRB" {
					Link = (val.Val).(string)
				}
			}
			hands = append(hands, Hand{URL: prod.Permalink, Link: Link})
		}
		TecalPage++
	}

	return hands, nil
}
