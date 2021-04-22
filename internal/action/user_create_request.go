package action

import (
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"fmt"
)

type UserCreateRequest struct {
	adapterRequest         AdapterUserRequest
	adapterUserRequirement AdapterUserRequirement
	adapterTierRequirement AdapterTierRequirement
	adapterNotifications   AdapterNotifications
}

func NewUserCreateRequest(
	adapterRequest AdapterUserRequest,
	adapterUserRequirement AdapterUserRequirement,
	adapterTierRequirement AdapterTierRequirement,
	adapterNotifications AdapterNotifications,
) *UserCreateRequest {
	return &UserCreateRequest{
		adapterRequest:         adapterRequest,
		adapterUserRequirement: adapterUserRequirement,
		adapterTierRequirement: adapterTierRequirement,
		adapterNotifications:   adapterNotifications,
	}
}

// Create a Tier Request from user
func (t *UserCreateRequest) Do(ctx context.Context, user model.User, tier model.Tier) (*model.UserRequest, error) {

	request, err := t.adapterRequest.FindByTierIdAndUserId(ctx, tier.ID, user.ID)
	if err == nil {
		if request.Status == model.RequestStatus.Approved {
			return nil, internal_errors.CreateError(nil, internal_errors.RequestIsApproved, "Request already approved")
		}

		if request.Status == model.RequestStatus.Pending {
			return nil, internal_errors.CreateError(nil, internal_errors.RequestIsPending, "Request status pending")
		}
	}

	var errors internal_errors.Errors

	//Check that all requirements for this level are filled
	var requirements []model.TierRequirement
	requirements, err = t.adapterTierRequirement.FindByTierId(ctx, tier.ID)
	if err != nil {
		return nil, err
	}

	for _, requirementObject := range requirements {
		userRequirement, err := t.adapterUserRequirement.FindByRequirementIdAndUserId(ctx, requirementObject.ID, user.ID)
		if err != nil {
			errors = errors.Add(err)
		} else if userRequirement.Status != model.RequirementStatus.Waiting &&
			userRequirement.Status != model.RequirementStatus.Approved &&
			userRequirement.Status != model.RequirementStatus.Pending {
			errors = errors.Add(fmt.Errorf("Requirement must be waiting"))
		}
	}
	if len(errors) > 0 {
		return nil, internal_errors.CreateError(errors, internal_errors.RequirementsFilled, "Not all requirements filled")
	}

	// Change statuses for requirements
	for _, requirementObject := range requirements {
		userRequirement, err := t.adapterUserRequirement.FindByRequirementIdAndUserId(ctx, requirementObject.ID, user.ID)
		if err == nil && (userRequirement.Status == model.RequirementStatus.Waiting) {
			userRequirement.Status = model.RequestStatus.Pending
			err := t.adapterUserRequirement.Updates(ctx, userRequirement)
			if err != nil {
				errors = errors.Add(err)
			}
		}
	}
	if len(errors) > 0 {
		return nil, internal_errors.CreateError(errors, internal_errors.RequirementsFilled, "Some requirements not established")
	}

	if request == nil {
		request = &model.UserRequest{
			TierId: tier.ID,
			UserId: user.ID,
			Status: model.RequestStatus.Pending,
		}
		err = t.adapterRequest.Create(ctx, request)
		if err != nil {
			return nil, err
		}
	} else {
		request.Status = model.RequestStatus.Pending

		err = t.adapterRequest.Updates(ctx, request)
		if err != nil {
			return nil, err
		}
	}

	t.adapterNotifications.SendNewUpgradeRequest(ctx)

	return request, nil
}
