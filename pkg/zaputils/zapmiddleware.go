package zaputils

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type logObject struct {
	Status    int
	Method    string
	Uri       string
	IP        string
	UserAgent string
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
				Uri:       req.RequestURI,
				IP:        c.RealIP(),
				UserAgent: req.UserAgent(),
			}

			switch {
			case res.Status >= 500:
				log.Error(logMessage)
			case res.Status >= 400:
				log.Info(logMessage)
			case res.Status >= 300:
				log.Info(logMessage)
			default:
				log.Info(logMessage)
			}

			return nil
		}
	}
}
