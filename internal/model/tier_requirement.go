package model

type TierRequirement struct {
	ID               uint64 `json:"id" gorm:"primary_key"`
	TierId           uint64
	Tier             *Tier  `gorm:"foreignkey:TierId"`
	Name             string `json:"name"`
	FormIndex        string
	Options          string
	Elements         []*TierRequirementElement `json:"elements"`
	UserRequirements []*UserRequirement        `json:"userRequirements"`
}
