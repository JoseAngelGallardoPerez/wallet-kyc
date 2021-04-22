package action

import (
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type AdminRequirementUpdateStatus struct {
	adapterUserRequirement AdapterUserRequirement
	adapterTierRequirement AdapterTierRequirement
	adapterUser            AdapterUser
	adapterNotifications   AdapterNotifications
	adapterLog             AdapterLog
}

func NewAdminRequirementUpdateStatus(
	adapterUserRequirement AdapterUserRequirement,
	adapterTierRequirement AdapterTierRequirement,
	adapterUser AdapterUser,
	adapterNotifications AdapterNotifications,
	adapterLog AdapterLog,
) *AdminRequirementUpdateStatus {
	return &AdminRequirementUpdateStatus{
		adapterUserRequirement: adapterUserRequirement,
		adapterTierRequirement: adapterTierRequirement,
		adapterUser:            adapterUser,
		adapterNotifications:   adapterNotifications,
		adapterLog:             adapterLog,
	}
}

// Status update for user defined requirements
func (t *AdminRequirementUpdateStatus) Do(ctx context.Context, editor model.User, requirementId uint64, userId string, status string) error {
	requirement, err := t.adapterTierRequirement.FindById(ctx, requirementId)
	if err != nil {
		return err
	}

	user, err := t.adapterUser.GetUserById(ctx, userId)
	if err != nil {
		return err
	}

	modelUserRequirement, err := t.adapterUserRequirement.FindByRequirementIdAndUserId(ctx, requirementId, userId)
	if err != nil {
		return err
	}

	if modelUserRequirement.Status == model.RequirementStatus.Approved {
		return internal_errors.CreateError(nil, internal_errors.RequestAlreadyExists, "This requirement has already been approved.")
	}

	oldUserRequirement, err := t.adapterUserRequirement.FindByRequirementIdAndUserId(ctx, requirementId, userId)
	if err != nil {
		return err
	}

	if modelUserRequirement.Status == status {
		return nil
	}

	modelUserRequirement.Status = status

	t.adapterNotifications.SendStatusDocChanged(ctx, *user, requirement.Name, status)

	err = t.adapterUserRequirement.Updates(ctx, modelUserRequirement)
	if err != nil {
		return err
	}

	oldUserRequirement.User = *user

	t.adapterLog.CreateRequirementUpdateStatus(ctx, editor.ID, *oldUserRequirement, *modelUserRequirement)
	return nil
}
