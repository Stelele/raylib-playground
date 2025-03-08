package stuff

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"

	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var planeWidth = 100
var planeHeight = 100

var heightMap = make([]float32, planeWidth*planeHeight)
var pSource2 = perlin.NewPerlin(2, 2, 3, rand.Int63())
var timer = 0
var timerOff = 0.
var freq = 3.2

var image *rl.Image
var texture rl.Texture2D
var model rl.Model

var mesh rl.Mesh

func PerlinNoise3DInit() {
	curFrame = 0
	t = 0

	image = rl.GenImageColor(planeWidth, planeHeight, rl.Black)
	texture = rl.LoadTextureFromImage(image)
	mesh = rl.GenMeshPlane(float32(planeWidth), float32(planeHeight), planeWidth-1, planeHeight-1)

	model = rl.LoadModelFromMesh(mesh)
}

func PerlinNoise3DDraw() {
	for y := range planeHeight {
		for x := range planeWidth {
			val := heightMap[y*planeWidth+x] * 20

			pos := rl.Vector3{
				X: float32(x),
				Y: val,
				Z: float32(y),
			}

			rl.DrawCube(pos, 1, 1, 1, rl.White)
			rl.DrawCubeWires(pos, 1, 1, 1, rl.Black)

		}
	}
}

func PerlinNoise3DUpdate() {

	t += 0.02
	timer += 1

	if timer%10 == 0 {
		// diff := pSource2.Noise1D(float64(timerOff))
		// freq += float64(rl.Remap(float32(diff), -1, 1, -0.1, 0.1))
		timerOff += 0.01
		timer = 0
	}

	yOffset := 0.
	for y := range planeHeight {
		xOffset := 0.
		for x := range planeWidth {

			heightMap[y*planeWidth+x] = float32(pSource2.Noise3D(xOffset*freq, yOffset*freq, t)) +
				0.5*float32(pSource2.Noise3D(xOffset*freq*2, yOffset*freq*2, t)) +
				0.25*float32(pSource2.Noise3D(xOffset*freq*3, yOffset*freq*3, t))
			xOffset += 0.01
		}
		yOffset += 0.01
	}
}

func float32ToByteArray(f float32) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
