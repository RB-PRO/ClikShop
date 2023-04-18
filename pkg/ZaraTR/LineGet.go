package zaratr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Загрузить все товары категории
func LoadLine(id string) (lin Line, ErrorLine error) {

	// Делаем запрос на получение категорий
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(LineURL, id), nil)
	if err != nil {
		return Line{}, err
	}

	// Добавляем необходимые атрибуты
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("accept-encoding", "application/json; charset=utf-8")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("if-none-match", "W/\"c2b35-nPn7rjU78OjmncJadKH5VqRw6b8\"")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.895 Yowser/2.5 Safari/537.36")

	// Выполнить запрос
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Line{}, err
	}

	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return Line{}, fmt.Errorf("LoadLine: wrong status code: %d", res.StatusCode)
	}

	// Читаем ответ в массив байтов
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Line{}, err
	}

	// Декодируем полученный json и получаем данные
	err = json.Unmarshal(bodyBytes, &lin)
	if err != nil {
		return Line{}, err
	}

	return lin, nil
}
