package shape

type Triangle struct {
	a Point
	b Point
	c Point
}

func NewTriangle(a, b, c Point) Triangle {
	return Triangle{
		a: a,
		b: b,
		c: c,
	}
}

func (t Triangle) GetVertexA() Point {
	return t.a
}

func (t Triangle) GetVertexB() Point {
	return t.b
}

func (t Triangle) GetVertexC() Point {
	return t.c
}
