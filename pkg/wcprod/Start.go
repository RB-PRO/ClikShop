// Отдельно вынесенный пакет для загрузки товаров на WordPress
// Использует кривую библиотеку.
// Кривая она из-за того, что некоторые параметры не соотносятся с документацией Woocommerce.
// Для упрощения написания кода, локально исправил некоторые строки в скаченной библиотеке. В идеале локально развернуть библиотеку и провести необходимые манипуляции.
package wcprod

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RB-PRO/ClikShop/pkg/imgbb"
	"github.com/RB-PRO/ClikShop/pkg/transrb"
	"github.com/RB-PRO/ClikShop/pkg/woocommerce"
	wc "github.com/hiscaler/woocommerce-go"
	config "github.com/hiscaler/woocommerce-go/config"
	"github.com/hiscaler/woocommerce-go/entity"
	"github.com/imagekit-developer/imagekit-go"
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
	Imgbb          *imgbb.ImgbbUser       // Сервис картинок
	IK             *imagekit.ImageKit     // Сервис картинок 2 - https://imagekit.io/dashboard/developer/api-keys
	LogFile        *os.File               // Логгирование

	// ID аттрибутов в WordPress.
	IdAttrColor int
	IdAttrSize  int
	IdManuf     int

	Delivery map[string]int // Мапа цен доставки для товаров

	WooClient *wc.WooCommerce // Клиент пользовательской библиотеки, с помощью которой добавляю товар

	Cat3 map[int]*Category3Base // Мапа категории

	AllProdSKU map[string]bool // Список всех товаров по Артиклу
}

// Полная инициализация базовой структуры загрузки товара
func New() (*WcAdd, error) {

	// Логгирование
	FileLog := "logs/" + time.Now().Format("15h04m 02Jan2006") + ".log"
	LogFile, ErrLogFile := os.OpenFile(FileLog, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644) // open log file
	if ErrLogFile != nil {
		return nil, fmt.Errorf("wcprod: New:  %w", ErrLogFile)
	}
	defer LogFile.Close()
	log.SetOutput(LogFile)                       // set log out put
	log.SetFlags(log.Lshortfile | log.LstdFlags) // optional: log date-time, filename, and line number

	// Загрузка конфига
	ConfigRBFileName := "config_rb.json"
	ConfigRB, ErrOpenConfigRB := LoadConfig(ConfigRBFileName)
	if ErrOpenConfigRB != nil {
		return nil, fmt.Errorf("wcprod: New: Read config file '%s' error: %v", ConfigRBFileName, ErrOpenConfigRB)
	}
	log.Println("New: Загрузил конфиг " + ConfigRBFileName)

	// Клиент от сторонней библиотеки(пользовательской)
	ConfigFileName := "config_wp.json"
	b, ErrReadFile := os.ReadFile(ConfigFileName)
	if ErrReadFile != nil {
		return nil, errors.New("wcprod: New: Read config error: " + ErrReadFile.Error())
	}
	log.Println("New: Загрузил конфиг " + ConfigFileName)

	var c config.Config
	ErrUnmarshal := json.Unmarshal(b, &c)
	if ErrUnmarshal != nil {
		return nil, errors.New("wcprod: New: Parse config file error: " + ErrUnmarshal.Error())
	}
	wooClient := wc.NewClient(c)
	log.Println("New: wc.NewClient: Распарсил конфиг " + ConfigFileName)

	// Мой клиент
	userWC, _ := woocommerce.New(ConfigRB.ConsumerKey, ConfigRB.SecretKey) // Авторизация
	if okErr := userWC.IsOrder(); okErr != nil {                           // Проверка на авторизацию
		return nil, okErr
	}
	log.Println("New: woocommerce.New: Создал клиент userWC")

	// Мапа цены доставки
	Delivery, ErrorDelivery := XlsxDelivery()
	if ErrorDelivery != nil {
		return nil, ErrorDelivery
	}
	log.Println("New: XlsxDelivery: Загрузил xlsx доставки")

	// Юзер для работы с сервисом катинок
	imgbbUser := imgbb.NewImgbbUser(ConfigRB.Imgbb)
	log.Println("New: NewImgbbUser: Создаю пользователя imgbb")

	// Теги
	tags, tagsError := userWC.AllTags_WC()
	if tagsError != nil {
		return nil, tagsError
	}
	log.Println("New: AllTags_WC: Получаю теги tags")

	// Создать Мапу тэгов
	tagMap := woocommerce.MapTags(tags)

	// Получить дерево категорий
	plc, errPLC := userWC.ProductsCategories()
	if errPLC != nil {
		return nil, errPLC
	}
	log.Println("New: ProductsCategories: Получаю дерево категорий")

	// Переводчик
	tr, ErrTranslate := transrb.New(ConfigRB.FolderID, ConfigRB.OAuthToken)
	if ErrTranslate != nil {
		return nil, ErrTranslate
	}
	log.Println("New: transrb.New: Создаю пользователя Яндекс переводчка")

	// Using environment variables IMAGEKIT_PRIVATE_KEY, IMAGEKIT_PUBLIC_KEY and IMAGEKIT_ENDPOINT_URL
	// ik, err := imagekit.New()

	// Using keys in argument
	IK := imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey:  ConfigRB.IKPrivateKey,
		PublicKey:   ConfigRB.IKPublicKey,
		UrlEndpoint: ConfigRB.IKUrlEndpoint,
	})
	log.Println("New: imagekit.NewFromParams: Создаю пользователя imagekit")

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
	log.Println("New: woocommerce.NewCategoryes: Создаю дерево категорий")

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
	log.Println("New: userWC.ProductsAttributes: Получить аттрибусы цвета, размера, производителя")

	Cat3 := make(map[int]*Category3Base)
	Cat3[0] = &Category3Base{}
	Cat3[0].Cat3 = make(map[int]*Category3Base)

	// Получить все артикулы товаров
	AllProdSKU := make(map[string]bool)
	var IsLastPages bool
	var TecalPage int
	for !IsLastPages {
		var ErrorAllProd error
		var items []entity.Product
		items, _, _, IsLastPages, ErrorAllProd = wooClient.Services.Product.All(wc.ProductsQueryParams{}, TecalPage, 100)
		if ErrorAllProd != nil {
			return nil, ErrorAllProd
		}
		for _, prod := range items {
			// AllProdSKU = append(AllProdSKU, prod.SKU)
			AllProdSKU[prod.SKU] = true
		}
		TecalPage++
	}
	log.Println("New: wooClient.Services.Product.All: Получены все SKU товаров")

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
		AllProdSKU:     AllProdSKU,
		Imgbb:          imgbbUser,
		IK:             IK,
		LogFile:        LogFile,
	}, nil
}

// Вторичная инициализация без переводчика и списка всех товаров, цены доставки, тегов и прочего.
func New2() (*WcAdd, error) {

	ConfigRBFileName := "config_rb.json"
	ConfigRB, ErrOpenConfigRB := LoadConfig(ConfigRBFileName)
	if ErrOpenConfigRB != nil {
		return nil, fmt.Errorf("wcprod: New: Read config file '%s' error: %v", ConfigRBFileName, ErrOpenConfigRB.Error())
	}

	// Клиент от сторонней библиотеки(пользовательской)
	b, ErrReadFile := os.ReadFile("config_wp.json")
	if ErrReadFile != nil {
		return nil, errors.New("wcprod: New: Read config error: " + ErrReadFile.Error())
	}

	var c config.Config
	ErrUnmarshal := json.Unmarshal(b, &c)
	if ErrUnmarshal != nil {
		return nil, errors.New("wcprod: New: Parse config file error: " + ErrUnmarshal.Error())
	}
	wooClient := wc.NewClient(c)

	// Мой клиент
	userWC, _ := woocommerce.New(ConfigRB.ConsumerKey, ConfigRB.SecretKey) // Авторизация
	if okErr := userWC.IsOrder(); okErr != nil {                           // Проверка на авторизацию
		return nil, okErr
	}

	return &WcAdd{
		WooClient: wooClient,
		UserWC:    userWC,
	}, nil
}

// Новый клиент, содержащий ТОЛЬКО переводчик
// На входе кушает только мой локальный файл "config_rb.json"
func NewTranslate() (*WcAdd, error) {

	// Загрузка конфига
	ConfigRBFileName := "config_rb.json"
	ConfigRB, ErrOpenConfigRB := LoadConfig(ConfigRBFileName)
	if ErrOpenConfigRB != nil {
		return nil, fmt.Errorf("wcprod: New: Read config file '%s' error: %v", ConfigRBFileName, ErrOpenConfigRB.Error())
	}

	// Переводчик
	tr, ErrTranslate := transrb.New(ConfigRB.FolderID, ConfigRB.OAuthToken)
	if ErrTranslate != nil {
		return nil, ErrTranslate
	}
	log.Println("New: transrb.New: Создаю пользователя Яндекс переводчка")

	return &WcAdd{
		Tr: tr,
	}, nil
}
