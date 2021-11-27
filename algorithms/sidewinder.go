package algorithms

import (
	"math/rand"

	"github.com/alex-petrov-vt/mazes/grid"
)

func CreateSidewinderMaze(g *grid.Grid, seed int64) *grid.Grid {
	rand.Seed(seed)

	for row := range g.Cells {
		run := make([]*grid.Cell, 0, g.Cols)
		for col := range g.Cells[row] {
			currCell := g.Cells[row][col]
			run = append(run, currCell)

			isAtEastBoundary := (currCell.East == nil)
			isAtSouthBoundary := (currCell.South == nil)

			shouldCloseRun := isAtEastBoundary || (!isAtSouthBoundary && rand.Intn(2) == 0)

			if shouldCloseRun {
				index := rand.Intn(len(run))
				chosenCell := run[index]
				if chosenCell.South != nil {
					chosenCell.Link(chosenCell.South)
				}
				// clear slice but keep allocated memory
				run = run[:0]
			} else {
				currCell.Link(currCell.East)
			}
		}
	}

	return g
}
