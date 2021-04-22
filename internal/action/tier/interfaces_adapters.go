package tier

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type AdapterTier interface {
	FindLastApprovedByUserIdCode(ctx context.Context, userId string, countryCode string) (model.Tier, error)
	GetByCountryCodeAndLevel(ctx context.Context, countryCode string, level int) (model.Tier, error)
}

type AdapterTierRequirement interface {
	FindByTierId(ctx context.Context, tierId uint64) ([]model.TierRequirement, error)
}

type AdapterUserRequest interface {
	Create(ctx context.Context, request *model.UserRequest) error
}

type AdapterUserRequirement interface {
	Create(ctx context.Context, model *model.UserRequirement) error
}

type AdapterTierLimitation interface {
	FindByTierId(ctx context.Context, tierId uint64) ([]*model.TierLimitation, error)
	FindByTierIdIndex(ctx context.Context, tierId uint64, index string) (*model.TierLimitation, error)
	Updates(ctx context.Context, object *model.TierLimitation) error
}
