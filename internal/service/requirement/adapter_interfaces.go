package requirement

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type AdapterUserRequirementValue interface {
	Update(ctx context.Context, model *model.UserRequirementValue) error
	Create(ctx context.Context, model *model.UserRequirementValue) error
	FindByUserRequirementIdAndIndex(ctx context.Context, userRequirementId uint64, index string) (*model.UserRequirementValue, error)
	FindByUserRequirementId(ctx context.Context, userRequirementId uint64) (objects []model.UserRequirementValue, err error)
	CreateOrUpdate(ctx context.Context, userRequirementId uint64, index string, value string) (*model.UserRequirementValue, error)
}

type AdapterUser interface {
	GetUserById(ctx context.Context, userId string) (*model.User, error)
}

type AdapterUserRequirement interface {
	Updates(ctx context.Context, model *model.UserRequirement) error
}

type AdapterTier interface {
	FindById(ctx context.Context, id uint64) (tier model.Tier, err error)
}

type AdapterFile interface {
	FindById(ctx context.Context, id uint64) (*model.File, error)
}
