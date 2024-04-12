package wcprod

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	wc "github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity"
)

// Функция добавления товара
func (woo *WcAdd) AddProduct(product bases.Product2) error {
	log.SetOutput(woo.LogFile)                   // set log out put
	log.SetFlags(log.Lshortfile | log.LstdFlags) // optional: log date-time, filename, and line number

	if product.Article == "" {
		return errors.New("AddProduct: В товаре нет Артикула")
	}
	product.Size = bases.EditProdSize(product)
	if len(product.Size) == 0 {
		return errors.New("AddProduct: В товаре нет рамеров")
	}
	if len(product.Manufacturer) == 0 {
		return errors.New("AddProduct: В товаре нет производителя")
	}
	if len(product.Item) == 0 {
		return errors.New("AddProduct: В товаре нет вариантов по цветам")
	}

	/*
		// Создать категории для товаров и получить её ID
		idCat, AddNewId2 := woo.UserWC.AddCat2(&woo.Plc, product.Cat)
		if AddNewId2 != nil {
			return AddNewId2
		}
		fmt.Println("ID категории", idCat)
	*/

	// Создать категории для товаров и получить её ID по версии 3
	idCat, AddNewId3 := woo.AddCategoryWC(product.Cat)
	if AddNewId3 != nil {
		return AddNewId3
	}
	// fmt.Println("ID категории", idCat)

	// ManufrId, ManufName, ManufSlug := AddAttr(woo.WooClient, woo.IdAttrColor, "Производитель", product.Manufacturer)
	// fmt.Println("Для данного товара Аттрибуты Производителя:", ManufrId, ManufName, ManufSlug)

	// // Создаём аттрибуты товара для цвета
	// for key := range product.Item {
	// 	tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug := AddAttr(woo.WooClient, woo.IdAttrColor, product.Item[key].ColorEng, product.Item[key].ColorEng)
	// 	fmt.Println("Для данного товара Аттрибуты цвета будут:", tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug)
	// }
	// // Создаём аттрибуты товара для Размера
	// for _, valSize := range product.Size {
	// 	tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug := AddAttr(woo.WooClient, woo.IdAttrSize, valSize, bases.FormingColorEng(valSize))
	// 	fmt.Println("Для данного товара Аттрибуты размера будут:", tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug)
	// }

	// Собираем гендер для загрузки в теги товара
	idGender, _, isGenderSlug := bases.GenderBook(product.GenderLabel, "")
	if !isGenderSlug {
		log.Println("AddProduct: Не найден гендр из базы")
	}
	log.Println("AddProduct: Гендр:", idGender)

	// Создаём массив цветов с полными назвавниями
	var colors []string
	for _, colorSet := range product.Item {
		colors = append(colors, colorSet.ColorEng)
	}
	// Сделаю массив со всеми изображениями
	imageInput := make([]entity.ProductImage, 0)
	var chet int
	for _, colorItemValue := range product.Item {
		for indexImage, valueImage := range colorItemValue.Image {
			valueImage = strings.ReplaceAll(valueImage, "///", "/")
			// fmt.Println(">"+valueImage+"<", colorItemValue.ColorEng+"_"+strconv.Itoa(indexImage))
			imageInput = append(imageInput, entity.ProductImage{
				Src:  valueImage,
				Name: colorItemValue.ColorEng + "_" + strconv.Itoa(indexImage),
				Alt:  valueImage,
			})
			chet++
		}
	}

	// СОздаём метаданные с ссылками на все товары
	metas := []entity.Meta{
		{
			Key: "linkRB",
			Val: product.Link,
		},
	}
	for _, val := range product.Item {
		metas = append(metas, entity.Meta{
			Key: "linkRB" + "_" + val.ColorCode,
			Val: val.Link,
		})
	}

	// Структура с исходным товаром
	paramVariableProduct := wc.CreateProductRequest{
		Name:             product.Name,
		Type:             "variable",
		SKU:              product.Article,
		Description:      product.Description.Rus,
		Tags:             []entity.ProductTag{{Name: idGender, Slug: product.GenderLabel}},
		ShortDescription: product.FullName,
		RegularPrice:     200.0,
		Slug:             bases.FormingColorEng(product.Name),
		// MetaData: []entity.Meta{ // Ссылка на товар
		// 	{
		// 		Key: "linkRB",
		// 		// Value: product.Link,
		// 		Val: product.Link,
		// 	},
		// },
		MetaData:   metas,
		Images:     imageInput,
		Categories: []entity.ProductCategory{{ID: idCat}},
		Attributes: []entity.ProductAttribute{
			{
				ID:      woo.IdManuf,
				Options: []string{product.Manufacturer},
				Visible: true,
			},
			{
				ID:        woo.IdAttrColor,
				Variation: true,
				Visible:   true,
				Options:   colors,
			},
			{
				ID:        woo.IdAttrSize,
				Variation: true,
				Visible:   true,
				Options:   product.Size,
			},
		},
	}

	var item entity.Product
	var errCreate error
	var itemID int
	item, errCreate = woo.WooClient.Services.Product.Create(paramVariableProduct)
	if errCreate != nil {
		log.Println("AddProduct: При создании товара произошла ошибка:", errCreate)
		return errCreate
	}
	itemID = item.ID
	log.Println("AddProduct: Создал товар с ID:", itemID)

	// Редактирвоание вариационных товаров.
	//
	// Создаём массив из обновлений вариационных товаров, в частности создании новых вариаций товаров.
	//
	// Загружаем их одним запросом [Batch]
	//
	// [Batch]: https://woocommerce.github.io/woocommerce-rest-api-docs/#batch-update-product-variations
	var CreateVariations []wc.BatchProductVariationsCreateItem
	for colorKey, colorItemValue := range product.Item {
		// fmt.Println("colorItemValue.Size", colorItemValue.Size)
		for sizeKey, SizeValue := range colorItemValue.Size {
			log.Printf("AddProduct: ProductVariation.Create[%d:%d]: Добавляю вар. товар с цветом '%s' и размером '%s'\n", colorKey+1, sizeKey+1, colorItemValue.ColorEng, SizeValue.Val)

			// Создаём элемент создания вариационного товара
			CreateVariation := wc.CreateProductVariationRequest{
				SKU:          product.Article + "_" + colorItemValue.ColorCode + "_" + SizeValue.Val,
				RegularPrice: colorItemValue.Price,
				SalePrice:    colorItemValue.Price,
				Description:  "Цвет: " + colorItemValue.ColorEng + "\n" + product.Description.Rus,
				Attributes: []entity.ProductVariationAttribute{ // Аттрибусы товара
					{
						ID:     woo.IdAttrColor,
						Name:   "Цвет",
						Option: colorItemValue.ColorEng,
					},
					{
						ID:     woo.IdAttrSize,
						Name:   "Размер",
						Option: SizeValue.Val,
					},
				},
				Status: statusCodeForVarientProd(SizeValue.IsExit), // Переменная, которая содержит статус наличия товара.
			}

			// Поиск ID картинки
			// Мы ищем ID картинки, чтобы добавить именно ID в фотографию вариации, а не ссылку на ту же картинку. Это сокращает потребление дискового пространства.
			var ID_Image int
			for _, FindImage := range item.Images {
				if FindImage.Alt == strings.ReplaceAll(colorItemValue.Image[0], "///", "/") {
					ID_Image = FindImage.ID
				}
			}
			log.Printf("AddProduct: Добавлена картинка с ID_Image = %d\n", ID_Image)

			if len(colorItemValue.Image) != 0 {
				CreateVariation.Image = &entity.ProductImage{
					// Src:  colorItemValue.Image[0],
					ID:   ID_Image,
					Name: colorItemValue.ColorEng,
					Alt:  colorItemValue.ColorEng,
				}
			}
			// fmt.Println("CreateVariation.RegularPrice", CreateVariation.RegularPrice)
			CreateVariations = append(CreateVariations, CreateVariation)
		}
	}
	// fmt.Printf("CreateVariations: %+#v\n\n", CreateVariations)

	// Выполняем запрос на создание вариационных товаров
	_, ErrBatch := woo.WooClient.Services.ProductVariation.Batch(itemID, wc.BatchProductVariationsRequest{Create: CreateVariations})
	if ErrBatch != nil {
		log.Println("Error Add variation:", ErrBatch)
	}

	// PostSmartImageErr := woo.UserWC.PostSmartImage(itemID)
	// if PostSmartImageErr != nil {
	// 	fmt.Println(PostSmartImageErr)
	// }

	return nil
}

// Обработка значение, передаваемого в качестве контроля к-ва позиция по каждой вариации товара
func statusCodeForVarientProd(IsExit bool) string {
	if IsExit {
		return "publish"
	} else {
		return "private"
	}
}

// Сличение данных и возврат актуальной цены товара для данной категории товара
func (woo *WcAdd) EditDelivery(categorys []bases.Cat, delivery int) int {
	if len(categorys) > 2 {
		if val, ok := woo.Delivery[categorys[2].Name]; ok {
			return val
		}
	}
	return delivery
}
