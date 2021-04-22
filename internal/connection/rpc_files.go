package connection

import (
	filespb "github.com/Confialink/wallet-files/rpc/files"
	"github.com/Confialink/wallet-kyc/internal/srvdiscovery"
	"net/http"
)

type RpcFiles struct {
	Client filespb.ServiceFiles
}

var rpcFiles *RpcFiles

func InitRpcFiles() (err error) {
	filesUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameFiles)
	if nil != err {
		return err
	}

	client, err := filespb.NewServiceFilesProtobufClient(filesUrl.String(), http.DefaultClient), nil

	rpcFiles = &RpcFiles{
		Client: client,
	}
	return
}

func GetRpcFiles() *RpcFiles {
	return rpcFiles
}
