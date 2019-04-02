package validator_test

import (
	"testing"

	"github.com/seize/go-utils/validator"
	"github.com/stretchr/testify/assert"
	v "gopkg.in/go-playground/validator.v9"
)

var serverItems = []struct {
	have string
	want bool
}{
	{"rtmp:///xxxxx", false},
	{"rtmp://a.rtmp.youtube.com/live2", true},
	{"rtmps://a.rtmp.youtube.com/live2", true},
}

// TestValidateURLStream ...
func TestValidateURLStream(t *testing.T) {
	validate := v.New()
	validate.RegisterValidation("url_stream", validator.ValidateURLStream)

	for _, item := range serverItems {
		err := validate.Var(item.have, "url_stream")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err, "Bad server", item.have)
		}
	}
}
