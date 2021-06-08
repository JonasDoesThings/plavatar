package utils

import (
	"image/color"
	"math/rand"
)

var ColorWhite = &color.RGBA{R: 255, G: 255, B: 255, A: 255}

func GetRandomColor(rng *rand.Rand) color.Color {
	return &color.RGBA{R: uint8(rng.Intn(255)), G: uint8(rng.Intn(255)), B: uint8(rng.Intn(255)), A: 255}
}
