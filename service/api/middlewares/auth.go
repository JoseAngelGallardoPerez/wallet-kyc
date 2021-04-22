package middlewares

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/model"
	errorsPkg "github.com/Confialink/wallet-pkg-errors"
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Auth middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokens := c.Request.Header.Get("Authorization")
		if len(tokens) < 8 || !strings.EqualFold(tokens[0:7], "Bearer ") {
			c.Header("Authentication", `Bearer realm="private"`)
			_ = c.Error(&errorsPkg.PublicError{
				Title:      "Access token not found",
				Code:       "ACCESS_TOKEN_NOT_FOUND",
				HttpStatus: http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		accessToken := tokens[7:]

		client := connection.GetRpcUsers()

		res, err := client.Client.ValidateAccessToken(context.Background(), &userpb.Request{AccessToken: accessToken})
		if nil != err {
			c.Header("Authentication", `Bearer realm="private"`)
			_ = c.Error(&errorsPkg.PublicError{
				Title:      "Access token is invalid",
				Code:       "ACCESS_TOKEN_INVALID",
				HttpStatus: http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		authUser := model.User{
			ID:    res.User.UID,
			Phone: res.User.PhoneNumber,
			Role:  res.User.RoleName,
		}

		authUser.RecognizeCountry()

		c.Set("auth_user", authUser)
	}
}
