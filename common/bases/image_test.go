package bases_test

import (
	"testing"

	"ClikShop/common/bases"
)

func TestWebp2Jpg(t *testing.T) {
	bases.Webp2Jpg("local.webp", "local.jpg")
}
