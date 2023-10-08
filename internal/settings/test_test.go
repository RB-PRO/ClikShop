package settings

import (
	"fmt"
	"testing"
)

func TestRussianSymb(t *testing.T) {
	fmt.Println(RussianSymb("qweasdzxc"))
	fmt.Println(RussianSymb("йцуфывячс"))

	str := "Короткая джинсовая куртка"
	str = "3 Parçalı Pamuklu Set"

	fmt.Println(RussianSymb(str))

	fmt.Println(str, RussianSymb(str), "<", len([]rune(str))/2, RussianSymb(str) < len([]rune(str))/2)

}
