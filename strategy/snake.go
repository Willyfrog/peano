package strategy

import (
	"github.com/Willyfrog/peano/matrix"
	"github.com/Willyfrog/peano/point"
	"github.com/Willyfrog/peano/square"
)

type SnakeStrategy struct{}

func (s SnakeStrategy) OrderPoints(sq square.Square) {
	s.orderPoints(sq, point.SortXY)
}

func (s SnakeStrategy) orderPoints(sq square.Square, sortFunc point.SortFunction) {
	point.PointList(sq.Points).Polyline(sortFunc)
}

func (s SnakeStrategy) ConnectSquares(m matrix.Matrix) [][2]*point.Point {
	pl := make([][2]*point.Point, 2)
	for x, row := range m.Squares {
		for y, sq := range row {
			if y != 0 {
				result := connect(m.Squares[x][y-1], sq)
				pl = append(pl, result)
			}
		}
	}
	return pl
}

func connect(sq1, sq2 square.Square) [2]*point.Point {
	pl := new([2]*point.Point)
	if sq1.Y == sq2.Y {
		pl[0] = sq1.Points[len(sq1.Points)-1]
		pl[1] = sq2.Points[0]
	} else if sq1.X == 0 { // firts column
		pl[0] = sq1.Points[0]
		pl[1] = sq2.Points[0]

	} else {
		pl[0] = sq1.Points[len(sq1.Points)-1]
		pl[1] = sq2.Points[len(sq2.Points)-1]
	}
	return *pl
}
