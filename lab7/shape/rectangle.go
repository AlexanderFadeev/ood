package shape

import (
	"github.com/AlexanderFadeev/ood/lab7/canvas"
	"github.com/AlexanderFadeev/ood/lab7/point"
	"github.com/AlexanderFadeev/ood/lab7/rect"
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

func (r *Rectangle) Draw(c canvas.Canvas) {
	rightTop := point.Point{X: r.Rect.RightBottom.X, Y: r.Rect.LeftTop.Y}
	leftBottom := point.Point{X: r.Rect.LeftTop.X, Y: r.Rect.RightBottom.Y}

	r.styles.apply(c)
	c.MoveTo(r.Rect.LeftTop)
	c.LineTo(rightTop)
	c.LineTo(r.Rect.RightBottom)
	c.LineTo(leftBottom)
	c.LineTo(r.Rect.LeftTop)
}
