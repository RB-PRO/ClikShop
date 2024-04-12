package trendyol

import (
	"fmt"
	"testing"
)

func TestPage(t *testing.T) {
	pg, err := ParsePage(332585, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(pg.Result.Products))
}
