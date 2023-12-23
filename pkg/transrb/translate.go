package transrb

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// Экземпляр переводчика с внутренними данными
type Translate struct {
	FolderID   string // Идентификатор https://console.cloud.yandex.ru/cloud/b1gok977v247t4a0f6qk
	IAM        string // Токен IAM https://cloud.yandex.ru/docs/iam/operations/iam-token/create
	OAuthToken string // https://cloud.yandex.ru/docs/iam/operations/iam-token/create
}

// Стукртура ответа сервера
type Response struct {
	// Done:
	Translations []struct {
		Text                 string `json:"text"`
		DetectedLanguageCode string `json:"detectedLanguageCode"`
	} `json:"translations"`

	// Error:
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details []struct {
		Type      string `json:"@type"`
		RequestID string `json:"requestId"`
	} `json:"details"`
}

// Ошибка при которой входные данные неверны
var ErrorInput error = errors.New("Translate: New: Invalid input data")

// Создать экземпляр переводчика
func New(FolderID, OAuthToken string) (*Translate, error) {
	IAM_token, ErrorIAM := IAM(OAuthToken)

	if ErrorIAM != nil {
		return nil, ErrorIAM
	}

	if FolderID == "" || IAM_token == "" || OAuthToken == "" {
		return nil, ErrorInput
	}
	return &Translate{FolderID: FolderID, OAuthToken: OAuthToken, IAM: IAM_token}, nil
}

// Перевести
func (cli *Translate) Trans(InputStrs []string) ([]string, error) {
	// Формируем запрос
	type TransPost struct {
		FolderID           string   `json:"folderId"`
		Texts              []string `json:"texts"`
		TargetLanguageCode string   `json:"targetLanguageCode"`
	}
	RawData := TransPost{
		FolderID:           cli.FolderID,
		Texts:              InputStrs,
		TargetLanguageCode: "ru",
	}

	// Конвертация из структуры в буффер
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(RawData)
	if err != nil {
		log.Fatal(err)
	}

	// Создаём запрос
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, "https://translate.api.cloud.yandex.net/translate/v2/translate", &buf) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+cli.IAM)

	// Выполняем запрос
	resp, ErrorDo := client.Do(r)
	if ErrorDo != nil {
		return nil, ErrorDo
	}
	defer resp.Body.Close()

	// Вывод результатов
	body, ErrReadAll := io.ReadAll(resp.Body) // response body is []byte
	if ErrReadAll != nil {
		return nil, ErrReadAll
	}

	// fmt.Println(string(body))

	// Распарсить ответ
	var result Response
	if ErrUnmarshal := json.Unmarshal(body, &result); ErrUnmarshal != nil { // Parse []byte to go struct pointer
		return nil, ErrUnmarshal
	}

	// Если есть сообщение об ошибке
	if result.Message != "" {
		return nil, errors.New(result.Message)
	}

	// ЗАписываем массив на вывод
	var OutPuts []string
	for _, val := range result.Translations {
		OutPuts = append(OutPuts, val.Text)
	}

	return OutPuts, nil
}

func (cli *Translate) TransENG(InputStrs string) (string, error) {
	// Формируем запрос
	type TransPost struct {
		FolderID           string   `json:"folderId"`
		Texts              []string `json:"texts"`
		TargetLanguageCode string   `json:"targetLanguageCode"`
	}
	RawData := TransPost{
		FolderID:           cli.FolderID,
		Texts:              []string{InputStrs},
		TargetLanguageCode: "en",
	}

	// Конвертация из структуры в буффер
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(RawData)
	if err != nil {
		log.Fatal(err)
	}

	// Создаём запрос
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, "https://translate.api.cloud.yandex.net/translate/v2/translate", &buf) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+cli.IAM)

	// Выполняем запрос
	resp, ErrorDo := client.Do(r)
	if ErrorDo != nil {
		return "", ErrorDo
	}
	defer resp.Body.Close()

	// Вывод результатов
	body, ErrReadAll := io.ReadAll(resp.Body) // response body is []byte
	if ErrReadAll != nil {
		return "", ErrReadAll
	}

	// fmt.Println(string(body))

	// Распарсить ответ
	var result Response
	if ErrUnmarshal := json.Unmarshal(body, &result); ErrUnmarshal != nil { // Parse []byte to go struct pointer
		return "", ErrUnmarshal
	}

	// Если есть сообщение об ошибке
	if result.Message != "" {
		return "", errors.New(result.Message)
	}

	// // ЗАписываем массив на вывод
	// var OutPuts []string
	// for _, val := range result.Translations {
	// 	OutPuts = append(OutPuts, val.Text)
	// }

	return result.Translations[0].Text, nil
}
