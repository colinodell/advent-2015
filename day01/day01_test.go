package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTravel(t *testing.T) {
	assert.Equal(t, 0, Travel("(())"))
	assert.Equal(t, 0, Travel("()()"))
	assert.Equal(t, 3, Travel("((("))
	assert.Equal(t, 3, Travel("(()(()("))
	assert.Equal(t, 3, Travel("))((((("))
	assert.Equal(t, -1, Travel("())"))
	assert.Equal(t, -1, Travel("))("))
	assert.Equal(t, -3, Travel(")))"))
	assert.Equal(t, -3, Travel(")())())"))
}

func TestTravelUntilBasementReached(t *testing.T) {
	assert.Equal(t, 1, TravelUntilBasementReached(")"))
	assert.Equal(t, 5, TravelUntilBasementReached("()())"))
}
