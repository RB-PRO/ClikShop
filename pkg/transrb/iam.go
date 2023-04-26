package transrb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func TRTR(Translate []string) {
	// y0_AgAAAAANmH2FAATuwQAAAADdzO0BpgX9OWjQTlKHbutb6RmAe2HmEII

	OAuthToken := "y0_AgAAAAANmH2FAATuwQAAAADdzO0BpgX9OWjQTlKHbutb6RmAe2HmEII"
	IAM_token, ErrorIAM := IAM(OAuthToken)
	if ErrorIAM != nil {
		fmt.Println(ErrorIAM)
	}

	fmt.Println(IAM_token)
	//  "AQVN21irl_QaeKLFY7fExYIMZ8Y1pvYU19Ci4504"

}

// Ошибка нулевого
var ErrorTokenNil error = errors.New("IAM: nil Token")

// Получить токен IAM в соответствии с https://cloud.yandex.ru/docs/iam/operations/iam-token/create
func IAM(OAuthToken string) (string, error) {

	// Структура ответа от сервера
	type iam_struct struct {
		IamToken  string    `json:"iamToken"`
		ExpiresAt time.Time `json:"expiresAt"`
	}

	// Создаём запрос
	req, ErrReq := http.NewRequest(http.MethodPost, "https://iam.api.cloud.yandex.net/iam/v1/tokens",
		strings.NewReader(`{"yandexPassportOauthToken":"`+OAuthToken+`"}`))
	if ErrReq != nil {
		return "", ErrReq
	}
	req.Header.Add("Content-Type", "Application/json")

	// Выполняем запрос
	client := &http.Client{}
	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return "", ErrDo
	}
	defer res.Body.Close()

	// Читаем ответ запроса
	body, ErrRead := io.ReadAll(res.Body)
	if ErrRead != nil {
		return "", ErrRead
	}

	//fmt.Println(string(body))

	// Распарсить ответ
	var result iam_struct
	if ErrUnmarshal := json.Unmarshal(body, &result); ErrUnmarshal != nil { // Parse []byte to go struct pointer
		return "", ErrUnmarshal
	}

	// Проверка на нулевой ответ от сервера
	if result.IamToken == "" {
		return "", ErrorTokenNil
	}

	return result.IamToken, nil
}
