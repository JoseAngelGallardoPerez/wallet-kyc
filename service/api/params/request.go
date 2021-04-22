package params

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-pkg-list_params"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateParams(ctx *gin.Context, object interface{}) *list_params.ListParams {
	listParams := list_params.NewListParamsFromQuery(ctx.Request.URL.RawQuery, object)
	listParams.AllowPagination()
	return listParams
}

func CreateQuery(ctx *gin.Context, params *list_params.ListParams) *gorm.DB {
	str, arguments := params.GetWhereCondition()

	var query *gorm.DB

	_ = connection.GetDbConnect().ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		query = db.Where(str, arguments...)

		query = query.Order(params.GetOrderByString())
		if params.GetLimit() != 0 {
			query = query.Limit(params.GetLimit())
		}

		query = query.Offset(params.GetOffset())

		query = query.Joins(params.GetJoinCondition())

		for _, preloadName := range params.GetPreloads() {
			query = query.Preload(preloadName)
		}
		return nil
	})
	return query
}

func CreateQueryCount(ctx *gin.Context, params *list_params.ListParams) *gorm.DB {
	str, arguments := params.GetWhereCondition()

	var query *gorm.DB

	_ = connection.GetDbConnect().ExtractTx(ctx, func(ctx context.Context, db *gorm.DB) error {
		query = db.Where(str, arguments...)
		return nil
	})
	return query
}
