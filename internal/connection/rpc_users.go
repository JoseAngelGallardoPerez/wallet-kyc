package connection

import (
	"github.com/Confialink/wallet-kyc/internal/srvdiscovery"
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
	"net/http"
)

type RpcUsers struct {
	Client userpb.UserHandler
}

var rpcUsers *RpcUsers

func InitRpcUsers() (err error) {
	usersUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameUsers)
	if nil != err {
		return err
	}

	client, err := userpb.NewUserHandlerProtobufClient(usersUrl.String(), http.DefaultClient), nil

	rpcUsers = &RpcUsers{
		Client: client,
	}
	return
}

func GetRpcUsers() *RpcUsers {
	return rpcUsers
}
