package grid

type Cell struct {
	Row, Col                 int
	North, East, South, West *Cell
	Links                    map[*Cell]bool
}

func NewCell(row, col int) *Cell {
	links := make(map[*Cell]bool)
	return &Cell{Row: row, Col: col, Links: links}
}

func (c *Cell) Link(other *Cell) {
	c.Links[other] = true
	other.Links[c] = true
}

func (c *Cell) Unlink(other *Cell) {
	delete(c.Links, other)
	delete(other.Links, c)
}

func (c *Cell) IsLinked(other *Cell) bool {
	_, hasLink := c.Links[other]
	return hasLink
}
