package utils

import (
	"math/rand"
	"strings"
	"testing"
)

func TestRandomColorHex(t *testing.T) {
	rng := rand.New(rand.NewSource(123))
	hexColor := RandomColorHex(rng)

	if !strings.HasPrefix(hexColor, "#") {
		t.Fatal("hexColor not starting with #")
	}

	if len(hexColor) != 7 {
		t.Fatal("hexColor not 7 chars (# + 3x2 digits) long")
	}

	// using the given seed (123) the first generated color should always be this one:
	// this makes sure that the same avatar is returned for the same input seed each time
	if hexColor != "#D7093A" {
		t.Error("returned hexColor not deterministic. is this intended?")
	}
}
