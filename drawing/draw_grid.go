package drawing

import (
	"image"

	"github.com/alex-petrov-vt/mazes/grid"
)

func DrawGrid(g *grid.Grid) *image.RGBA {
	IMAGE_WIDTH := 2*EXTERNAL_PADDING + (CELL_WIDTH * g.Cols)
	IMAGE_HEIGHT := 2*EXTERNAL_PADDING + (CELL_HEIGHT * g.Rows)
	i := image.NewRGBA(image.Rect(0, 0, IMAGE_WIDTH, IMAGE_HEIGHT))

	setBackgroundColor(i)
	drawExternalBoundary(i)
	drawCells(i, g)

	return i
}

func setBackgroundColor(i *image.RGBA) {
	b := i.Bounds()
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			i.SetRGBA(x, y, BACKGROUND_COLOR)
		}
	}
}

func drawCells(i *image.RGBA, g *grid.Grid) {
	for row := 0; row < g.Rows; row++ {
		for col := 0; col < g.Cols; col++ {
			c := g.Cells[row][col]
			drawCell(i, c)
		}
	}
}

// Each cell is responsible for drawing its right
// and bottom walls
func drawCell(i *image.RGBA, c *grid.Cell) {
	if _, linkedToBot := c.Links[c.South]; !linkedToBot {
		drawBottomCellWall(i, c)
	}
	if _, linkedToRight := c.Links[c.East]; !linkedToRight {
		drawRightCellWall(i, c)
	}
}

func drawRightCellWall(i *image.RGBA, c *grid.Cell) {
	b := i.Bounds()
	x1 := b.Min.X + EXTERNAL_PADDING + (CELL_WIDTH)*(c.Col+1) - WALL_WIDTH
	x2 := b.Min.X + EXTERNAL_PADDING + (CELL_WIDTH)*(c.Col+1) + WALL_WIDTH
	y1 := b.Min.Y + EXTERNAL_PADDING + (CELL_HEIGHT)*(c.Row) - WALL_WIDTH
	y2 := b.Min.Y + EXTERNAL_PADDING + (CELL_HEIGHT)*(c.Row+1) + WALL_WIDTH

	drawWall(i, x1, x2, y1, y2)
}

func drawBottomCellWall(i *image.RGBA, c *grid.Cell) {
	b := i.Bounds()
	x1 := b.Min.X + EXTERNAL_PADDING + (CELL_WIDTH)*(c.Col) - WALL_WIDTH
	x2 := b.Min.X + EXTERNAL_PADDING + (CELL_WIDTH)*(c.Col+1) + WALL_WIDTH
	y1 := b.Min.Y + EXTERNAL_PADDING + (CELL_HEIGHT)*(c.Row+1) - WALL_WIDTH
	y2 := b.Min.Y + EXTERNAL_PADDING + (CELL_HEIGHT)*(c.Row+1) + WALL_WIDTH

	drawWall(i, x1, x2, y1, y2)
}

func drawExternalBoundary(i *image.RGBA) {
	drawTopWall(i)
	drawLeftWall(i)
	drawBottomWall(i)
	drawRightWall(i)
}

func drawTopWall(i *image.RGBA) {
	b := i.Bounds()
	x1 := b.Min.X + EXTERNAL_PADDING - WALL_WIDTH
	x2 := b.Max.X - EXTERNAL_PADDING + WALL_WIDTH
	y1 := b.Min.Y + EXTERNAL_PADDING - WALL_WIDTH
	y2 := b.Min.Y + EXTERNAL_PADDING + WALL_WIDTH

	drawWall(i, x1, x2, y1, y2)
}

func drawLeftWall(i *image.RGBA) {
	b := i.Bounds()

	x1 := b.Min.X + EXTERNAL_PADDING - WALL_WIDTH
	x2 := b.Min.X + EXTERNAL_PADDING + WALL_WIDTH
	y1 := b.Min.Y + EXTERNAL_PADDING - WALL_WIDTH
	y2 := b.Max.Y - EXTERNAL_PADDING + WALL_WIDTH

	drawWall(i, x1, x2, y1, y2)
}

func drawRightWall(i *image.RGBA) {
	b := i.Bounds()

	x1 := b.Max.X - EXTERNAL_PADDING - WALL_WIDTH
	x2 := b.Max.X - EXTERNAL_PADDING + WALL_WIDTH
	y1 := b.Min.Y + EXTERNAL_PADDING - WALL_WIDTH
	y2 := b.Max.Y - EXTERNAL_PADDING + WALL_WIDTH

	drawWall(i, x1, x2, y1, y2)
}

func drawBottomWall(i *image.RGBA) {
	b := i.Bounds()

	x1 := b.Min.X + EXTERNAL_PADDING - WALL_WIDTH
	x2 := b.Max.X - EXTERNAL_PADDING + WALL_WIDTH
	y1 := b.Max.Y - EXTERNAL_PADDING - WALL_WIDTH
	y2 := b.Max.Y - EXTERNAL_PADDING + WALL_WIDTH

	drawWall(i, x1, x2, y1, y2)
}

func drawWall(i *image.RGBA, x1, x2, y1, y2 int) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			i.SetRGBA(x, y, WALL_COLOR)
		}
	}
}
