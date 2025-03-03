package main

import (
	"math"
	c "ray-random/constants"
	"ray-random/stuff"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var screenWidth int32 = c.VIRTUAL_WIDTH
var screenHeight int32 = c.VIRTUAL_HEIGHT

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Random Stuff")
	defer rl.CloseWindow()

	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)

	stuff.Init()

	for !rl.WindowShouldClose() {
		stuff.Update()

		rl.BeginDrawing()
		camera := scaleContentsToWindow()
		rl.BeginMode2D(camera)

		rl.ClearBackground(rl.Black)
		stuff.Draw()

		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func scaleContentsToWindow() rl.Camera2D {
	scaleX := float64(rl.GetScreenWidth()) / float64(c.VIRTUAL_WIDTH)
	scaleY := float64(rl.GetScreenHeight()) / float64(c.VIRTUAL_HEIGHT)
	scale := math.Min(scaleX, scaleY)

	offsetX := (float64(rl.GetScreenWidth()) - (float64(c.VIRTUAL_WIDTH) * scale)) * 0.5
	offsetY := (float64(rl.GetScreenHeight()) - (float64(c.VIRTUAL_HEIGHT) * scale)) * 0.5

	camera := rl.Camera2D{
		Offset:   rl.Vector2{X: float32(offsetX), Y: float32(offsetY)},
		Target:   rl.Vector2{X: 0, Y: 0},
		Rotation: 0,
		Zoom:     float32(scale),
	}

	return camera
}
