package validutil_test

import (
	"testing"

	"github.com/seize/go-utils/validator"
	"github.com/stretchr/testify/assert"
	validator "gopkg.in/go-playground/validator.v9"
)

var usernameItems = []struct {
	have string
	want bool
}{
	{"jbx", true},
	{"julien-breux", true},
	{"C3PO", true},
	{"R2-D2", true},
	{"julien_breux", false},
	{"jb", false},
	{"BB-8", false},
	{"julien-", false},
	{"-breux-", false},
	{"-breux", false},
}

// TestValidateUsername ...
func TestValidateUsername(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", validutil.ValidateUsername)

	for _, item := range usernameItems {
		err := validate.Var(item.have, "username")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
