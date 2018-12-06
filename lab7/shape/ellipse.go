package shape

import (
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

func (e *Ellipse) Accept(v Visitor) {
	v.VisitEllipse(e)
}
