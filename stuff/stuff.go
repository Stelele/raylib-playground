package stuff

import c "ray-random/constants"

var tiles []bool = make([]bool, c.COLS*c.ROWS)

func Init() {
	// ConwayInit()
	// RandomWalkInit()
	// NormalDistInit()
	// PerlinInit()
	// PerlinNoise2DInit()
	PerlinNoise3DInit()
}

func Draw() {
	// ConwayDraw()
	// RandomWalkDraw()
	// NormalDistDraw()
	// PerlinDraw()
	// PerlinNoise2DDraw()
	PerlinNoise3DDraw()
}

func Update() {
	// ConwayUpdate()
	// RandomWalkUpdate()
	// NormalDistUpdate()
	// PerlinWalkUpdate()
	// PerlinNoise2DUpdate()
	PerlinNoise3DUpdate()
}
