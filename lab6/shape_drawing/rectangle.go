package shape_drawing

import "ood/lab6/graphics"

type rectangle struct {
	leftTop       Point
	width, height int
	color         uint32
}

func NewRectangle(leftTop Point, width, height int, color uint32) Drawable {
	return &rectangle{
		leftTop: leftTop,
		width:   width,
		height:  height,
		color:   color,
	}
}

func (r rectangle) Draw(canvas graphics.Canvas) {
	canvas.SetColor(r.color)
	canvas.MoveTo(r.leftTop.X, r.leftTop.Y)
	canvas.LineTo(r.leftTop.X+r.width, r.leftTop.Y)
	canvas.LineTo(r.leftTop.X+r.width, r.leftTop.Y+r.height)
	canvas.LineTo(r.leftTop.X, r.leftTop.Y+r.height)
	canvas.LineTo(r.leftTop.X, r.leftTop.Y)
}
