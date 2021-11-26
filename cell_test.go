package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCell(t *testing.T) {
	s := newCell(1, 2)
	assert.Equal(t, 1, s.row)
	assert.Equal(t, 2, s.col)

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
	assert.True(t, c1.isLinked(c2))
	assert.True(t, c2.isLinked(c1))

	c1.unlink(c2)
	assert.False(t, c1.isLinked(c2))
	assert.False(t, c2.isLinked(c1))
}
