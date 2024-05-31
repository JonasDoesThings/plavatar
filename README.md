# plavatar
A library for generating **pla**ceholder **avatar**s (=plavatars).

`plavatar`: A library for generating **pla**ceholder **avatar**s (=plavatars).  
`plavatar-rest` A stateless REST microservice wrapping plavatar for you (docker image available) (https://github.com/jonasdoesthings/plavatar-rest)

![docs/assets/readme-demo.png](docs/assets/readme-demo.png)

## Install
`go get github.com/jonasdoesthings/plavatar/v3`  
Then you can import the `"github.com/jonasdoesthings/plavatar/v3"` package.

## Usage
Full Docs: https://pkg.go.dev/github.com/jonasdoesthings/plavatar/v3

Basic Example with a built-in generatorFunc:
```go
import (
    "bytes"
    "github.com/jonasdoesthings/plavatar/v3"
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
    "github.com/jonasdoesthings/plavatar/v3"
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

## **If possible, use format=SVG.**
Not only is format=SVG extremely faster, if you transfer the image to your user, SVG also saves you a lot of bandwidth and latency (A generated SVG is only ~2% the size of a 512px PNG)

## Testing
To run the go tests, use `go test -v ./...` in the root directory of the project.

## Support my work
<a href="https://www.buymeacoffee.com/JonasDoesThings" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>  

If this library helps you, you can use https://www.buymeacoffee.com/JonasDoesThings to support my work.
