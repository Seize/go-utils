package validator_test

import (
	"testing"

	"github.com/seize/go-utils/validator"
	"github.com/stretchr/testify/assert"
	v "gopkg.in/go-playground/validator.v9"
)

var idItems = []struct {
	have string
	want bool
}{
	{"507f1f77bcf86cd799439011", true},
	{"507g1f77bcf86cd799439011", false},
	{"1234567890QWERTYUIOP", false},
	{"@zerty", false},
	{"azerty-", false},
	{"_azerty", false},
}

// TestValidateID ...
func TestValidateID(t *testing.T) {
	validate := v.New()
	validate.RegisterValidation("id", validator.ValidateID)

	for _, item := range idItems {
		err := validate.Var(item.have, "id")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
