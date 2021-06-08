package caching

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"io"
	"net/http"
)

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
	statusCode int
}

func (w *bodyDumpResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyDumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func CacheMiddleware(avatarCache *cache.Cache) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if cachedAvatar, found := avatarCache.Get(context.Request().RequestURI); found {
				return context.Blob(http.StatusOK, "image/png", cachedAvatar.([]byte))
			}

			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(context.Response().Writer, resBody)
			writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: context.Response().Writer}
			context.Response().Writer = writer

			err := next(context)
			if err != nil {
				context.Error(err)
			}

			if writer.statusCode == http.StatusOK && context.Response().Header().Get("Content-Type") == "image/png" {
				avatarCache.SetDefault(context.Request().RequestURI, resBody.Bytes())
			}

			return nil
		}
	}
}
