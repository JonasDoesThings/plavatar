package plavatar

import (
	"bytes"
	"errors"
	svg "github.com/ajstarks/svgo"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"github.com/srwiley/scanFT"
	"golang.org/x/image/draw"
	"hash/fnv"
	"image"
	"image/png"
	"io"
	"math/rand"
	"strconv"
)

const CanvasSize = 512

type Generator struct{}

type Shape = int

const (
	Circle Shape = iota
	Square
)

type Format = int

const (
	PNG Format = iota
	SVG
)

type Options struct {
	Name         string
	OutputSize   int
	OutputFormat Format
	OutputShape  Shape
}

func GetAvatarCanvas(targetWriter io.Writer) *svg.SVG {
	canvas := svg.New(targetWriter)
	canvas.Startview(CanvasSize, CanvasSize, -CanvasSize/2, -CanvasSize/2, CanvasSize, CanvasSize)

	return canvas
}

func DrawCanvasBackground(canvas *svg.SVG, options *Options) {
	if options.OutputShape == Square {
		canvas.Square(-CanvasSize/2, -CanvasSize/2, CanvasSize, "fill: url(#bg)")
	} else {
		canvas.Circle(0, 0, CanvasSize/2, "fill: url(#bg)")
	}
}

func RasterizeSVGToPNG(svg io.Reader, imageSize int) (*bytes.Buffer, error) {
	icon, err := oksvg.ReadIconStream(svg, oksvg.WarnErrorMode)
	if err != nil {
		return nil, err
	}

	icon.SetTarget(0, 0, CanvasSize, CanvasSize)
	rgba := image.NewRGBA(image.Rect(0, 0, CanvasSize, CanvasSize))
	icon.Draw(rasterx.NewDasher(CanvasSize, CanvasSize, scanFT.NewScannerFT(CanvasSize, CanvasSize, scanFT.NewRGBAPainter(rgba))), 1)

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

func HashString(s string) (int64, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0, err
	}
	return int64(h.Sum32()), nil
}

func GetRNGFromName(name string) (*rand.Rand, int64, string, error) {
	var rawSeed string
	var seed int64

	if name != "" {
		rawSeed = name
	} else {
		rawSeed = strconv.FormatInt(rand.Int63n(2147483647), 10)
	}

	seed, err := HashString(rawSeed)
	if err != nil {
		return nil, -1, rawSeed, errors.New(`{"error": "hashing name"}`)
	}

	rng := rand.New(rand.NewSource(seed))
	return rng, seed, rawSeed, nil
}

func (generator *Generator) GenerateAvatar(generatorFunc func(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *Options), generatorOptions *Options) (*bytes.Buffer, string, error) {
	imageBuffer := bytes.NewBuffer([]byte{})
	svgCanvas := GetAvatarCanvas(imageBuffer)
	rng, rngSeed, rawSeed, err := GetRNGFromName(generatorOptions.Name)
	if err != nil {
		return nil, rawSeed, err
	}

	if generatorOptions.OutputSize < 1 {
		return nil, rawSeed, errors.New(`{"error": "invalid size"}`)
	}

	generatorFunc(svgCanvas, rng, rngSeed, generatorOptions)
	svgCanvas.End()

	if generatorOptions.OutputFormat == SVG {
		return imageBuffer, rawSeed, nil
	}

	pngBuffer, err := RasterizeSVGToPNG(imageBuffer, generatorOptions.OutputSize)
	if err != nil {
		return nil, rawSeed, errors.New(`{"error": "error encoding image to png"}`)
	}

	return pngBuffer, rawSeed, nil
}
