// package matrix
package matrix

import (
	"fmt"
	"github.com/Willyfrog/peano/drawing"
	"github.com/Willyfrog/peano/point"
	"github.com/Willyfrog/peano/square"
	"image/color"
)

type Matrix struct {
	Width   float32
	Squares [][]square.Square
}

func (m Matrix) allCellsAreFull() bool {
	for _, line := range m.Squares {
		for _, sq := range line {
			if sq.Empty() {
				return false
			}
		}
	}
	return true
}

// GetSquare
// get square in position (i,j)
func (m Matrix) GetSquare(i, j int) square.Square {
	return m.Squares[i][j]
}

// FindSmallerCellSize
// find the matrix with the smaller cell size
func FindSmallerCellSize(pointList *[]point.Point) Matrix {
	squares := make([][]square.Square, 1)
	squares[0] = make([]square.Square, 1)
	squares[0][0] = square.Square{
		Width: 1.0,
		X:     0,
		Y:     0,
		// points must be normalized into a range [0.0, 1.0)
		Points: make([]*point.Point, len(*pointList)),
	}
	for i := range *pointList {
		squares[0][0].Points[i] = &((*pointList)[i])
	}
	m := Matrix{
		Width:   1.0,
		Squares: squares[:],
	}
	return findSmallerCellSize(m)
}

func findSmallerCellSize(m Matrix) Matrix {
	current := &m
	previous := &m
	for current.allCellsAreFull() {
		previous = current
		tmp := current.Subdivide()
		current = tmp
	}
	return *previous
}

// Subdivide
// get the next iteration of matrix if every square
// of it was subidivided into a 4x4 submatrix
func (m Matrix) Subdivide() *Matrix {
	cap := len(m.Squares) * 2
	nm := New(m.Width/2.0, cap)
	for _, row := range m.Squares {
		for _, oldSq := range row {
			result := oldSq.Partition()
			for i := range result {
				for j := range result[i] {
					sq := result[i][j]
					nm.Squares[sq.X][sq.Y] = sq
				}
			}
		}
	}
	return &nm
}

// New
// creates a new matrix with a given size
func New(width float32, size int) Matrix {
	nm := Matrix{
		Width:   width,
		Squares: make([][]square.Square, size, size),
	}
	for i := range nm.Squares {
		nm.Squares[i] = make([]square.Square, size, size)
	}
	return nm
}

func (m *Matrix) Draw(canvas *drawing.Canvas) {
	fmt.Println("drawing the matrix")
	path := canvas.GetContext()
	path.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	path.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	path.SetLineWidth(5)
	drawing.DrawSquare(0.0, 0.0, 1.0, 1.0, path)
	path.Close()
	path.FillStroke()
	// for _, row := range m.Squares {
	// 	for _, sq := range row {
	// 		sq.Draw(canvas)
	// 	}
	// }
}
