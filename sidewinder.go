package main

import "math/rand"

func createSidewinderMaze(g *grid, seed int64) *grid {
	rand.Seed(seed)

	for row := range g.cells {
		run := make([]*cell, 0, g.cols)
		for col := range g.cells[row] {
			currCell := g.cells[row][col]
			run = append(run, currCell)

			if currCell.east == nil || (currCell.south != nil && rand.Intn(2) == 0) {
				index := rand.Intn(len(run))
				chosenCell := run[index]
				if chosenCell.south != nil {
					chosenCell.link(chosenCell.south)
				}
				// clear slice but keep allocated memory
				run = run[:0]
			} else {
				currCell.link(currCell.east)
			}
		}
	}

	return g
}
