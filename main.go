package main

import (
	"fmt"
	"github.com/Willyfrog/peano/drawing"
	"github.com/Willyfrog/peano/matrix"
	"github.com/Willyfrog/peano/point"
)

func main() {
	fmt.Println("Starting program")
	pl := point.RandomSlice(100)
	m := matrix.FindSmallerCellSize(&pl)
	//fmt.Println("Initial matrix: ", m)
	canvas := drawing.NewCanvas(1024, "hello.png")
	m.Draw(canvas)
	canvas.Save()
	fmt.Println("Finished program")
}
