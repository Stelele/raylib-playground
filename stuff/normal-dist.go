package stuff

import (
	"math/rand"
	c "ray-random/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const BINS_NUM = 20
const BIN_WIDTH = c.VIRTUAL_WIDTH / BINS_NUM

var bins = make([]int, BINS_NUM)

func NormalDistInit() {
	for bin := range BINS_NUM {
		bins[bin] = 0
	}
}

func NormalDistDraw() {
	for idx, binVal := range bins {
		xPos := idx * int(BIN_WIDTH)
		yPos := c.VIRTUAL_HEIGHT - int32(binVal)
		rec := rl.Rectangle{
			X:      float32(xPos),
			Y:      float32(yPos),
			Width:  float32(BIN_WIDTH),
			Height: float32(binVal),
		}
		rl.DrawRectangleRec(rec, rl.White)
		rl.DrawRectangleLines(rec.ToInt32().X, rec.ToInt32().Y, rec.ToInt32().Width, int32(rec.Height), rl.Black)
	}
}

func NormalDistUpdate() {
	idx := rand.Intn(BINS_NUM)
	bins[idx] = int(clamp(0, float64(c.VIRTUAL_HEIGHT), float64(bins[idx]+1)))
}
