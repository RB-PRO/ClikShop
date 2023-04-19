package zaratr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Загрузить все товары категории
func LoadTouch(id string) (tou Touch, ErrorLine error) {

	// Делаем запрос на получение категорий
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(TouchURL, id), nil)
	if err != nil {
		return Touch{}, err
	}

	// Добавляем необходимые атрибуты
	req.Header.Add("accept-encoding", "application/json; charset=utf-8")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.895 Yowser/2.5 Safari/537.36")

	// Выполнить запрос
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Touch{}, err
	}

	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return Touch{}, fmt.Errorf("LoadTouch: wrong status code: %d", res.StatusCode)
	}

	// Читаем ответ в массив байтов
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Touch{}, err
	}

	// Декодируем полученный json и получаем данные
	err = json.Unmarshal(bodyBytes, &tou)
	if err != nil {
		return Touch{}, err
	}

	return tou, nil
}
