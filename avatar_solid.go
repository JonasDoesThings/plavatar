package plavatar

import (
	svg "github.com/ajstarks/svgo"
	"github.com/jonasdoesthings/plavatar/v3/utils"
	"math/rand"
)

func (generator *Generator) Solid(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *Options) {
	backgroundColor := utils.RandomColorHex(rng)

	canvas.Def()
	gradientColors := []svg.Offcolor{{0, backgroundColor, 1}}
	canvas.LinearGradient("bg", 0, 0, 100, 100, gradientColors)
	canvas.DefEnd()

	DrawCanvasBackground(canvas, options)
}
