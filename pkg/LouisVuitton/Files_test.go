package louisvuitton

import "testing"

func TestFolder(t *testing.T) {
	dir := NewDir("")
	errMakedir := dir.MakeDir("REAL/REAL")
	if errMakedir != nil {
		t.Error(errMakedir)
	}
}
