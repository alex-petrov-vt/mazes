package main

import (
	"math/rand"
	"strings"
)

type grid struct {
	rows, cols int
	cells      [][]*cell
}

func newGrid(rows, cols int) *grid {
	cells := prepareCells(rows, cols)
	return &grid{
		rows,
		cols,
		cells,
	}
}

func prepareCells(rows, cols int) [][]*cell {
	cells := createCells(rows, cols)
	cells = connectCells(rows, cols, cells)
	return cells
}

func createCells(rows, cols int) [][]*cell {
	var cells [][]*cell
	for row := 0; row < rows; row++ {
		var newRow []*cell
		for col := 0; col < cols; col++ {
			newRow = append(newRow, newCell(row, col))
		}
		cells = append(cells, newRow)
	}

	return cells
}

func connectCells(rows, cols int, cells [][]*cell) [][]*cell {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if row != 0 {
				cells[row][col].north = cells[row-1][col]
			}

			if row != rows-1 {
				cells[row][col].south = cells[row+1][col]
			}

			if col != 0 {
				cells[row][col].west = cells[row][col-1]
			}

			if col != cols-1 {
				cells[row][col].east = cells[row][col+1]
			}
		}
	}

	return cells
}

func (g *grid) getSize() int {
	return g.cols * g.rows
}

func (g *grid) getRandomCell(seed int64) *cell {
	rand.Seed(seed)
	row := rand.Intn(g.rows)
	col := rand.Intn(g.cols)
	return g.cells[row][col]
}

func (g *grid) String() string {
	result := strings.Builder{}
	result.WriteString("\n+" + strings.Repeat("---+", g.cols) + "\n")

	for row := 0; row < g.rows; row++ {
		rowLine := strings.Builder{}
		borderLine := strings.Builder{}
		rowLine.WriteString("|")
		borderLine.WriteString("+")
		for col := 0; col < g.cols; col++ {
			cell := g.cells[row][col]

			// Each cell is three spaces wide
			rowLine.WriteString("   ")

			// Check for cell below
			if row == g.rows-1 || !cell.isLinked(g.cells[row+1][col]) {
				borderLine.WriteString("---")
			} else {
				borderLine.WriteString("   ")
			}

			// Check for cell to the right
			if col == g.cols-1 || !cell.isLinked(g.cells[row][col+1]) {
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
