package forms

import (
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
)

type AdminUpdateTier struct {
	CountryCode string                      `json:"countryCode"`
	Limitations []AdminUpdateTierLimitation `json:"limitations" binding:"required,dive"`
}

type AdminUpdateTierLimitation struct {
	Index string           `json:"index" binding:"required"`
	Value *decimal.Decimal `json:"value" binding:""`
}

func (s AdminUpdateTier) Unserialize(ctx context.Context, adminUpdateTier AdminUpdateTier) (tier model.Tier, err error) {

	publicJson, err := json.Marshal(adminUpdateTier)

	if nil != err {
		err = fmt.Errorf("can't marshal json")
		return
	}

	err = json.Unmarshal(publicJson, &tier)
	if nil != err {
		err = fmt.Errorf("can't unmarshal json")
		return
	}

	var errors internal_errors.Errors

	for _, v := range adminUpdateTier.Limitations {
		public, err := json.Marshal(v)
		if err != nil {
			errors = errors.Add(err)
		}

		lim := model.TierLimitation{}
		err = json.Unmarshal(public, &lim)
		if err != nil {
			errors = errors.Add(err)
		}

		tier.Limitations = append(tier.Limitations, &lim)
	}

	if len(errors) > 0 {
		err = errors
	}
	return
}
