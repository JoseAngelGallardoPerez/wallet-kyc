package request

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type GetByTierUserId struct {
	adapterUserRequest AdapterUserRequest
	adapterTier        AdapterTier
}

func NewGetByTierUserId(
	adapterUserRequest AdapterUserRequest,
	adapterTier AdapterTier,
) *GetByTierUserId {
	return &GetByTierUserId{
		adapterUserRequest: adapterUserRequest,
		adapterTier:        adapterTier,
	}
}

func (s *GetByTierUserId) Do(ctx context.Context, tier model.Tier, userId string) (*model.UserRequest, error) {
	request, err := s.adapterUserRequest.FindByTierIdAndUserId(ctx, tier.ID, userId)

	if err != nil {
		previousTier, err := s.adapterTier.GetByCountryCodeAndLevel(ctx, tier.CountryCode, tier.Level-1)
		if err != nil {
			return nil, err
		}

		request, err := s.adapterUserRequest.FindByTierIdAndUserId(ctx, previousTier.ID, userId)
		if err != nil {
			return &model.UserRequest{
				UserId: userId,
				TierId: tier.ID,
				Status: model.RequestStatus.NotAvailable,
			}, nil
		} else {
			if request.Status == model.RequestStatus.Approved {
				return &model.UserRequest{
					UserId: userId,
					TierId: tier.ID,
					Status: model.RequestStatus.NotFilled,
				}, nil
			} else {
				return &model.UserRequest{
					UserId: userId,
					TierId: tier.ID,
					Status: model.RequestStatus.NotAvailable,
				}, nil
			}
		}
	} else {
		return request, nil
	}
}
