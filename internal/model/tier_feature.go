package model

type TierFeature struct {
	ID    uint64 `gorm:"primary_key"`
	Index string
}
