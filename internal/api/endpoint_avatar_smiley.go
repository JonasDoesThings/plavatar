package api

import (
	"bytes"
	"github.com/fogleman/gg"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"plavatar/internal/utils"
)

func (server *Server) HandleGetCutieAvatar() echo.HandlerFunc {
	return func(context echo.Context) error {
		imageContext, err := server.getAvatarImageContext(context)
		if err != nil {
			return err
		}
		if imageContext == nil {
			return nil
		}

		size := float64(imageContext.Image().Bounds().Max.X)

		name := context.Param("name")
		seed := int64(rand.Intn(2147483647))
		if name != "" {
			seed = int64(server.hashString(name))
		}
		rng := rand.New(rand.NewSource(seed))

		gradient := gg.NewLinearGradient(0, 0, size, size)
		gradient.AddColorStop(0, utils.GetRandomColor(rng))
		rng.Seed(seed + 128)
		gradient.AddColorStop(1, utils.GetRandomColor(rng))

		imageContext.SetFillStyle(gradient)
		imageContext.DrawRectangle(0, 0, size, size)
		imageContext.Fill()

		eyeOffset1 := utils.RandomRangeFloat(rng, -size/20, size/20)
		eyeSizeOffset := utils.RandomRangeFloat(rng, -size/50, size/50)
		rng.Seed(seed + 256)
		eyeOffset2 := utils.RandomRangeFloat(rng, -size/20, size/20)
		mouthRotationOffset := utils.RandomRangeFloat(rng, -15, 15)

		imageContext.SetColor(utils.ColorWhite)
		imageContext.DrawCircle(float64(size)*(0.75/4)+eyeOffset1, size/2-float64(size)*(1.0/10)+eyeOffset2, size/9+eyeSizeOffset)
		imageContext.DrawCircle(float64(size)*(3.25/4)+eyeOffset2, size/2-float64(size)*(1.0/10)+eyeOffset1, size/9+eyeSizeOffset)
		imageContext.Fill()
		imageContext.DrawEllipticalArc(size/2, size/2+float64(size)*(1.0/6), size/4, size/4, gg.Radians(180+mouthRotationOffset), gg.Radians(0))
		imageContext.Fill()

		imageBuffer := bytes.NewBuffer([]byte{})
		if imageContext.EncodePNG(imageBuffer) != nil {
			server.logger.Error("error encoding image to buffer", err)
			return context.Blob(http.StatusInternalServerError, "application/json", []byte(`{"error": "error encoding image to buffer"}`))
		}
		return context.Blob(http.StatusOK, "image/png", imageBuffer.Bytes())
	}
}
