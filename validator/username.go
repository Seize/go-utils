package validator

import (
	"regexp"

	v "gopkg.in/go-playground/validator.v9"
)

// ValidateUsername is used to help to validate username in a struct
func ValidateUsername(fl v.FieldLevel) bool {
	m, _ := regexp.MatchString("^([A-Za-z0-9]+(-[A-Za-z0-9]+)*){3,38}$", fl.Field().String())
	return m
}
