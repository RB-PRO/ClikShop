package wcprod

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

// Структура конфигурационного файла
// Данные, которые необходимы для запуска приложения
type Config struct {
	ConsumerKey string `json:"consumer_key"`
	SecretKey   string `json:"secret_key"`
	FolderID    string `json:"folder_id"`
	OAuthToken  string `json:"oauth_token"`
	Proxy       string `json:"proxy"`
	Token       string `json:"token"`
	YandexToken string `json:"yandex_token"`
	Imgbb       string `json:"imgbb"`

	IKPrivateKey  string `json:"ik_private_key"`
	IKPublicKey   string `json:"ik_public_key"`
	IKUrlEndpoint string `json:"ik_url_endpoint"`
}

// Загрузить данные из файла
func LoadConfig(filename string) (config Config, ErrorFIle error) {
	// Открыть файл
	jsonFile, ErrorFIle := os.Open(filename)
	if ErrorFIle != nil {
		return config, ErrorFIle
	}
	defer jsonFile.Close()

	// Прочитать файл и получить массив byte
	jsonData, ErrorFIle := io.ReadAll(jsonFile)
	if ErrorFIle != nil {
		return config, ErrorFIle
	}

	// Распарсить массив byte в структуру
	if ErrorFIle := json.Unmarshal(jsonData, &config); ErrorFIle != nil {
		return config, ErrorFIle
	}
	return config, ErrorFIle
}

// Проверить данные на корректность
func (config *Config) IsСorrect() (bool, error) {
	if config.ConsumerKey == "" {
		return false, errors.New("IsСorrect: Нет данных в config.ConsumerKey")
	}
	if config.SecretKey == "" {
		return false, errors.New("IsСorrect: Нет данных в config.SecretKey")
	}
	if config.FolderID == "" {
		return false, errors.New("IsСorrect: Нет данных в config.FolderID")
	}
	if config.OAuthToken == "" {
		return false, errors.New("IsСorrect: Нет данных в config.OAuthToken")
	}
	if config.Proxy == "" {
		return false, errors.New("IsСorrect: Нет данных в config.Proxy")
	}
	if config.Token == "" {
		return false, errors.New("IsСorrect: Нет данных в config.Token")
	}
	if config.YandexToken == "" {
		return false, errors.New("IsСorrect: Нет данных в config.YandexToken")
	}
	if config.Imgbb == "" {
		return false, errors.New("IsСorrect: Нет данных в config.Imgbb")
	}
	return true, nil
}
