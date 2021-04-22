package adapter

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/jinzhu/gorm"
)

type TierLimitation struct {
	db *connection.DbConnect
}

func NewTierLimitation() *TierLimitation {
	return &TierLimitation{
		db: connection.GetDbConnect(),
	}
}

func (s *TierLimitation) FindByTierId(ctx context.Context, tierId uint64) ([]*model.TierLimitation, error) {
	var objects []*model.TierLimitation

	err := s.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where("tier_id = ?", tierId).Find(&objects).Error
	})

	return objects, err
}

func (s *TierLimitation) FindByTierIdIndex(ctx context.Context, tierId uint64, index string) (*model.TierLimitation, error) {
	object := model.TierLimitation{}

	err := s.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where(model.TierLimitation{TierId: tierId, Index: index}).First(&object).Error
	})

	return &object, err
}

func (r *TierLimitation) Updates(ctx context.Context, object *model.TierLimitation) error {
	return r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Model(model.TierLimitation{ID: object.ID}).Select("value").Updates(map[string]interface{}{"value": object.Value}).Error
	})
}
