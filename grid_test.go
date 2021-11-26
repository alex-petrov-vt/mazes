package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGrid(t *testing.T) {

	// Predicatbly seed rand to test getting random cell
	var randomSeed int64 = 1

	g := newGrid(1, 1)
	assert.Equal(t, g.rows, 1)
	assert.Equal(t, g.cols, 1)
	assert.Equal(t, len(g.cells), 1)
	assert.Equal(t, len(g.cells[0]), 1)
	assert.Nil(t, g.cells[0][0].north)
	assert.Nil(t, g.cells[0][0].east)
	assert.Nil(t, g.cells[0][0].south)
	assert.Nil(t, g.cells[0][0].west)
	assert.Equal(t, g.getSize(), 1)
	assert.Equal(t, g.getRandomCell(randomSeed), g.cells[0][0])

	g = newGrid(100, 100)
	assert.Equal(t, g.rows, 100)
	assert.Equal(t, g.cols, 100)
	assert.Equal(t, len(g.cells), 100)
	assert.Equal(t, len(g.cells[0]), 100)

	// Make sure the connection is bidirectional
	assert.Equal(t, g.cells[50][50].north, g.cells[49][50])
	assert.Equal(t, g.cells[50][50].east, g.cells[50][51])
	assert.Equal(t, g.cells[50][50].south, g.cells[51][50])
	assert.Equal(t, g.cells[50][50].west, g.cells[50][49])

	assert.Equal(t, g.cells[49][50].south, g.cells[50][50])
	assert.Equal(t, g.cells[50][51].west, g.cells[50][50])
	assert.Equal(t, g.cells[51][50].north, g.cells[50][50])
	assert.Equal(t, g.cells[50][49].east, g.cells[50][50])
	assert.Equal(t, g.getSize(), 10000)
	assert.Equal(t, g.getRandomCell(randomSeed), g.cells[81][87])
}

func TestGridDrawingASCII(t *testing.T) {
	grid := newGrid(3, 3)

	grid.cells[0][0].link(grid.cells[0][1])
	grid.cells[0][0].link(grid.cells[1][0])

	grid.cells[0][1].link(grid.cells[0][2])

	grid.cells[0][2].link(grid.cells[1][2])

	grid.cells[1][1].link(grid.cells[1][2])
	grid.cells[1][1].link(grid.cells[2][1])

	grid.cells[1][2].link(grid.cells[2][2])

	grid.cells[2][0].link(grid.cells[2][1])

	expected := `
+---+---+---+
|           |
+   +---+   +
|   |       |
+---+   +   +
|       |   |
+---+---+---+
`
	assert.Equal(t, expected, grid.String())
}
