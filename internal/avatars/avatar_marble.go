package avatars

import (
	svg "github.com/ajstarks/svgo"
	"math/rand"
	"plavatar/internal/utils"
)

func (generator *Generator) Marble(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *Options) {
	startColor := utils.GetRandomColorHex(rng)
	rng.Seed(rngSeed + 128)
	stopColor := utils.GetRandomColorHex(rng)

	canvas.Def()
	gradientColors := []svg.Offcolor{{0, stopColor, 1}, {100, startColor, 1}}
	canvas.RadialGradient("bg", 50, 50, 100, 50, 50, gradientColors)
	canvas.DefEnd()

	generator.DrawCanvasBackground(canvas, options)
}
