package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsConfigs struct {
	Methods []string
	Origins []string
	Headers []string
}

// Cors middleware
func Cors(corsConfig CorsConfigs) gin.HandlerFunc {
	corsDefaultConfig := cors.DefaultConfig()

	corsDefaultConfig.AllowMethods = corsConfig.Methods
	for _, origin := range corsConfig.Origins {
		if origin == "*" {
			corsDefaultConfig.AllowAllOrigins = true
		}
	}
	if !corsDefaultConfig.AllowAllOrigins {
		corsDefaultConfig.AllowOrigins = corsConfig.Origins
	}
	corsDefaultConfig.AllowHeaders = corsConfig.Headers

	return cors.New(corsDefaultConfig)
}
