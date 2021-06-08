package api

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"plavatar/internal/utils"
)

func (server *Server) HandleGetPixelAvatar() echo.HandlerFunc {
	return func(context echo.Context) error {
		imageContext, err := server.getAvatarImageContext(context)
		if err != nil {
			return err
		}
		if imageContext == nil {
			return nil
		}

		size := imageContext.Image().Bounds().Max.X
		pixelSize := size / 8
		pixels := size / pixelSize

		name := context.Param("name")
		seed := int64(rand.Intn(2147483647))
		if name != "" {
			seed = int64(server.hashString(name))
		}
		rng := rand.New(rand.NewSource(seed))

		for i := 0; i < pixels; i++ {
			for j := 0; j < pixels; j++ {
				imageContext.SetColor(utils.GetRandomColor(rng))
				imageContext.DrawRectangle(float64(i*pixelSize), float64(j*pixelSize), float64(pixelSize), float64(pixelSize))
				imageContext.Fill()
			}
		}

		imageBuffer := bytes.NewBuffer([]byte{})
		if imageContext.EncodePNG(imageBuffer) != nil {
			server.logger.Error("error encoding image to buffer", err)
			return context.Blob(http.StatusInternalServerError, "application/json", []byte(`{"error": "error encoding image to buffer"}`))
		}
		return context.Blob(http.StatusOK, "image/png", imageBuffer.Bytes())
	}
}
