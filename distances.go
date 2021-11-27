package main

type distances struct {
	root  *cell
	cells map[*cell]int
}

func newDistances(c *cell) *distances {
	cells := make(map[*cell]int)
	cells[c] = 0
	return &distances{root: c, cells: cells}
}

func findDistances(c *cell) *distances {
	d := newDistances(c)

	nextCells := make([]*cell, 0)
	nextCells = append(nextCells, c)

	for len(nextCells) != 0 {
		newNextCells := make([]*cell, 0)

		for _, currCell := range nextCells {
			for link := range currCell.links {
				_, alreadyCounted := d.cells[link]
				if alreadyCounted {
					continue
				} else {
					d.cells[link] = d.cells[currCell] + 1
					newNextCells = append(newNextCells, link)
				}
			}
		}

		nextCells = newNextCells
	}

	return d
}

func (d *distances) findShortestPath(c *cell) *distances {
	currCell := c
	path := newDistances(d.root)
	path.cells[currCell] = d.cells[currCell]

	for currCell != d.root {
		for link := range currCell.links {
			if d.cells[link] < d.cells[currCell] {
				path.cells[link] = d.cells[link]
				currCell = link
				break
			}
		}
	}

	return path
}

func (d *distances) getMostDistantPoint() *cell {
	maxDistance := 0
	maxCell := d.root

	for currCell := range d.cells {
		if d.cells[currCell] > maxDistance {
			maxCell = currCell
			maxDistance = d.cells[currCell]
		}
	}

	return maxCell
}

func (d *distances) findLongestPath() *distances {
	mostDistantPoint := d.getMostDistantPoint()

	secondDistances := findDistances(mostDistantPoint)

	goal := secondDistances.getMostDistantPoint()

	longestPath := secondDistances.findShortestPath(goal)
	return longestPath
}
