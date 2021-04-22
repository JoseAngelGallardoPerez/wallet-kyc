package action

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
)

type AdminCountries struct {
	adapterCountries AdapterCountries
}

func NewAdminCountries(
	adapterCountries AdapterCountries,
) *AdminCountries {
	return &AdminCountries{
		adapterCountries: adapterCountries,
	}
}

func (s *AdminCountries) Do(ctx context.Context) (objects []model.Country, err error) {
	return s.adapterCountries.FindAll(ctx)
}
