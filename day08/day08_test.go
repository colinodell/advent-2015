package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	assert.Equal(t, 2, Decode(`""`))
	assert.Equal(t, 2, Decode(`"abc"`))
	assert.Equal(t, 3, Decode(`"aaa\"aaa"`))
	assert.Equal(t, 5, Decode(`"\x27"`))
}

func TestEncode(t *testing.T) {
	assert.Equal(t, 4, Encode(`""`))
	assert.Equal(t, 4, Encode(`"abc"`))
	assert.Equal(t, 6, Encode(`"aaa\"aaa"`))
	assert.Equal(t, 5, Encode(`"\x27"`))
}
