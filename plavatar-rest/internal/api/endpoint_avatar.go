package api

import (
	svg "github.com/ajstarks/svgo"
	"github.com/jonasdoesthings/plavatar"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func (server *Server) HandleGetAvatar(generatorFunc func(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *plavatar.Options)) echo.HandlerFunc {
	return func(context echo.Context) error {
		outputSize, err := strconv.Atoi(context.Param("size"))
		if err != nil || outputSize < minSize || outputSize > maxSize {
			return context.Blob(http.StatusBadRequest, "application/json", []byte(`{"error": "invalid size"}`))
		}

		outputFormat := plavatar.PNG
		mimeType := "image/png"
		if strings.ToLower(context.QueryParam("format")) == "svg" {
			outputFormat = plavatar.SVG
			mimeType = "image/svg+xml"
		}

		outputShape := plavatar.Circle
		if strings.ToLower(context.QueryParam("shape")) == "square" {
			outputShape = plavatar.Square
		}

		generatedAvatar, rngSeed, err := server.avatarGenerator.GenerateAvatar(generatorFunc, &plavatar.Options{
			Name:         context.Param("name"),
			OutputSize:   outputSize,
			OutputFormat: outputFormat,
			OutputShape:  outputShape,
		})

		context.Response().Header().Add("Rng-Seed", rngSeed)

		if err != nil {
			return context.Blob(http.StatusInternalServerError, "application/json", []byte(err.Error()))
		}

		return context.Blob(http.StatusOK, mimeType, generatedAvatar.Bytes())
	}
}
