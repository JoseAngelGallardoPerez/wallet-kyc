package api

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/service/api/middlewares"
	"github.com/gin-gonic/gin"
)

var ApiEnvMap = map[string]string{
	"production":  gin.ReleaseMode,
	"development": gin.DebugMode,
	"stage":       gin.DebugMode,
	"test":        gin.TestMode,
}

type Configs struct {
	Port string
	Env  string
	Cors middlewares.CorsConfigs
}

func (c Configs) Validate() (isValid bool) {
	isValid = true
	if c.Port == "" {
		isValid = false
		connection.Logger.Warn("API Parameter not configured", "ApiConfigs.Port", c.Port)
	}

	if len(c.Cors.Methods) == 0 {
		isValid = false
		connection.Logger.Warn("API Parameter not configured", "CorsConfigs.Cors.Methods", c.Cors.Methods)
	}

	if len(c.Cors.Origins) == 0 {
		isValid = false
		connection.Logger.Warn("API Parameter not configured", "CorsConfigs.Cors.Origins", c.Cors.Origins)
	}

	if len(c.Cors.Headers) == 0 {
		isValid = false
		connection.Logger.Warn("API Parameter not configured", "CorsConfigs.Cors.Headers", c.Cors.Headers)
	}
	return
}
