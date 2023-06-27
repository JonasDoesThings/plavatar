package avatars

import (
	"bytes"
	svg "github.com/ajstarks/svgo"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"github.com/srwiley/scanFT"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"io"
)

type Generator struct {
}

type Options struct {
	OutputShape Shape
}

type Shape = int

const (
	Circle Shape = iota
	Square
)

const CanvasSize = 512

func (generator *Generator) GetAvatarCanvas(targetWriter io.Writer) *svg.SVG {
	canvas := svg.New(targetWriter)
	canvas.Startview(CanvasSize, CanvasSize, -CanvasSize/2, -CanvasSize/2, CanvasSize, CanvasSize)

	return canvas
}

func (generator *Generator) DrawCanvasBackground(canvas *svg.SVG, options *Options) {
	if options.OutputShape == Square {
		canvas.Square(-CanvasSize/2, -CanvasSize/2, CanvasSize, "fill: url(#bg)")
	} else {
		canvas.Circle(0, 0, CanvasSize/2, "fill: url(#bg)")
	}
}

func (generator *Generator) RasterizeSVGToPNG(svg io.Reader, imageSize int) (*bytes.Buffer, error) {
	icon, err := oksvg.ReadIconStream(svg, oksvg.WarnErrorMode)
	if err != nil {
		return nil, err
	}

	icon.SetTarget(0, 0, CanvasSize, CanvasSize)
	rgba := image.NewRGBA(image.Rect(0, 0, CanvasSize, CanvasSize))
	icon.Draw(rasterx.NewDasher(imageSize, imageSize, scanFT.NewScannerFT(imageSize, imageSize, scanFT.NewRGBAPainter(rgba))), 1)

	if imageSize != CanvasSize {
		scaledOutput := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))
		draw.NearestNeighbor.Scale(scaledOutput, scaledOutput.Bounds(), rgba, rgba.Bounds(), draw.Over, nil)
		rgba = scaledOutput
	}

	outBuffer := bytes.NewBuffer([]byte{})
	err = png.Encode(outBuffer, rgba)
	if err != nil {
		return nil, err
	}

	return outBuffer, nil
}
