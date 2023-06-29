package utils

import (
	"fmt"
	"math/rand"
)

func GetRandomColorHex(rng *rand.Rand) string {
	return fmt.Sprintf("#%02X%02X%02X", rng.Intn(255), rng.Intn(255), rng.Intn(255))
}
