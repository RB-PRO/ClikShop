package woocommerce

import (
	"encoding/json"
	"errors"

	"github.com/RB-PRO/ClikShop/pkg/bases"
)

// Структура ответа API на создание категории
type AddCatResponse struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Parent      int         `json:"parent"`
	Description string      `json:"description"`
	Display     string      `json:"display"`
	Image       interface{} `json:"image"`
	MenuOrder   int         `json:"menu_order"`
	Count       int         `json:"count"`
	Links       struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
	} `json:"_links"`

	// Если ошибка, то заполняются эти данные:
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status     int `json:"status"`
		ResourceID int `json:"resource_id"`
	} `json:"data"`
}

// Метод [create-a-product-category] поможет Вам добавить категорию
//
// # Использую для добавления категории товаров
// Возвращает ID категории
//
// [create-a-product-category]: http://woocommerce.github.io/woocommerce-rest-api-docs/?shell#create-a-product-category
func (user *User) AddCat_WC(valetCat MeCat) (int, error) {

	// Сделать json добавления категории
	bytesData, errMarshal := json.Marshal(valetCat)
	if errMarshal != nil {
		return 0, errMarshal
	}

	// Выполнить запрос
	bodyBytes, errData := user.quering("POST", "/products/categories", bytesData)
	if errData != nil {
		return 0, errData
	}

	// Получить ответ
	var AddCatRes AddCatResponse
	errUnmarshal := json.Unmarshal(bodyBytes, &AddCatRes)
	if errUnmarshal != nil { // Если ошибка распарсивания в структуру данных
		return 0, errors.New("AddCat_WC: Не удалось распарсить ответ сервера: " + string(bodyBytes))
	}
	if AddCatRes.Code == "term_exists" { // Обработка случая, когда существует такая категория
		return AddCatRes.Data.ResourceID, nil
	}

	// Если всё верно сработало и произошло добавление
	return AddCatRes.ID, nil
}

// Функция добавления категории с обновлением домашней структуры данных
//
// Используется в качестве внешнего интерфейса для добавления категории товара по методике - добавил - проверил - получил ID
func (user *User) AddCat(NodeCategoryes *Node, NewCategory []bases.Cat) (CatIDcreate int, err error) {
	// var CatIDcreate int // ID новой или старой категории
	for _, NewCat := range NewCategory {
		findNode, findNodeBool := NodeCategoryes.FindSlug(NewCat.Slug)
		if !findNodeBool { // Если категория не добавлена
			// То добавляем её в WC
			// Добавить категорию на WP
			CatIDcreate, err = user.AddCat_WC(MeCat{
				Name:        NewCat.Name,
				Slug:        NewCat.Slug,
				Description: "Создано автоматически при добавлении товара",
			})
			if err != nil {
				return 0, err
			}

			// Добавляем в дерево категорий
			NodeCategoryes.Add(findNode.Id, MeCat{
				Id:   CatIDcreate,
				Name: NewCat.Name,
				Slug: NewCat.Slug,
			})

		} else {
			CatIDcreate = findNode.Id
		}
	}
	// fmt.Println("ID новой актуальной категории товара - ", CatIDcreate)
	return CatIDcreate, nil
}

// ******************************************************************

// Создать дерево категорий по полученому массиву категорий
func (plc Categorys) FormingNode(NodeCategoryes *Node) (*Node, error) {
	var errorAddWithParent error
	NodeCategoryes, errorAddWithParent = plc.addWithParent(NodeCategoryes, 0)
	if errorAddWithParent != nil {
		return nil, errorAddWithParent
	}

	// Цикл по потомкам
	for i := 0; i < 4; i++ {
		for indexCat, cat := range plc.Category {
			if !plc.Category[indexCat].IsAdd {
				FindNode, errorFInd := NodeCategoryes.find(cat.ID)
				if errorFInd == nil { // Если найдено
					errorAddNode := NodeCategoryes.Add(FindNode.ParentID, MeCat{
						Id:       cat.ID,
						Name:     cat.Name,
						Slug:     cat.Slug,
						ParentID: FindNode.ParentID,
					})
					if errorAddNode != nil {
						return nil, errorAddNode
					}
					plc.Category[indexCat].IsAdd = true
				}
			}
		}
	}
	return NodeCategoryes, nil
}

// Добавить данные в категорию, учитывая ID Родительского элемента,
// в котором и происходит добавление категории
func (plc *Categorys) addWithParent(NodeCategoryes *Node, parent int) (*Node, error) {
	for _, categ := range plc.Category { // Цикл по структуре категории WC
		if categ.Parent == parent { // Если это тот самый родитель
			errorAddNode := NodeCategoryes.Add(parent, MeCat{
				Id:   categ.ID,
				Name: categ.Name,
				Slug: categ.Slug,
			})
			if errorAddNode != nil {
				return nil, errorAddNode
			}
		}
	}
	return NodeCategoryes, nil
}

// Проблема в том, что когда я добавляю товары через эту залупень, то товары не учитывают,
// что они могут отноиться к категории, которая ещё не добавлена в мою структуру.
// Поэтому программа падает в случае, когда добавлена 4-я категория, с родителем в 3 категории, который ещё не добавлен
// Решение: добавление в структуру товаров с родителем ID 0.
