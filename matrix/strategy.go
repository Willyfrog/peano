package matrix

import (
	"github.com/Willyfrog/peano/point"
	"github.com/Willyfrog/peano/square"
)

type Strategy interface {
	OrderPoints(sq square.Square)
	ConnectSquares(m Matrix) [][2]*point.Point
}
