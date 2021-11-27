package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGrid(t *testing.T) {

	// Predicatbly seed rand to test getting random cell
	var randomSeed int64 = 1

	g := NewGrid(1, 1)
	assert.Equal(t, 1, g.Rows)
	assert.Equal(t, 1, g.Cols)
	assert.Equal(t, 1, len(g.Cells))
	assert.Equal(t, 1, len(g.Cells[0]))
	assert.Nil(t, g.Cells[0][0].North)
	assert.Nil(t, g.Cells[0][0].East)
	assert.Nil(t, g.Cells[0][0].South)
	assert.Nil(t, g.Cells[0][0].West)
	assert.Equal(t, 1, g.getSize())
	assert.Equal(t, g.Cells[0][0], g.getRandomCell(randomSeed))

	g = NewGrid(100, 100)
	assert.Equal(t, 100, g.Rows)
	assert.Equal(t, 100, g.Cols)
	assert.Equal(t, 100, len(g.Cells))
	assert.Equal(t, 100, len(g.Cells[0]))

	// Make sure the connection is bidirectional
	assert.Equal(t, g.Cells[50][50].North, g.Cells[49][50])
	assert.Equal(t, g.Cells[50][50].East, g.Cells[50][51])
	assert.Equal(t, g.Cells[50][50].South, g.Cells[51][50])
	assert.Equal(t, g.Cells[50][50].West, g.Cells[50][49])

	assert.Equal(t, g.Cells[49][50].South, g.Cells[50][50])
	assert.Equal(t, g.Cells[50][51].West, g.Cells[50][50])
	assert.Equal(t, g.Cells[51][50].North, g.Cells[50][50])
	assert.Equal(t, g.Cells[50][49].East, g.Cells[50][50])
	assert.Equal(t, 10000, g.getSize())
	assert.Equal(t, g.Cells[81][87], g.getRandomCell(randomSeed))
}

func TestGridDrawingASCII(t *testing.T) {
	grid := NewGrid(3, 3)

	grid.Cells[0][0].Link(grid.Cells[0][1])
	grid.Cells[0][0].Link(grid.Cells[1][0])

	grid.Cells[0][1].Link(grid.Cells[0][2])

	grid.Cells[0][2].Link(grid.Cells[1][2])

	grid.Cells[1][1].Link(grid.Cells[1][2])
	grid.Cells[1][1].Link(grid.Cells[2][1])

	grid.Cells[1][2].Link(grid.Cells[2][2])

	grid.Cells[2][0].Link(grid.Cells[2][1])

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
