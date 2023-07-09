package plavatar

import (
	svg "github.com/ajstarks/svgo"
	"github.com/jonasdoesthings/plavatar/utils"
	"math/rand"
)

func (generator *Generator) Marble(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *Options) {
	startColor := utils.RandomColorHex(rng)
	rng.Seed(rngSeed + 128)
	stopColor := utils.RandomColorHex(rng)

	canvas.Def()
	gradientColors := []svg.Offcolor{{0, stopColor, 1}, {100, startColor, 1}}
	canvas.RadialGradient("bg", 50, 50, 100, 50, 50, gradientColors)
	canvas.DefEnd()

	DrawCanvasBackground(canvas, options)
}
