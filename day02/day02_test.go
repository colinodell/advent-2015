package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPresent_CalculateWrappingPaper(t *testing.T) {
	p1 := Present{2, 3, 4}
	assert.Equal(t, 58, p1.CalculateWrappingPaper())

	p2 := Present{1, 1, 10}
	assert.Equal(t, 43, p2.CalculateWrappingPaper())
}

func TestPresent_CalculateRibbon(t *testing.T) {
	p1 := Present{2, 3, 4}
	assert.Equal(t, 34, p1.CalculateRibbon())

	p2 := Present{1, 1, 10}
	assert.Equal(t, 14, p2.CalculateRibbon())
}
