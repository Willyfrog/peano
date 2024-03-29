package strategy

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/Willyfrog/peano/matrix"
	"github.com/Willyfrog/peano/point"
	"github.com/Willyfrog/peano/square"
)

// Implements matrix.Strategy interface
type SnakeStrategy struct {
	size int
}

// NewSnake get a pointer to the snake strategy
func NewSnake(size int) *SnakeStrategy {
	return &SnakeStrategy{size - 1}
}

// Order points inside a square
func (s SnakeStrategy) OrderPoints(sq square.Square) {
	modY := sq.Y % 2
	if sq.X == 0 && modY == 0 {
		s.orderPoints(sq, point.SortCounterRadial(sq.BottomLeft()))
	} else if sq.X == 0 && modY == 1 {
		s.orderPoints(sq, point.SortRadial(sq.End()))
	} else if sq.X == s.size && modY == 0 {
		s.orderPoints(sq, point.SortRadial(sq.TopRight()))
	} else if sq.X == s.size && modY == 1 {
		s.orderPoints(sq, point.SortCounterRadial(sq.Origin()))
	} else {
		s.orderPoints(sq, point.SortXY)
	}
}

// Order squares inside a matrix
func (s SnakeStrategy) ConnectSquares(m matrix.Matrix) [][]*point.Point {
	pl := make([][]*point.Point, 0)
	var curr, prev square.Square
	for x := 0; x <= s.size; x++ {
		for y := 0; y <= s.size; y++ {
			px, py, err := getPrevious(x, y, s.size)
			if err != nil {
				continue
			} else {
				curr = m.Squares[x][y]
				prev = m.Squares[px][py]
			}
			log.Debug(fmt.Sprintf("conecting [%d][%d]-[%d][%d]", prev.X, prev.Y, curr.X, curr.Y))
			result := connect(prev, curr)
			pl = append(pl, result)
		}
	}
	return pl
}

func getPrevious(x, y, max int) (int, int, error) {
	if x == 0 && y == 0 {
		return -1, -1, fmt.Errorf("No previous item")
	} else if x == 0 && y%2 == 0 {
		return 0, y - 1, nil
	} else if x == max && y%2 == 1 {
		return max, y - 1, nil
	} else if y%2 == 0 {
		return x - 1, y, nil
	} else {
		return x + 1, y, nil
	}
}

func (s SnakeStrategy) orderPoints(sq square.Square, sortFunc point.SortFunction) {
	point.PointList(sq.Points).Polyline(sortFunc)
}

func connect(sq1, sq2 square.Square) []*point.Point {
	pl := make([]*point.Point, 2)
	if sq1.Y == sq2.Y {
		if sq1.X < sq2.X {
			pl[0] = sq1.Points[len(sq1.Points)-1]
			pl[1] = sq2.Points[0]
		} else {
			pl[0] = sq1.Points[0]
			pl[1] = sq2.Points[len(sq2.Points)-1]
		}
	} else if sq1.X == 0 { // first column
		pl[0] = sq1.Points[0]
		pl[1] = sq2.Points[0]

	} else {
		pl[0] = sq1.Points[len(sq1.Points)-1]
		pl[1] = sq2.Points[len(sq2.Points)-1]
	}
	//log.Debug(fmt.Sprintf("Line is %v", pl))
	return pl
}
