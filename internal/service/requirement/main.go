package requirement

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"log"
	"reflect"
	"strconv"
)

type Requirement interface {
	GetElements() []*model.TierRequirementElement
	BindValues(ctx context.Context, userRequirement *model.UserRequirement) error
	SetValues(ctx context.Context, userRequirement *model.UserRequirement, UserRequirementValue []model.UserRequirementValue) error
}

const (
	Email                        = "email"
	Phone                        = "phone"
	EmailOrPhone                 = "email_or_phone"
	FullName                     = "full_name"
	MotherMaidenName             = "mother_maiden_name"
	SelfiePhoto                  = "selfie_photo"
	DateBirth                    = "date_birth"
	IdentificationDocument       = "identification_document"
	BankNumber                   = "bank_number"
	Address                      = "address"
	BusinessRegistrationDocument = "business_registration_document"
	DirectorsAndBeneficial       = "directors_and_beneficial"
	SixMonthsStatement           = "six_months_statement"
	BankGuaranty                 = "bank_guaranty"
	IncomeTaxCertificates        = "income_tax_certificates"
	ShareholdersDocuments        = "shareholders_documents"
)

// Interface is a validator interface
type Interface interface {
	Struct(current interface{}) error
	StructExcept(current interface{}, fields ...string) error
	RegisterValidation(key string, fn validator.Func) error
	RegisterStructValidation(fn validator.StructLevelFunc, types ...interface{})
}

type Service struct {
	adapterUserRequirementValue AdapterUserRequirementValue
	adapterUserRequirement      AdapterUserRequirement
	adapterUser                 AdapterUser
	adapterTier                 AdapterTier
	adapterFile                 AdapterFile
	valid                       *validator.Validate
}

func NewService(
	adapterUserRequirementValue AdapterUserRequirementValue,
	adapterUserRequirement AdapterUserRequirement,
	adapterUser AdapterUser,
	adapterTier AdapterTier,
	adapterFile AdapterFile,
) *Service {

	service := &Service{
		adapterUserRequirementValue: adapterUserRequirementValue,
		adapterUserRequirement:      adapterUserRequirement,
		adapterUser:                 adapterUser,
		adapterTier:                 adapterTier,
		adapterFile:                 adapterFile,
	}

	service.valid = binding.Validator.Engine().(*validator.Validate)

	err := service.valid.RegisterValidation("isFile", func(fl validator.FieldLevel) bool {
		if fl.Field().Kind() != reflect.String {
			return false
		}
		fieldStr := fl.Field().Interface().(string)

		fieldInt, _ := strconv.ParseInt(fieldStr, 10, 64)

		_, err := service.adapterFile.FindById(context.Background(), uint64(fieldInt))
		return err == nil
	})

	if err != nil {
		log.Println(err)
	}

	return service
}

func (s *Service) newRequirement(requirementObject *model.TierRequirement) Requirement {
	switch requirementObject.FormIndex {
	case DateBirth:
		return s.createDateBirth()
	case EmailOrPhone:
		return s.createEmailOrPhone()
	case FullName:
		return s.createFullName()
	case Email:
		return s.createEmail()
	case IdentificationDocument:
		return s.createIdentificationDocument(requirementObject)
	case MotherMaidenName:
		return s.createMotherMaidenName()
	case Phone:
		return s.createPhone()
	case SelfiePhoto:
		return s.createSelfiePhoto()
	case BankNumber:
		return s.createBankNumber()
	case Address:
		return s.createAddress(requirementObject)
	case BusinessRegistrationDocument:
		return s.createBusinessRegistrationDocument()
	case DirectorsAndBeneficial:
		return s.createDirectorsAndBeneficialt()
	case SixMonthsStatement:
		return s.createSixMonthsStatement()
	case BankGuaranty:
		return s.createBankGuaranty()
	case IncomeTaxCertificates:
		return s.createIncomeTaxCertificates()
	case ShareholdersDocuments:
		return s.createShareholdersDocuments()
	default:
		return nil
	}
}

func (s *Service) BindElements(ctx context.Context, requirementObject *model.TierRequirement) error {
	err := s.bindTier(ctx, requirementObject)
	if err != nil {
		return err
	}

	requirement := s.newRequirement(requirementObject)
	if requirement != nil {
		requirementObject.Elements = requirement.GetElements()
	}
	return nil
}

func (s *Service) SetValue(ctx context.Context, userRequirement *model.UserRequirement, userRequirementValue []model.UserRequirementValue) error {
	err := s.bindTier(ctx, userRequirement.TierRequirement)
	if err != nil {
		return err
	}

	requirement := s.newRequirement(userRequirement.TierRequirement)
	if requirement != nil {
		err := requirement.SetValues(ctx, userRequirement, userRequirementValue)

		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) BindValues(ctx context.Context, userRequirement *model.UserRequirement) error {
	err := s.bindTier(ctx, userRequirement.TierRequirement)
	if err != nil {
		return err
	}

	requirement := s.newRequirement(userRequirement.TierRequirement)
	if requirement != nil {
		err := requirement.BindValues(ctx, userRequirement)

		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) bindTier(ctx context.Context, tierRequirement *model.TierRequirement) error {
	if tierRequirement.Tier == nil {
		tier, err := s.adapterTier.FindById(ctx, tierRequirement.TierId)
		if err != nil {
			return err
		}

		tierRequirement.Tier = &tier
	}
	return nil
}

type Validate struct {
}

func (s *Validate) ValidateForm(valid *validator.Validate, form interface{}, values []model.UserRequirementValue) error {
	for i := 0; i < len(values); i++ {
		v := reflect.ValueOf(form)
		if v.Kind() != reflect.Ptr {
			name := values[i].Index
			return errors.New(
				fmt.Sprintf(
					"%s[%s]: %s",
					"internal/service/requirement/main.go.ValidateForm",
					name,
					"the 'form' element must be addressable (passed by reference)",
				),
			)
		}
		v = reflect.Indirect(v).FieldByName(strcase.ToCamel(values[i].Index))
		v.SetString(values[i].Value)
	}
	errs := valid.Struct(form)
	if errs != nil {
		return errs
	}
	return nil
}
