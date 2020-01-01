package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindShortestAndLongestRoute(t *testing.T) {
	input := `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`

	shortest, longest := FindShortestAndLongestRoute(input)

	assert.Equal(t, 605, shortest)
	assert.Equal(t, 982, longest)
}
