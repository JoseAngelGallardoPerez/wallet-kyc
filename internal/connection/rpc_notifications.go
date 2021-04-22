package connection

import (
	"github.com/Confialink/wallet-kyc/internal/srvdiscovery"
	notificationspb "github.com/Confialink/wallet-notifications/rpc/proto/notifications"
	"net/http"
)

type RpcNotifications struct {
	Client notificationspb.NotificationHandler
}

var rpcNotifications *RpcNotifications

func InitRpcNotifications() (err error) {
	notificationsUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNameNotifications)
	if nil != err {
		return err
	}

	client, err := notificationspb.NewNotificationHandlerProtobufClient(notificationsUrl.String(), http.DefaultClient), nil

	rpcNotifications = &RpcNotifications{
		Client: client,
	}
	return
}

func GetRpcNotifications() *RpcNotifications {
	return rpcNotifications
}
