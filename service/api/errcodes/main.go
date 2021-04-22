package errcodes

import (
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/internal_errors"
	"github.com/Confialink/wallet-pkg-errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddError(c *gin.Context, err error) {
	if internalErrors, ok := err.(internal_errors.Errors); ok {
		for _, internalError := range internalErrors {
			AddError(c, internalError)
		}
	} else {
		if internalError, ok := err.(*internal_errors.Error); ok {
			code := internalError.Code
			errors.AddErrors(c, &errors.PublicError{
				Code:       code,
				HttpStatus: statusCodes[code],
				Title:      internalError.PublicMessages,
			})

			if internalError.Parent != nil {
				connection.Logger.Error(internalError.Parent.Error())
			}

		} else if validationErrors, ok := err.(validator.ValidationErrors); ok {
			AddFormError(c, validationErrors)
		} else {
			AddErrorCode(c, InternalServerError)
			connection.Logger.Error(err.Error())
		}
	}
}

func AddErrorLog(c *gin.Context, code string, err error) {
	errors.AddErrors(c, CreatePublicError(code))
	connection.Logger.Error(err.Error())
}

func AddErrorCode(c *gin.Context, code string) {
	errors.AddErrors(c, CreatePublicError(code))

}

func AddFormError(ctx *gin.Context, err error) {
	errors.AddShouldBindError(ctx, err)
}

func CreatePublicError(code string) *errors.PublicError {
	return &errors.PublicError{
		Code:       code,
		HttpStatus: statusCodes[code],
	}
}

func AddErrorMeta(c *gin.Context, code string, meta interface{}) {
	publicErr := &errors.PublicError{
		Code:       code,
		HttpStatus: statusCodes[code],
		Meta:       meta,
	}
	errors.AddErrors(c, publicErr)
}
