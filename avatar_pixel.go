package plavatar

import (
	svg "github.com/ajstarks/svgo"
	"github.com/jonasdoesthings/plavatar/v3/utils"
	"math/rand"
)

// Pixels
// TODO:
// Currently this always outputs a [ShapeCircle] in PNG output mode.
// oksvg does not support clippaths at the moment, so the whole square gets rasterized
// see: https://github.com/srwiley/oksvg/issues/10
func (generator *Generator) Pixels(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *Options) {
	pixels := CanvasSize / 8

	canvas.Group("clip-path=\"url(#clip)\"")
	for x := -CanvasSize / 2; x < CanvasSize/2; x += pixels {
		for y := -CanvasSize / 2; y < CanvasSize/2; y += pixels {
			canvas.Rect(x, y, pixels, pixels, "fill:"+utils.RandomColorHex(rng))
		}
	}

	canvas.Gend()

}
