package utils

import (
	"math/rand"
)

func RandomRangeInt(rng *rand.Rand, min, max int) int {
	return min + rng.Int()*(max-min)
}

func RandomRangeFloat(rng *rand.Rand, min, max float64) float64 {
	return min + rng.Float64()*(max-min)
}
