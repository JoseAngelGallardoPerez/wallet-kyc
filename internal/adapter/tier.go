package adapter

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/jinzhu/gorm"
)

type Tier struct {
	db *connection.DbConnect
}

func NewTier() *Tier {
	return &Tier{
		db: connection.GetDbConnect(),
	}
}

func (r Tier) FindByCountryCode(ctx context.Context, countryCode string) (tiers []model.Tier, err error) {

	err = r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where("country_code = ?", countryCode).
			Preload("Requirements").
			Preload("Limitations").
			Order("level").
			Find(&tiers).Error
	})
	return tiers, err
}

func (r Tier) GetByCountryCodeAndLevel(ctx context.Context, countryCode string, level int) (model.Tier, error) {
	object := model.Tier{}

	err := r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where("country_code = ?", countryCode).
			Where("level = ?", level).
			Preload("Limitations").
			Order("level desc").
			First(&object).Error
	})

	return object, err
}

func (r Tier) FindById(ctx context.Context, id uint64) (tier model.Tier, err error) {
	err = r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.
			Where("id = ?", id).
			Preload("Requirements").
			Preload("Limitations").
			First(&tier).
			Error
	})

	if err != nil {
		return tier, internal_errors.CreateError(err, internal_errors.TierNotFound, "")
	}
	return tier, err
}

func (r *Tier) FindLastApprovedByUserIdCode(ctx context.Context, userId string, countryCode string) (model.Tier, error) {
	object := model.Tier{}

	err := r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.
			Select("tiers.*").
			Joins("left join user_requests on user_requests.tier_id = tiers.id").
			Where("user_requests.user_id = ?", userId).
			Where("user_requests.status = ?", model.RequestStatus.Approved).
			Where("tiers.country_code = ?", countryCode).
			Order("tiers.level DESC").
			First(&object).Error
	})

	return object, err
}
