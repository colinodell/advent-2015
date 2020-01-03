package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLookAndSay(t *testing.T) {
	assert.Equal(t, "11", LookAndSay("1"))
	assert.Equal(t, "21", LookAndSay("11"))
	assert.Equal(t, "1211", LookAndSay("21"))
	assert.Equal(t, "111221", LookAndSay("1211"))
	assert.Equal(t, "312211", LookAndSay("111221"))
}

func TestLookAndSayMultipleTimes(t *testing.T) {
	assert.Equal(t, "312211", LookAndSayMultipleTimes("1", 5))
}
