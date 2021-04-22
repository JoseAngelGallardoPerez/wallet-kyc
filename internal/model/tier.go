package model

type Tier struct {
	ID           uint64             `json:"id" gorm:"primary_key:true"`
	CountryCode  string             `json:"countryCode"`
	Level        int                `json:"level"`
	Name         string             `json:"name"`
	Features     []*TierFeature     `gorm:"many2many:tier_features;association_jointable_foreignkey:feature_id;jointable_foreignkey:tier_id;"`
	Requirements []*TierRequirement `json:"requirements" gorm:"foreignkey:TierId"`
	Limitations  []*TierLimitation  `json:"limitations" gorm:"foreignkey:TierId"`
	Requests     []*UserRequest     `json:"requests" gorm:"foreignkey:TierId"`
}
