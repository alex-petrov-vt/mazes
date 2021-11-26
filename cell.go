package main

type cell struct {
	row, col                 int
	north, east, south, west *cell
	links                    map[*cell]bool
}

func newCell(row, col int) *cell {
	links := make(map[*cell]bool)
	return &cell{row: row, col: col, links: links}
}

func (c *cell) link(other *cell) {
	c.links[other] = true
	other.links[c] = true
}

func (c *cell) unlink(other *cell) {
	delete(c.links, other)
	delete(other.links, c)
}
