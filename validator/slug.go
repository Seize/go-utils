package validator

import (
	"regexp"

	v "gopkg.in/go-playground/validator.v9"
)

// ValidateSlug is used to help to validate slug in a struct
func ValidateSlug(fl v.FieldLevel) bool {
	m, _ := regexp.MatchString("^[a-z0-9]+[a-z0-9-]+[a-z0-9]+$", fl.Field().String())
	return m
}
