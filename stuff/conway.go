package stuff

import (
	"math/rand"
	c "ray-random/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const MAX_FRAME_NUM = 3

var curFrame = 0

func ConwayInit() {
	for y := range c.ROWS {
		for x := range c.COLS {
			if 10*rand.ExpFloat64() >= 50 {
				tiles[y*c.COLS+x] = true
			} else {
				tiles[y*c.COLS+x] = false
			}
		}
	}
}

func ConwayDraw() {
	for y := range c.ROWS {
		for x := range c.COLS {
			if tiles[y*c.COLS+x] {
				xPos := x * c.TILE_WIDTH
				yPos := y * c.TILE_HEIGHT
				rl.DrawRectangle(xPos, yPos, c.TILE_WIDTH, c.TILE_HEIGHT, rl.White)
			}
		}
	}
}

func ConwayUpdate() {
	handleInput()
	curFrame += 1

	if curFrame >= 60/MAX_FRAME_NUM && !rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		updateTiles()
		curFrame = 0
	}
}

func handleInput() {
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		pos := rl.GetMousePosition()

		x := int32(clamp(0, float64(c.VIRTUAL_WIDTH), float64(pos.X/float32(c.TILE_WIDTH))))
		y := int32(clamp(0, float64(c.VIRTUAL_HEIGHT), float64(pos.Y/float32(c.TILE_HEIGHT))))

		tiles[y*c.COLS+x] = true
	}
}

func clamp(min float64, max float64, num float64) float64 {
	if min > num {
		return min
	}

	if max < num {
		return max
	}

	return num
}

func updateTiles() {
	tmpTiles := make([]bool, c.COLS*c.ROWS)
	for y := range c.ROWS {
		for x := range c.COLS {
			neighbours := 0
			for i := y - 1; i <= y+1; i++ {
				if i < 0 || i >= c.ROWS {
					continue
				}

				for j := x - 1; j <= x+1; j++ {
					if j < 0 || j >= c.COLS {
						continue
					}
					if tiles[i*c.COLS+j] {
						neighbours += 1
					}
				}
			}

			idx := y*c.COLS + x
			if tiles[idx] && neighbours < 2 {
				tmpTiles[idx] = false
				continue
			}
			if tiles[idx] && (neighbours == 2 || neighbours == 3) {
				tmpTiles[idx] = true
				continue
			}
			if tiles[idx] && neighbours > 3 {
				tmpTiles[idx] = false
			}
			if !tiles[idx] && neighbours == 3 {
				tmpTiles[idx] = true
			}
		}
	}

	tiles = tmpTiles
}
