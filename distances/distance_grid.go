package distances

import (
	"strconv"

	"github.com/alex-petrov-vt/mazes/grid"
)

type DistanceGrid struct {
	Grid      *grid.Grid
	Distances *Distances
}

func NewDistanceGrid(g *grid.Grid) *DistanceGrid {
	return &DistanceGrid{Grid: g, Distances: nil}
}

func (dg *DistanceGrid) String() string {
	return dg.Grid.PrintWithCellContent(dg.getCellContent)
}

func (dg *DistanceGrid) getCellContent(c *grid.Cell) string {
	distance, found := dg.Distances.Cells[c]
	if dg.Distances != nil && found {
		return strconv.FormatInt(int64(distance), 36)
	}

	return " "
}
