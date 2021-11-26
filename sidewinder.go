package main

import "math/rand"

func createSidewinderMaze(g *grid, seed int64) *grid {
	rand.Seed(seed)

	for row := range g.cells {
		run := make([]*cell, 0, g.cols)
		for col := range g.cells[row] {
			currCell := g.cells[row][col]
			run = append(run, currCell)

			isAtEastBoundary := (currCell.east == nil)
			isAtSouthBoundary := (currCell.south == nil)

			shouldCloseRun := isAtEastBoundary || (!isAtSouthBoundary && rand.Intn(2) == 0)

			if shouldCloseRun {
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
