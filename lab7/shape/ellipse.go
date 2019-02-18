package shape

import (
	"ood/lab7/canvas"
	"ood/lab7/rect"
)

type Ellipse struct {
	leaf
	styles
	frame
}

func NewEllipse(rect rect.Rect) Shape {
	return &Ellipse{
		styles: *NewStyles(),
		frame:  *NewFrame(rect),
	}
}

func (e *Ellipse) Draw(c canvas.Canvas) {
	e.styles.apply(c)
	c.DrawEllipse(e.Rect)
}
