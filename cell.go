package main

type cell struct {
	row, col                 int
	north, east, south, west *cell
}

func newCell(row, col int) *cell {
	return &cell{row: row, col: col}
}
