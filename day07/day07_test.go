package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRunCircuit(t *testing.T) {
	instructions := `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
`
	c := NewCircuit(strings.Split(strings.TrimSpace(instructions), "\n"))

	assert.Equal(t, uint16(72), c.GetValue("d"))
	assert.Equal(t, uint16(507), c.GetValue("e"))
	assert.Equal(t, uint16(492), c.GetValue("f"))
	assert.Equal(t, uint16(114), c.GetValue("g"))
	assert.Equal(t, uint16(65412), c.GetValue("h"))
	assert.Equal(t, uint16(65079), c.GetValue("i"))
	assert.Equal(t, uint16(123), c.GetValue("x"))
	assert.Equal(t, uint16(456), c.GetValue("y"))
}
