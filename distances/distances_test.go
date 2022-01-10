package distances

import (
	"testing"

	"github.com/alex-petrov-vt/mazes/algorithms"
	"github.com/alex-petrov-vt/mazes/grid"
	"github.com/stretchr/testify/assert"
)

func TestNewDistances(t *testing.T) {
	c := grid.NewCell(0, 0)
	d := newDistances(c)

	assert.Equal(t, c, d.Root)
	assert.Equal(t, 1, len(d.Cells))
	assert.Equal(t, 0, d.Cells[c])
}

func TestFindingDistances(t *testing.T) {
	g := grid.NewGrid(2, 2)
	g.Cells[0][0].Link(g.Cells[0][1])
	g.Cells[0][0].Link(g.Cells[1][0])
	g.Cells[0][1].Link(g.Cells[1][1])
	g.Cells[1][0].Link(g.Cells[1][1])

	// +---+---+
	// | 0   1 |
	// +---+---+
	// | 1   2 |
	// +---+---+

	d := FindDistances(g.Cells[0][0])
	assert.Equal(t, 1, d.Cells[g.Cells[0][1]])
	assert.Equal(t, 1, d.Cells[g.Cells[1][0]])
	assert.Equal(t, 2, d.Cells[g.Cells[1][1]])

	g = grid.NewGrid(3, 3)
	g.Cells[0][0].Link(g.Cells[0][1])
	g.Cells[0][0].Link(g.Cells[1][0])

	g.Cells[0][1].Link(g.Cells[0][2])
	g.Cells[0][2].Link(g.Cells[1][2])

	g.Cells[1][1].Link(g.Cells[1][2])
	g.Cells[1][1].Link(g.Cells[2][1])

	g.Cells[1][2].Link(g.Cells[2][2])

	g.Cells[2][0].Link(g.Cells[2][1])
	// +---+---+---+
	// | 0   1   2 |
	// +   +---+   +
	// | 1 | 4   3 |
	// +---+       +
	// | 6   5 | 4 |
	// +---+---+---+

	d = FindDistances(g.Cells[0][0])
	assert.Equal(t, 0, d.Cells[g.Cells[0][0]])
	assert.Equal(t, 1, d.Cells[g.Cells[0][1]])
	assert.Equal(t, 2, d.Cells[g.Cells[0][2]])

	assert.Equal(t, 1, d.Cells[g.Cells[1][0]])
	assert.Equal(t, 4, d.Cells[g.Cells[1][1]])
	assert.Equal(t, 3, d.Cells[g.Cells[1][2]])

	assert.Equal(t, 6, d.Cells[g.Cells[2][0]])
	assert.Equal(t, 5, d.Cells[g.Cells[2][1]])
	assert.Equal(t, 4, d.Cells[g.Cells[2][2]])
}

func TestFindingShortestPath(t *testing.T) {
	var randomSeed int64 = 1

	grid := grid.NewGrid(4, 4)
	grid = algorithms.CreateBinaryTreeMaze(grid, randomSeed)

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

	distanceGrid := NewDistanceGrid(grid)

	distances := FindDistances(grid.Cells[0][0])
	shortestPath := distances.FindShortestPath(grid.Cells[3][0])
	distanceGrid.Distances = shortestPath

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

func TestFindingLongestPath(t *testing.T) {
	var randomSeed int64 = 1

	grid := grid.NewGrid(4, 4)
	grid = algorithms.CreateBinaryTreeMaze(grid, randomSeed)

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

	// we want to find a longest path in the maze
	distanceGrid := NewDistanceGrid(grid)

	distances := FindDistances(grid.Cells[0][0])
	distanceGrid.Distances = distances
	assert.Equal(t, grid.Cells[1][0], distances.getMostDistantPoint())

	distances = distances.FindLongestPath()
	distanceGrid.Distances = distances

	expected := `
+---+---+---+---+
| b   a   9   8 |
+---+---+---+   +
| 0   1 |     7 |
+---+   +---+   +
|   | 2 |   | 6 |
+   +   +   +   +
|     3   4   5 |
+---+---+---+---+
`

	assert.Equal(t, expected, distanceGrid.String())

}
