package main

import (
	"github.com/Confialink/wallet-kyc/config"
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/service/api"
)

func main() {

	if err := config.Validate(); err != nil {
		panic(err)
	}

	configs := config.GetConfigs()

	if err := connection.InitDb(configs.Db); err != nil {
		panic(err)
	}

	if err := connection.InitRpcUsers(); err != nil {
		panic(err)
	}

	if err := connection.InitRpcLimit(); err != nil {
		panic(err)
	}

	if err := connection.InitRpcCurrencies(); err != nil {
		panic(err)
	}

	if err := connection.InitRpcNotifications(); err != nil {
		panic(err)
	}

	if err := connection.InitRpcFiles(); err != nil {
		panic(err)
	}

	if err := connection.InitRpcLogs(); err != nil {
		panic(err)
	}

	if err := api.InitServer(configs.Api); err != nil {
		panic(err)
	}
}
