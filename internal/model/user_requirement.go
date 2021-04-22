package model

import "time"

var RequirementStatus = struct {
	NotFilled string
	Canceled  string
	Pending   string
	Approved  string
	Waiting   string
}{
	"not-filled",
	"canceled",
	"pending",
	"approved",
	"waiting",
}

type UserRequirement struct {
	ID                uint64 `gorm:"primary_key"`
	Status            string
	UserId            string
	User              User `json:"user"`
	TierRequirementId uint64
	TierRequirement   *TierRequirement       `gorm:"foreignkey:TierRequirementId"`
	Values            []UserRequirementValue `json:"values" gorm:"foreignkey:UserRequirementId"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
