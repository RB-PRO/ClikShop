package settings

import (
	"fmt"
	"testing"

	"ClikShop/common/apibitrix"
)

func TestGetdatafromsizes(t *testing.T) {
	bx, _ := apibitrix.NewBitrixUser()
	fmt.Println(getdatafromsizes(bx, "418079"))
}
