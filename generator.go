// Package plavatar implements the core avatar generation functionality.
//
// It contains the [Generator] struct, which implements the library's main method: [Generator.GenerateAvatar].
//
// [Generator.GenerateAvatar] is called in combination with a GeneratorFunction like [Generator.Smiley]
// (see the matching avatar_XXX.go file for details on the GeneratorFunction's implementations)
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

// A Generator is used to generate avatars using its [Generator.GenerateAvatar] method.
type Generator struct{}

// Shape the output image should have.
type Shape = int

const (
	ShapeCircle Shape = iota // Instructs the generator to return a circle-shaped avatar. (default)
	ShapeSquare              // Instructs the generator to return a square-shaped avatar.
)

// Format is the file-format the output image should be encoded in.
type Format = int

const (
	FormatPNG Format = iota
	FormatSVG
)

// Options contains the generation instructions like seed (Name) or OutputSize, passed to the generation method
type Options struct {
	Name         string
	OutputSize   int
	OutputFormat Format
	OutputShape  Shape
}

func getAvatarCanvas(targetWriter io.Writer) *svg.SVG {
	canvas := svg.New(targetWriter)
	canvas.Startview(CanvasSize, CanvasSize, -CanvasSize/2, -CanvasSize/2, CanvasSize, CanvasSize)

	return canvas
}

// DrawCanvasBackground fills the canvas with the fitting background.
// Important: when using this method, the canvas must already contain definitions for the bg color gradient.
//
// Example definition in your custom generatorFunc:
//
//	func MyCustomGenerator(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *Options) {
//		backgroundColor := utils.RandomColorHex(rng)
//
//		canvas.Def()
//		gradientColors := []svg.Offcolor{{0, backgroundColor, 1}}
//		canvas.LinearGradient("bg", 0, 0, 100, 100, gradientColors)
//		canvas.DefEnd()
//
//		DrawCanvasBackground(canvas, options)
//		...
//	}
//
// See avatar_solid.go's source code for a full example on how to define the bg color.
func DrawCanvasBackground(canvas *svg.SVG, options *Options) {
	if options.OutputShape == ShapeSquare {
		canvas.Square(-CanvasSize/2, -CanvasSize/2, CanvasSize, "fill: url(#bg)")
	} else {
		canvas.Circle(0, 0, CanvasSize/2, "fill: url(#bg)")
	}
}

// RasterizeSVGToPNG rasterizes the SVG file to a PNG image of the given imageSize in the form of a [bytes.Buffer].
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

func hashString(s string) (int64, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0, err
	}
	return int64(h.Sum32()), nil
}

func getRNGFromName(name string) (*rand.Rand, int64, string, error) {
	var rawSeed string
	var seed int64

	if name != "" {
		rawSeed = name
	} else {
		rawSeed = strconv.FormatInt(rand.Int63n(2147483647), 10)
	}

	seed, err := hashString(rawSeed)
	if err != nil {
		return nil, -1, rawSeed, errors.New("failed hashing name")
	}

	rng := rand.New(rand.NewSource(seed))
	return rng, seed, rawSeed, nil
}

// GenerateAvatar generates an avatar by setting-up the image canvas and then calling the passed generatorFunc.
// It uses the passed generatorOptions to instruct the avatar generation.
//
// The passed generatorFunc can either be a built-in one like [Generator.Smiley], [Generator.Solid], or a custom one written by you.
//
// A successful generation, returns err == nil, a string with the used rng seed, and a buffer filled with the image data.
//
// Usage Example:
//
//	func generateMyAvatar() (*bytes.Buffer, string) {
//		avatarGenerator := plavatar.Generator{}
//		options := &plavatar.Options{
//			Name:         "exampleSeed",
//			OutputSize:   256,
//			OutputFormat: plavatar.FormatSVG,
//			OutputShape:  plavatar.ShapeSquare,
//		}
//		avatar, rngSeed, err := avatarGenerator.GenerateAvatar(avatarGenerator.Smiley, options)
//		if err != nil {
//			panic(err)
//		}
//
//		return avatar, rngSeed
//	}
func (generator *Generator) GenerateAvatar(generatorFunc func(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *Options), generatorOptions *Options) (*bytes.Buffer, string, error) {
	imageBuffer := bytes.NewBuffer([]byte{})
	svgCanvas := getAvatarCanvas(imageBuffer)
	rng, rngSeed, rawSeed, err := getRNGFromName(generatorOptions.Name)
	if err != nil {
		return nil, rawSeed, err
	}

	if generatorOptions.OutputSize < 1 && generatorOptions.OutputFormat != FormatSVG {
		return nil, rawSeed, errors.New("invalid size")
	}

	generatorFunc(svgCanvas, rng, rngSeed, generatorOptions)
	svgCanvas.End()

	if generatorOptions.OutputFormat == FormatSVG {
		return imageBuffer, rawSeed, nil
	}

	pngBuffer, err := RasterizeSVGToPNG(imageBuffer, generatorOptions.OutputSize)
	if err != nil {
		return nil, rawSeed, errors.New("error encoding image to png")
	}

	return pngBuffer, rawSeed, nil
}
