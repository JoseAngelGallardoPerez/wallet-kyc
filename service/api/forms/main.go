package forms

import (
	"github.com/Confialink/wallet-kyc/service/api/errcodes"
	"github.com/gin-gonic/gin"
)

func Bind(ctx *gin.Context, obj interface{}) bool {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		errcodes.AddFormError(ctx, err)
		return false
	}
	return true
}
