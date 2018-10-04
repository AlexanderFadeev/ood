package shape

import (
	"fmt"
	"math"

	"ood/lab4/canvas"
	"ood/lab4/color"
	"ood/lab4/point"

	"github.com/pkg/errors"
)

type RegularPolygon struct {
	shapeColor

	vertices uint
	center   point.Point
	radius   float64
}

func NewRegularPolygon(vertices uint, center point.Point, radius float64, color color.Color) (*RegularPolygon, error) {
	if vertices < 3 {
		return nil, errors.New("Too few vertices")
	}
	if radius < 0 {
		return nil, errors.New("Negative radius")
	}

	return &RegularPolygon{
		shapeColor: shapeColor(color),
		vertices:   vertices,
		center:     center,
		radius:     radius,
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

func (rp RegularPolygon) Draw(canvas canvas.Canvas) {
	canvas.SetColor(rp.GetColor())

	for index := uint(0); index < rp.vertices; index++ {
		from := rp.indexToPoint(index)
		to := rp.indexToPoint(index + 1)
		canvas.DrawLine(from, to)
	}
}

func (rp RegularPolygon) indexToPoint(index uint) point.Point {
	angle := float64(index) / float64(rp.vertices) * 2 * math.Pi
	return point.Point{
		X: rp.center.X + rp.radius*math.Cos(angle),
		Y: rp.center.Y - rp.radius*math.Sin(angle),
	}
}

func (rp RegularPolygon) String() string {
	return fmt.Sprintf("%s %d-sided regular polygon: O=%s, R=%.1f", rp.GetColor(), rp.vertices, rp.center, rp.radius)
}
