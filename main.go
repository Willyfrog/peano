package main

import (
	"flag"
	"fmt"
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/Willyfrog/peano/drawing"
	"github.com/Willyfrog/peano/matrix"
	"github.com/Willyfrog/peano/point"
	"github.com/Willyfrog/peano/strategy"
)

func main() {
	// let the user enter the number of points to be used
	numPoints := flag.Int("points", 1000, "Number of points to be generated")
	filename := flag.String("output", "curve", "name of the file where the image will be written")
	size := flag.Int("size", 1024, "Size of the file to be generated")
	debug := flag.Bool("d", false, "Print debug information")
	flag.Parse()
	if *debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug enabled")
	} else {
		log.SetLevel(log.InfoLevel)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	file := fmt.Sprintf("%s.png", *filename)
	log.Info("Starting program")
	pl := point.RandomSlice(*numPoints)
	m := matrix.FindSmallerCellSize(&pl)
	//fmt.Println("Initial matrix: ", m)
	canvas := drawing.NewCanvas(*size, file)
	m.Draw(canvas, strategy.SnakeStrategy{})
	log.Info("Saving the image (this process might take a few minutes)")
	canvas.Save()
	log.Info("Finished program")
}
