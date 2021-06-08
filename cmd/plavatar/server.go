package main

import (
	"github.com/spf13/viper"
	"plavatar/internal/api"
)

func main() {
	viper.SetDefault("webserver", map[string]interface{}{
		"gzip": false,
		"http": map[string]interface{}{
			"enabled": true,
			"host":    "0.0.0.0",
			"port":    7331,
		},
		"https": map[string]interface{}{
			"enabled": false,
			"host":    "0.0.0.0",
			"port":    7332,
			"cert":    "testing.crt",
			"key":     "testing.key",
		},
	})
	viper.SetDefault("caching", map[string]interface{}{
		"enabled": true,
		"ttl":     "8h",
	})

	api.StartServer()
}
