package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDistanceGrid(t *testing.T) {
	dg := newDistanceGrid(2, 2)
	assert.Equal(t, 2, dg.grid.rows)
	assert.Equal(t, 2, dg.grid.cols)
	assert.Equal(t, 2, len(dg.grid.cells))
	assert.Equal(t, 2, len(dg.grid.cells[0]))

	assert.Nil(t, dg.distances)
}

func TestDistanceGridDrawingASCII(t *testing.T) {
	dg := newDistanceGrid(3, 3)

	dg.grid.cells[0][0].link(dg.grid.cells[0][1])
	dg.grid.cells[0][0].link(dg.grid.cells[1][0])

	dg.grid.cells[0][1].link(dg.grid.cells[0][2])

	dg.grid.cells[0][2].link(dg.grid.cells[1][2])

	dg.grid.cells[1][1].link(dg.grid.cells[1][2])
	dg.grid.cells[1][1].link(dg.grid.cells[2][1])

	dg.grid.cells[1][2].link(dg.grid.cells[2][2])

	dg.grid.cells[2][0].link(dg.grid.cells[2][1])

	dg.distances = findDistances(dg.grid.cells[0][0])

	expected := `
+---+---+---+
| 0   1   2 |
+   +---+   +
| 1 | 4   3 |
+---+   +   +
| 6   5 | 4 |
+---+---+---+
`
	assert.Equal(t, expected, dg.String())
}
