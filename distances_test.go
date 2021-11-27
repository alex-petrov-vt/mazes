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

func TestFindingShortestPath(t *testing.T) {
	var randomSeed int64 = 1

	grid := newGrid(4, 4)
	grid = createBinaryTreeMaze(grid, randomSeed)

	// binary tree algorithm with a seed of 1 and size of 4x4
	// will produce the following maze
	// +---+---+---+---+
	// |               |
	// +---+---+---+   +
	// |       |       |
	// +---+   +---+   +
	// |   |   |   |   |
	// +   +   +   +   +
	// |               |
	// +---+---+---+---+

	// we want to find the shortest path from southwest corner to
	// root in the northwest corner

	distanceGrid := newDistanceGrid(4, 4)
	distanceGrid.grid = grid

	distances := findDistances(grid.cells[0][0])
	shortestPath := distances.findShortestPath(grid.cells[3][0])
	distanceGrid.distances = shortestPath

	expected := `
+---+---+---+---+
| 0   1   2   3 |
+---+---+---+   +
|       |     4 |
+---+   +---+   +
|   |   |   | 5 |
+   +   +   +   +
| 9   8   7   6 |
+---+---+---+---+
`

	assert.Equal(t, expected, distanceGrid.String())

}
