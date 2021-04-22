package model

import "github.com/shopspring/decimal"

type TierLimitation struct {
	ID     uint64           `json:"id" gorm:"primary_key"`
	Value  *decimal.Decimal `json:"value"`
	Index  string           `json:"index"`
	Name   string           `json:"name"`
	TierId uint64
	Tier   *Tier `gorm:"foreignkey:TierId"`
}
