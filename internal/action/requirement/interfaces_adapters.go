package requirement

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type AdapterUserRequirement interface {
	FindByRequirementIdAndUserId(ctx context.Context, requirementId uint64, userId string) (*model.UserRequirement, error)
	Create(ctx context.Context, model *model.UserRequirement) error
}
