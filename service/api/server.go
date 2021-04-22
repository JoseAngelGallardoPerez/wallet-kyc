package api

import (
	"github.com/Confialink/wallet-kyc/service/api/validators"
	"github.com/gin-gonic/gin"
)

var server Server

type Server struct {
	Configs Configs
	Gin     *gin.Engine
}

func InitServer(configs Configs) error {

	server.Configs = configs

	gin.SetMode(server.Configs.Env)

	server.Gin = gin.New()

	validators.InitValidate()

	server.initRoutes()

	return server.Gin.Run(":" + server.Configs.Port)
}
