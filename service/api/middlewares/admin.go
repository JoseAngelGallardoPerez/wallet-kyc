package middlewares

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/Confialink/wallet-kyc/service/api/errcodes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminOrRoot() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.Get("auth_user")
		if !ok {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		roleName := user.(model.User).Role
		if roleName != "admin" && roleName != "root" {
			ctx.Status(http.StatusForbidden)
			_ = ctx.Error(errcodes.CreatePublicError(errcodes.CodeForbidden))
			ctx.Abort()
			return
		}
	}
}
