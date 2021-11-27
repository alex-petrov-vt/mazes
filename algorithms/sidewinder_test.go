package algorithms

import (
	"fmt"
	"testing"

	"github.com/alex-petrov-vt/mazes/grid"
	"github.com/stretchr/testify/assert"
)

func TestCreatingSidewinderMaze(t *testing.T) {
	var randomSeed int64 = 1

	g := grid.NewGrid(2, 2)
	g = CreateSidewinderMaze(g, randomSeed)

	expected := `
+---+---+
|       |
+---+   +
|       |
+---+---+
`
	assert.Equal(t, expected, fmt.Sprintf("%s", g))

	g = grid.NewGrid(4, 4)
	g = CreateSidewinderMaze(g, randomSeed)

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
	assert.Equal(t, expected, fmt.Sprintf("%s", g))
}
