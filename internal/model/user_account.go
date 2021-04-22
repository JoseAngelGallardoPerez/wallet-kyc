package model

import (
	"github.com/shopspring/decimal"
)

type UserAccount struct {
	UserId       string
	Balance      decimal.Decimal
	CurrencyCode string
}
