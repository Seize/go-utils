package validutil

import (
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

// ValidateKey is used to help to validate key in a struct
func ValidateKey(fl validator.FieldLevel) bool {
	m, _ := regexp.MatchString("^[a-zA-Z0-9]{24}$", fl.Field().String())

	return m
}
