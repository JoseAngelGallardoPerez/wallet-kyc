package action

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type ServiceRequirement interface {
	BindValues(ctx context.Context, userRequirement *model.UserRequirement) error
	SetValue(ctx context.Context, userRequirement *model.UserRequirement, userRequirementValue []model.UserRequirementValue) error
}
