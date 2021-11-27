package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCell(t *testing.T) {
	s := NewCell(1, 2)
	assert.Equal(t, 1, s.Row)
	assert.Equal(t, 2, s.Col)
	assert.Nil(t, s.North)
	assert.Nil(t, s.East)
	assert.Nil(t, s.South)
	assert.Nil(t, s.West)
	assert.Empty(t, s.Links)
}

func TestLinkingCells(t *testing.T) {
	c1 := NewCell(1, 1)
	c2 := NewCell(1, 2)

	c1.Link(c2)
	// Make sure linking is bidirectional
	assert.True(t, c1.IsLinked(c2))
	assert.True(t, c2.IsLinked(c1))

	c1.Unlink(c2)
	assert.False(t, c1.IsLinked(c2))
	assert.False(t, c2.IsLinked(c1))
}
