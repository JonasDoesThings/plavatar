package utils

import (
	"fmt"
	"math/rand"
)

func RandomColorHex(rng *rand.Rand) string {
	return fmt.Sprintf("#%02X%02X%02X", rng.Intn(255), rng.Intn(255), rng.Intn(255))
}

func RandomRangeInt(rng *rand.Rand, min, max int) int {
	return min + rng.Intn(max-min)
}

func RandomRangeFloat(rng *rand.Rand, min, max float64) float64 {
	return min + rng.Float64()*(max-min)
}
