package requirement

import (
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/go-playground/validator/v10"
)

type DateBirthRequirement struct {
	elements                    []*model.TierRequirementElement
	adapterUserRequirementValue AdapterUserRequirementValue
	valid                       *validator.Validate
	Validate
}

type DateBirthForm struct {
	DateBirth string `json:"dateBirth" binding:""`
}

func (s *DateBirthRequirement) GetElements() []*model.TierRequirementElement {
	return s.elements
}

func (s *DateBirthRequirement) BindValues(ctx context.Context, userRequirement *model.UserRequirement) error {
	values, err := s.adapterUserRequirementValue.FindByUserRequirementId(ctx, userRequirement.ID)
	if err != nil {
		return err
	}
	userRequirement.Values = values
	return nil
}

func (s *DateBirthRequirement) SetValues(ctx context.Context, userRequirement *model.UserRequirement, values []model.UserRequirementValue) error {
	form := &DateBirthForm{}
	formErrors := s.ValidateForm(s.valid, form, values)
	if formErrors != nil {
		return formErrors
	}

	var errors internal_errors.Errors

	for i := 0; i < len(values); i++ {
		_, err := s.adapterUserRequirementValue.CreateOrUpdate(ctx, userRequirement.ID, values[i].Index, values[i].Value)
		if err != nil {
			errors = errors.Add(err)
		}
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

func (s *Service) createDateBirth() *DateBirthRequirement {
	var elements []*model.TierRequirementElement

	elements = append(elements, &model.TierRequirementElement{
		Name:  "Date of Birth",
		Type:  "date",
		Index: "date_birth",
	})

	return &DateBirthRequirement{
		elements:                    elements,
		adapterUserRequirementValue: s.adapterUserRequirementValue,
		valid:                       s.valid,
	}
}
