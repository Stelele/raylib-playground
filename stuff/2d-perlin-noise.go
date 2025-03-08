package stuff

import (
	"math/rand"
	c "ray-random/constants"

	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var pixelVals = make([]uint8, c.VIRTUAL_WIDTH*c.VIRTUAL_HEIGHT)
var pGen = perlin.NewPerlin(2, 2, 3, int64(rand.Int31()))
var t = 0.

func PerlinNoise2DInit() {
}

func PerlinNoise2DDraw() {
	for y := range c.VIRTUAL_HEIGHT {
		for x := range c.VIRTUAL_WIDTH {
			color := rl.Color{
				R: pixelVals[y*c.VIRTUAL_WIDTH+x],
				G: pixelVals[y*c.VIRTUAL_WIDTH+x],
				B: pixelVals[y*c.VIRTUAL_WIDTH+x],
				A: 255,
			}
			rl.DrawPixel(x, y, color)
		}
	}

	t += 0.11
}

func PerlinNoise2DUpdate() {
	yoffset := 0.

	for y := range c.VIRTUAL_HEIGHT {
		xoffset := 0.
		for x := range c.VIRTUAL_WIDTH {
			noise := pGen.Noise3D(float64(xoffset), float64(yoffset), t)
			pixelVals[y*c.VIRTUAL_WIDTH+x] = uint8(rl.Remap(float32(noise), -1, 1, 0, 255))
			xoffset += 0.01
		}
		yoffset += 0.01
	}
}
