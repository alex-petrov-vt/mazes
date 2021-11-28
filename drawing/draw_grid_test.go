package drawing

import (
	"image/png"
	"os"
	"testing"

	"github.com/alex-petrov-vt/mazes/grid"
	"github.com/stretchr/testify/assert"
)

func TestGridDrawing(t *testing.T) {
	g := grid.NewGrid(4, 4)

	g.Cells[0][0].Link(g.Cells[0][1])
	g.Cells[0][1].Link(g.Cells[0][2])
	g.Cells[0][2].Link(g.Cells[1][2])
	g.Cells[1][2].Link(g.Cells[2][2])
	g.Cells[2][2].Link(g.Cells[3][2])
	g.Cells[3][2].Link(g.Cells[3][3])

	i := DrawGrid(g)

	handle, err := os.Open("testdata/grid-test.png")
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	expected, err := png.Decode(handle)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, expected, i)
}
