package validator

import (
	"regexp"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

// ValidateGA helps to validate Google Analytics tracking code in a struct
func ValidateGA(fl validator.FieldLevel) bool {
	code := strings.ToLower(fl.Field().String())
	m, _ := regexp.MatchString("^ua-\\d{4,9}-\\d{1,4}$", code)

	return m
}
