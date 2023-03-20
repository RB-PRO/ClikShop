package transrb_test

import (
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/transrb"
)

func TestTrans(t *testing.T) {

	inputStr := "Hello world"
	outputStr := "Привет мир"

	answerTranslate, errorTranslate := transrb.Trans(inputStr)
	if errorTranslate != nil {
		t.Error(errorTranslate)
	}
	if outputStr != answerTranslate {
		t.Error("Неверный перевод. Получено: " + answerTranslate)
	}

	/*
		yt := translate.New("t1.9euelZqMys6SlcuLnMaRzsyNyJHPju3rnpWay5LJkY2QipGby5WSzs-QlM3l9PcAICFf-e8eeDXF3fT3QE4eX_nvHng1xQ.u6HjMY2PJsJMaqFgHBqunCpQCiW7xrfzlO9JF6sQhD9eUV3letPSpTkse2KMPfOTyRvClof2HMax_fvZeMJFBg") // get the key from https://translate.yandex.com/developers/keys

		texts, err := yt.Translate([]string{"Test", "Hello"}, "en-ru", "plain")
		fmt.Println(err)
		fmt.Println(texts)
	*/
}
