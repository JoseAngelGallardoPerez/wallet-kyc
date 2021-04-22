package action

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type AdminTiersByCountryCode struct {
	adapterTier AdapterTier
}

func NewAdminTiersByCountryCode(
	adapterTier AdapterTier,
) *AdminTiersByCountryCode {
	return &AdminTiersByCountryCode{
		adapterTier: adapterTier,
	}
}

func (s *AdminTiersByCountryCode) Do(ctx context.Context, countryCode string) (objects []model.Tier, err error) {
	return s.adapterTier.FindByCountryCode(ctx, countryCode)
}
