package square

import (
	"fmt"
	"github.com/Willyfrog/peano/drawing"
	"github.com/Willyfrog/peano/point"
	"github.com/Willyfrog/peano/utils"
	"image/color"
)

type Square struct {
	X      int
	Y      int
	Width  float32
	Points []*point.Point
}

// fitsIn
// given a point.Point, check if it's inside it's boundaries
func (sq Square) fitsIn(pnt point.Point) bool {
	xo, yo := sq.Origin()
	xe, ye := sq.End()
	return utils.Between(pnt.X, xo, xe) && utils.Between(pnt.Y, yo, ye)
}

// Origin
// get the equivalent of (0,0) for this square
func (sq Square) Origin() (x, y float32) {
	x = (float32(sq.X)) * sq.Width
	y = (float32(sq.Y)) * sq.Width
	return
}

// End
// get the equivalent of (1,1) for this square
func (sq Square) End() (x, y float32) {
	x = (float32(sq.X + 1)) * sq.Width
	y = (float32(sq.Y + 1)) * sq.Width
	return
}

// Empty
// are any points inside?
func (sq Square) Empty() bool {
	return len(sq.Points) == 0
}

// Partition
// given a square subidivide it into 4
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
		panic(fmt.Sprintf("We missed some points while subdividing %d!=%d", pointsAssigned, len(sq.Points))) // don't want to find out later...
	}
	return sub
}

func (sq *Square) Draw(canvas *drawing.Canvas) {
	path := canvas.GetContext()
	path.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	path.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	path.SetLineWidth(5)
	xo, yo := sq.Origin()
	xe, ye := sq.End()
	drawing.DrawSquare(xo, yo, xe, ye, path)
	path.Close()
	path.FillStroke()
	for _, pt := range sq.Points {
		pt.Draw(canvas)
	}
}