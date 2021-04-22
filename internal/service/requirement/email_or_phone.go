package requirement

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type EmailOrPhoneRequirement struct {
	elements               []*model.TierRequirementElement
	adapterUser            AdapterUser
	adapterUserRequirement AdapterUserRequirement
}

func (s *EmailOrPhoneRequirement) GetElements() []*model.TierRequirementElement {
	return s.elements
}

func (s *EmailOrPhoneRequirement) BindValues(ctx context.Context, userRequirement *model.UserRequirement) error {
	user, err := s.adapterUser.GetUserById(ctx, userRequirement.UserId)
	if err != nil {
		return err
	}
	var value string
	if user.IsEmailConfirmed {
		value = user.Email
	} else {
		value = user.Phone
	}

	userRequirement.Values = append(userRequirement.Values, model.UserRequirementValue{
		Index: "email_or_phone",
		Value: value,
	})

	if userRequirement.Status != model.RequirementStatus.Approved {
		userRequirement.Status = model.RequirementStatus.Approved
		err = s.adapterUserRequirement.Updates(ctx, userRequirement)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *EmailOrPhoneRequirement) SetValues(ctx context.Context, userRequirement *model.UserRequirement, values []model.UserRequirementValue) error {
	return nil
}

func (s *Service) createEmailOrPhone() *EmailOrPhoneRequirement {
	var elements []*model.TierRequirementElement

	elements = append(elements, &model.TierRequirementElement{
		Name:  "Email or Phone",
		Type:  "input",
		Index: "email_or_phone",
	})

	return &EmailOrPhoneRequirement{
		elements:               elements,
		adapterUser:            s.adapterUser,
		adapterUserRequirement: s.adapterUserRequirement,
	}
}
