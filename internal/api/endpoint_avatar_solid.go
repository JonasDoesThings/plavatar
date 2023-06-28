package api

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"plavatar/internal/utils"
)

func (server *Server) HandleGetSolidAvatar() echo.HandlerFunc {
	return func(context echo.Context) error {
		imageContext, err := server.getAvatarImageContext(context)
		if err != nil {
			return err
		}
		if imageContext == nil {
			return nil
		}

		size := imageContext.Image().Bounds().Max.X

		name := context.Param("name")
		seed := int64(rand.Intn(2147483647))
		if name != "" {
			seed = int64(server.hashString(name))
		}
		rng := rand.New(rand.NewSource(seed))

		backgroundColor := utils.GetRandomColor(rng)
		imageContext.SetColor(backgroundColor)
		imageContext.DrawRectangle(0, 0, float64(size), float64(size))
		imageContext.Fill()

		imageBuffer := bytes.NewBuffer([]byte{})
		if imageContext.EncodePNG(imageBuffer) != nil {
			server.logger.Error("error encoding image to buffer", err)
			return context.Blob(http.StatusInternalServerError, "application/json", []byte(`{"error": "error encoding image to buffer"}`))
		}
		return context.Blob(http.StatusOK, "image/png", imageBuffer.Bytes())
	}
}
