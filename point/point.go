package point

import (
	"github.com/Willyfrog/peano/drawing"
	"image/color"
	"math/rand"
)

type Point struct {
	X float32
	Y float32
}

// Random
// Get a random point whose coordinates are in the range: [0.0, 1.0)
func Random() Point {
	return Point{X: rand.Float32(), Y: rand.Float32()}
}

// RandomSlice
// get a length of random points
func RandomSlice(length int) []Point {
	pl := make([]Point, length)
	for i := range pl {
		pl[i] = Random()
	}
	return pl
}

func (pt *Point) Draw(canvas *drawing.Canvas) {
	path := canvas.GetContext()
	path.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	path.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	path.SetLineWidth(5)
	drawing.DrawPoint(0.0, 0.0, path)
	path.Close()
	path.FillStroke()
}
