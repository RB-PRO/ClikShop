package clicker

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RB-PRO/ClikShop/pkg/wcprod"
	"github.com/cheggaaa/pb"
)

// Получить [Ссылки]+[Ссылки на источники] на все товары
func Click() {
	// token, ErrorDataLoad := bases.DataFile("token")
	// if ErrorDataLoad != nil {
	// 	log.Fatalln(ErrorDataLoad)
	// }

	// not, ErrorNotif := notification.NewNotification(token, "-768253730", "Отчёт кликера", "Clicker")
	// if ErrorNotif != nil {
	// 	log.Fatalln(ErrorNotif)
	// }

	Adding, errorInitWcAdd := wcprod.New2() // Создаём экземпляр загрузчика данных
	if errorInitWcAdd != nil {
		log.Fatalln(errorInitWcAdd)
	}

	fmt.Println("Получаем все ссылки на товары")
	// not.Sends("Получаем все ссылки на товары")
	// Получаем все ссылки
	hands, ErrorHand := Hands(Adding)
	if ErrorHand != nil {
		log.Fatalln(ErrorHand)
	}

	// Кликаем все-все товары
	bar := pb.StartNew(len(hands))
	client := http.Client{Timeout: 10 * time.Second}
	for _, hand := range hands {
		client.Get(hand.URL)
		bar.Increment()
	}
	bar.Finish()

}
