package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMineAdventCoin5(t *testing.T) {
	assert.Equal(t, 609043, MineAdventCoin5("abcdef"))
	assert.Equal(t, 1048970, MineAdventCoin5("pqrstuv"))

	assert.Equal(t, 254575, MineAdventCoin5("bgvyzdsv"))
}

func TestMineAdventCoin6(t *testing.T) {
	assert.Equal(t, 1038736, MineAdventCoin6("bgvyzdsv"))
}
