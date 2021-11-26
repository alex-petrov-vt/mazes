package main

import "math/rand"

func createBinaryTreeMaze(g *grid, seed int64) *grid {
	rand.Seed(seed)
	neighbors := make([]*cell, 0, 2)

	for row := range g.cells {
		for col := range g.cells[row] {
			currCell := g.cells[row][col]
			if currCell.south != nil {
				neighbors = append(neighbors, currCell.south)
			}

			if currCell.east != nil {
				neighbors = append(neighbors, currCell.east)
			}

			// The last cell in the corner will not have any neighbors
			if len(neighbors) > 0 {
				index := rand.Intn(len(neighbors))
				neighbor := neighbors[index]
				currCell.link(neighbor)
			}

			neighbors = neighbors[:0]
		}
	}

	return g
}
