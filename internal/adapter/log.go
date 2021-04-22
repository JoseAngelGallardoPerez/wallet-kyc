package adapter

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/model"
	logspb "github.com/Confialink/wallet-logs/rpc/logs"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Log struct {
	rpcLogs *connection.RpcLogs
}

func NewLog() *Log {
	return &Log{
		rpcLogs: connection.GetRpcLogs(),
	}
}

const (
	ModifyUserRequirement       = "Modify User Requirement"
	UpdateStatusUserRequirement = "Update User Requirement Status"
	UpdateStatusUserRequest     = "Update User Request Status"
)

func (s *Log) CreateRequirementUpdate(ctx context.Context, editorId string, old model.UserRequirement, new model.UserRequirement) {
	go func() {
		data, err := json.Marshal(map[string]interface{}{
			"old": old,
			"new": new,
		})

		if err != nil {
			connection.Logger.Error(err.Error())
		}

		_, err = s.rpcLogs.Client.CreateLog(ctx, &logspb.CreateLogReq{
			Subject:    ModifyUserRequirement,
			UserId:     editorId,
			LogTime:    time.Now().Format(time.RFC3339),
			DataTitle:  fmt.Sprintf("Requirement %s for user %s", old.TierRequirement.Name, old.User.FirstName+" "+old.User.LastName),
			DataFields: data,
		})

		if err != nil {
			connection.Logger.Error(err.Error())
		}
	}()
}

func (s *Log) CreateRequirementUpdateStatus(ctx context.Context, editorId string, old model.UserRequirement, new model.UserRequirement) {
	go func() {
		data, err := json.Marshal(map[string]interface{}{
			"old": old,
			"new": new,
		})

		if err != nil {
			connection.Logger.Error(err.Error())
		}

		_, err = s.rpcLogs.Client.CreateLog(ctx, &logspb.CreateLogReq{
			Subject:    UpdateStatusUserRequirement,
			UserId:     editorId,
			LogTime:    time.Now().Format(time.RFC3339),
			DataTitle:  fmt.Sprintf("Update requirement status %s for user %s", old.TierRequirement.Name, old.User.FirstName+" "+old.User.LastName),
			DataFields: data,
		})

		if err != nil {
			connection.Logger.Error(err.Error())
		}
	}()
}

func (s *Log) CreateRequestUpdateStatus(ctx context.Context, editorId string, old model.UserRequest, new model.UserRequest) {
	go func() {
		data, err := json.Marshal(map[string]interface{}{
			"old": old,
			"new": new,
		})

		if err != nil {
			connection.Logger.Error(err.Error())
		}

		_, err = s.rpcLogs.Client.CreateLog(ctx, &logspb.CreateLogReq{
			Subject:    UpdateStatusUserRequest,
			UserId:     editorId,
			LogTime:    time.Now().Format(time.RFC3339),
			DataTitle:  fmt.Sprintf("Update request status %s for user %s", old.Tier.Name, old.User.FirstName+" "+old.User.LastName),
			DataFields: data,
		})

		if err != nil {
			connection.Logger.Error(err.Error())
		}
	}()
}
