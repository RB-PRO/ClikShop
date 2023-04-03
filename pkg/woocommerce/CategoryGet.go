package woocommerce

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Categorys struct {
	Category []ProductListCategory
}

type ProductListCategory struct {
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
		Up []struct {
			Href string `json:"href"`
		} `json:"up"`
	} `json:"_links,omitempty"`
	IsAdd  bool `json:"-"` // Фиксирование того, что это
	IsAdd3 bool `json:"-"` // Если true, то такой товар добавлен в категорию товаров
}

// Метод [product/categories] позволяет Вам извлекать все категории продуктов.
//
// # Использую для создания структуры всех категорий
//
// [Orders]: http://woocommerce.github.io/woocommerce-rest-api-docs/?shell#list-all-product-categories
func (user *User) ProductsCategories() (Categorys, error) {
	var category Categorys // Структура по категории
	var TotalPages int = 2 // /!!!

	for i := 1; i <= TotalPages; i++ {
		var categ []ProductListCategory
		var bodyBytes []byte
		var errData error
		bodyBytes, TotalPages, errData = user.queringProductsCategories("GET", "/products/categories", i)
		if errData != nil {
			return Categorys{}, errData
		}
		errUnmarshal := json.Unmarshal(bodyBytes, &categ)
		if errUnmarshal != nil { // Если ошибка
			return Categorys{}, errUnmarshal
		}
		category.Category = append(category.Category, categ...)
	}
	// Если всё верно сработало
	return category, nil
}

// Ядро запроса с входными параметрами:
// - methodURL - Метод GET, POST, PUT, ...
// - methodSite - Метод API
// - data - Массив byte с передаваемыми данным
func (user *User) queringProductsCategories(methodURL, methodApi string, page int) ([]byte, int, error) {
	var TotalPages int
	client := &http.Client{}
	req, errReq := http.NewRequest(methodURL, URL+REQ+methodApi+"?per_page=100&page="+strconv.Itoa(page), nil)
	if errReq != nil {
		return nil, 0, errReq
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(user.consumer_key, user.secret_key))
	res, errRes := client.Do(req)
	if errRes != nil {
		return nil, 0, errRes
	}
	defer res.Body.Close()

	for name, values := range res.Header {
		if name == "X-Wp-Totalpages" {
			var errorTotalPages error
			TotalPages, errorTotalPages = strconv.Atoi(values[0])
			if errorTotalPages != nil {
				return nil, TotalPages, errorTotalPages
			}
		}
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		return nil, 0, errBody
	}

	return body, TotalPages, nil
}
