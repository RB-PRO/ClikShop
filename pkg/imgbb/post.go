package imgbb

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const URL string = "https://api.imgbb.com/1/upload"

type ImgbbUser struct {
	API_key string // API Ключ сервиса https://api.imgbb.com/
}

func NewImgbbUser(key string) *ImgbbUser {
	return &ImgbbUser{key}
}

// Стуктура ответа от загрузки на сервер
type ImgbbResponse struct {
	Data struct {
		ID         string `json:"id"`
		Title      string `json:"title"`
		URLViewer  string `json:"url_viewer"`
		URL        string `json:"url"`
		DisplayURL string `json:"display_url"`
		Width      int    `json:"width"`
		Height     int    `json:"height"`
		Size       int    `json:"size"`
		Time       int    `json:"time"`
		Expiration int    `json:"expiration"`
		Image      struct {
			Filename  string `json:"filename"`
			Name      string `json:"name"`
			Mime      string `json:"mime"`
			Extension string `json:"extension"`
			URL       string `json:"url"`
		} `json:"image"`
		Thumb struct {
			Filename  string `json:"filename"`
			Name      string `json:"name"`
			Mime      string `json:"mime"`
			Extension string `json:"extension"`
			URL       string `json:"url"`
		} `json:"thumb"`
		DeleteURL string `json:"delete_url"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`

	StatusCode int `json:"status_code"`
	Error      struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error"`
	StatusTxt string `json:"status_txt"`
}

// Функция загрузки изображения на сервер.
func (img ImgbbUser) Upload(pictureBase64, name string, expiration int) (ImgbbResponse, error) {
	// Ответ от сервера
	var imgbbRes ImgbbResponse

	var UrlValue url.Values
	if expiration != 0 {
		UrlValue = url.Values{"key": {img.API_key}, "image": {pictureBase64}, "name": {name}, "expiration": {strconv.Itoa(expiration)}}
	} else {
		UrlValue = url.Values{"key": {img.API_key}, "image": {pictureBase64}, "name": {name}}
	}

	// Выполнить запрос
	resp, responseError := http.PostForm(URL, UrlValue)
	if responseError != nil {
		return ImgbbResponse{}, responseError
	}
	defer resp.Body.Close()

	//Считываем ответ запроса
	body, bodyRead := io.ReadAll(resp.Body)
	if bodyRead != nil {
		return ImgbbResponse{}, bodyRead
	}

	fmt.Println(string(body))

	// Распарсить данные
	responseErrorUnmarshal := json.Unmarshal(body, &imgbbRes)
	if responseErrorUnmarshal != nil {
		return ImgbbResponse{}, responseErrorUnmarshal
	}

	return imgbbRes, responseError
}

// Convert Picture in base64
func PicToBase64(filename string) (string, error) {
	imgFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()

	fInfo, fInfoError := imgFile.Stat()
	if fInfoError != nil {
		return "", fInfoError
	}
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	_, ReadError := fReader.Read(buf)
	if ReadError != nil {
		return "", ReadError
	}

	return base64.StdEncoding.EncodeToString(buf), nil
}
