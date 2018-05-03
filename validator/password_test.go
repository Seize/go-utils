package validator_test

import (
	"testing"

	"github.com/seize/go-utils/validator"
	"github.com/stretchr/testify/assert"
	v "gopkg.in/go-playground/validator.v9"
)

var passwordItems = []struct {
	have string
	want bool
}{
	{"secret", false},
	{"Secrets1", true},
}

// TestValidatePassword ...
func TestValidatePassword(t *testing.T) {
	validate := v.New()
	validate.RegisterValidation("password", validator.ValidatePassword)

	for _, item := range passwordItems {
		err := validate.Var(item.have, "password")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err, "Bad password", item.have)
		}
	}
}
