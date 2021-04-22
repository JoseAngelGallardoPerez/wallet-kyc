package tier

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/Confialink/wallet-kyc/internal/service/limit"
	"context"
)

var limitNameMap = map[string]string{
	limit.MaxTotalBalance:       "Max total account balance (all time)",
	limit.MaxTotalDebitPerDay:   "Max total debited amount per day",
	limit.MaxTotalDebitPerMonth: "Max total debited amount per month",
	limit.MaxDebitPerTransfer:   "Max amount that could be debited within single transfer",
	limit.MaxCreditPerTransfer:  "Max amount that could be credited within single transfer",
}

type BindLimitations struct {
	adapterTierLimitation AdapterTierLimitation
}

func NewBindLimitations(
	adapterTierLimitation AdapterTierLimitation,
) *BindLimitations {
	return &BindLimitations{
		adapterTierLimitation: adapterTierLimitation,
	}
}

func (s *BindLimitations) Do(ctx context.Context, tier *model.Tier) error {

	limitations, err := s.adapterTierLimitation.FindByTierId(ctx, tier.ID)
	tier.Limitations = limitations

	for i := 0; i < len(tier.Limitations); i++ {
		limitation := tier.Limitations[i]
		limitation.Name = limitNameMap[limitation.Index]
	}

	return err
}
