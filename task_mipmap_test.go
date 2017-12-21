package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// configがパースできる
func TestParseConfigure(t *testing.T) {
	cfg := Configure{}

	assert.NoError(t, ParseYamlFromFile("examples/mipmap/config.yaml", &cfg))

	assert.NotEqual(t, cfg.Mipmap.Requests[0].Path, "")
	assert.NotEqual(t, cfg.Mipmap.Requests[0].Platform, "")
	assert.NotEqual(t, cfg.Mipmap.Requests[0].Type, "")
	assert.NotEqual(t, cfg.Mipmap.Requests[0].OutputPath, "")
}
