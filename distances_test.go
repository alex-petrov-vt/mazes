package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDistances(t *testing.T) {
	c := newCell(0, 0)
	d := newDistances(c)

	assert.Equal(t, c, d.root)
	assert.Equal(t, 1, len(d.cells))
	assert.Equal(t, 0, d.cells[c])
}

func TestFindingDistances(t *testing.T) {
	g := newGrid(2, 2)
	g.cells[0][0].link(g.cells[0][1])
	g.cells[0][0].link(g.cells[1][0])
	g.cells[0][1].link(g.cells[1][1])
	g.cells[1][0].link(g.cells[1][1])

	// +---+---+
	// | 0   1 |
	// +---+---+
	// | 1   2 |
	// +---+---+

	d := findDistances(g.cells[0][0])
	assert.Equal(t, 1, d.cells[g.cells[0][1]])
	assert.Equal(t, 1, d.cells[g.cells[1][0]])
	assert.Equal(t, 2, d.cells[g.cells[1][1]])

	g = newGrid(3, 3)
	g.cells[0][0].link(g.cells[0][1])
	g.cells[0][0].link(g.cells[1][0])

	g.cells[0][1].link(g.cells[0][2])
	g.cells[0][2].link(g.cells[1][2])

	g.cells[1][1].link(g.cells[1][2])
	g.cells[1][1].link(g.cells[2][1])

	g.cells[1][2].link(g.cells[2][2])

	g.cells[2][0].link(g.cells[2][1])
	// +---+---+---+
	// | 0   1   2 |
	// +   +---+   +
	// | 1 | 4   3 |
	// +---+       +
	// | 6   5 | 4 |
	// +---+---+---+

	d = findDistances(g.cells[0][0])
	assert.Equal(t, 0, d.cells[g.cells[0][0]])
	assert.Equal(t, 1, d.cells[g.cells[0][1]])
	assert.Equal(t, 2, d.cells[g.cells[0][2]])

	assert.Equal(t, 1, d.cells[g.cells[1][0]])
	assert.Equal(t, 4, d.cells[g.cells[1][1]])
	assert.Equal(t, 3, d.cells[g.cells[1][2]])

	assert.Equal(t, 6, d.cells[g.cells[2][0]])
	assert.Equal(t, 5, d.cells[g.cells[2][1]])
	assert.Equal(t, 4, d.cells[g.cells[2][2]])
}
