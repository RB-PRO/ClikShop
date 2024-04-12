package settings

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/apibitrix"
)

func TestGetdatafromsizes(t *testing.T) {
	bx, _ := apibitrix.NewBitrixUser()
	fmt.Println(getdatafromsizes(bx, "418079"))
}
