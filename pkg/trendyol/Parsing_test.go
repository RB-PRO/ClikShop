package trendyol

import (
	"fmt"
	"testing"
)

func TestPage(t *testing.T) {
	pg, err := ParsePage("106871", 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(pg.Result.Products))
}
