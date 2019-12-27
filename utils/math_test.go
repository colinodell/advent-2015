package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 0, Min(0, 1))
	assert.Equal(t, 4, Min(4, 8, 15, 16, 23, 42))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 1, Max(0, 1))
	assert.Equal(t, 42, Max(4, 8, 15, 16, 23, 42))
}
