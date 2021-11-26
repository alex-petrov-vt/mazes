package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingSidewinderMaze(t *testing.T) {
	var randomSeed int64 = 1

	grid := newGrid(2, 2)
	grid = createSidewinderMaze(grid, randomSeed)

	expected := `
+---+---+
|       |
+---+   +
|       |
+---+---+
`
	assert.Equal(t, expected, grid.String())

	grid = newGrid(4, 4)
	grid = createSidewinderMaze(grid, randomSeed)

	expected = `
+---+---+---+---+
|               |
+---+---+---+   +
|       |   |   |
+---+   +   +   +
|   |   |   |   |
+   +   +   +   +
|               |
+---+---+---+---+
`
	assert.Equal(t, expected, grid.String())
}
