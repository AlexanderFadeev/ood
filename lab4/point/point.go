package point

import (
	"fmt"
	"math"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (p Point) DistTo(other Point) float64 {
	return math.Hypot(p.X-other.X, p.Y-other.Y)
}

func (p Point) String() string {
	return fmt.Sprintf("(%.1f, %.1f)", p.X, p.Y)
}
