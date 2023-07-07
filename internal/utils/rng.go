package utils

import (
	"fmt"
	"math/rand"
)

func RandomColorHex(rng *rand.Rand) string {
	// rng.Intn(256) would be correct, since n is exclusive, but to not change all previously generated avatars
	// using the same seed, we will keep this fault in.
	return fmt.Sprintf("#%02X%02X%02X", rng.Intn(255), rng.Intn(255), rng.Intn(255))
}

func RandomRangeInt(rng *rand.Rand, min, max int) int {
	// rng.Intn(max-min+1) would be correct, since n is exclusive, but to not change all previously generated avatars
	// using the same seed, we will keep this fault in.
	return min + rng.Intn(max-min)
}

func RandomRangeFloat(rng *rand.Rand, min, max float64) float64 {
	return min + rng.Float64()*(max-min)
}
