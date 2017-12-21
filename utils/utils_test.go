package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestListDirectories(t *testing.T) {
	_gopath := ListDirectories("../.gopath")
	assert.Equal(t, len(_gopath), 2)

	utils := ListDirectories(".")
	assert.Equal(t, len(utils), 0)
}
