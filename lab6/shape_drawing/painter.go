package shape_drawing

import "ood/lab6/graphics"

type Painter interface {
	Draw(Drawable)
}

type painter struct {
	canvas graphics.Canvas
}

func NewPainter(canvas graphics.Canvas) Painter {
	return &painter{
		canvas: canvas,
	}
}

func (p *painter) Draw(drawable Drawable) {
	drawable.Draw(p.canvas)
}
