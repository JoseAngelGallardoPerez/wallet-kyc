package request

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type AdapterUser interface {
	GetUsersByIds(ctx context.Context, userIds []string) ([]model.User, error)
}

type AdapterUserRequest interface {
	FindByTierIdAndUserId(ctx context.Context, tierId uint64, userId string) (*model.UserRequest, error)
}

type AdapterTier interface {
	GetByCountryCodeAndLevel(ctx context.Context, countryCode string, level int) (model.Tier, error)
}
