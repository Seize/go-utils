package validator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	v "gopkg.in/go-playground/validator.v9"

	"github.com/seize/go-utils/validator"
)

var localeItems = []struct {
	have string
	want bool
}{
	{"fr-FR", true},
	{"en-us", false},
	{"12-45", false},
	{"frr-FRR", false},
}

// TestValidateLocale ...
func TestValidateLocale(t *testing.T) {
	validate := v.New()
	validate.RegisterValidation("locale", validator.ValidateLocale)

	for _, item := range localeItems {
		err := validate.Var(item.have, "locale")
		if item.want {
			assert.Nil(t, err, item.have)
		} else {
			assert.Error(t, err)
		}
	}
}
