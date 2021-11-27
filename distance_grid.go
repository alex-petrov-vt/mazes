package main

import "strconv"

type distanceGrid struct {
	grid      *grid
	distances *distances
}

func newDistanceGrid(rows, cols int) *distanceGrid {
	grid := newGrid(rows, cols)
	return &distanceGrid{grid: grid, distances: nil}
}

func (dg *distanceGrid) String() string {
	return dg.grid.printWithCellContent(dg.getCellContent)
}

func (dg *distanceGrid) getCellContent(c *cell) string {
	distance, found := dg.distances.cells[c]
	if dg.distances != nil && found {
		return strconv.FormatInt(int64(distance), 36)
	}

	return " "
}
