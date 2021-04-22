package forms

import "github.com/Confialink/wallet-kyc/internal/model"

type UpdateRequirement struct {
	Values []model.UserRequirementValue `json:"values" binding:"required"`
}
