package wcprod_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/wcprod"
)

// go test -timeout 90s -run ^TestAddProduct$ github.com/RB-PRO/ClikShop/pkg/wcprod
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

// go test -v -run ^TestAddProduct_2$ github.com/RB-PRO/ClikShop/pkg/wcprod
func TestAddProduct_2(t *testing.T) {
	Adding, errorInitWcAdd := wcprod.New() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		log.Fatalln(errorInitWcAdd)
	}

	variet := Variety2_2() // Получаем товар

	fmt.Println("Спарсили товар с параметрами:\n", bases.ProdStr(variet.Product[0]))

	errAdd := Adding.AddProduct(variet.Product[0])
	if errAdd != nil {
		t.Error(errAdd)
	}
}

// Тестовый товар с сайта ZARA
// https://www.zara.com/tr/en/knit-crop-top-p01750002.html
func Variety2_2() bases.Variety2 {
	prods := bases.Variety2{Product: make([]bases.Product2, 1)}
	prods.Product[0].GenderLabel = "woman"
	prods.Product[0].Name = "KNIT CROP TOP"
	prods.Product[0].FullName = "KNIT CROP TOP"
	prods.Product[0].Link = "https://www.zara.com/tr/en/knit-crop-top-p01750002.html"
	prods.Product[0].Article = "1750/002a"
	prods.Product[0].Description.Eng = `We work with monitoring programmes to ensure compliance with our social, environmental and health and safety standards for our garments.
To assess compliance, we have developed a programme of audits and continuous improvement plans.
OUTER SHELL
54% viscose
46% polyamide`
	prods.Product[0].Manufacturer = "zara"
	prods.Product[0].Size = []string{"XS", "S", "M", "L"}

	prods.Product[0].Cat = append(prods.Product[0].Cat, bases.Cat{Name: "Женщины", Slug: "woman"})
	prods.Product[0].Cat = append(prods.Product[0].Cat, bases.Cat{Name: "KNITWEAR", Slug: "knitwear"})

	prods.Product[0].Item = make([]bases.ColorItem, 2)
	prods.Product[0].Item[0].ColorCode = "pink"
	prods.Product[0].Item[0].ColorEng = "Pink"
	prods.Product[0].Item[0].Price = 1211.8
	prods.Product[0].Item[0].Image = []string{"https://static.zara.net/photos///2023/V/0/1/p/1750/002/620/2/w/563/1750002620_1_1_1.jpg?ts=1683123062129", "https://static.zara.net/photos///2023/V/0/1/p/1750/002/620/2/w/522/1750002620_2_1_1.jpg?ts=1683123058519"}
	prods.Product[0].Item[0].Size = []bases.Size{
		{
			Val:    "XS",
			IsExit: true,
		},
		{
			Val:    "S",
			IsExit: true,
		},
		{
			Val:    "M",
			IsExit: true,
		},
		{
			Val:    "L",
			IsExit: true,
		},
	}

	prods.Product[0].Item[1].ColorCode = "light-blue"
	prods.Product[0].Item[1].ColorEng = "Light blue"
	prods.Product[0].Item[1].Price = 1211.8
	prods.Product[0].Item[1].Image = []string{"https://static.zara.net/photos///2023/V/0/1/p/1750/002/406/2/w/563/1750002406_1_1_1.jpg?ts=1677230425889", "https://static.zara.net/photos///2023/V/0/1/p/1750/002/406/2/w/546/1750002406_2_1_1.jpg?ts=1677230425404"}
	prods.Product[0].Item[1].Size = []bases.Size{
		{
			Val:    "XS",
			IsExit: true,
		},
		{
			Val:    "S",
			IsExit: false,
		},
		{
			Val:    "M",
			IsExit: false,
		},
		{
			Val:    "L",
			IsExit: true,
		},
	}
	return prods
}
