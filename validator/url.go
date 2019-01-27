package validator

import (
	"net/url"

	validator "gopkg.in/go-playground/validator.v9"
)

// ValidateURL helps to validate URL in a struct
func ValidateURL(fl validator.FieldLevel) bool {
	_, err := url.ParseRequestURI(fl.Field().String())
	return err == nil
}
