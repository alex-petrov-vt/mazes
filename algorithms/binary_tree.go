package algorithms

import (
	"math/rand"

	"github.com/alex-petrov-vt/mazes/grid"
)

func CreateBinaryTreeMaze(g *grid.Grid, seed int64) *grid.Grid {
	rand.Seed(seed)
	neighbors := make([]*grid.Cell, 0, 2)

	for row := range g.Cells {
		for col := range g.Cells[row] {
			currCell := g.Cells[row][col]
			if currCell.South != nil {
				neighbors = append(neighbors, currCell.South)
			}

			if currCell.East != nil {
				neighbors = append(neighbors, currCell.East)
			}

			// The last cell in the corner will not have any neighbors
			if len(neighbors) > 0 {
				index := rand.Intn(len(neighbors))
				neighbor := neighbors[index]
				currCell.Link(neighbor)
			}

			neighbors = neighbors[:0]
		}
	}

	return g
}
