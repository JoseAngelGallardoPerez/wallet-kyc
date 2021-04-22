package limit

import (
	"github.com/Confialink/wallet-accounts/rpc/limit"
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/pkg/errors"
)

const (
	MaxTotalBalance       = "max_total_balance"
	MaxCreditPerTransfer  = "max_credit_per_transfer"
	MaxDebitPerTransfer   = "max_debit_per_transfer"
	MaxTotalDebitPerDay   = "max_total_debit_per_day"
	MaxTotalDebitPerMonth = "max_total_debit_per_month"
)

var tierLimitIndexToLimitName = map[string]limit.LimitName{
	MaxTotalBalance:       limit.LimitName_MAX_TOTAL_BALANCE,
	MaxCreditPerTransfer:  limit.LimitName_MAX_CREDIT_PER_TRANSFER,
	MaxDebitPerTransfer:   limit.LimitName_MAX_DEBIT_PER_TRANSFER,
	MaxTotalDebitPerDay:   limit.LimitName_MAX_TOTAL_DEBIT_PER_DAY,
	MaxTotalDebitPerMonth: limit.LimitName_MAX_TOTAL_DEBIT_PER_MONTH,
}

type Service struct {
	rpcLimit *connection.RpcLimit
}

// NewService is Service constructor
func NewService(rpcLimit *connection.RpcLimit) *Service {
	return &Service{rpcLimit: rpcLimit}
}

// Set creates/updates limit for a given user based on the tier limitations
func (s *Service) Set(user *model.User, tierLimits []*model.TierLimitation) error {
	requestLimits := make([]*limit.LimitWithId, len(tierLimits))
	for i := 0; i < len(tierLimits); i++ {
		tierLimit := tierLimits[i]
		id := &limit.LimitId{
			Name:     tierLimitIndexToLimitName[tierLimit.Index],
			Entity:   "user",
			EntityId: user.ID,
		}
		lim := &limit.Limit{}
		if tierLimit.Value == nil {
			lim.NoLimit = true
		} else {
			lim.Amount = tierLimit.Value.String()
			lim.CurrencyCode = user.CurrencyCode
		}
		requestLimits[i] = &limit.LimitWithId{
			Limit:   lim,
			LimitId: id,
		}
	}

	request := &limit.SetLimitsRequest{
		Limits: requestLimits,
	}
	_, err := s.rpcLimit.Client.Set(context.Background(), request)
	if err != nil {
		return errors.Wrapf(
			err,
			"failed to set limits for user %s",
			user.ID,
		)
	}
	return nil
}
