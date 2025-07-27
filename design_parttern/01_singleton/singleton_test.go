package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert.NotNil(t, GetInstance())
	assert.Equal(t, GetInstance(), singleton)
	assert.True(t, GetInstance() == GetInstance())
	assert.False(t, GetInstance() == GetLazyInstance())
}
