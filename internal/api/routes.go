package api

import (
	"crypto/subtle"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func (server *Server) routes() {
	server.echoRouter.GET("/pixel/:size/:name", server.HandleGetAvatar(server.avatarGenerator.Pixels))
	server.echoRouter.GET("/pixel/:size", server.HandleGetAvatar(server.avatarGenerator.Pixels))
	server.echoRouter.GET("/solid/:size/:name", server.HandleGetAvatar(server.avatarGenerator.Solid))
	server.echoRouter.GET("/solid/:size", server.HandleGetAvatar(server.avatarGenerator.Solid))
	server.echoRouter.GET("/gradient/:size/:name", server.HandleGetAvatar(server.avatarGenerator.Gradient))
	server.echoRouter.GET("/gradient/:size", server.HandleGetAvatar(server.avatarGenerator.Gradient))
	server.echoRouter.GET("/marble/:size/:name", server.HandleGetAvatar(server.avatarGenerator.Marble))
	server.echoRouter.GET("/marble/:size", server.HandleGetAvatar(server.avatarGenerator.Marble))
	server.echoRouter.GET("/laughing/:size/:name", server.HandleGetAvatar(server.avatarGenerator.Laughing))
	server.echoRouter.GET("/laughing/:size", server.HandleGetAvatar(server.avatarGenerator.Laughing))
	server.echoRouter.GET("/smiley/:size/:name", server.HandleGetAvatar(server.avatarGenerator.Smiley))
	server.echoRouter.GET("/smiley/:size", server.HandleGetAvatar(server.avatarGenerator.Smiley))
	server.echoRouter.GET("/happy/:size/:name", server.HandleGetAvatar(server.avatarGenerator.Happy))
	server.echoRouter.GET("/happy/:size", server.HandleGetAvatar(server.avatarGenerator.Happy))
}

func (server *Server) enablePrometheus() {
	server.echoRouter.Use(echoprometheus.NewMiddleware("plavatar"))

	if viper.GetBool("metrics.auth.enabled") {
		metricsAuthUsername := []byte(viper.GetString("metrics.auth.username"))
		metricsAuthPassword := []byte(viper.GetString("metrics.auth.password"))
		basicAuthMiddleware := middleware.BasicAuth(func(username, password string, context echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(username), metricsAuthUsername) != 1 {
				return false, nil
			}
			if subtle.ConstantTimeCompare([]byte(password), metricsAuthPassword) != 1 {
				return false, nil
			}

			return true, nil
		})

		server.echoRouter.GET("/metrics", echoprometheus.NewHandler(), basicAuthMiddleware)
	} else {
		server.echoRouter.GET("/metrics", echoprometheus.NewHandler())
	}
}
