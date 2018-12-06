package shape

import (
	"ood/lab7/rect"
)

type Rectangle struct {
	leaf
	styles
	frame
}

func NewRectangle(rect rect.Rect) Shape {
	return &Rectangle{
		styles: *NewStyles(),
		frame:  *NewFrame(rect),
	}
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}
