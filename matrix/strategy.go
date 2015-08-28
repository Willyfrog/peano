package matrix

import (
	"github.com/Willyfrog/peano/point"
	"github.com/Willyfrog/peano/square"
)

// Strategy interface defines the order inside the square and inside the matrix.
// this order will define how the matrix will be drawn
// there are examples of use in the `github.com/Willyfrog/peano/strategy` module
type Strategy interface {
	OrderPoints(sq square.Square)
	ConnectSquares(m Matrix) [][]*point.Point
}
