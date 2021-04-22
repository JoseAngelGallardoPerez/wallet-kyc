package connection

import (
	"github.com/Confialink/wallet-accounts/rpc/limit"
	"github.com/Confialink/wallet-kyc/internal/srvdiscovery"

	"net/http"
)

type RpcLimit struct {
	Client limit.Limits
}

var rpcLimit *RpcLimit

func InitRpcLimit() (err error) {
	accountsUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameAccounts)
	if nil != err {
		return err
	}

	client, err := limit.NewLimitsProtobufClient(accountsUrl.String(), http.DefaultClient), nil

	rpcLimit = &RpcLimit{
		Client: client,
	}
	return
}

func GetRpcLimit() *RpcLimit {
	return rpcLimit
}
