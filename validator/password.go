package validutil

import (
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

// ValidatePassword is used to help to validate password in a struct
func ValidatePassword(fl validator.FieldLevel) bool {
	m0, _ := regexp.MatchString("[A-Z]+", fl.Field().String())
	m1, _ := regexp.MatchString("[0-9]+", fl.Field().String())
	return m0 && m1
}
