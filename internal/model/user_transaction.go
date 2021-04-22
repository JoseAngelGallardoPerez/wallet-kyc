package model

import (
	"github.com/shopspring/decimal"
)

type UserTransaction struct {
	UserId       string
	User         User
	CurrencyCode string
	Amount       decimal.Decimal
}
