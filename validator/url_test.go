package validator_test

import (
	"testing"

	"github.com/seize/go-utils/validator"
	"github.com/stretchr/testify/assert"
	v "gopkg.in/go-playground/validator.v9"
)

var urlItems = []struct {
	have string
	want bool
}{
	{"http://google.com/", true},
	{"http//google.com", false},
	{"google.com", false},
}

// TestValidateURL ...
func TestValidateURL(t *testing.T) {
	validate := v.New()
	validate.RegisterValidation("url", validator.ValidateURL)

	for _, item := range urlItems {
		err := validate.Var(item.have, "url")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err, "Bad server", item.have)
		}
	}
}
