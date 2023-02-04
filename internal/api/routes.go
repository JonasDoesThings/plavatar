package api

import (
	"crypto/subtle"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func (server *Server) routes() {
	server.echoRouter.GET("/pixel/:size/:name", server.HandleGetPixelAvatar())
	server.echoRouter.GET("/pixel/:size", server.HandleGetPixelAvatar())
	server.echoRouter.GET("/solid/:size/:name", server.HandleGetSolidAvatar())
	server.echoRouter.GET("/solid/:size", server.HandleGetSolidAvatar())
	server.echoRouter.GET("/gradient/:size/:name", server.HandleGetGradientAvatar())
	server.echoRouter.GET("/gradient/:size", server.HandleGetGradientAvatar())
	server.echoRouter.GET("/marble/:size/:name", server.HandleGetMarbleAvatar())
	server.echoRouter.GET("/marble/:size", server.HandleGetMarbleAvatar())
	server.echoRouter.GET("/laughing/:size/:name", server.HandleGetLaughingAvatar())
	server.echoRouter.GET("/laughing/:size", server.HandleGetLaughingAvatar())
	server.echoRouter.GET("/smiley/:size/:name", server.HandleGetSmileyAvatar())
	server.echoRouter.GET("/smiley/:size", server.HandleGetSmileyAvatar())
	server.echoRouter.GET("/happy/:size/:name", server.HandleGetHappyAvatar())
	server.echoRouter.GET("/happy/:size", server.HandleGetHappyAvatar())
}

func (server *Server) enablePrometheus() {
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
		server.echoRouter.GET("/metrics", server.HandleMetrics(), basicAuthMiddleware)
	} else {
		server.echoRouter.GET("/metrics", server.HandleMetrics())
	}

	prometheusMiddleware := prometheus.NewPrometheus("echo", nil)
	server.echoRouter.Use(prometheusMiddleware.HandlerFunc)
}
