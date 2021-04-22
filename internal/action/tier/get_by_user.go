package tier

import (
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type GetByUser struct {
	adapterTier            AdapterTier
	adapterTierRequirement AdapterTierRequirement
	adapterUserRequest     AdapterUserRequest
	adapterUserRequirement AdapterUserRequirement
}

func NewGetByUser(
	adapterTier AdapterTier,
	adapterTierRequirement AdapterTierRequirement,
	adapterUserRequest AdapterUserRequest,
	adapterUserRequirement AdapterUserRequirement,
) *GetByUser {
	return &GetByUser{
		adapterTier:            adapterTier,
		adapterTierRequirement: adapterTierRequirement,
		adapterUserRequest:     adapterUserRequest,
		adapterUserRequirement: adapterUserRequirement,
	}
}

func (s *GetByUser) Do(ctx context.Context, user model.User) (*model.Tier, error) {

	var errors internal_errors.Errors

	userTier, err := s.adapterTier.FindLastApprovedByUserIdCode(ctx, user.ID, user.CountryCode)
	if err != nil {

		userTier, err = s.adapterTier.GetByCountryCodeAndLevel(ctx, user.CountryCode, 0)
		if err != nil {
			return nil, err
		}
		var requirements []model.TierRequirement
		requirements, err = s.adapterTierRequirement.FindByTierId(ctx, userTier.ID)
		if err != nil {
			return nil, err
		}

		for r := 0; r < len(requirements); r++ {
			err = s.adapterUserRequirement.Create(ctx, &model.UserRequirement{
				TierRequirementId: requirements[r].ID,
				UserId:            user.ID,
				Status:            model.RequirementStatus.Approved,
			})
			if err != nil {
				errors = errors.Add(err)
			}
		}

		err = s.adapterUserRequest.Create(ctx, &model.UserRequest{
			TierId: userTier.ID,
			UserId: user.ID,
			Status: model.RequestStatus.Approved,
		})
	}

	return &userTier, err
}
