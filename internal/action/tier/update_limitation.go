package tier

import (
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type UpdateLimitation struct {
	adapterTierLimitation AdapterTierLimitation
}

func NewTierUpdateLimitation(
	adapterTierLimitation AdapterTierLimitation,
) *UpdateLimitation {
	return &UpdateLimitation{
		adapterTierLimitation: adapterTierLimitation,
	}
}

func (s *UpdateLimitation) Do(ctx context.Context, tierId uint64, limitations []*model.TierLimitation) error {
	var errors internal_errors.Errors
	for _, v := range limitations {

		limitation, err := s.adapterTierLimitation.FindByTierIdIndex(ctx, tierId, v.Index)
		if err != nil {
			errors = errors.Add(err)
		} else {
			limitation.Value = v.Value
			err = s.adapterTierLimitation.Updates(ctx, limitation)
			if err != nil {
				errors = errors.Add(err)
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}
