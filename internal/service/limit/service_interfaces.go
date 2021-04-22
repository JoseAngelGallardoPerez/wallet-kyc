package limit

import (
	"github.com/shopspring/decimal"
)

type ServiceCurrencies interface {
	Convert(amount decimal.Decimal, currencyCodeFrom, currencyCodeTo string) (decimal.Decimal, error)
}
