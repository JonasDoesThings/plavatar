package api

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (server *Server) HandleMetrics() echo.HandlerFunc {
	h := promhttp.Handler()
	return func(context echo.Context) error {
		h.ServeHTTP(context.Response(), context.Request())
		return nil
	}
}
