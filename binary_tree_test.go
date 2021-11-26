package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingBinaryTreeMaze(t *testing.T) {
	var randomSeed int64 = 1

	grid := newGrid(2, 2)
	grid = createBinaryTreeMaze(grid, randomSeed)

	expected := `
+---+---+
|       |
+---+   +
|       |
+---+---+
`
	assert.Equal(t, expected, grid.String())

	grid = newGrid(4, 4)
	grid = createBinaryTreeMaze(grid, randomSeed)

	expected = `
+---+---+---+---+
|               |
+---+---+---+   +
|       |       |
+---+   +---+   +
|   |   |   |   |
+   +   +   +   +
|               |
+---+---+---+---+
`
	assert.Equal(t, expected, grid.String())
}
