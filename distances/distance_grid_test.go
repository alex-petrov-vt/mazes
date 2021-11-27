package distances

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDistanceGrid(t *testing.T) {
	dg := NewDistanceGrid(2, 2)
	assert.Equal(t, 2, dg.Grid.Rows)
	assert.Equal(t, 2, dg.Grid.Cols)
	assert.Equal(t, 2, len(dg.Grid.Cells))
	assert.Equal(t, 2, len(dg.Grid.Cells[0]))

	assert.Nil(t, dg.Distances)
}

func TestDistanceGridDrawingASCII(t *testing.T) {
	dg := NewDistanceGrid(3, 3)

	dg.Grid.Cells[0][0].Link(dg.Grid.Cells[0][1])
	dg.Grid.Cells[0][0].Link(dg.Grid.Cells[1][0])

	dg.Grid.Cells[0][1].Link(dg.Grid.Cells[0][2])

	dg.Grid.Cells[0][2].Link(dg.Grid.Cells[1][2])

	dg.Grid.Cells[1][1].Link(dg.Grid.Cells[1][2])
	dg.Grid.Cells[1][1].Link(dg.Grid.Cells[2][1])

	dg.Grid.Cells[1][2].Link(dg.Grid.Cells[2][2])

	dg.Grid.Cells[2][0].Link(dg.Grid.Cells[2][1])

	dg.Distances = FindDistances(dg.Grid.Cells[0][0])

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
