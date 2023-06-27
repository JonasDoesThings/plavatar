package api

import (
	"bytes"
	svg "github.com/ajstarks/svgo"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"plavatar/internal/avatars"
	"strings"
)

func (server *Server) HandleGetAvatar(generatorFunc func(canvas *svg.SVG, rng *rand.Rand, rngSeed int64, options *avatars.Options)) echo.HandlerFunc {
	return func(context echo.Context) error {
		imageBuffer := bytes.NewBuffer([]byte{})
		svgCanvas := server.avatarGenerator.GetAvatarCanvas(imageBuffer)
		rng, rngSeed := server.getRNGFromRequest(context)

		avatarShape := avatars.Circle
		if strings.ToLower(context.QueryParam("shape")) == "square" {
			avatarShape = avatars.Square
		}

		generatorFunc(svgCanvas, rng, rngSeed, &avatars.Options{
			OutputShape: avatarShape,
		})
		svgCanvas.End()

		if strings.ToLower(context.QueryParam("format")) == "svg" {
			return context.Blob(http.StatusOK, "image/svg+xml", imageBuffer.Bytes())
		}

		imageSize := server.getSizeFromRequest(context)
		if imageSize < 1 {
			return context.Blob(http.StatusBadRequest, "application/json", []byte(`{"error": "invalid size"}`))
		}

		pngBuffer, err := server.avatarGenerator.RasterizeSVGToPNG(imageBuffer, imageSize)
		if err != nil {
			server.logger.Error("error encoding image to png", err)
			return context.Blob(http.StatusInternalServerError, "application/json", []byte(`{"error": "error encoding image to png"}`))
		}

		return context.Blob(http.StatusOK, "image/png", pngBuffer.Bytes())
	}
}
