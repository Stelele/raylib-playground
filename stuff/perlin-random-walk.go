package stuff

import (
	"math/rand"
	c "ray-random/constants"

	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var pSource = rl.Vector2{X: 0, Y: 1000}
var pPos = make([]rl.Vector2, 0)
var p = perlin.NewPerlin(2, 2, 3, rand.Int63())

func PerlinWalkUpdate() {

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

	if len(pPos) > 100 {
		pPos = pPos[1:]
	}
}

func PerlinDraw() {
	for _, pos := range pPos {
		rl.DrawCircle(int32(pos.X), int32(pos.Y), float32(2*c.TILE_WIDTH), rl.White)
		rl.DrawCircleLines(int32(pos.X), int32(pos.Y), float32(2*c.TILE_WIDTH), rl.Black)
	}
}

func PerlinInit() {

}
