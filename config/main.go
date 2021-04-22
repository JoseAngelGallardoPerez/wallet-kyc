package config

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/service/api"
	"github.com/Confialink/wallet-kyc/service/api/middlewares"
	"fmt"
	"os"
	"strings"
)

const defaultEnv = "development"

type Configs struct {
	App AppConfigs
	Api api.Configs
	Db  connection.DbConfigs
}

var configs Configs

func init() {
	configs.App = AppConfigs{
		Env: env("ENV", defaultEnv),
	}

	configs.Api = api.Configs{
		Port: os.Getenv("VELMIE_WALLET_KYC_PORT"),
		Env:  api.ApiEnvMap[configs.App.Env],
		Cors: middlewares.CorsConfigs{
			Methods: strings.Split(env("VELMIE_WALLET_KYS_CORS_METHODS", "GET,POST,PUT,PATCH,DELETE,OPTIONS"), ","),
			Origins: strings.Split(env("VELMIE_WALLET_KYS_CORS_ORIGINS", "*"), ","),
			Headers: strings.Split(env("VELMIE_WALLET_KYC_CORS_HEADERS", "*"), ","),
		},
	}

	configs.Db = connection.DbConfigs{
		IsDebugMode: env("VELMIE_WALLET_KYC_DB_IS_DEBUG_MODE", "false") == "true",
		Driver:      env("VELMIE_WALLET_KYC_DB_DRIV", "mysql"),
		Port:        env("VELMIE_WALLET_KYC_DB_PORT", "3306"),
		Schema:      env("VELMIE_WALLET_KYC_DB_NAME", "velmie-wallet-kys"),

		Host:     os.Getenv("VELMIE_WALLET_KYC_DB_HOST"),
		User:     os.Getenv("VELMIE_WALLET_KYC_DB_USER"),
		Password: os.Getenv("VELMIE_WALLET_KYC_DB_PASS"),
	}
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Validate() error {
	if configs.App.Validate() ||
		configs.Api.Validate() ||
		configs.Db.Validate() {
		return nil
	}
	return fmt.Errorf("Configs invalid")
}

func GetConfigs() Configs {
	return configs
}
