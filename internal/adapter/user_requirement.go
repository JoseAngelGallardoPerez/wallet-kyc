package adapter

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/jinzhu/gorm"
)

type UserRequirement struct {
	db *connection.DbConnect
}

func NewUserRequirement() *UserRequirement {
	return &UserRequirement{
		db: connection.GetDbConnect(),
	}
}

func (t *UserRequirement) FindByRequirementIdAndUserId(ctx context.Context, requirementId uint64, userId string) (*model.UserRequirement, error) {
	object := model.UserRequirement{}

	err := t.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where(model.UserRequirement{TierRequirementId: requirementId, UserId: userId}).
			Preload("TierRequirement").
			First(&object).Error
	})

	if err != nil {
		return nil, err
	}

	return &object, nil
}

func (r *UserRequirement) Create(ctx context.Context, model *model.UserRequirement) error {
	return r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Create(model).Error
	})
}

func (r *UserRequirement) Updates(ctx context.Context, modelObject *model.UserRequirement) error {
	return r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Model(model.UserRequirement{}).
			Where(&model.UserRequirement{ID: modelObject.ID}).
			Updates(modelObject).Error
	})
}

func (r *UserRequirement) FindById(ctx context.Context, id uint64) (model *model.UserRequirement, err error) {
	err = r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where("id = ?", id).First(model).Error
	})

	if err != nil {
		return model, internal_errors.CreateError(err, internal_errors.RequirementNotFound, "")
	}
	return model, err
}
