package adapter

import (
	//accountspb "github.com/Confialink/wallet-accounts/rpc/accounts"
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/shopspring/decimal"
	"time"
)

type UserAccount struct {
	rpcAccount *connection.RpcLimit
}

func NewUserAccount() *UserAccount {
	return &UserAccount{
		rpcAccount: connection.GetRpcLimit(),
	}
}

// Get user balance in local currency
func (s *UserAccount) GetBalanceByUser(ctx context.Context, user model.User) (decimal.Decimal, error) {
	return decimal.Zero, nil
	/*balanceStr, err := s.rpcAccount.Client.GetBalanceByUser(ctx, &accountspb.GetBalanceByUserReq{
		UserID:       user.ID,
		CurrencyCode: user.CurrencyCode,
	})
	if err != nil {
		return decimal.Decimal{}, err
	}

	balance, err := decimal.NewFromString(balanceStr.Balance)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return balance, nil*/
}

// Get the amount of outgoing user transactions for a given time filter in local currency
func (s *UserAccount) GetSumOutgoingByUserFromDateToDate(ctx context.Context, user model.User, from time.Time, to time.Time) (decimal.Decimal, error) {
	return decimal.Zero, nil
	/*sumStr, err := s.rpcAccount.Client.GetSumOutgoingByUserFromDateToDate(ctx, &accountspb.GetSumOutgoingByUserFromDateToDateReq{
		UserID:       user.ID,
		CurrencyCode: user.CurrencyCode,
		TimeFrom:     from.Format(time.RFC3339),
		TimeTo:       to.Format(time.RFC3339),
	})
	if err != nil {
		return decimal.Decimal{}, err
	}

	sum, err := decimal.NewFromString(sumStr.Sum)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return sum, nil*/
}
