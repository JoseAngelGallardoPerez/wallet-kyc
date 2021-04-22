package requirement

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type EmailRequirement struct {
	elements               []*model.TierRequirementElement
	adapterUserRequirement AdapterUserRequirement
	adapterUser            AdapterUser
}

func (s *EmailRequirement) GetElements() []*model.TierRequirementElement {
	return s.elements
}

func (s *EmailRequirement) BindValues(ctx context.Context, userRequirement *model.UserRequirement) error {
	user, err := s.adapterUser.GetUserById(ctx, userRequirement.UserId)

	if err != nil {
		return err
	}

	userRequirement.Values = append(userRequirement.Values, model.UserRequirementValue{
		Index: "email",
		Value: user.Email,
	})

	if user.IsEmailConfirmed && userRequirement.Status != model.RequirementStatus.Approved {
		userRequirement.Status = model.RequirementStatus.Approved
		err = s.adapterUserRequirement.Updates(ctx, userRequirement)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *EmailRequirement) SetValues(ctx context.Context, userRequirement *model.UserRequirement, values []model.UserRequirementValue) error {
	return nil
}

func (s *Service) createEmail() *EmailRequirement {
	var elements []*model.TierRequirementElement

	elements = append(elements, &model.TierRequirementElement{
		Name:  "Email",
		Type:  "input",
		Index: "email",
	})

	return &EmailRequirement{
		elements:               elements,
		adapterUserRequirement: s.adapterUserRequirement,
		adapterUser:            s.adapterUser,
	}
}
