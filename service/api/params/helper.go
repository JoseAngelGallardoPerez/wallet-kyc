package params

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/gin-gonic/gin"
)

type helper struct {
}

func (h *helper) SetAuthUser(ctx *gin.Context, user model.User) {
	ctx.Set("auth_user", user)
}

func (h *helper) GetAuthUser(ctx *gin.Context) model.User {
	user, ok := ctx.Get("auth_user")

	if !ok {
		connection.Logger.Error("No authorized user")
	}

	return user.(model.User)
}

var Helper = helper{}
