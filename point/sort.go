package point

import (
	"fmt"
)

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

// Quicksort sorts a list of points based on one of the sorting functions
func (pnts *PointList) QuickSort(sortFunc SortFunction) PointList {
	fmt.Println(fmt.Sprintf("Quicksorting %d elements", len(*pnts)))
	pnts.quicksort(0, len(*pnts)-1, sortFunc)
	return *pnts
}
func (pnts *PointList) quicksort(lo, hi int, sortFunc SortFunction) {
	fmt.Println(fmt.Sprintf("QS [%d, %d]", lo, hi))

	if hi > lo {
		p := pnts.partition(lo, hi, sortFunc)
		fmt.Println(fmt.Sprintf("QS'p is %d", p))
		pnts.quicksort(lo, p, sortFunc)
		pnts.quicksort(p+1, hi, sortFunc)
	}
}

func (pnts *PointList) partition(lo, hi int, sortFunc SortFunction) int {
	A := *pnts
	pivot := A[lo]
	i := lo - 1
	j := hi + 1
	for {

		j = j - 1
		for sortFunc(A[j], pivot) {
			j = j - 1
		}

		i = i + 1
		for sortFunc(pivot, A[i]) {
			i = i + 1
		}

		if i < j {
			A[j], A[i] = A[i], A[j]
		} else {
			return j
		}
	}
}

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
