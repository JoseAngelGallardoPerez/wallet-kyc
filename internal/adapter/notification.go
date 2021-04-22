package adapter

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type Notification struct {
	rpcNotifications *connection.RpcNotifications
}

func NewNotification() *Notification {
	return &Notification{
		rpcNotifications: connection.GetRpcNotifications(),
	}
}

func (s *Notification) SendNewUpgradeRequest(ctx context.Context) {
	/*go func() {
		_, err := s.rpcNotifications.Client.Dispatch(context.Background(), &notificationspb.Request{
			EventName:    "KycNewUpgradeRequest",
			TemplateData: &notificationspb.TemplateData{},
			Notifiers: []string{
				"email",
			},
		})
		if err != nil {
			connection.Logger.Error(err.Error())
		}
	}()*/
}

func (s *Notification) SendRequestApproved(ctx context.Context, user model.User, tierName string) {
	/*go func() {
		_, err := s.rpcNotifications.Client.Dispatch(context.Background(), &notificationspb.Request{
			To:        user.ID,
			EventName: "KycRequestApproved",
			TemplateData: &notificationspb.TemplateData{
				TierName: tierName,
			},
			Notifiers: []string{
				"email",
				"sms",
				"push_notification",
			},
		})
		if err != nil {
			connection.Logger.Error(err.Error())
		}
	}()*/
}

func (s *Notification) SendRequestCanceled(ctx context.Context, user model.User, tierName string) {
	/*go func() {
		_, err := s.rpcNotifications.Client.Dispatch(context.Background(), &notificationspb.Request{
			To:        user.ID,
			EventName: "KycRequestCanceled",
			TemplateData: &notificationspb.TemplateData{
				TierName: tierName,
			},
			Notifiers: []string{
				"email",
				"sms",
				"push_notification",
			},
		})
		if err != nil {
			connection.Logger.Error(err.Error())
		}
	}()*/
}

func (s *Notification) SendStatusDocChanged(ctx context.Context, user model.User, documentName string, status string) {
	/*go func() {
		_, err := s.rpcNotifications.Client.Dispatch(context.Background(), &notificationspb.Request{
			To:        user.ID,
			EventName: "KycStatusDocChanged",
			TemplateData: &notificationspb.TemplateData{
				DocumentName: documentName,
				Status:       status,
			},
			Notifiers: []string{
				"email",
				"sms",
				"push_notification",
			},
		})

		if err != nil {
			connection.Logger.Error(err.Error())
		}
	}()*/
}
