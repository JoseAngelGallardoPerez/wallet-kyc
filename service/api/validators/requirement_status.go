package validators

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func requirementStatus(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	fieldStr := fl.Field().Interface().(string)
	return fieldStr == model.RequirementStatus.Approved || fieldStr == model.RequirementStatus.Canceled
}
