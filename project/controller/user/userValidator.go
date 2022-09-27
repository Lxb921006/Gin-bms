package user

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ValidateData struct {
	validate *validator.Validate
}

func (v *ValidateData) ValidateStruct(s interface{}) (err error) {
	if err = v.validate.Struct(s); err != nil {
		return
	}
	return
}

func (v *ValidateData) ValidatorNumber(fl validator.FieldLevel) bool {
	num := fl.Field().Interface()
	switch num := num.(type) {
	case int:
		if num <= 0 {
			return false
		} else {
			return true
		}
	case uint:
		if num <= 0 {
			return false
		} else {
			return true
		}
	default:
		return false
	}
}

func NewValidateData(v *validator.Validate) *ValidateData {
	return &ValidateData{
		validate: v,
	}
}

func RegisterValidator() {
	var vd ValidateData
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ValidatorNumber", vd.ValidatorNumber)
	}
}
