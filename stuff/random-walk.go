package stuff

import (
	"math/rand"
	c "ray-random/constants"

	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var curX = c.COLS / 2
var curY = c.ROWS / 2

var pSource = rl.Vector2{X: 0, Y: 1000}
var pPos = make([]rl.Vector2, 0)
var p = perlin.NewPerlin(2, 2, 3, rand.Int63())

func RandomWalkInit() {
	// uniformInit()
	perlinInit()
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
	// ConwayDraw()
	perlinDraw()
}

func RandomWalkUpdate() {
	// uniformWalkUpdate()
	// biasedUniformWalkUpdate()
	perlinWalkUpdate()
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

func perlinWalkUpdate() {

	dX := p.Noise1D(float64(pSource.X))
	dY := p.Noise1D(float64(pSource.Y))

	pSource.X += 0.01
	pSource.Y += 0.01

	dX = float64(rl.Remap(float32(dX), -1, 1, 0, float32(c.VIRTUAL_WIDTH-1)))
	dY = float64(rl.Remap(float32(dY), -1, 1, 0, float32(c.VIRTUAL_HEIGHT-1)))
	val := rl.Vector2{
		X: float32(dX),
		Y: float32(dY),
	}

	// if len(pPos) == 0 {
	// 	dX = float64(rl.Remap(float32(dX), -1, 1, 0, float32(c.VIRTUAL_WIDTH-1)))
	// 	dY = float64(rl.Remap(float32(dY), -1, 1, 0, float32(c.VIRTUAL_HEIGHT-1)))
	// 	pPos = append(pPos, rl.Vector2{
	// 		X: float32(dX),
	// 		Y: float32(dY),
	// 	})
	// 	return
	// }

	// dX = float64(rl.Remap(float32(dX), -1, 1, -10, 10))
	// dY = float64(rl.Remap(float32(dY), -1, 1, -10, 10))

	// val := rl.Vector2{
	// 	X: float32(pPos[len(pPos)-1].X) + float32(dX),
	// 	Y: float32(pPos[len(pPos)-1].Y) + float32(dY),
	// }

	pPos = append(pPos, val)

	if len(pPos) > 3000 {
		pPos = pPos[1:]
	}

}

func perlinDraw() {
	for _, pos := range pPos {
		rl.DrawCircle(int32(pos.X), int32(pos.Y), float32(2*c.TILE_WIDTH), rl.White)
		rl.DrawCircleLines(int32(pos.X), int32(pos.Y), float32(2*c.TILE_WIDTH), rl.Black)
	}
}

func perlinInit() {

}
