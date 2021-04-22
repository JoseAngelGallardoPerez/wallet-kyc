package adapter

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/model"
	"context"
	"github.com/jinzhu/gorm"
)

type Country struct {
	db *connection.DbConnect
}

func NewCountry() *Country {
	return &Country{
		db: connection.GetDbConnect(),
	}
}

func (r Country) FindAll(ctx context.Context) (objects []model.Country, err error) {
	err = r.db.ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		return db.Find(&objects).Error
	})
	return objects, err
}
