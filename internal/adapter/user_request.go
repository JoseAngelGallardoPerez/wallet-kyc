package adapter

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/jinzhu/gorm"
)

type UserRequest struct {
	db *connection.DbConnect
}

func NewUserRequest() *UserRequest {
	return &UserRequest{
		db: connection.GetDbConnect(),
	}
}

func (r *UserRequest) Create(ctx context.Context, request *model.UserRequest) error {
	return r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Create(request).Error
	})
}

func (r *UserRequest) FindByUserIdAndTierIdAndStatuses(ctx context.Context, userId string, tierId uint64, statuses []string) ([]model.UserRequest, error) {
	var requirements []model.UserRequest

	err := r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where("tier_id = ?", tierId).
			Where("user_id = ?", userId).
			Where("status in (?)", statuses).
			Find(&requirements).Error
	})
	return requirements, err
}

func (r *UserRequest) GetByUserIdAndCountryCodeAndStatus(ctx context.Context, userId string, countryCode string, status string) (model.UserRequest, error) {
	var userRequest model.UserRequest

	err := r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.
			Select("user_requests.*").
			Joins("left join tiers on user_requests.tier_id = tiers.id").
			Where("tiers.country_code = ?", countryCode).
			Where("user_requests.user_id = ?", userId).
			Where("user_requests.status = ?", status).
			Preload("Tier").
			First(&userRequest).Error
	})
	return userRequest, err
}

func (r *UserRequest) FindByQuery(ctx context.Context, query *gorm.DB) ([]model.UserRequest, error) {
	var requirements []model.UserRequest

	err := r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return query.Find(&requirements).Error
	})
	return requirements, err
}

func (r *UserRequest) FindByQueryCount(ctx context.Context, query *gorm.DB) (uint64, error) {

	var count uint64

	err := r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return query.Model(&model.UserRequest{}).Count(&count).Error
	})
	return count, err
}

func (r *UserRequest) FindById(ctx context.Context, id uint64) (model model.UserRequest, err error) {
	err = r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where("id = ?", id).First(&model).Error
	})

	if err != nil {
		return model, internal_errors.CreateError(err, internal_errors.RequestNotFound, "")
	}
	return model, err
}

func (r *UserRequest) Updates(ctx context.Context, request *model.UserRequest) error {
	return r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Model(request).Updates(request).Error
	})
}

func (t *UserRequest) FindByTierIdAndUserId(ctx context.Context, tierId uint64, userId string) (*model.UserRequest, error) {
	object := model.UserRequest{}

	err := t.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where(model.UserRequest{TierId: tierId, UserId: userId}).First(&object).Error
	})

	if err != nil {
		return nil, err
	}

	return &object, nil
}
