package avatars

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"math"
	"math/rand"
	"plavatar/internal/utils"
)

func (generator *Generator) Happy(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *Options) {
	startColor := utils.GetRandomColorHex(rng)
	rng.Seed(rngSeed + 128)
	stopColor := utils.GetRandomColorHex(rng)

	canvas.Def()
	gradientColors := []svg.Offcolor{{0, startColor, 1}, {100, stopColor, 1}}
	canvas.LinearGradient("bg", 0, 0, 100, 100, gradientColors)
	canvas.DefEnd()

	DrawCanvasBackground(canvas, options)

	eyePositionY := -utils.RandomRangeInt(rng, CanvasSize/6, CanvasSize/4)
	rightEyePositionY := eyePositionY + utils.RandomRangeInt(rng, -CanvasSize/20, CanvasSize/20)
	leftEyePositionY := eyePositionY + utils.RandomRangeInt(rng, -CanvasSize/20, CanvasSize/20)

	leftEyePositionX := -utils.RandomRangeInt(rng, CanvasSize/4, int(math.Round(CanvasSize/2*0.6)))
	rightEyePositionX := utils.RandomRangeInt(rng, CanvasSize/4, int(math.Round(CanvasSize/2*0.6)))

	mouthPositionY := CanvasSize / 5

	eyeSizeX := CanvasSize / 14
	eyeSizeY := CanvasSize / 10

	rng.Seed(rngSeed + 256)

	canvas.Polyline(
		[]int{leftEyePositionX - eyeSizeX, leftEyePositionX, leftEyePositionX + eyeSizeX},
		[]int{leftEyePositionY + eyeSizeY, leftEyePositionY, leftEyePositionY + eyeSizeY},
		fmt.Sprintf("fill: none; stroke: white; transform: rotate(%d); stroke-width: %d; stroke-linecap: round; stroke-linejoin: round;", utils.RandomRangeInt(rng, -5, 5), utils.RandomRangeInt(rng, 15, 20)),
	)
	canvas.Polyline(
		[]int{rightEyePositionX - eyeSizeX, rightEyePositionX, rightEyePositionX + eyeSizeX},
		[]int{rightEyePositionY + eyeSizeY, rightEyePositionY, rightEyePositionY + eyeSizeY},
		fmt.Sprintf("fill: none; stroke: white; transform: rotate(%d); stroke-width: %d; stroke-linecap: round; stroke-linejoin: round;", utils.RandomRangeInt(rng, -5, 5), utils.RandomRangeInt(rng, 15, 20)),
	)
	canvas.Arc(
		leftEyePositionX+utils.RandomRangeInt(rng, 40, 80),
		mouthPositionY,
		150,
		200,
		180+utils.RandomRangeInt(rng, -10, +10),
		false,
		false,
		rightEyePositionX+utils.RandomRangeInt(rng, -80, -40),
		mouthPositionY,
		fmt.Sprintf("fill: none; stroke: white; transform: rotate(%d); stroke-width: %d; stroke-linecap: round;", utils.RandomRangeInt(rng, -5, 5), utils.RandomRangeInt(rng, 20, 24)),
	)
}
