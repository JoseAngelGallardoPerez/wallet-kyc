package action

import (
	"github.com/Confialink/wallet-kyc/internal/action/request"
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type UserRequirementUpdate struct {
	adapterTierRequirement       AdapterTierRequirement
	adapterUserRequirementValue  AdapterUserRequirementValue
	adapterUserRequirement       AdapterUserRequirement
	serviceRequirement           ServiceRequirement
	actionRequestGetByTierUserId *request.GetByTierUserId
}

func NewUserRequirementUpdate(
	adapterTierRequirement AdapterTierRequirement,
	adapterUserRequirementValue AdapterUserRequirementValue,
	adapterUserRequirement AdapterUserRequirement,
	serviceRequirement ServiceRequirement,
	actionRequestGetByTierUserId *request.GetByTierUserId,
) *UserRequirementUpdate {
	return &UserRequirementUpdate{
		adapterTierRequirement:       adapterTierRequirement,
		adapterUserRequirementValue:  adapterUserRequirementValue,
		adapterUserRequirement:       adapterUserRequirement,
		serviceRequirement:           serviceRequirement,
		actionRequestGetByTierUserId: actionRequestGetByTierUserId,
	}
}

// Updating User Requirement Values
func (s *UserRequirementUpdate) Do(ctx context.Context, user model.User, requirementId uint64, values []model.UserRequirementValue) error {

	requirement, err := s.adapterTierRequirement.FindById(ctx, requirementId)
	if err != nil {
		return err
	}

	if requirement.Tier.CountryCode != user.CountryCode {
		return internal_errors.CreateError(nil, internal_errors.TierIsNotAvailable, "This tier is not available")
	}

	userRequirement, err := s.adapterUserRequirement.FindByRequirementIdAndUserId(ctx, requirementId, user.ID)

	if err != nil {
		return err
	} else {
		if userRequirement.Status == model.RequirementStatus.Approved {
			return internal_errors.CreateError(nil, internal_errors.RequirementIsApproved, "This requirement is approved")
		}

		if userRequirement.Status == model.RequirementStatus.Pending {
			return internal_errors.CreateError(nil, internal_errors.RequirementIsPending, "This requirement is pending")
		}
	}

	userRequest, err := s.actionRequestGetByTierUserId.Do(ctx, *requirement.Tier, user.ID)
	if err != nil {
		return err
	} else {
		if userRequest.Status == model.RequestStatus.Approved {
			return internal_errors.CreateError(nil, internal_errors.RequestIsApproved, "This request is approved")
		}

		if userRequest.Status == model.RequestStatus.Pending {
			return internal_errors.CreateError(nil, internal_errors.RequestIsPending, "This request is pending")
		}

		if userRequest.Status == model.RequestStatus.NotAvailable {
			return internal_errors.CreateError(nil, internal_errors.RequirementIsNotAvailable, "This requirement not available")
		}
	}

	err = s.serviceRequirement.SetValue(ctx, userRequirement, values)
	if err != nil {
		return err
	}

	userRequirement.Status = model.RequirementStatus.Waiting
	err = s.adapterUserRequirement.Updates(ctx, userRequirement)
	if err != nil {
		return err
	}

	return nil
}
