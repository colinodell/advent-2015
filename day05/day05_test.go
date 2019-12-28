package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNiceV1(t *testing.T) {
	assert.True(t, IsNiceV1("ugknbfddgicrmopn"))
	assert.True(t, IsNiceV1("aaa"))

	assert.False(t, IsNiceV1("jchzalrnumimnmhp"))
	assert.False(t, IsNiceV1("haegwjzuvuyypxyu"))
	assert.False(t, IsNiceV1("dvszwmarrgswjxmb"))
}

func TestIsNiceV2(t *testing.T) {
	assert.True(t, IsNiceV2("qjhvhtzxzqqjkmpb"))
	assert.True(t, IsNiceV2("xxyxx"))

	assert.False(t, IsNiceV2("uurcxstgmygtbstg"))
	assert.False(t, IsNiceV2("ieodomkazucvgmuy"))
}
