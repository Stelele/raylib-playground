package stuff

import c "ray-random/constants"

var tiles []bool = make([]bool, c.COLS*c.ROWS)

func Init() {
	// ConwayInit()
	RandomWalkInit()
	// NormalDistInit()
}

func Draw() {
	// ConwayDraw()
	RandomWalkDraw()
	// NormalDistDraw()
}

func Update() {
	// ConwayUpdate()
	RandomWalkUpdate()
	// NormalDistUpdate()
}
