package action

import (
	"github.com/Confialink/wallet-kyc/internal/action/requirement"
	"github.com/Confialink/wallet-kyc/internal/action/tier"
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type UserTier struct {
	adapterTier                          AdapterTier
	actionRequirementBindUserRequirement *requirement.BindRequirementsAndElements
	actionTierGetByUser                  *tier.GetByUser
}

func NewUserTier(
	adapterTier AdapterTier,
	actionRequirementBindUserRequirement *requirement.BindRequirementsAndElements,
	actionTierGetByUser *tier.GetByUser,
) *UserTier {
	return &UserTier{
		adapterTier:                          adapterTier,
		actionRequirementBindUserRequirement: actionRequirementBindUserRequirement,
		actionTierGetByUser:                  actionTierGetByUser,
	}
}

// Obtaining information of a separate Tier for the user
func (t *UserTier) Do(ctx context.Context, user model.User, tierId uint64) (*model.Tier, error) {
	_, err := t.actionTierGetByUser.Do(ctx, user)
	if err != nil {
		return nil, err
	}

	tierObject, err := t.adapterTier.FindById(ctx, tierId)
	if err != nil {
		return nil, err
	}

	if tierObject.CountryCode != user.CountryCode {
		return nil, internal_errors.CreateError(nil, internal_errors.TierIsNotAvailable, "This tier is not available")
	}

	var errors internal_errors.Errors

	for _, requirementObject := range tierObject.Requirements {
		err := t.actionRequirementBindUserRequirement.Do(ctx, requirementObject, &user)
		if err != nil {
			errors = errors.Add(err)
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	return &tierObject, nil
}
