package adapter

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/jinzhu/gorm"
)

type TierRequirement struct {
	db *connection.DbConnect
}

func NewTierRequirement() *TierRequirement {
	return &TierRequirement{
		db: connection.GetDbConnect(),
	}
}

func (t *TierRequirement) FindByTierId(ctx context.Context, tierId uint64) ([]model.TierRequirement, error) {
	var requirements []model.TierRequirement

	err := t.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where("tier_id = ?", tierId).Find(&requirements).Error
	})

	return requirements, err
}

func (r *TierRequirement) FindById(ctx context.Context, id uint64) (model model.TierRequirement, err error) {
	err = r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where("id = ?", id).
			Preload("Tier").
			First(&model).Error
	})

	if err != nil {
		return model, internal_errors.CreateError(err, internal_errors.RequirementNotFound, "")
	}
	return model, err
}
