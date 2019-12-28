package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountHouses(t *testing.T) {
	assert.Equal(t, 2, CountHouses(">"))
	assert.Equal(t, 4, CountHouses("^>v<"))
	assert.Equal(t, 2, CountHouses("^v^v^v^v^v"))
}

func TestCountHousesWithRoboSanta(t *testing.T) {
	assert.Equal(t, 3, CountHousesWithRoboSanta("^v"))
	assert.Equal(t, 3, CountHousesWithRoboSanta("^>v<"))
	assert.Equal(t, 11, CountHousesWithRoboSanta("^v^v^v^v^v"))
}
