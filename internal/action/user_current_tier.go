package action

import (
	"github.com/Confialink/wallet-kyc/internal/action/tier"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type UserCurrentTier struct {
	actionTierGetByUser       *tier.GetByUser
	actionTierBindLimitations *tier.BindLimitations
}

func NewUserCurrentTier(
	actionTierGetByUser *tier.GetByUser,
	actionTierBindLimitations *tier.BindLimitations,
) *UserCurrentTier {
	return &UserCurrentTier{
		actionTierGetByUser:       actionTierGetByUser,
		actionTierBindLimitations: actionTierBindLimitations,
	}
}

// Get current user tier
func (s *UserCurrentTier) Do(ctx context.Context, user model.User) (*model.Tier, error) {

	tierObject, err := s.actionTierGetByUser.Do(ctx, user)
	if err != nil {
		return nil, err
	}

	err = s.actionTierBindLimitations.Do(ctx, tierObject)
	if err != nil {
		return nil, err
	}

	return tierObject, nil
}
