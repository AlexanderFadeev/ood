package shape_drawing

import "ood/lab6/graphics"

type triangle struct {
	a, b, c Point
}

func NewTriangle(a, b, c Point) Drawable {
	return &triangle{
		a: a,
		b: b,
		c: c,
	}
}

func (t triangle) Draw(canvas graphics.Canvas) {
	canvas.MoveTo(t.a.X, t.a.Y)
	canvas.LineTo(t.b.X, t.b.Y)
	canvas.LineTo(t.c.X, t.c.Y)
	canvas.LineTo(t.a.X, t.a.Y)
}
