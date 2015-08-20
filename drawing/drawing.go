package drawing

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"math"
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

// DrawSquare will get a canvascontext and draw an empty square
func DrawSquare(xo, yo, xe, ye float32, path *CanvasContext) {
	xoi := path.ToScale(xo)
	yoi := path.ToScale(yo)
	xei := path.ToScale(xe)
	yei := path.ToScale(ye)
	path.MoveTo(xoi, yoi)
	path.LineTo(xei, yoi)
	path.LineTo(xei, yei)
	path.LineTo(xoi, yei)
	path.LineTo(xoi, yoi)
}

func DrawPoint(x, y float32, path *CanvasContext) {
	xi := path.ToScale(x)
	yi := path.ToScale(y)
	path.MoveTo(xi, yi)
	path.ArcTo(xi, yi, 1, 1, 0.0, 2.0*math.Pi)
}

func NewCanvas(size int, fileName string) *Canvas {
	var cnv Canvas
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	cnv = Canvas{*img, fileName, size}
	return &cnv
}

func (c *Canvas) Save() {
	draw2dimg.SaveToPngFile(c.filename, c)
}

func (c *Canvas) GetContext() *CanvasContext {
	gc := draw2dimg.NewGraphicContext(&(c.RGBA))
	cnv := CanvasContext{*gc, c.size}
	return &cnv
}
