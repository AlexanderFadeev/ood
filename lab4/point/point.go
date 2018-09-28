package point

import "math"

type Point struct {
	X float64
	Y float64
}

func (p Point) DistTo(other Point) float64 {
	return math.Hypot(p.X-other.X, p.Y-other.Y)
}
