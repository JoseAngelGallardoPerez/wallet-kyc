package action

import (
	"github.com/Confialink/wallet-kyc/internal/action/tier"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type AdminTier struct {
	adapterTier               AdapterTier
	actionTierBindLimitations *tier.BindLimitations
}

func NewAdminTier(
	adapterTier AdapterTier,
	actionTierBindLimitations *tier.BindLimitations,
) *AdminTier {
	return &AdminTier{
		adapterTier:               adapterTier,
		actionTierBindLimitations: actionTierBindLimitations,
	}
}

func (s *AdminTier) Do(ctx context.Context, id uint64) (tierObject model.Tier, err error) {
	tierObject, err = s.adapterTier.FindById(ctx, id)
	if err != nil {
		return
	}

	err = s.actionTierBindLimitations.Do(ctx, &tierObject)
	return
}
