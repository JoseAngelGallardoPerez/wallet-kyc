package adapter

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/jinzhu/gorm"
)

type UserRequirementValue struct {
	db *connection.DbConnect
}

func NewUserRequirementValue() *UserRequirementValue {
	return &UserRequirementValue{
		db: connection.GetDbConnect(),
	}
}

func (r *UserRequirementValue) Create(ctx context.Context, model *model.UserRequirementValue) error {
	return r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Create(model).Error
	})
}

func (r *UserRequirementValue) Update(ctx context.Context, model *model.UserRequirementValue) error {
	return r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Model(model).Select("value").Updates(map[string]interface{}{"value": model.Value}).Error
	})
}

func (t *UserRequirementValue) FindByUserRequirementIdAndIndex(ctx context.Context, userRequirementId uint64, index string) (*model.UserRequirementValue, error) {
	object := model.UserRequirementValue{}

	err := t.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where(model.UserRequirementValue{
			UserRequirementId: userRequirementId,
			Index:             index,
		}).First(&object).Error
	})

	if err != nil {
		return nil, err
	}

	return &object, nil
}

func (t *UserRequirementValue) FindByUserRequirementId(ctx context.Context, userRequirementId uint64) (objects []model.UserRequirementValue, err error) {

	err = t.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Where(model.UserRequirementValue{
			UserRequirementId: userRequirementId,
		}).Find(&objects).Error
	})

	if err != nil {
		return nil, err
	}

	return objects, nil
}

func (s *UserRequirementValue) CreateOrUpdate(ctx context.Context, userRequirementId uint64, index string, value string) (*model.UserRequirementValue, error) {
	userRequirementValue, err := s.FindByUserRequirementIdAndIndex(ctx, userRequirementId, index)
	if err == nil {
		userRequirementValue.Value = value
		err = s.Update(ctx, userRequirementValue)
	} else {
		userRequirementValue = &model.UserRequirementValue{
			UserRequirementId: userRequirementId,
			Index:             index,
			Value:             value,
		}
		err = s.Create(ctx, userRequirementValue)
	}
	if err != nil {
		return nil, err
	}
	return userRequirementValue, nil
}
