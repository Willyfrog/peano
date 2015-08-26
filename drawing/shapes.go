package drawing

import kit "github.com/llgcode/draw2d/draw2dkit"

// DrawSquare will get a canvascontext and draw an empty square
func DrawSquare(xo, yo, xe, ye float32, path *CanvasContext) {
	xoi := path.ToScale(xo)
	yoi := path.ToScale(yo)
	xei := path.ToScale(xe)
	yei := path.ToScale(ye)
	kit.Rectangle(path, xoi, yoi, xei, yei)
}

// Drawpoint will draw a point into the canvas
func DrawPoint(x, y float32, path *CanvasContext) {
	xi := path.ToScale(x)
	yi := path.ToScale(y)
	kit.Circle(path, xi, yi, 0.1)
	//path.Set(xi, yi, color.RGBA{0x88, 0xff, 0x88, 0xff})
}

// DrawLine from (xo, yo) to (xe, ye)
func DrawLine(xo, yo, xe, ye float32, path *CanvasContext) {
	xoi := path.ToScale(xo)
	yoi := path.ToScale(yo)
	xei := path.ToScale(xe)
	yei := path.ToScale(ye)
	path.MoveTo(xoi, yoi)
	path.LineTo(xei, yei)
	path.Close()
}
