package plavatar

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/jonasdoesthings/plavatar/v3/utils"
	"math"
	"math/rand"
)

func (generator *Generator) Laughing(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *Options) {
	startColor := utils.RandomColorHex(rng)
	rng.Seed(rngSeed + 128)
	stopColor := utils.RandomColorHex(rng)

	canvas.Def()
	gradientColors := []svg.Offcolor{{0, startColor, 1}, {100, stopColor, 1}}
	canvas.LinearGradient("bg", 0, 0, 100, 100, gradientColors)
	canvas.DefEnd()

	DrawCanvasBackground(canvas, options)

	rightEyePositionY := -utils.RandomRangeInt(rng, 0, CanvasSize/6)
	leftEyePositionY := -utils.RandomRangeInt(rng, 0, CanvasSize/6)

	leftEyePositionX := -utils.RandomRangeInt(rng, CanvasSize/4, int(math.Round(CanvasSize/2*0.6)))
	rightEyePositionX := utils.RandomRangeInt(rng, CanvasSize/4, int(math.Round(CanvasSize/2*0.6)))

	mouthPositionY := CanvasSize / 5

	eyeSize := CanvasSize/9 + int(utils.RandomRangeFloat(rng, -CanvasSize/50, CanvasSize/80))

	rng.Seed(rngSeed + 256)

	canvas.Circle(leftEyePositionX, leftEyePositionY, eyeSize, "fill: white")
	canvas.Circle(rightEyePositionX, rightEyePositionY, eyeSize, "fill: white")
	canvas.Arc(
		leftEyePositionX+utils.RandomRangeInt(rng, 0, 1),
		mouthPositionY,
		155,
		200,
		180+utils.RandomRangeInt(rng, -10, +10),
		false,
		false,
		rightEyePositionX+utils.RandomRangeInt(rng, 0, 1),
		mouthPositionY,
		fmt.Sprintf("fill: white; transform: rotate(%d); stroke-width: %d; stroke-linecap: round;", utils.RandomRangeInt(rng, -5, 5), utils.RandomRangeInt(rng, 20, 24)),
	)
}
