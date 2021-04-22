package config

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
)

type AppConfigs struct {
	Env string
}

func (c AppConfigs) Validate() (isValid bool) {
	isValid = true
	if len(c.Env) == 0 {
		isValid = false
		connection.Logger.Warn("Parameter not configured", "AppConfigs.Env", c.Env)
	}
	return
}
