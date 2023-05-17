// Отдельно вынесенный пакет для загрузки товаров на WordPress
// Использует кривую библиотеку.
// Кривая она из-за того, что некоторые параметры не соотносятся с документацией Woocommerce.
// Для упрощения написания кода, локально исправил некоторые строки в скаченной библиотеке. В идеале локально развернуть библиотеку и провести необходимые манипуляции.
package wcprod

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/transrb"
	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
	wc "github.com/hiscaler/woocommerce-go"
	config "github.com/hiscaler/woocommerce-go/config"
	"github.com/hiscaler/woocommerce-go/entity"
)

// Созовая структура, которая объединяет в себе все необходимые данные для работы с библиотекой и для загрузки товаров
type WcAdd struct {
	NodeCategoryes *woocommerce.Node      // Дерево категорий собственной разработки
	UserWC         *woocommerce.User      // структура пользователя из своей библиотеки
	Tags           []woocommerce.Tag      // Массив тегов, которые присутствуют в WordPress
	TagMap         map[string]int         // Мапа тегов. Вообще бы её вывести отсюда нахрен
	Sttr           woocommerce.Attributes // Структура аттрибутов, которые лежат на WP
	Plc            woocommerce.Categorys  // Массив категорий товара
	Tr             *transrb.Translate     // Переводчик

	// ID аттрибутов в WordPress.
	IdAttrColor int
	IdAttrSize  int
	IdManuf     int

	Delivery map[string]int // Мапа цен доставки для товаров

	WooClient *wc.WooCommerce // Клиент пользовательской библиотеки, с помощью которой добавляю товар

	Cat3 map[int]*Category3Base // Мапа категории
}

// Инициализации базовой структуры загрузки товара
func New() (*WcAdd, error) {
	// Клиент от сторонней библиотеки(пользовательской)
	b, err := os.ReadFile("config_test.json")
	if err != nil {
		return nil, errors.New("wcprod: New: Read config error: " + err.Error())
	}

	var c config.Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		return nil, errors.New("wcprod: New: Parse config file error: " + err.Error())
	}
	wooClient := wc.NewClient(c)

	// Мой клиент
	consumer_key, _ := DataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := DataFile("secret_key")     // Секретный код пользователя

	userWC, _ := woocommerce.New(consumer_key, secret_key) // Авторизация
	if okErr := userWC.IsOrder(); okErr != nil {           // Проверка на авторизацию
		return nil, okErr
	}

	// Мапа цены доставки
	Delivery, ErrorDelivery := XlsxDelivery()
	if ErrorDelivery != nil {
		return nil, ErrorDelivery
	}

	// Теги
	tags, tagsError := userWC.AllTags_WC()
	if tagsError != nil {
		return nil, tagsError
	}

	// Создать Мапу тэгов
	tagMap := woocommerce.MapTags(tags)

	// Получить дерево категорий
	plc, errPLC := userWC.ProductsCategories()
	if errPLC != nil {
		return nil, errPLC
	}

	// Переводчик

	FolderID, _ := DataFile("FolderID")
	OAuthToken, _ := DataFile("OAuthToken")

	tr, ErrTranslate := transrb.New(FolderID, OAuthToken)
	if ErrTranslate != nil {
		return nil, ErrTranslate
	}

	// Дерево категорий - Формирование внутренней структуры
	NodeCategoryes := woocommerce.NewCategoryes()
	for _, categ := range plc.Category {
		addingCategory := woocommerce.MeCat{
			Id:   categ.ID,
			Name: categ.Name,
			Slug: categ.Slug,
		}
		NodeCategoryes.Add(categ.Parent, addingCategory)
	}

	// Аттрибуты
	attr, errAttr := userWC.ProductsAttributes()
	if errAttr != nil {
		return nil, errAttr
	}
	idAttrColor, isFind_AttrColor := attr.Find_id_of_name("Цвет")
	if isFind_AttrColor != nil {
		return nil, isFind_AttrColor
	}
	idAttrSize, isFind_AttrSize := attr.Find_id_of_name("Размер")
	if isFind_AttrSize != nil {
		return nil, isFind_AttrSize
	}
	idManuf, isFind_AttrManuf := attr.Find_id_of_name("Производитель")
	if isFind_AttrManuf != nil {
		return nil, isFind_AttrManuf
	}

	Cat3 := make(map[int]*Category3Base)
	Cat3[0] = &Category3Base{}
	Cat3[0].Cat3 = make(map[int]*Category3Base)

	return &WcAdd{
		WooClient:      wooClient,
		UserWC:         userWC,
		Tags:           tags,
		TagMap:         tagMap,
		NodeCategoryes: NodeCategoryes,
		IdAttrColor:    idAttrColor,
		IdAttrSize:     idAttrSize,
		IdManuf:        idManuf,
		Delivery:       Delivery,
		Plc:            plc,
		Cat3:           Cat3,
		Tr:             tr,
	}, nil
}

// Функция добавления товара
func (woo *WcAdd) AddProduct(product bases.Product2) error {

	if product.Article == "" {
		return errors.New("AddProduct: нет в товаре артикула")
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

	/*
		ManufrId, ManufName, ManufSlug := AddAttr(woo.WooClient, woo.IdAttrColor, "Производитель", product.Manufacturer)
		fmt.Println("Для данного товара Аттрибуты Производителя:", ManufrId, ManufName, ManufSlug)

		// Создаём аттрибуты товара для цвета
		for key := range product.Item {
			tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug := AddAttr(woo.WooClient, woo.IdAttrColor, product.Item[key].ColorEng, key)
			fmt.Println("Для данного товара Аттрибуты цвета будут:", tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug)
		}
		// Создаём аттрибуты товара для Размера
		for _, valSize := range product.Size {
			tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug := AddAttr(woo.WooClient, woo.IdAttrSize, valSize, bases.FormingColorEng(valSize))
			fmt.Println("Для данного товара Аттрибуты размера будут:", tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug)
		}
	*/
	/**/

	// Собираем гендер для загрузки в теги товара
	idGender, _, isGenderSlug := bases.GenderBook(product.GenderLabel, "")
	if !isGenderSlug {
		fmt.Print("Не найден гендер.", idGender)
	}
	fmt.Println(" Гендр: " + idGender + ". ")

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
			if chet == 0 {
				imageInput = append(imageInput, entity.ProductImage{
					Src:  valueImage,
					Name: valueImage + strconv.Itoa(indexImage) + ".jpg",
					Alt:  valueImage + strconv.Itoa(indexImage),
				})
			}
			imageInput = append(imageInput, entity.ProductImage{
				Src:  valueImage,
				Name: valueImage + strconv.Itoa(indexImage) + ".jpg",
				Alt:  valueImage + strconv.Itoa(indexImage),
			})
			chet++
		}
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

		MetaData: []entity.Meta{ // Ссылка на товар
			{
				Key:   "linkRB",
				Value: product.Link,
			},
		},

		Images: imageInput,

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
		fmt.Println("Product.Create:", errCreate)
		return errCreate
	}
	itemID = item.ID
	fmt.Println("Product.Create: Done itemID =", itemID)

	// Редактирвоание вариационных товаров.
	//
	// Создаём массив из обновлений вариационных товаров, в частности создании новых вариаций товаров.
	//
	// Загружаем их одним запросом [Batch]
	//
	// [Batch]: https://woocommerce.github.io/woocommerce-rest-api-docs/#batch-update-product-variations
	var CreateVariations []wc.BatchProductVariationsCreateItem
	for colorKey, colorItemValue := range product.Item {
		fmt.Println("colorItemValue.Size", colorItemValue.Size)
		for sizeKey, SizeValue := range colorItemValue.Size {
			fmt.Printf("ProductVariation.Create[%v:%v]: Добавляю вар. товар с цветом '%v' и размером '%v'\n", colorKey+1, sizeKey+1, colorItemValue.ColorEng, SizeValue.Val)

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
			if len(colorItemValue.Image) != 0 {
				CreateVariation.Image = &entity.ProductImage{
					Src:  colorItemValue.Image[0],
					Name: colorItemValue.ColorEng + ".jpg",
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
		fmt.Println("Error Add variation:", ErrBatch)
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
