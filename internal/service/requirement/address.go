package requirement

import (
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/go-playground/validator/v10"
)

var AddressType = struct {
	UtilityBill          string
	TenancyAgreement     string
	PropertyDocument     string
	IncomeTaxCertificate string
	BankStatements       string
	ReferenceLetter      string
}{
	"utility_bill",
	"tenancy_agreement",
	"property_document",
	"income_tax_certificate",
	"bank_statements",
	"reference_letter",
}

type AddressRequirement struct {
	elements                    []*model.TierRequirementElement
	adapterUserRequirementValue AdapterUserRequirementValue
	valid                       *validator.Validate
	Validate
}

type AddressForm struct {
	AddressCountry string `json:"addressCountry" binding:"required,max=100"`
	AddressZipCode string `json:"addressZipCode" binding:"required,max=10"`
	Address1       string `json:"address1" binding:"required,max=100"`
	Address2       string `json:"address2" binding:"max=100"`
	AddressCity    string `json:"addressCity" binding:"required,max=100"`
	AddressDocType string `json:"addressDocType" binding:"required"`
	AddressScanned string `json:"addressScanned" binding:"omitempty,isFile"`
}

func (s *AddressRequirement) GetElements() []*model.TierRequirementElement {
	return s.elements
}

func (s *AddressRequirement) BindValues(ctx context.Context, userRequirement *model.UserRequirement) error {
	values, err := s.adapterUserRequirementValue.FindByUserRequirementId(ctx, userRequirement.ID)
	if err != nil {
		return err
	}
	userRequirement.Values = values
	return nil
}

func (s *AddressRequirement) SetValues(ctx context.Context, userRequirement *model.UserRequirement, values []model.UserRequirementValue) error {

	form := &AddressForm{}

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

func (s *Service) createAddress(requirementObject *model.TierRequirement) *AddressRequirement {
	var elements []*model.TierRequirementElement

	elements = append(elements, &model.TierRequirementElement{
		Name:  "Country",
		Type:  "input",
		Index: "address_country",
	})

	elements = append(elements, &model.TierRequirementElement{
		Name:  "Zip/Postal Code",
		Type:  "input",
		Index: "address_zip_code",
	})

	elements = append(elements, &model.TierRequirementElement{
		Name:  "Address",
		Type:  "input",
		Index: "address_1",
	})

	elements = append(elements, &model.TierRequirementElement{
		Name:  "Address (2nd line)",
		Type:  "input",
		Index: "address_2",
	})

	elements = append(elements, &model.TierRequirementElement{
		Name:  "City",
		Type:  "input",
		Index: "address_city",
	})

	var options []*model.TierRequirementElementOption

	code := requirementObject.Tier.CountryCode

	if code == "GHA" || code == "NGA" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "Utility Bill",
			Value: AddressType.UtilityBill,
		})
	}

	if code == "GHA" || code == "NGA" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "Tenancy Agreement",
			Value: AddressType.TenancyAgreement,
		})
	}

	if code == "NGA" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "Property Document",
			Value: AddressType.PropertyDocument,
		})
	}

	if code == "GHA" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "Income Tax Certificate",
			Value: AddressType.IncomeTaxCertificate,
		})
	}

	if code == "GHA" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "Bank statements",
			Value: AddressType.BankStatements,
		})
	}

	if code == "GHA" {
		options = append(options, &model.TierRequirementElementOption{
			Name:  "Reference letter or Employee’s reference letter",
			Value: AddressType.ReferenceLetter,
		})
	}

	elements = append(elements, &model.TierRequirementElement{
		Name:    "Document Type",
		Type:    "options",
		Index:   "address_doc_type",
		Options: options,
	})

	elements = append(elements, &model.TierRequirementElement{
		Name:  "Scanned Copy",
		Type:  "file",
		Index: "address_scanned",
	})

	return &AddressRequirement{
		elements:                    elements,
		adapterUserRequirementValue: s.adapterUserRequirementValue,
		valid:                       s.valid,
	}
}
