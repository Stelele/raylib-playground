package stuff

import (
	"math/rand"
	c "ray-random/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var curX = c.COLS / 2
var curY = c.ROWS / 2

func RandomWalkInit() {
	uniformInit()
}

func uniformInit() {
	for y := range c.ROWS {
		for x := range c.COLS {
			tiles[y*c.COLS+x] = false
		}
	}

	tiles[curY*c.COLS+curX] = true
}

func RandomWalkDraw() {
	ConwayDraw()
}

func RandomWalkUpdate() {
	// uniformWalkUpdate()
	biasedUniformWalkUpdate()
}

func uniformWalkUpdate() {
	dX := rand.Intn(3) - 1
	dY := rand.Intn(3) - 1

	curX = int32(clamp(0, float64(c.COLS-1), float64(curX+int32(dX))))
	curY = int32(clamp(0, float64(c.ROWS-1), float64(curY+int32(dY))))

	tiles[curY*c.COLS+curX] = true
}

func biasedUniformWalkUpdate() {
	num := rand.Float32()

	if num < 0.05 {
		tX := int32(clamp(0, float64(c.COLS-1), float64(rl.GetMouseX())/float64(c.TILE_WIDTH)))
		tY := int32(clamp(0, float64(c.COLS-1), float64(rl.GetMouseY())/float64(c.TILE_HEIGHT)))

		dX := tX - curX
		dY := tY - curY

		if dX < 0 {
			dX = -1
		} else {
			dX = 1
		}
		if dY < 0 {
			dY = -1
		} else {
			dY = 1
		}

		curX = int32(clamp(0, float64(c.COLS-1), float64(curX+dX)))
		curY = int32(clamp(0, float64(c.ROWS-1), float64(curY+dY)))
	} else {
		dX := rand.Intn(3) - 1
		dY := rand.Intn(3) - 1
		curX = int32(clamp(0, float64(c.COLS-1), float64(curX+int32(dX))))
		curY = int32(clamp(0, float64(c.ROWS-1), float64(curY+int32(dY))))
	}

	tiles[curY*c.COLS+curX] = true
}
