// package point
// the smallest unit in this project.
package point

import (
	"fmt"
	"image/color"
	"math/rand"
	"strings"
	"time"

	"github.com/Willyfrog/peano/drawing"
)

type Point struct {
	X float32
	Y float32
}

type PointList []*Point

type SortFunction func(*Point, *Point) bool

// Random
// Get a random point whose coordinates are in the range: [0.0, 1.0)
// if it's not being called from RandomSlice, remember to call
// `rand.Seed` previously so you get different results each time.
func Random() Point {
	return Point{X: rand.Float32(), Y: rand.Float32()}
}

// RandomSlice
// get a length of random points
func RandomSlice(length int) []Point {
	rand.Seed(time.Now().UTC().UnixNano())
	pl := make([]Point, length)
	for i := range pl {
		pl[i] = Random()
	}
	return pl
}

// Draw the point into the canvas
func (pt *Point) Draw(canvas *drawing.Canvas) {
	path := canvas.GetContext()
	path.SetFillColor(color.RGBA{0x88, 0xff, 0x88, 0xff})
	path.SetStrokeColor(color.RGBA{0x88, 0xff, 0x88, 0xff})
	path.SetLineWidth(5)
	drawing.DrawPoint(pt.X, pt.Y, path)
	path.FillStroke()
}

// Polyline define a polyline by ordering an slice of points
func (ps PointList) Polyline(sortFunc SortFunction) PointList {
	//ps.QuickSort(sortFunc)
	QuickSort(ps, sortFunc)
	return ps
}

//String generate a string form a list of points
//mostly for debugging purposes
func (pl *PointList) String() string {
	st := make([]string, len(*pl))
	for _, p := range *pl {
		st = append(st, p.String())
	}
	return strings.Join(st, ", ")
}

// String marshalls a point into a string for reading
func (p *Point) String() string {
	return fmt.Sprintf("[%f, %f]", p.X, p.Y)
}
