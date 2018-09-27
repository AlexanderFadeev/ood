package shape

import (
	"ood/lab4/point"

	"github.com/pkg/errors"
)

type RegularPolygon struct {
	vertices uint
	center   point.Point
	radius   float64
}

func NewRegularPolygon(vertices uint, center point.Point, radius float64) (*RegularPolygon, error) {
	if vertices < 3 {
		return nil, errors.New("Too few vertices")
	}
	if radius < 0 {
		return nil, errors.New("Negative radius")
	}

	return &RegularPolygon{
		vertices: vertices,
		center:   center,
		radius:   radius,
	}, nil
}

func (rp RegularPolygon) GetVerticesCount() uint {
	return rp.vertices
}

func (rp RegularPolygon) GetCenter() point.Point {
	return rp.center
}

func (rp RegularPolygon) GetRadius() float64 {
	return rp.radius
}
