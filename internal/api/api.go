package api

import (
	"bufio"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"plavatar/internal/avatars"
	"plavatar/internal/caching"
	"plavatar/internal/utils"
	"strconv"
)

type Server struct {
	logger          *zap.SugaredLogger
	echoRouter      *echo.Echo
	avatarGenerator *avatars.Generator
}

var minSize, maxSize int

func StartServer() {
	logger := utils.InitLogger()

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

	minSize = viper.GetInt("dimensions.min")
	maxSize = viper.GetInt("dimensions.max")

	echoRouter := echo.New()
	echoRouter.HideBanner = true
	echoRouter.Use(utils.ZapLogger(logger))
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
		logger:          logger,
		echoRouter:      echoRouter,
		avatarGenerator: &avatars.Generator{},
	}
	apiServer.routes()

	if viper.GetBool("metrics.enabled") {
		apiServer.enablePrometheus()
	}

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

func (server *Server) getSizeFromRequest(context echo.Context) int {
	size, err := strconv.Atoi(context.Param("size"))
	if err != nil {
		return -1
	}

	if size < minSize || size > maxSize {
		return -1
	}

	return size
}
