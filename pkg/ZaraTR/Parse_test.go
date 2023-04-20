package zaratr

import (
	"fmt"
	"testing"
)

func TestCatCycle(t *testing.T) {
	Items := CatCycle() // Наполнить цикл
	fmt.Println(len(Items))
}
