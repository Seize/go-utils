package validator_test

import (
	"testing"

	"github.com/seize/go-utils/validator"
	"github.com/stretchr/testify/assert"
	v "gopkg.in/go-playground/validator.v9"
)

var slugItems = []struct {
	have string
	want bool
}{
	{"good", true},
	{"good-good-good", true},
	{"124", true},
	{"ba", false},
	{"Bad", false},
	{"-bad", false},
	{"bad-", false},
	{"bad-", false},
}

// TestValidateSlug ...
func TestValidateSlug(t *testing.T) {
	validate := v.New()
	validate.RegisterValidation("slug", validator.ValidateSlug)

	for _, item := range slugItems {
		err := validate.Var(item.have, "slug")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
