package wcprod_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/wcprod"
)

// go test -timeout 90s -run ^TestAddProduct$ github.com/RB-PRO/SanctionedClothing/pkg/wcprod
func TestAddProduct(t *testing.T) {
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		log.Fatalln(errorInitWcAdd)
	}

	variet := varietBasesVariety2() // Получаем товар

	errAdd := Adding.AddProduct(variet.Product[0])
	if errAdd != nil {
		t.Error(errAdd)
	}
}

func TestEditDelivery(t *testing.T) {
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		log.Fatalln(errorInitWcAdd)
	}
	variet := varietBasesVariety2() // Получаем товар

	delivery := Adding.EditDelivery(variet.Product[0].Cat, 500)
	fmt.Println("Получил цену доставки", delivery, ". Для категории:\n", variet.Product[0].Cat[2])
}

// Создать тестовый товар
func varietBasesVariety2() bases.Variety2 {
	AddCat := make([]bases.Cat, 4)
	AddCat[0].Name = "Женщины"
	AddCat[0].Slug = "women"
	AddCat[1].Name = "Clothing"
	AddCat[1].Slug = "clothing"
	AddCat[2].Name = "Sweaters"
	AddCat[2].Slug = "sweaters"
	AddCat[3].Name = "1.STATE"
	AddCat[3].Slug = "1-state"
	return bases.Variety2{
		[]bases.Product2{
			bases.Product2{
				Manufacturer: "1.STATE",
				Name:         "Balloon Sleeve Crew Neck Sweater",
				FullName:     "Complete your cool-weather look with the soft and cozy 1.STATE™ Balloon Sleeve Crew Neck Sweater.",
				Link:         "/p/1-state-balloon-sleeve-crew-neck-sweater-antique-white/product/9621708/color/26216",
				Article:      "9621708",
				//Cat3:            bases.Cat{{"Женщины", "women"}, {"Clothing", "clothing"}, {"Sweaters", "sweaters"}, {"1.STATE", "1-state"}},
				Cat:            AddCat,
				GenderLabel:    "women",
				Specifications: map[string]string{"Length": "23 in"},
				Size:           []string{"SM", "LG", "XL"},
				Description: struct {
					Eng string
					Rus string
				}{Eng: `Complete your cool-weather look with the soft and cozy 1.STATE™ Balloon Sleeve Crew Neck Sweater.
			SKU: #9621708
			Pull-over design with ribbed crew neckline.
			Long balloon sleeves with elongated, ribbed cuffs.
			Classic fit with straight hemline.
			73% acrylic, 24% polyester, 3% spandex.
			Hand wash, dry flat.
			Imported.
			Product measurements were taken using size SM. Please note that measurements may vary by size.
			 Length: 23 in`},
				// Item: map[string]bases.ProdParam{
				// 	"wild-oak": bases.ProdParam{
				// 		Link:     "/product/9621708/color/836781",
				// 		ColorEng: "Wild Oak",
				// 		Price:    42.0,
				// 		Size:     []string{"SM", "LG", "XL"},
				// 		Image:    []string{"https://m.media-amazon.com/images/I/91GJ2hRcTeL.jpg", "https://m.media-amazon.com/images/I/91WQzGVObeL.jpg", "https://m.media-amazon.com/images/I/913KXCLH1lL.jpg", "https://m.media-amazon.com/images/I/71a8c4Fw+uL.jpg"},
				// 	},
				// 	"antique-white": bases.ProdParam{
				// 		Link:     "/product/9621708/color/26216",
				// 		ColorEng: "Antique White",
				// 		Price:    31.58,
				// 		Size:     []string{"SM", "LG", "XL"},
				// 		Image:    []string{"https://m.media-amazon.com/images/I/71Mf94kDFvL.jpg", "https://m.media-amazon.com/images/I/71EOOcBc+bL.jpg", "https://m.media-amazon.com/images/I/81PeCItuTmL.jpg", "https://m.media-amazon.com/images/I/71+cz20ouIL.jpg"},
				// 	},
				// },
			},
		},
	}
}
