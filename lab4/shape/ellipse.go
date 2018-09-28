package shape

import (
	"ood/lab4/canvas"
	"ood/lab4/point"

	"github.com/pkg/errors"
)

type Ellipse struct {
	center           point.Point
	horizontalRadius float64
	verticalRadius   float64
}

func NewEllipse(center point.Point, horizontalRadius, verticalRadius float64) (*Ellipse, error) {
	if horizontalRadius < 0 {
		return nil, errors.New("Negative horizontal radius value")
	}
	if verticalRadius < 0 {
		return nil, errors.New("Negative vertical radius value")
	}

	return &Ellipse{
		center:           center,
		horizontalRadius: horizontalRadius,
		verticalRadius:   verticalRadius,
	}, nil
}

func (e Ellipse) GetCenter() point.Point {
	return e.center
}

func (e Ellipse) GetHorizontalRadius() float64 {
	return e.horizontalRadius
}

func (e Ellipse) GetVerticalRadius() float64 {
	return e.verticalRadius
}

func (e Ellipse) Draw(canvas canvas.Canvas) {
	topLeft := point.Point{e.center.X - e.horizontalRadius, e.center.Y - e.verticalRadius}
	canvas.DrawEllipse(topLeft, e.horizontalRadius*2, e.verticalRadius*2)
}
