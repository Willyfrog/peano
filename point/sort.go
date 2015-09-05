package point

// Several similar sort functions. A SortXY will give more importance
// to X then to Y. If any of those characters is lowercase, it'll mean
// that the order in that case is reversed.
func SortXY(p1, p2 *Point) bool {
	return (p1.X > p2.X) || (p1.X == p2.X && p1.Y >= p2.Y)
}

func SortXy(p1, p2 *Point) bool {
	return (p1.X > p2.X) || (p1.X == p2.X && p1.Y <= p2.Y)
}

func SortxY(p1, p2 *Point) bool {
	return (p1.X < p2.X) || (p1.X == p2.X && p1.Y >= p2.Y)
}

func Sortxy(p1, p2 *Point) bool {
	return (p1.X < p2.X) || (p1.X == p2.X && p1.Y <= p2.Y)
}

func SortYX(p1, p2 *Point) bool {
	return (p1.Y > p2.Y) || (p1.Y == p2.Y && p1.X >= p2.X)
}

func SortYx(p1, p2 *Point) bool {
	return (p1.Y > p2.Y) || (p1.Y == p2.Y && p1.X <= p2.X)
}

func SortyX(p1, p2 *Point) bool {
	return (p1.Y < p2.Y) || (p1.Y == p2.Y && p1.X >= p2.X)
}

func Sortyx(p1, p2 *Point) bool {
	return (p1.Y < p2.Y) || (p1.Y == p2.Y && p1.X <= p2.X)
}

// SortFromRadial Sorts using a tangential from a center point
// this sorter is different from the rest in that it returns the actual sort
// as it needs an origin
func SortRadial(x, y float32) func(*Point, *Point) bool {
	return func(p1, p2 *Point) bool {
		return tan(p1, x, y) >= tan(p2, x, y)
	}
}

func SortCounterRadial(x, y float32) func(*Point, *Point) bool {
	return func(p1, p2 *Point) bool {
		return tan(p1, x, y) <= tan(p2, x, y)
	}
}

func tan(p1 *Point, x, y float32) float32 {
	var tan1 float32
	x1, y1 := movePoint(p1, x, y)
	if x1 != 0.0 {
		tan1 = y1 / x1
	} else {
		tan1 = 1.0
	}
	return tan1
}

func movePoint(p1 *Point, x, y float32) (float32, float32) {
	return p1.X - x, p1.Y - y
}

// SortDiagonal \
func SortDiagonal(p1, p2 *Point) bool {
	return (p1.X + p1.Y) >= (p2.X + p2.Y)
}

// SortDiagonal2 /
func SortDiagonal2(p1, p2 *Point) bool {
	return (p1.X + (1.0 - p1.Y)) >= (p2.X + (1.0 - p2.Y))
}

// Quicksort implementation. This one takes a sorting function
// instead of asuming that the less than is the only posible ordering.
// It's recursive, so it may fail for very big datasets
func QuickSort(data PointList, sortFunc SortFunction) {
	if len(data) < 2 {
		return
	}
	pivot := data[0]
	l, r := 1, len(data)-1
	for l <= r {
		for l <= r && sortFunc(pivot, data[l]) {
			l++
		}
		for r >= l && !sortFunc(pivot, data[r]) {
			r--
		}
		if l < r {
			data[l], data[r] = data[r], data[l]
		}
	}
	if r > 0 {
		data[0], data[r] = data[r], data[0]
		QuickSort(data[0:r], sortFunc)
	}
	QuickSort(data[l:], sortFunc)
}
