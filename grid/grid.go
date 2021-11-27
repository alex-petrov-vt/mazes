package grid

import (
	"math/rand"
	"strings"
)

type Grid struct {
	Rows, Cols int
	Cells      [][]*Cell
}

func NewGrid(rows, cols int) *Grid {
	cells := prepareCells(rows, cols)
	return &Grid{
		rows,
		cols,
		cells,
	}
}

func prepareCells(rows, cols int) [][]*Cell {
	cells := createCells(rows, cols)
	cells = connectCells(rows, cols, cells)
	return cells
}

func createCells(rows, cols int) [][]*Cell {
	var cells [][]*Cell
	for row := 0; row < rows; row++ {
		var newRow []*Cell
		for col := 0; col < cols; col++ {
			newRow = append(newRow, NewCell(row, col))
		}
		cells = append(cells, newRow)
	}

	return cells
}

func connectCells(rows, cols int, cells [][]*Cell) [][]*Cell {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if row != 0 {
				cells[row][col].North = cells[row-1][col]
			}

			if row != rows-1 {
				cells[row][col].South = cells[row+1][col]
			}

			if col != 0 {
				cells[row][col].West = cells[row][col-1]
			}

			if col != cols-1 {
				cells[row][col].East = cells[row][col+1]
			}
		}
	}

	return cells
}

func (g *Grid) getSize() int {
	return g.Cols * g.Rows
}

func (g *Grid) getRandomCell(seed int64) *Cell {
	rand.Seed(seed)
	row := rand.Intn(g.Rows)
	col := rand.Intn(g.Cols)
	return g.Cells[row][col]
}

func (g *Grid) String() string {
	return g.PrintWithCellContent(g.getCellContent)
}

func (g *Grid) PrintWithCellContent(getCellContent func(*Cell) string) string {
	result := strings.Builder{}
	result.WriteString("\n+" + strings.Repeat("---+", g.Cols) + "\n")

	for row := 0; row < g.Rows; row++ {
		rowLine := strings.Builder{}
		borderLine := strings.Builder{}
		rowLine.WriteString("|")
		borderLine.WriteString("+")
		for col := 0; col < g.Cols; col++ {
			cell := g.Cells[row][col]

			// Each cell is three spaces wide
			rowLine.WriteString(" " + getCellContent(cell) + " ")

			// Check for cell below
			if row == g.Rows-1 || !cell.IsLinked(g.Cells[row+1][col]) {
				borderLine.WriteString("---")
			} else {
				borderLine.WriteString("   ")
			}

			// Check for cell to the right
			if col == g.Cols-1 || !cell.IsLinked(g.Cells[row][col+1]) {
				rowLine.WriteString("|")
			} else {
				rowLine.WriteString(" ")
			}
			borderLine.WriteString("+")
		}
		result.WriteString(rowLine.String() + "\n")
		result.WriteString(borderLine.String() + "\n")
	}

	return result.String()
}

func (g *Grid) getCellContent(c *Cell) string {
	return " "
}
