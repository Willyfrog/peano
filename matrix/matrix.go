// package matrix
package matrix

import (
	"fmt"
	"image/color"

	log "github.com/Sirupsen/logrus"
	"github.com/Willyfrog/peano/drawing"
	"github.com/Willyfrog/peano/point"
	"github.com/Willyfrog/peano/square"
)

type Matrix struct {
	Width   float32
	Squares [][]square.Square
}

type position struct {
	X int
	Y int
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

func (m *Matrix) Draw(canvas *drawing.Canvas, strat Strategy) {
	path := canvas.GetContext()
	path.SetFillColor(color.RGBA{0xaa, 0xaa, 0xaa, 0xff})
	path.SetStrokeColor(color.RGBA{0x99, 0x99, 0x99, 0xff})
	path.SetLineWidth(2)
	drawing.DrawSquare(0.0, 0.0, 1.0, 1.0, path)
	log.Info("filling the background of the image, this might take some seconds")
	path.FillStroke()
	m.drawSquares(canvas, strat)
}

func (m *Matrix) drawSquares(canvas *drawing.Canvas, strat Strategy) {
	finished := make(chan position)
	numSquares := len(m.Squares) * len(m.Squares)
	sentSquares := 0
	for i, row := range m.Squares {
		for j := range row {
			go fillSquare(m.Squares[i][j], finished, strat)
			sentSquares++
			log.Debug(fmt.Sprintf("Sent [%d, %d] %d/%d", i, j, sentSquares, numSquares))
		}
	}
	log.Debug("Waiting for squares to be filled.")
	wait := make(chan bool)
	go drawEach(finished, wait, numSquares, *m, canvas)
	_ = <-wait //synchronize
	path := canvas.GetContext()
	//path.SetFillColor(color.RGBA{0xaa, 0xaa, 0xaa, 0xff})
	path.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	path.SetLineWidth(5)
	drawing.DrawSquare(0.0, 0.0, 1.0, 1.0, path)
	//path.FillStroke()
	log.Debug("About to draw line connections")
	for i, line := range strat.ConnectSquares(*m) {
		//log.Debug(fmt.Sprintf("Drawing the %dth line: %v", i, line))
		if len(line) == 2 {
			p1 := line[0]
			p2 := line[1]
			path2 := canvas.GetContext()
			//path2.SetFillColor(color.RGBA{0xaa, 0xaa, 0xaa, 0xff})
			path2.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
			path2.SetLineWidth(1)
			drawing.DrawLine(p1.X, p1.Y, p2.X, p2.Y, path2)
			path2.FillStroke()
		} else {
			log.Error(fmt.Sprintf("Line %d didn't contain 2 points", i)) // shouldn't happen :/
		}
	}
}

// fillSquare ...
func fillSquare(sq square.Square, finished chan position, strat Strategy) {
	pos := position{sq.X, sq.Y}
	log.Debug(fmt.Sprintf("Connecting square [%d][%d]: %v", sq.X, sq.Y, pos))
	strat.OrderPoints(sq)
	finished <- pos
}

// drawSquare ...
func drawEach(squares chan position, finish chan bool, numSquares int, m Matrix, canvas *drawing.Canvas) {
	log.Debug("Waiting for squares to be filled.")
	for posFilled := range squares {
		squee := m.Squares[posFilled.X][posFilled.Y]
		squee.Draw(canvas)
		numSquares--
		if (numSquares) < 1 {
			log.Debug(fmt.Sprintf("Draw %v, waiting for %d squares left", posFilled, numSquares))
			finish <- true
			return
		}
		log.Debug(fmt.Sprintf("Draw %v, waiting for %d squares left", posFilled, numSquares))
	}
	log.Debug("Notify of the end of drawing")
	finish <- true
}
