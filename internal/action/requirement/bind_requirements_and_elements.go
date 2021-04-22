package requirement

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type BindRequirementsAndElements struct {
	adapterUserRequirement AdapterUserRequirement
	serviceRequirement     ServiceRequirement
}

func NewBindRequirementsAndElements(
	adapterUserRequirement AdapterUserRequirement,
	serviceRequirement ServiceRequirement,
) *BindRequirementsAndElements {
	return &BindRequirementsAndElements{
		adapterUserRequirement: adapterUserRequirement,
		serviceRequirement:     serviceRequirement,
	}
}

func (s *BindRequirementsAndElements) Do(ctx context.Context, requirement *model.TierRequirement, user *model.User) (err error) {

	err = s.serviceRequirement.BindElements(ctx, requirement)
	if err != nil {
		return err
	}

	userRequirement, err := s.adapterUserRequirement.FindByRequirementIdAndUserId(ctx, requirement.ID, user.ID)

	if err != nil {
		userRequirement = &model.UserRequirement{
			TierRequirementId: requirement.ID,
			UserId:            user.ID,
			Status:            model.RequirementStatus.NotFilled,
		}

		err := s.adapterUserRequirement.Create(ctx, userRequirement)
		if err != nil {
			return err
		}
	}

	if err == nil {
		userRequirement.TierRequirement = requirement
		err = s.serviceRequirement.BindValues(ctx, userRequirement)

		if err != nil {
			return
		}
		requirement.UserRequirements = append(requirement.UserRequirements, userRequirement)
	}
	return nil
}
