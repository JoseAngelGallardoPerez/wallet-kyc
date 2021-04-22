package requirement

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type PhoneRequirement struct {
	elements               []*model.TierRequirementElement
	adapterUserRequirement AdapterUserRequirement
	adapterUser            AdapterUser
}

func (s *PhoneRequirement) GetElements() []*model.TierRequirementElement {
	return s.elements
}

func (s *PhoneRequirement) BindValues(ctx context.Context, userRequirement *model.UserRequirement) error {
	user, err := s.adapterUser.GetUserById(ctx, userRequirement.UserId)
	if err != nil {
		return err
	}

	userRequirement.Values = append(userRequirement.Values, model.UserRequirementValue{
		Index: "phone",
		Value: user.Phone,
	})

	if user.IsPhoneNumberConfirmed && userRequirement.Status != model.RequirementStatus.Approved {
		userRequirement.Status = model.RequirementStatus.Approved
		err = s.adapterUserRequirement.Updates(ctx, userRequirement)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *PhoneRequirement) SetValues(ctx context.Context, userRequirement *model.UserRequirement, values []model.UserRequirementValue) error {
	return nil
}

func (s *Service) createPhone() *PhoneRequirement {
	var elements []*model.TierRequirementElement

	elements = append(elements, &model.TierRequirementElement{
		Name:  "Phone",
		Type:  "input",
		Index: "phone",
	})

	return &PhoneRequirement{
		elements:               elements,
		adapterUserRequirement: s.adapterUserRequirement,
		adapterUser:            s.adapterUser,
	}
}
