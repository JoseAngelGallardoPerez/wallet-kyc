package adapter

import (
	currenciespb "github.com/Confialink/wallet-currencies/rpc/currencies"
	"github.com/Confialink/wallet-kyc/internal/connection"
	"context"
	"github.com/shopspring/decimal"
)

type Currency struct {
	rpcCurrencies *connection.RpcCurrencies
}

func NewCurrency() *Currency {
	return &Currency{
		rpcCurrencies: connection.GetRpcCurrencies(),
	}
}

func (s *Currency) GetCurrenciesRateValueByCodes(currencyCodeFrom, currencyCodeTo string) (decimal.Decimal, error) {
	request := currenciespb.CurrenciesRateValueRequest{
		CurrencyCodeFrom: currencyCodeFrom,
		CurrencyCodeTo:   currencyCodeTo,
	}
	client := s.rpcCurrencies.Client

	if response, err := client.GetCurrenciesRateValueByCodes(context.Background(), &request); err != nil {
		return decimal.Zero, err
	} else {
		return decimal.NewFromString(response.Value)
	}
}
