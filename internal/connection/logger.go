package connection

import "github.com/inconshreveable/log15"

var Logger log15.Logger

func init() {
	Logger = log15.New("Microservice", "Kyc")
}

func GetLogger() log15.Logger {
	return Logger
}
