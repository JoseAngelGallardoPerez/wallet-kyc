package requirement

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type ServiceRequirement interface {
	BindElements(ctx context.Context, requirementObject *model.TierRequirement) error
	BindValues(ctx context.Context, userRequirement *model.UserRequirement) error
}
