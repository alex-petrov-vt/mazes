package drawing

import "image/color"

const (
	EXTERNAL_PADDING = 100
	WALL_WIDTH       = 3
	CELL_WIDTH       = 50
	CELL_HEIGHT      = 50
)

var (
	WALL_COLOR       = color.RGBA{0, 0, 0, 255}
	BACKGROUND_COLOR = color.RGBA{255, 255, 255, 255}
)
