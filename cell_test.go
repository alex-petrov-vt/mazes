package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCell(t *testing.T) {
	s := newCell(1, 2)
	assert.Equal(t, s.row, 1)
	assert.Equal(t, s.col, 2)

	assert.Nil(t, s.north)
	assert.Nil(t, s.east)
	assert.Nil(t, s.south)
	assert.Nil(t, s.west)
	assert.Empty(t, s.links)
}

func TestLinkingCells(t *testing.T) {
	c1 := newCell(1, 1)
	c2 := newCell(1, 2)

	c1.link(c2)
	// Make sure linking is bidirectional
	assert.True(t, c1.links[c2])
	assert.True(t, c2.links[c1])

	c1.unlink(c2)
	_, check := c1.links[c2]
	assert.False(t, check)
	_, check = c2.links[c1]
	assert.False(t, check)
}
