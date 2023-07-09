package avatars

import (
	svg "github.com/ajstarks/svgo"
	"github.com/jonasdoesthings/plavatar/internal/utils"
	"math/rand"
)

func (generator *Generator) Gradient(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *Options) {
	startColor := utils.RandomColorHex(rng)
	rng.Seed(rngSeed + 128)
	stopColor := utils.RandomColorHex(rng)

	canvas.Def()
	gradientColors := []svg.Offcolor{{0, startColor, 1}, {100, stopColor, 1}}
	canvas.LinearGradient("bg", 0, 0, 100, 100, gradientColors)
	canvas.DefEnd()

	DrawCanvasBackground(canvas, options)
}
