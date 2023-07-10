# plavatar
A library for generating **pla**ceholder **avatar**s (=plavatars).

`plavatar`: A library for generating **pla**ceholder **avatar**s (=plavatars). 
`plavatar-rest` A stateless REST microservice wrapping plavatar for you (docker image available) (https://github.com/jonasdoesthings/plavatar-rest)

![docs/assets/readme-demo.png](docs/assets/readme-demo.png)

## Install
`go get -u github.com/jonasdoesthings/plavatar`  
Then you can import the `"github.com/jonasdoesthings/plavatar"` package.

## Usage
Full Docs: https://pkg.go.dev/github.com/jonasdoesthings/plavatar

Basic Example with a built-in generatorFunc:
```go
import (
    "bytes"
    "github.com/jonasdoesthings/plavatar"
)

func generateMyAvatar() (*bytes.Buffer, string) {
    // Set-up a plavatar Generator instance
    avatarGenerator := plavatar.Generator{}
    
    // Configure the plavatar you want to generate 
    options := &plavatar.Options{
        Name:         "exampleSeed", // the seed to use
        OutputShape:  plavatar.ShapeSquare, // ShapeSquare or ShapeCircle
        OutputFormat: plavatar.FormatSVG,
        // OR if you want a PNG with the size of 512x512px:
        // OutputFormat: plavatar.FormatPNG,
        // OutputSize: 512,
    }
    
    // generate the avatar using the built-in Smiley generatorFunc and pass the options from above
    avatar, rngSeed, err := avatarGenerator.GenerateAvatar(avatarGenerator.Smiley, options)
    if err != nil {
        panic(err)
    }

    // returns the avatar as *bytes.Buffer and the used rngSeed as string
    return avatar, rngSeed
}
```

The plavatar Generator uses a passed generatorFunc to generate the avatar graphic.
The generatorFunc takes a svg canvas, a rng, the used rngSeed, and the generation options.
The generatorFunc then modifies the passed svg canvas.

Basic example with a custom generatorFunc:
```go
import (
    "bytes"
    svg "github.com/ajstarks/svgo"
    "github.com/jonasdoesthings/plavatar"
    "math/rand"
)

func CustomAvatar(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *plavatar.Options) {
    canvas.Def()
    gradientColors := []svg.Offcolor{{0, "#FF0000", 1}}
    canvas.LinearGradient("bg", 0, 0, 100, 100, gradientColors)
    canvas.DefEnd()

    plavatar.DrawCanvasBackground(canvas, options)
    canvas.Line(-100, -10, 100, 10, "stroke: white; stroke-width: 23")
}

func generateMyCustomAvatar() (*bytes.Buffer, string) {
    avatarGenerator := plavatar.Generator{}
    options := &plavatar.Options{
        Name:         "exampleSeed",
        OutputSize:   256,
        OutputFormat: plavatar.FormatSVG,
        OutputShape:  plavatar.ShapeSquare,
    }
    avatar, rngSeed, err := avatarGenerator.GenerateAvatar(CustomAvatar, options)
    if err != nil {
        panic(err)
    }

    return avatar, rngSeed
}
```

## Testing
To run the go tests, use `go test -v ./...` in the root directory of the project.
