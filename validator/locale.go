package validator

import (
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

// ValidateLocale helps to validate locale field in a struct
func ValidateLocale(fl validator.FieldLevel) bool {
	m, _ := regexp.MatchString("^([a-z]{2})-([A-Z]{2})$", fl.Field().String())

	return m
}
