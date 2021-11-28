package distances

import "github.com/alex-petrov-vt/mazes/grid"

type Distances struct {
	Root  *grid.Cell
	Cells map[*grid.Cell]int
}

func newDistances(c *grid.Cell) *Distances {
	cells := make(map[*grid.Cell]int)
	cells[c] = 0
	return &Distances{Root: c, Cells: cells}
}

func FindDistances(c *grid.Cell) *Distances {
	d := newDistances(c)

	nextCells := make([]*grid.Cell, 0)
	nextCells = append(nextCells, c)

	for len(nextCells) != 0 {
		newNextCells := make([]*grid.Cell, 0)

		for _, currCell := range nextCells {
			for link := range currCell.Links {
				_, alreadyCounted := d.Cells[link]
				if alreadyCounted {
					continue
				} else {
					d.Cells[link] = d.Cells[currCell] + 1
					newNextCells = append(newNextCells, link)
				}
			}
		}

		nextCells = newNextCells
	}

	return d
}

func (d *Distances) FindShortestPath(c *grid.Cell) *Distances {
	currCell := c
	path := newDistances(d.Root)
	path.Cells[currCell] = d.Cells[currCell]

	for currCell != d.Root {
		for link := range currCell.Links {
			if d.Cells[link] < d.Cells[currCell] {
				path.Cells[link] = d.Cells[link]
				currCell = link
				break
			}
		}
	}

	return path
}

func (d *Distances) getMostDistantPoint() *grid.Cell {
	maxDistance := 0
	maxCell := d.Root

	for currCell := range d.Cells {
		if d.Cells[currCell] > maxDistance {
			maxCell = currCell
			maxDistance = d.Cells[currCell]
		}
	}

	return maxCell
}

func (d *Distances) FindLongestPath() *Distances {
	mostDistantPoint := d.getMostDistantPoint()

	secondDistances := FindDistances(mostDistantPoint)

	goal := secondDistances.getMostDistantPoint()

	longestPath := secondDistances.FindShortestPath(goal)
	return longestPath
}
