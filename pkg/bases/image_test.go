package bases_test

import (
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

func TestWebp2Jpg(t *testing.T) {
	bases.Webp2Jpg("local.webp", "local.jpg")
}
