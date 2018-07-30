package validator_test

import (
	"testing"

	"github.com/seize/go-utils/validator"
	"github.com/stretchr/testify/assert"
	v "gopkg.in/go-playground/validator.v9"
)

var gaItems = []struct {
	have string
	want bool
}{
	{"ua-1234-1", true},
	{"ua-123456789-1", true},
	{"ua-1234-1234", true},
	{"1234-1", false},
	{"ua-123-1", false},
	{"ua-123456789-12345", false},
}

// TestValidateGA ...
func TestValidateGA(t *testing.T) {
	validate := v.New()
	validate.RegisterValidation("ga", validator.ValidateGA)

	for _, item := range gaItems {
		err := validate.Var(item.have, "ga")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
