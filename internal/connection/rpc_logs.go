package connection

import (
	"github.com/Confialink/wallet-kyc/internal/srvdiscovery"
	logspb "github.com/Confialink/wallet-logs/rpc/logs"
	"net/http"
)

type RpcLogs struct {
	Client logspb.LogsService
}

var rpcLogs *RpcLogs

func InitRpcLogs() (err error) {
	logsUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameLogs)
	if nil != err {
		return err
	}

	client, err := logspb.NewLogsServiceProtobufClient(logsUrl.String(), http.DefaultClient), nil

	rpcLogs = &RpcLogs{
		Client: client,
	}
	return
}

func GetRpcLogs() *RpcLogs {
	return rpcLogs
}
