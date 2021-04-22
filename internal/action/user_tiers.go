package action

import (
	"github.com/Confialink/wallet-kyc/internal/action/requirement"
	"github.com/Confialink/wallet-kyc/internal/action/tier"
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type UserTiers struct {
	adapterTier        AdapterTier
	adapterUserRequest AdapterUserRequest

	actionRequirementBindUserRequirement *requirement.BindRequirementsAndElements
	actionTierBindLimitations            *tier.BindLimitations
	actionTierGetByUser                  *tier.GetByUser
}

func NewUserTiers(
	adapterTier AdapterTier,
	adapterUserRequest AdapterUserRequest,

	actionRequirementBindUserRequirement *requirement.BindRequirementsAndElements,
	actionTierBindLimitations *tier.BindLimitations,
	actionTierGetByUser *tier.GetByUser,
) *UserTiers {
	return &UserTiers{
		adapterTier:        adapterTier,
		adapterUserRequest: adapterUserRequest,

		actionRequirementBindUserRequirement: actionRequirementBindUserRequirement,
		actionTierBindLimitations:            actionTierBindLimitations,
		actionTierGetByUser:                  actionTierGetByUser,
	}
}

// Get the Tier list for the user; if there is no initial Tier, it will be created automatically
func (t *UserTiers) Do(ctx context.Context, user model.User) (tiers []model.Tier, err error) {
	_, err = t.actionTierGetByUser.Do(ctx, user)
	if err != nil {
		return nil, err
	}

	var errors internal_errors.Errors

	tiers, err = t.adapterTier.FindByCountryCode(ctx, user.CountryCode)
	if err != nil {
		return nil, err
	}

	for a, tierObject := range tiers {
		for _, requirementObject := range tierObject.Requirements {
			err := t.actionRequirementBindUserRequirement.Do(ctx, requirementObject, &user)
			if err != nil {
				errors = errors.Add(err)
			}
		}

		userRequest, err := t.adapterUserRequest.FindByTierIdAndUserId(ctx, tierObject.ID, user.ID)

		if err != nil {
			previousTier := tiers[a-1]

			if len(previousTier.Requests) > 0 && previousTier.Requests[0].Status == model.RequestStatus.Approved {
				userRequest = &model.UserRequest{
					UserId: user.ID,
					TierId: tierObject.ID,
					Status: model.RequestStatus.NotFilled,
				}
			}
		}

		if userRequest != nil {
			tiers[a].Requests = append(tiers[a].Requests, userRequest)
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	return tiers, nil
}
