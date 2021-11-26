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
}
