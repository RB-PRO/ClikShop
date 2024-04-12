package woocommerce_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/bases"
	"github.com/RB-PRO/ClikShop/pkg/woocommerce"
)

func TestAddCat2(t *testing.T) {
	consumer_key, _ := dataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := dataFile("secret_key")     // Секретный код пользователя
	// Авторизация
	userWC, _ := woocommerce.New(consumer_key, secret_key)
	// [{Name:Men Slug:men} {Name:Clothing Slug:clothing} {Name:Hoodies & Sweatshirts Slug:hoodies-sweatshirts} {Name:'47 NHL Slug:47-nhl}]

	categ := make([]bases.Cat, 4)
	categ[0].Name = "Test0"
	categ[0].Slug = "test0"
	categ[1].Name = "Test1"
	categ[1].Slug = "test1"
	categ[2].Name = "Test2"
	categ[2].Slug = "test2"
	categ[3].Name = "Test3"
	categ[3].Slug = "test3"

	// Получить дерева категорий
	plc, errPLC := userWC.ProductsCategories()
	if errPLC != nil {
		t.Error(errPLC)
	}
	fmt.Println("Сделал массив товаров plc, длиной", len(plc.Category))

	fmt.Println("Добавляю впервые")
	NewId, AddNewId := userWC.AddCat2(&plc, categ)
	if AddNewId != nil {
		t.Error(AddNewId)
	}
	fmt.Println("Новый добавленный ID1 =", NewId)
	fmt.Println("Длина массива plc =", len(plc.Category))

	fmt.Println("Добавляю второй раз категорию")
	NewId2, AddNewId2 := userWC.AddCat2(&plc, categ)
	if AddNewId2 != nil {
		t.Error(AddNewId2)
	}
	fmt.Println("Новый добавленный ID2 =", NewId2)
	fmt.Println("Длина массива plc =", len(plc.Category))

	if NewId2 != NewId {
		t.Error("NewId2 и NewId не совпали, что означает, что категория создатся в другом месте")
	}

	categ2 := make([]bases.Cat, 4)
	categ2[0].Name = "Test0"
	categ2[0].Slug = "test0"
	categ2[1].Name = "Test11"
	categ2[1].Slug = "test11"
	categ2[2].Name = "Test2"
	categ2[2].Slug = "test2"
	categ2[3].Name = "Test3"
	categ2[3].Slug = "test3"
	fmt.Println("Добавляю впервые")
	NewId3, AddNewId := userWC.AddCat2(&plc, categ)
	if AddNewId != nil {
		t.Error(AddNewId)
	}
	fmt.Println("Новый добавленный ID3 =", NewId3)
}
