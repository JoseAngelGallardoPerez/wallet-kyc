package action

import (
	"github.com/Confialink/wallet-kyc/internal/action/requirement"
	"github.com/Confialink/wallet-kyc/internal/action/tier"
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type AdminTiersForUser struct {
	actionTierGetByUser                  *tier.GetByUser
	actionTierBindLimitations            *tier.BindLimitations
	actionRequirementBindUserRequirement *requirement.BindRequirementsAndElements

	adapterTier        AdapterTier
	adapterUserRequest AdapterUserRequest
	adapterUser        AdapterUser
}

func NewAdminTiersForUser(
	actionTierGetByUser *tier.GetByUser,
	actionTierBindLimitations *tier.BindLimitations,
	actionRequirementBindUserRequirement *requirement.BindRequirementsAndElements,

	adapterTier AdapterTier,
	adapterUserRequest AdapterUserRequest,
	adapterUser AdapterUser,
) *AdminTiersForUser {
	return &AdminTiersForUser{
		actionTierGetByUser:                  actionTierGetByUser,
		actionTierBindLimitations:            actionTierBindLimitations,
		actionRequirementBindUserRequirement: actionRequirementBindUserRequirement,

		adapterTier:        adapterTier,
		adapterUserRequest: adapterUserRequest,
		adapterUser:        adapterUser,
	}
}

// Getting the list of Tiers for the user. The list includes the requirement,
// elements of the requirements and the meaning of the requirements.
func (s *AdminTiersForUser) Do(ctx context.Context, userId string) (*[]model.Tier, error) {

	user, err := s.adapterUser.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	_, err = s.actionTierGetByUser.Do(ctx, *user)
	if err != nil {
		return nil, err
	}

	var errors internal_errors.Errors

	tiers, err := s.adapterTier.FindByCountryCode(ctx, user.CountryCode)
	if err != nil {
		return nil, err
	}

	for a, tierObject := range tiers {
		for _, requirementObject := range tierObject.Requirements {
			err := s.actionRequirementBindUserRequirement.Do(ctx, requirementObject, user)
			if err != nil {
				errors = errors.Add(err)
			}
		}

		userRequest, err := s.adapterUserRequest.FindByTierIdAndUserId(ctx, tierObject.ID, user.ID)

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
	return &tiers, nil
}
