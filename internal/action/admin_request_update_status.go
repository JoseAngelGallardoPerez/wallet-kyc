package action

import (
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/Confialink/wallet-kyc/internal/service/limit"
	"context"
	"fmt"
)

type AdminRequestUpdateStatus struct {
	adapterRequest         AdapterUserRequest
	adapterNotifications   AdapterNotifications
	adapterUser            AdapterUser
	adapterTier            AdapterTier
	adapterLog             AdapterLog
	adapterTierRequirement AdapterTierRequirement
	adapterUserRequirement AdapterUserRequirement
	limitService           *limit.Service
}

func NewAdminRequestUpdateStatus(
	adapterRequest AdapterUserRequest,
	adapterNotifications AdapterNotifications,
	adapterUser AdapterUser,
	adapterTier AdapterTier,
	adapterLog AdapterLog,
	adapterTierRequirement AdapterTierRequirement,
	adapterUserRequirement AdapterUserRequirement,
	limitService *limit.Service,
) *AdminRequestUpdateStatus {
	return &AdminRequestUpdateStatus{
		adapterRequest:         adapterRequest,
		adapterNotifications:   adapterNotifications,
		adapterUser:            adapterUser,
		adapterTier:            adapterTier,
		adapterLog:             adapterLog,
		adapterTierRequirement: adapterTierRequirement,
		adapterUserRequirement: adapterUserRequirement,
		limitService:           limitService,
	}
}

// Status Update for Tier Request
func (t *AdminRequestUpdateStatus) Do(ctx context.Context, editor model.User, requestId uint64, status string) (*model.UserRequest, error) {
	modelRequest, err := t.adapterRequest.FindById(ctx, requestId)

	if err != nil {
		return nil, err
	}

	if modelRequest.Status == model.RequestStatus.Approved {
		return nil, internal_errors.CreateError(nil, internal_errors.RequestAlreadyExists, "This request has already been approved.")
	}

	if modelRequest.Status == status {
		return &modelRequest, nil
	}

	modelRequest.Status = status

	user, err := t.adapterUser.GetUserById(ctx, modelRequest.UserId)
	if err != nil {
		return nil, err
	}

	tier, err := t.adapterTier.FindById(ctx, modelRequest.TierId)
	if err != nil {
		return nil, err
	}

	var requirements []model.TierRequirement
	requirements, err = t.adapterTierRequirement.FindByTierId(ctx, tier.ID)
	if err != nil {
		return nil, err
	}

	setNewLimit := false

	switch status {
	case model.RequestStatus.Approved:
		var errors internal_errors.Errors
		for _, requirementObject := range requirements {
			userRequirement, err := t.adapterUserRequirement.FindByRequirementIdAndUserId(ctx, requirementObject.ID, user.ID)
			if err != nil {
				errors = errors.Add(err)
			} else if userRequirement.Status != model.RequirementStatus.Approved {
				errors = errors.Add(fmt.Errorf("Requirement must be approved"))
			}
		}
		if len(errors) > 0 {
			return nil, internal_errors.CreateError(errors, internal_errors.RequirementsNotApproved, "Not all requirements are approved.")
		}
		t.adapterNotifications.SendRequestApproved(ctx, *user, tier.Name)
		setNewLimit = true
	case model.RequestStatus.Canceled:
		var errors internal_errors.Errors
		for _, requirementObject := range requirements {
			userRequirement, err := t.adapterUserRequirement.FindByRequirementIdAndUserId(ctx, requirementObject.ID, user.ID)
			if err != nil {
				errors = errors.Add(err)
			} else if userRequirement.Status == model.RequirementStatus.Pending {
				errors = errors.Add(fmt.Errorf("The requirement should not be pending"))
			}
		}
		if len(errors) > 0 {
			return nil, internal_errors.CreateError(errors, internal_errors.RequirementsIsPending, "The requirement should not be pending.")
		}
		t.adapterNotifications.SendRequestCanceled(ctx, *user, tier.Name)
	}
	oldRequest, err := t.adapterRequest.FindById(ctx, requestId)
	if err != nil {
		return nil, err
	}

	err = t.adapterRequest.Updates(ctx, &modelRequest)
	if err != nil {
		return nil, err
	}

	if setNewLimit {
		err = t.limitService.Set(user, tier.Limitations)
		if err != nil {
			return nil, err
		}
	}

	t.adapterLog.CreateRequestUpdateStatus(ctx, editor.ID, oldRequest, modelRequest)
	return &modelRequest, nil
}
