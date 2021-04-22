package model

import "time"

type UserRequirementValue struct {
	ID                uint64
	UserRequirementId uint64
	UserRequirement   *UserRequirement `gorm:"foreignkey:UserRequirementId;association_foreignkey:ID"`
	Index             string           `json:"index"`
	Value             string           `json:"value"`
	CreatedAt         time.Time        `json:"createdAt"`
	UpdatedAt         time.Time        `json:"updatedAt"`
}
