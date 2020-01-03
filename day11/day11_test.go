package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNext(t *testing.T) {
	assert.Equal(t, "aab", nextPassword("aaa"))
	assert.Equal(t, "aba", nextPassword("aaz"))
	assert.Equal(t, "baa", nextPassword("azz"))
	assert.Equal(t, "aaa", nextPassword("zzz"))
}

func TestRequirements(t *testing.T) {
	assert.True(t, hasStraight("hijklmmn"))
	assert.False(t, hasOnlyValidCharacters("hijklmmn"))

	assert.False(t, hasStraight("abbceffg"))
	assert.True(t, hasTwoNonOverlappingPairs("abbceffg"))

	assert.False(t, hasTwoNonOverlappingPairs("abbcegjk"))
}

func TestGetNextValidPassword(t *testing.T) {
	assert.Equal(t, "abcdffaa", GetNextValidPassword("abcdefgh"))
	assert.Equal(t, "ghjaabcc", GetNextValidPassword("ghijklmn"))
}
