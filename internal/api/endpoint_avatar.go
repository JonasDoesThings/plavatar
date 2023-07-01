package api

import (
	svg "github.com/ajstarks/svgo"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"plavatar/internal/avatars"
	"strconv"
	"strings"
)

func (server *Server) HandleGetAvatar(generatorFunc func(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *avatars.Options)) echo.HandlerFunc {
	return func(context echo.Context) error {
		outputSize, err := strconv.Atoi(context.Param("size"))
		if err != nil || outputSize < minSize || outputSize > maxSize {
			return context.Blob(http.StatusBadRequest, "application/json", []byte(`{"error": "invalid size"}`))
		}

		outputFormat := avatars.PNG
		if strings.ToLower(context.QueryParam("format")) == "svg" {
			outputFormat = avatars.SVG
		}

		outputShape := avatars.Circle
		if strings.ToLower(context.QueryParam("shape")) == "square" {
			outputShape = avatars.Square
		}

		generatedAvatar, rngSeed, err := server.avatarGenerator.GenerateAvatar(generatorFunc, &avatars.Options{
			Name:         context.Param("name"),
			OutputSize:   outputSize,
			OutputFormat: outputFormat,
			OutputShape:  outputShape,
		})

		context.Response().Header().Add("RngSeed", rngSeed)
		context.Response().WriteHeader(http.StatusOK)

		if err != nil {
			return context.Blob(http.StatusInternalServerError, "application/json", []byte(err.Error()))
		}

		return context.Blob(http.StatusOK, "image/png", generatedAvatar.Bytes())
	}
}
