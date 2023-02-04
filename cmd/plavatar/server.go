package main

import (
	"github.com/spf13/viper"
	"plavatar/internal/api"
)

func main() {
	viper.SetDefault("dimensions", map[string]interface{}{
		"min": 128,
		"max": 512,
	})
	viper.SetDefault("webserver", map[string]interface{}{
		"gzip": true,
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
	viper.SetDefault("metrics", map[string]interface{}{
		"enabled": false,
		"auth": map[string]interface{}{
			"enabled":  true,
			"username": "",
			"password": "",
		},
	})
	viper.SetDefault("caching", map[string]interface{}{
		"enabled": true,
		"ttl":     "8h",
	})

	api.StartServer()
}
