package shape_drawing

import "ood/lab6/graphics"

type rectangle struct {
	leftTop       Point
	width, height int
}

func NewRectangle(leftTop Point, width, height int) Drawable {
	return &rectangle{
		leftTop: leftTop,
		width:   width,
		height:  height,
	}
}

func (r rectangle) Draw(canvas graphics.Canvas) {
	canvas.MoveTo(r.leftTop.X, r.leftTop.Y)
	canvas.LineTo(r.leftTop.X+r.width, r.leftTop.Y)
	canvas.LineTo(r.leftTop.X+r.width, r.leftTop.Y+r.height)
	canvas.LineTo(r.leftTop.X, r.leftTop.Y+r.height)
	canvas.LineTo(r.leftTop.X, r.leftTop.Y)
}
