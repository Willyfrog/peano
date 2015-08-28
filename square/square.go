// Squares are the components of a Matrix
package square

import (
	"fmt"
	"image/color"

	log "github.com/Sirupsen/logrus"
	"github.com/Willyfrog/peano/drawing"
	"github.com/Willyfrog/peano/point"
	"github.com/Willyfrog/peano/utils"
)

type Square struct {
	X      int
	Y      int
	Width  float32
	Points []*point.Point
}

// fitsIn Given a point.Point, check if it's inside it's boundaries
func (sq Square) fitsIn(pnt point.Point) bool {
	xo, yo := sq.Origin()
	xe, ye := sq.End()
	return utils.Between(pnt.X, xo, xe) && utils.Between(pnt.Y, yo, ye)
}

// Origin Get the equivalent of (0,0) for this square
func (sq Square) Origin() (x, y float32) {
	x = (float32(sq.X)) * sq.Width
	y = (float32(sq.Y)) * sq.Width
	return
}

// End Get the equivalent of (1,1) for this square
func (sq Square) End() (x, y float32) {
	x = (float32(sq.X + 1)) * sq.Width
	y = (float32(sq.Y + 1)) * sq.Width
	return
}

// Empty are any points inside?
func (sq Square) Empty() bool {
	return len(sq.Points) == 0
}

// Partition given a square subidivide it into 4
func (sq Square) Partition() [2][2]Square {
	width := sq.Width / 2.0
	sub := [2][2]Square{
		{
			Square{Width: width, Points: make([]*point.Point, 0)},
			Square{Width: width, Points: make([]*point.Point, 0)},
		},
		{
			Square{Width: width, Points: make([]*point.Point, 0)},
			Square{Width: width, Points: make([]*point.Point, 0)},
		},
	}
	pointsAssigned := 0
	for i, line := range sub {
		for j, subsq := range line {
			subsq.X = sq.X*2 + j
			subsq.Y = sq.Y*2 + i
			for _, p := range sq.Points {
				if subsq.fitsIn(*p) {
					subsq.Points = append(subsq.Points, p)
				}
			}
			sub[i][j] = subsq // WTF?
			pointsAssigned += len(subsq.Points)
		}
	}

	if pointsAssigned != len(sq.Points) {
		message := fmt.Sprintf("We missed some points while subdividing %d!=%d", pointsAssigned, len(sq.Points)) // don't want to find out later...
		log.Fatal(message)
		panic(message)
	}
	return sub
}

// Draw the square.
// Remember that the square should be ordered previously to make sure it
// prints in the intended order. Take a look at
// `github.com/Willyfrog/peano/matrix.Strategy` Interface
func (sq *Square) Draw(canvas *drawing.Canvas) {
	path := canvas.GetContext()
	if log.StandardLogger().Level == log.DebugLevel {
		if sq.X == 0 && sq.Y == 0 {
			path.SetStrokeColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
		} else if sq.X == 0 && sq.Y == 1 {
			path.SetStrokeColor(color.RGBA{0xff, 0x44, 0x44, 0xff})
		} else {
			path.SetStrokeColor(color.RGBA{0xcc, 0xcc, 0xcc, 0xff})
		}
	} else {
		path.SetStrokeColor(color.RGBA{0xcc, 0xcc, 0xcc, 0xff})
	}

	path.SetLineWidth(1)
	xo, yo := sq.Origin()
	xe, ye := sq.End()
	drawing.DrawSquare(xo, yo, xe, ye, path)
	path.Stroke()
	//path.FillStroke()
	//log.Debug(fmt.Sprintf("Result: %v", sq.Points))
	var origin *point.Point
	for _, pt := range sq.Points {
		pt.Draw(canvas)
		if origin != nil {
			linepath := canvas.GetContext()
			path.SetStrokeColor(color.RGBA{0x44, 0x44, 0x88, 0xff})
			path.SetLineWidth(5)
			drawing.DrawLine(origin.X, origin.Y, pt.X, pt.Y, linepath)
			linepath.Stroke()
		}
		origin = pt
	}
}

func (sq *Square) Connect() point.PointList {
	return point.PointList(sq.Points).Polyline(point.SortXY)
}

// String make a printable version of the square and its contents
func (sq *Square) String() string {
	pl := point.PointList(sq.Points)
	return fmt.Sprintf("[%d, %d]:\n%s", sq.X, sq.Y, (&pl).String())
}
