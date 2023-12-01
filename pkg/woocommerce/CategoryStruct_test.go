package woocommerce_test

import (
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/woocommerce"
)

func NewTestNode() *woocommerce.Node {
	Node := woocommerce.NewCategoryes()
	Node.Add(0, woocommerce.MeCat{Id: 1, Name: "Man", Slug: "man"})
	Node.Add(0, woocommerce.MeCat{Id: 2, Name: "Women", Slug: "women"})

	Node.Add(1, woocommerce.MeCat{Id: 3, Name: "Пиджаки", Slug: "jacket"})
	Node.Add(3, woocommerce.MeCat{Id: 4, Name: "Пиджаки с рукавами", Slug: "jacket-with-rukav"})
	Node.Add(4, woocommerce.MeCat{Id: 5, Name: "Сударь", Slug: "sudar"})

	Node.Add(2, woocommerce.MeCat{Id: 6, Name: "Шубы", Slug: "fur"})
	Node.Add(6, woocommerce.MeCat{Id: 7, Name: "Норковые", Slug: "iz-norki"})
	Node.Add(7, woocommerce.MeCat{Id: 8, Name: "Остин", Slug: "ostine"})

	Node.Add(3, woocommerce.MeCat{Id: 9, Name: "Рубашки", Slug: "rubaha"})
	Node.Add(9, woocommerce.MeCat{Id: 10, Name: "H&M", Slug: "hm"})

	Node.Add(2, woocommerce.MeCat{Id: 11, Name: "Флисовые", Slug: "flis"})
	Node.Add(11, woocommerce.MeCat{Id: 12, Name: "GJ", Slug: "gj"})
	return Node
}
func TestFindId(t *testing.T) {
	Node := NewTestNode() // Создать дерево категорий
	//Node.PrintInorderName("-")
	if tecNode, booleans := Node.FindId(7); !booleans {
		t.Error("Не найдена категория по ID")
	} else {
		if tecNode.Id != 7 {
			t.Error("ID Категории не равно 7")
		}
		if tecNode.Name != "Норковые" {
			t.Error("Name Категории не равно Норковые")
		}
		if tecNode.Slug != "iz-norki" {
			t.Error("Slug Категории не равно iz-norki")
		}
		if tecNode.ParentID != 6 {
			t.Error("Родительский ID не равен 6")
		}
	}

	if tecNode, booleans := Node.FindId(11); !booleans {
		t.Error("Не найдена категория по ID")
	} else {
		if tecNode.Id != 11 {
			t.Error("ID Категории не равно 7")
		}
		if tecNode.Name != "Флисовые" {
			t.Error("Name Категории не равно Флисовые")
		}
		if tecNode.Slug != "flis" {
			t.Error("Slug Категории не равно flis")
		}
		if tecNode.ParentID != 2 {
			t.Error("Родительский ID не равен 6")
		}
	}
}

func TestLen(t *testing.T) {
	Node := NewTestNode() // Создать дерево категорий
	//Node.PrintInorderName("-")
	lens := Node.Len()
	if lens != 12 {
		t.Error("Длина не соответствует. Получен ответ -", lens)
	}
}
