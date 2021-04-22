package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Interface is a validator interface
type Interface interface {
	Struct(current interface{}) error
	StructExcept(current interface{}, fields ...string) error
	RegisterValidation(key string, fn validator.Func) error
	RegisterStructValidation(fn validator.StructLevelFunc, types ...interface{})
}

var valid *validator.Validate

func InitValidate() {
	valid = binding.Validator.Engine().(*validator.Validate)
	_ = valid.RegisterValidation("requestStatus", requestStatus)
	_ = valid.RegisterValidation("requirementStatus", requirementStatus)
}
