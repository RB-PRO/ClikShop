package transrb

import (
	translate "github.com/sergei-svistunov/go-yandex-translate"
)

type Translate struct {
	tr *translate.Translate
}

func New() (*Translate, error) {
	yt := translate.New("t1.9euelZqZzZiMnpmSj4qPzo_Hz5mVjO3rnpWay5LJkY2QipGby5WSzs-QlM3l9Pd2MmVd-e8LfWeR3fT3NmFiXfnvC31nkQ.d21DBeuA2TfRVdsMc9RTHihMhq1yUE49olmPYPEw4N-wTnv05mf8xMg0Wxrs4MvXtgMew6sA1yKHFfzJ2TtAAg")

	return &Translate{yt}, nil
}

func (cli *Translate) Trans(inputStr string) (string, error) {
	texts, err := cli.tr.Translate([]string{"Test"}, "en-ru", "plain")
	if err != nil {
		return "", err
	}

	return texts[0], nil
}
