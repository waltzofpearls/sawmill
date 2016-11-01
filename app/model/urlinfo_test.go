package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUrlInfo(t *testing.T) {
	mdl := NewUrlInfo("url", "desc", false)

	assert.Equal(t, "url", mdl.Url)
	assert.Equal(t, "desc", mdl.Description)
	assert.False(t, mdl.HasMalware)

	assert.Equal(t, "url", mdl.GetId())
}
