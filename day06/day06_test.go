package main

import (
	"advent-2015/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrid(t *testing.T) {
	g := NewGrid()
	g.TurnOn(utils.Vector2{0, 0}, utils.Vector2{999, 999})

	assert.Equal(t, 1000000, g.CountLit())

	g.Toggle(utils.Vector2{0, 0}, utils.Vector2{999, 0})

	assert.Equal(t, 999000, g.CountLit())

	g.TurnOff(utils.Vector2{499, 499}, utils.Vector2{500, 500})

	assert.Equal(t, 998996, g.CountLit())
}
