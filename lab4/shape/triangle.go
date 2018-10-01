package shape

import (
	"ood/lab4/canvas"
	"ood/lab4/color"
	"ood/lab4/point"
)

type Triangle struct {
	shapeColor

	a point.Point
	b point.Point
	c point.Point
}

func NewTriangle(a, b, c point.Point, color color.Color) Triangle {
	return Triangle{
		shapeColor: shapeColor(color),
		a:          a,
		b:          b,
		c:          c,
	}
}

func (t Triangle) GetVertexA() point.Point {
	return t.a
}

func (t Triangle) GetVertexB() point.Point {
	return t.b
}

func (t Triangle) GetVertexC() point.Point {
	return t.c
}

func (t Triangle) Draw(canvas canvas.Canvas) {
	canvas.SetColor(t.GetColor())
	canvas.DrawLine(t.a, t.b)
	canvas.DrawLine(t.b, t.c)
	canvas.DrawLine(t.c, t.a)
}
