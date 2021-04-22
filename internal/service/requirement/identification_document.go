package requirement

import (
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/go-playground/validator/v10"
)

var IdentificationDocumentType = struct {
	NationalNumber             string
	PassportNumber             string
	ForeignPassportNumber      string
	VoterCardNumber            string
	MilitaryNumber             string
	DriverNumber               string
	ForeignerCertificateNumber string
	NationalHealthInsurance    string
}{
	"national_number",
	"passport_number",
	"foreign_passport_number",
	"voter_card_number",
	"military_number",
	"driver_number",
	"foreigner_certificate_number",
	"national_health_insurance",
}

type IdentificationDocumentRequirement struct {
	elements                    []*model.TierRequirementElement
	adapterUserRequirementValue AdapterUserRequirementValue
	valid                       *validator.Validate
	Validate
}

type IdentificationDocumentForm struct {
	IdentificationType    string `json:"identificationType" binding:""`
	IdentificationNumber  string `json:"identificationNumber" binding:""`
	IdentificationScanned string `json:"identificationScanned" binding:"omitempty,isFile"`
}

func (s *IdentificationDocumentRequirement) GetElements() []*model.TierRequirementElement {
	return s.elements
}

func (s *IdentificationDocumentRequirement) BindValues(ctx context.Context, userRequirement *model.UserRequirement) error {
	values, err := s.adapterUserRequirementValue.FindByUserRequirementId(ctx, userRequirement.ID)
	if err != nil {
		return err
	}
	userRequirement.Values = values
	return nil
}

func (s *IdentificationDocumentRequirement) SetValues(ctx context.Context, userRequirement *model.UserRequirement, values []model.UserRequirementValue) error {
	form := &IdentificationDocumentForm{}
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

func (s *Service) createIdentificationDocument(requirementObject *model.TierRequirement) *IdentificationDocumentRequirement {
	var elements []*model.TierRequirementElement

	var options []*model.TierRequirementElementOption

	code := requirementObject.Tier.CountryCode

	options = append(options, &model.TierRequirementElementOption{
		Name:  "National ID Number",
		Value: IdentificationDocumentType.NationalNumber,
	})

	options = append(options, &model.TierRequirementElementOption{
		Name:  "Passport Number",
		Value: IdentificationDocumentType.PassportNumber,
	})

	if code == "KEN" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "Foreign Passport Number",
			Value: IdentificationDocumentType.ForeignPassportNumber,
		})
	}

	if code == "GHA" || code == "NGA" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "Voter's Card Number",
			Value: IdentificationDocumentType.VoterCardNumber,
		})
	}

	if code == "KEN" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "Military ID Number",
			Value: IdentificationDocumentType.MilitaryNumber,
		})
	}

	if code == "GHA" || code == "NGA" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "Driver's License Number",
			Value: IdentificationDocumentType.DriverNumber,
		})
	}

	if code == "KEN" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "Foreigner Certificate Number",
			Value: IdentificationDocumentType.ForeignerCertificateNumber,
		})
	}

	if code == "GHA" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "National Health Insurance ID Number",
			Value: IdentificationDocumentType.NationalHealthInsurance,
		})
	}

	elements = append(elements, &model.TierRequirementElement{
		Name:    "Document Type",
		Type:    "options",
		Index:   "identification_type",
		Options: options,
	})

	elements = append(elements, &model.TierRequirementElement{
		Name:  "Number",
		Type:  "input",
		Index: "identification_number",
	})

	elements = append(elements, &model.TierRequirementElement{
		Name:  "Scanned Copy",
		Type:  "file",
		Index: "identification_scanned",
	})

	return &IdentificationDocumentRequirement{
		elements:                    elements,
		adapterUserRequirementValue: s.adapterUserRequirementValue,
		valid:                       s.valid,
	}
}
