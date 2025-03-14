package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
}

// ValidateURL валидирует входящую ссылку
func ValidateURL(url string) error {
    return validate.Var(url, "required,url")
}