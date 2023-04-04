package zaputils

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type logObject struct {
	Method    string
	URI       string
	IP        string
	UserAgent string
	Status    int
}

func ZapLogger(log *zap.SugaredLogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			logMessage := logObject{
				Status:    res.Status,
				Method:    req.Method,
				URI:       req.RequestURI,
				IP:        c.RealIP(),
				UserAgent: req.UserAgent(),
			}

			switch {
			case res.Status >= http.StatusInternalServerError:
				log.Error(logMessage)
			case res.Status >= http.StatusBadRequest:
				log.Info(logMessage)
			case res.Status >= http.StatusMultipleChoices:
				log.Info(logMessage)
			default:
				log.Info(logMessage)
			}

			return nil
		}
	}
}
