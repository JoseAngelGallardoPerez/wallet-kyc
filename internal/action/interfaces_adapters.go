package action

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type AdapterTier interface {
	FindByCountryCode(ctx context.Context, countryCode string) (tiers []model.Tier, err error)
	GetByCountryCodeAndLevel(ctx context.Context, countryCode string, level int) (model.Tier, error)
	FindById(ctx context.Context, id uint64) (tier model.Tier, err error)
}

type AdapterUserRequest interface {
	Create(ctx context.Context, request *model.UserRequest) error
	FindByUserIdAndTierIdAndStatuses(ctx context.Context, userId string, tierId uint64, statuses []string) ([]model.UserRequest, error)
	FindById(ctx context.Context, id uint64) (model model.UserRequest, err error)
	Updates(ctx context.Context, request *model.UserRequest) error
	FindByTierIdAndUserId(ctx context.Context, tierId uint64, userId string) (*model.UserRequest, error)
	GetByUserIdAndCountryCodeAndStatus(ctx context.Context, userId string, countryCode string, status string) (model.UserRequest, error)
}

type AdapterTierRequirement interface {
	FindByTierId(ctx context.Context, tierId uint64) ([]model.TierRequirement, error)
	FindById(ctx context.Context, id uint64) (model model.TierRequirement, err error)
}

type AdapterUserRequirement interface {
	FindByRequirementIdAndUserId(ctx context.Context, requirementId uint64, userId string) (*model.UserRequirement, error)
	Updates(ctx context.Context, model *model.UserRequirement) error
	Create(ctx context.Context, model *model.UserRequirement) error
	FindById(ctx context.Context, id uint64) (model *model.UserRequirement, err error)
}

type AdapterUserRequirementValue interface {
	Update(ctx context.Context, model *model.UserRequirementValue) error
	Create(ctx context.Context, model *model.UserRequirementValue) error
	FindByUserRequirementIdAndIndex(ctx context.Context, userRequirementId uint64, index string) (*model.UserRequirementValue, error)
}

type AdapterUser interface {
	GetUserById(ctx context.Context, userId string) (*model.User, error)
}

type AdapterTierLimitation interface {
	FindByTierId(ctx context.Context, tierId uint64) ([]*model.TierLimitation, error)
}

type AdapterCountries interface {
	FindAll(ctx context.Context) (objects []model.Country, err error)
}

type AdapterNotifications interface {
	SendNewUpgradeRequest(ctx context.Context)
	SendRequestApproved(ctx context.Context, user model.User, tierName string)
	SendRequestCanceled(ctx context.Context, user model.User, tierName string)
	SendStatusDocChanged(ctx context.Context, user model.User, documentName string, status string)
}

type AdapterLog interface {
	CreateRequirementUpdate(ctx context.Context, userId string, old model.UserRequirement, new model.UserRequirement)
	CreateRequestUpdateStatus(ctx context.Context, editorId string, old model.UserRequest, new model.UserRequest)
	CreateRequirementUpdateStatus(ctx context.Context, editorId string, old model.UserRequirement, new model.UserRequirement)
}
