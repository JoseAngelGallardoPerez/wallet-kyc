package limit

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/shopspring/decimal"
	"time"
)

type AdapterAccountsLimit interface {
	GetBalanceByUser(ctx context.Context, user model.User) (decimal.Decimal, error)
	GetSumOutgoingByUserFromDateToDate(ctx context.Context, user model.User, from time.Time, to time.Time) (decimal.Decimal, error)
}
