package url_test

import (
	"testing"

	"github.com/seize/go-utils/url"
	"github.com/stretchr/testify/assert"
)

// TestValuesToURLAll...
func TestValuesToURLAll(t *testing.T) {
	actual := url.Build("mongodb", "jb", "seize", "0.0.0.0", 27017)
	expected := "mongodb://jb:seize@0.0.0.0:27017"

	assert.Equal(t, expected, actual)
}

// TestValuesToURLNoPort...
func TestValuesToURLNoPort(t *testing.T) {
	actual := url.Build("mongodb", "jb", "seize", "0.0.0.0", 0)
	expected := "mongodb://jb:seize@0.0.0.0"

	assert.Equal(t, expected, actual)
}

// TestValuesToURLNoPassword...
func TestValuesToURLNoPassword(t *testing.T) {
	actual := url.Build("mongodb", "jb", "", "0.0.0.0", 27017)
	expected := "mongodb://jb@0.0.0.0:27017"

	assert.Equal(t, expected, actual)
}

// TestValuesToURLNoUsername...
func TestValuesToURLNoUsername(t *testing.T) {
	actual := url.Build("mongodb", "", "seize", "0.0.0.0", 27017)
	expected := "mongodb://0.0.0.0:27017"

	assert.Equal(t, expected, actual)
}
