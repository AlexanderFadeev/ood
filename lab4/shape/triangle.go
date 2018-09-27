package shape

import "ood/lab4/point"

type Triangle struct {
	a point.Point
	b point.Point
	c point.Point
}

func NewTriangle(a, b, c point.Point) Triangle {
	return Triangle{
		a: a,
		b: b,
		c: c,
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
