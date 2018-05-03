package validator_test

import (
	"testing"

	"github.com/seize/go-utils/validator"
	"github.com/stretchr/testify/assert"
	v "gopkg.in/go-playground/validator.v9"
)

var keyItems = []struct {
	have string
	want bool
}{
	{"QWERTYQWERTYQWERTYQWERTY", true},
	{"123456QWERTYQWERTYqwerty", true},
	{"1234 QWERTYQWERTYqwerty", false},
	{"12345QWERTY", false},
	{"qwerty123456", false},
	{"qawsed$44D$", false},
}

// TestValidateKey ...
func TestValidateKey(t *testing.T) {
	validate := v.New()
	validate.RegisterValidation("key", validator.ValidateKey)

	for _, item := range keyItems {
		err := validate.Var(item.have, "key")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
