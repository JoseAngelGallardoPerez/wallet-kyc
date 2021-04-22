package connection

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type DbConnect struct {
	db *gorm.DB
}

type DbConfigs struct {
	Driver string
	Host   string
	Port   string
	Schema string

	User     string
	Password string

	IsDebugMode bool
}

func (c DbConfigs) Validate() (isValid bool) {
	isValid = true
	if len(c.Driver) == 0 {
		isValid = false
		Logger.Warn("Parameter not configured", "DbConfigs.Driver", c.Driver)
	}

	if len(c.Host) == 0 {
		isValid = false
		Logger.Warn("Parameter not configured", "DbConfigs.Host", c.Host)
	}

	if len(c.Port) == 0 {
		isValid = false
		Logger.Warn("Parameter not configured", "DbConfigs.Port", c.Port)
	}

	if len(c.Schema) == 0 {
		isValid = false
		Logger.Warn("Parameter not configured", "DbConfigs.Schema", c.Schema)
	}

	if len(c.User) == 0 {
		isValid = false
		Logger.Warn("Parameter not configured", "DbConfigs.User", c.User)
	}

	if len(c.Password) == 0 {
		isValid = false
		Logger.Warn("Parameter not configured", "DbConfigs.Password")
	}
	return
}

var db *DbConnect

type ctxTxKey struct{}

func (db *DbConnect) WithTx(ctx context.Context, fn func(ctx context.Context) error) (err error) {

	tx, alreadyHasTx := ctx.Value(ctxTxKey{}).(*gorm.DB)
	if !alreadyHasTx {
		tx = db.db.BeginTx(ctx, nil)

		if tx.Error != nil {
			return errors.WithStack(tx.Error)
		}
		ctx = context.WithValue(ctx, ctxTxKey{}, tx)
	}

	err = errors.WithStack(fn(ctx))

	if alreadyHasTx {
		return err
	}

	if err == nil {
		tx.Commit()
		if tx.Error != nil {
			return errors.WithStack(tx.Error)
		}
	}

	tx.Rollback()
	if tx.Error != nil {
		return errors.WithStack(tx.Error)
	}

	return err
}

func (db *DbConnect) ExtractTx(ctx context.Context, fn func(context.Context, *gorm.DB) error) error {

	tx, alreadyHasTx := ctx.Value(ctxTxKey{}).(*gorm.DB)
	if !alreadyHasTx {
		tx = db.db.New()
		ctx = context.WithValue(ctx, ctxTxKey{}, tx)
	}

	err := errors.WithStack(fn(ctx, tx))
	if alreadyHasTx {
		return err
	}

	return err
}

func InitDb(configs DbConfigs) error {
	gormDb, err := gorm.Open(
		configs.Driver,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true", // username:password@protocol(host)/dbname?param=value
			configs.User, configs.Password, configs.Host, configs.Port, configs.Schema,
		),
	)
	if err != nil {
		return err
	}

	if configs.IsDebugMode {
		gormDb.LogMode(true)
	}

	db = &DbConnect{
		db: gormDb,
	}

	return nil
}

func GetDbConnect() *DbConnect {
	return db
}
