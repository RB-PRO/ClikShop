package actualizer

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/ClikShop/pkg/cbbank"
)

func TestUploadProduct(t *testing.T) {
	bx, ErrNewActualizer := NewActualizer()
	if ErrNewActualizer != nil {
		panic(fmt.Errorf("gol.NewGol: %v", ErrNewActualizer))
	}
	cb, ErrorCB := cbbank.New() // Цены ЦБ для получение актуального курса
	if ErrorCB != nil {
		panic(ErrorCB)
	}
	bx.BX.CB = cb

	// ***

	_, ErrCoasts := bx.BX.Coasts()
	if ErrCoasts != nil {
		panic(ErrCoasts)
	}

	Folder := "foldertest"
	ErrPush := bx.Push(Folder)
	if ErrPush != nil {
		bx.GLOG.Err(fmt.Sprintf("%v: bx.ErrPush: %v", Folder, ErrPush))
		return
	}

}
