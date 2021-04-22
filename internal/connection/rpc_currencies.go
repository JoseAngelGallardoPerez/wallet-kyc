package connection

import (
	currenciespb "github.com/Confialink/wallet-currencies/rpc/currencies"
	"github.com/Confialink/wallet-kyc/internal/srvdiscovery"
	"net/http"
)

type RpcCurrencies struct {
	Client currenciespb.CurrencyFetcher
}

var rpcCurrencies *RpcCurrencies

func InitRpcCurrencies() (err error) {
	currenciesUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameCurrencies)
	if nil != err {
		return err
	}

	client, err := currenciespb.NewCurrencyFetcherProtobufClient(currenciesUrl.String(), http.DefaultClient), nil

	rpcCurrencies = &RpcCurrencies{
		Client: client,
	}
	return
}

func GetRpcCurrencies() *RpcCurrencies {
	return rpcCurrencies
}
