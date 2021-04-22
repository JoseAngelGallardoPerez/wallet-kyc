package adapter

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
	"context"
)

type User struct {
	rpcUser *connection.RpcUsers
}

func NewUser() *User {
	return &User{
		rpcUser: connection.GetRpcUsers(),
	}
}

func (s *User) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	request := &userpb.Request{
		UID: userId,
	}

	response, err := s.rpcUser.Client.GetByUID(ctx, request)
	if err != nil {
		return nil, internal_errors.CreateError(err, internal_errors.UserNotFound, "")
	}

	user := s.convertUser(response.User)
	return &user, nil
}

func (s *User) GetUsersByIds(ctx context.Context, userIds []string) ([]model.User, error) {
	request := &userpb.Request{
		UIDs: userIds,
	}

	response, err := s.rpcUser.Client.GetByUIDs(ctx, request)
	if err != nil {
		return nil, internal_errors.CreateError(err, internal_errors.UserNotFound, "")
	}

	var users []model.User
	for _, userObject := range response.Users {
		user := s.convertUser(userObject)
		users = append(users, user)
	}

	return users, nil
}

func (s *User) convertUser(user *userpb.User) model.User {
	object := model.User{
		ID:                     user.UID,
		Phone:                  user.PhoneNumber,
		Email:                  user.Email,
		Role:                   user.RoleName,
		FirstName:              user.FirstName,
		LastName:               user.LastName,
		IsPhoneNumberConfirmed: user.IsPhoneNumberConfirmed,
		IsEmailConfirmed:       user.IsEmailConfirmed,
	}
	object.RecognizeCountry()
	return object
}
