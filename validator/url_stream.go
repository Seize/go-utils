package validator

import (
	"regexp"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

// ValidateURLStream helps to validate stream URL in a struct
func ValidateURLStream(fl validator.FieldLevel) bool {
	code := strings.ToLower(fl.Field().String())
	m, _ := regexp.MatchString(`^(((rtmp:(?:\/\/)?)(?:[\-;:&=\+\$,\w]+@)?[A-Za-z0-9\.\-]+|(?:www\.|[\-;:&=\+\$,\w]+@)[A-Za-z0-9\.\-]+)((?:\/[\+~%\/\.\w\-_]*)?\??(?:[\-\+=&;%@\.\w_]*)#?(?:[\.\!\/\\\w]*))?)`, code)

	return m
}
