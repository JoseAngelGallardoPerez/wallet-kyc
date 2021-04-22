package request

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type BindUsers struct {
	adapterUser AdapterUser
}

func NewBindUsers(adapterUser AdapterUser) *BindUsers {
	return &BindUsers{
		adapterUser: adapterUser,
	}
}

func (s *BindUsers) Do(ctx context.Context, requests []model.UserRequest) ([]model.UserRequest, error) {
	var ids []string
	for _, requestObject := range requests {
		ids = append(ids, requestObject.UserId)
	}

	users, err := s.adapterUser.GetUsersByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	userList := make(map[string]model.User)

	for _, userObject := range users {
		userList[userObject.ID] = userObject
	}

	var userRequests []model.UserRequest
	for _, requestObject := range requests {
		requestObject.User = userList[requestObject.UserId]
		userRequests = append(userRequests, requestObject)
	}

	return userRequests, nil
}
