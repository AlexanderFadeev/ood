package point

type Point struct {
	X float64
	Y float64
}

func (p *Point) Shift(vec Vector) Point {
	return Point{
		X: p.X + vec.X,
		Y: p.Y + vec.Y,
	}
}
