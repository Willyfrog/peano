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
