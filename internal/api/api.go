package api

import (
	"bufio"
	"flag"
	"github.com/fogleman/gg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"hash/fnv"
	"net/http"
	"os"
	"plavatar/internal/caching"
	"plavatar/pkg/zaputils"
	"strconv"
)

type Server struct {
	logger     *zap.SugaredLogger
	echoRouter *echo.Echo
}

func StartServer() {
	logger := zaputils.InitLogger()

	var configLocation = flag.String("config", "", "config file location")
	flag.Parse()

	viper.SetConfigName("plavatar")
	viper.SetConfigType("json")

	if *configLocation != "" {
		configFile, err := os.Open(*configLocation)
		if err != nil {
			logger.Fatal("failed opening config file at ", configLocation)
		}
		err = viper.ReadConfig(bufio.NewReader(configFile))
		if err != nil {
			logger.Fatal("failed reading config ", err)
		}
	} else {
		viper.AddConfigPath("./config/")
		if err := viper.ReadInConfig(); err != nil {
			logger.Warn("no config found, using default config!")
		}
	}

	echoRouter := echo.New()
	echoRouter.HideBanner = true
	echoRouter.Pre(middleware.RemoveTrailingSlash())
	echoRouter.Use(zaputils.ZapLogger(logger))
	echoRouter.Use(middleware.Recover())

	if viper.GetBool("webserver.gzip") {
		echoRouter.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Skipper: middleware.DefaultSkipper,
			Level:   4,
		}))
	}

	if viper.GetBool("caching.enabled") {
		avatarCache := cache.New(viper.GetDuration("caching.ttl"), viper.GetDuration("caching.ttl"))
		echoRouter.Use(caching.CacheMiddleware(avatarCache))
	}

	apiServer := Server{
		logger:     logger,
		echoRouter: echoRouter,
	}
	apiServer.routes()

	if viper.GetBool("webserver.https.enabled") {
		go func() {
			err := echoRouter.StartTLS(viper.GetString("webserver.https.host")+":"+viper.GetString("webserver.https.port"), viper.GetString("webserver.https.cert"), viper.GetString("webserver.https.key"))
			if err != nil {
				logger.Fatal("failed starting with https! ", err)
			}
		}()
	}

	if viper.GetBool("webserver.http.enabled") {
		go func() {
			err := echoRouter.Start(viper.GetString("webserver.http.host") + ":" + viper.GetString("webserver.http.port"))
			if err != nil {
				logger.Fatal("failed starting with http! ", err)
			}
		}()
	}

	select {}
}

func (server *Server) getAvatarImageContext(context echo.Context) (*gg.Context, error) {
	size, err := strconv.Atoi(context.Param("size"))
	if err != nil {
		return nil, context.Blob(http.StatusBadRequest, "application/json", []byte(`{"error": "invalid size"}`))
	}

	if size < 16 || size > 1024 {
		return nil, context.Blob(http.StatusBadRequest, "application/json", []byte(`{"error": "invalid size"}`))
	}

	imageContext := gg.NewContext(size, size)

	imageContext.DrawCircle(float64(size/2), float64(size/2), float64(size/2))
	imageContext.Clip()
	imageContext.AsMask()

	return imageContext, nil
}

func (server *Server) hashString(s string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		server.logger.Error("error hashing ", s, err)
	}
	return h.Sum32()
}
