package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestListDirectories(t *testing.T) {
	_gopath := ListDirectories(".gopath")
	assert.Equal(t, len(_gopath), 2)
}
