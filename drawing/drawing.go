package drawing

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
)

const SCALE = 1000

type Canvas struct {
	image.RGBA
	filename string
	size     int
}

type CanvasContext struct {
	draw2dimg.GraphicContext
	size int
}

type Drawable interface {
	Draw(canvas *Canvas)
}

// ToScale transforms the [0.0, 1.0) space into [0, SCALE)
func (cnv CanvasContext) ToScale(x float32) float64 {
	return float64(x * float32(cnv.size))
}

func NewCanvas(size int, fileName string) *Canvas {
	var cnv *Canvas
	cnv = new(Canvas)
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	cnv.RGBA = *img
	cnv.filename = fileName
	cnv.size = size
	return cnv
}

func (c *Canvas) Save() {
	draw2dimg.SaveToPngFile(c.filename, c)
}

func (c *Canvas) GetContext() *CanvasContext {
	gc := draw2dimg.NewGraphicContext(&(c.RGBA))
	cnv := CanvasContext{*gc, (*c).size}
	return &cnv
}
